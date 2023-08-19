package handler

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/ismdeep/message-center/app/server/store"
	"github.com/ismdeep/message-center/internal/model"
	"github.com/ismdeep/message-center/pkg/log"
	"github.com/ismdeep/message-center/pkg/timeutil"
)

// BucketCreate create a bucket
func BucketCreate(ctx context.Context, id string, name string) error {
	if id == "" {
		return errors.New("id is empty")
	}

	if name == "" {
		return errors.New("name is empty")
	}

	if len(id) > 36 {
		return errors.New("id too long")
	}

	if len(name) > 255 {
		return errors.New("name too long")
	}

	if err := store.BucketCreate(id, name); err != nil {
		log.WithContext(ctx).Error("failed to create bucket", zap.Error(err))
		return errors.New("failed to create bucket")
	}

	return nil
}

// BucketList get bucket list
func BucketList(ctx context.Context) ([]model.BucketInfoResp, error) {
	lst, err := store.BucketList()
	if err != nil {
		log.WithContext(ctx).Error("failed to get bucket list", zap.Error(err))
		return nil, errors.New("failed to get bucket list")
	}

	var results []model.BucketInfoResp
	for _, bucket := range lst {
		results = append(results, model.BucketInfoResp{
			ID:        bucket.ID,
			Name:      bucket.Name,
			CreatedAt: timeutil.UnixNano(bucket.CreatedAt),
		})
	}

	return results, nil
}

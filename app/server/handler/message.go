package handler

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/ismdeep/message-center/app/server/store"
	"github.com/ismdeep/message-center/pkg/log"
)

// MessageAdd push a message
func MessageAdd(ctx context.Context, bucketID string, title string, content string, tags []string) (uint64, error) {
	id, err := store.MessageAdd(bucketID, title, content, tags)
	if err != nil {
		log.WithContext(ctx).Error("failed to add message", zap.Error(err))
		return 0, errors.New("failed to add message")
	}
	return id, nil
}

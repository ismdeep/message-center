package store

import (
	"time"

	"gorm.io/gorm"

	"github.com/ismdeep/message-center/internal/model"
)

func BucketList() ([]model.Bucket, error) {
	var lst []model.Bucket
	if err := db.Find(&lst).Error; err != nil {
		return nil, err
	}

	return lst, nil
}

func BucketCreate(id string, name string) error {
	now := time.Now()
	b := model.Bucket{
		ID:        id,
		Name:      name,
		CreatedAt: now.UnixNano(),
		UpdatedAt: now.UnixNano(),
		DeletedAt: gorm.DeletedAt{}, // default
	}

	if err := db.Create(&b).Error; err != nil {
		return err
	}

	return nil
}

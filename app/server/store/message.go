package store

import (
	"strings"
	"time"

	"github.com/ismdeep/message-center/internal/model"
)

func MessageAdd(bucketID string, title string, content string, tags []string) (uint64, error) {
	m := model.Message{
		ID:        0,
		BucketID:  bucketID,
		Title:     title,
		Content:   content,
		Tag:       strings.Join(tags, ","),
		CreatedAt: time.Now().UnixNano(),
	}

	if err := db.Create(&m).Error; err != nil {
		return 0, err
	}

	return m.ID, nil
}

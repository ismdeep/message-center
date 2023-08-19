package model

import "gorm.io/gorm"

// Bucket model
type Bucket struct {
	ID        string `gorm:"type:varchar(36);primaryKey"`
	Name      string `gorm:"type:varchar(255);not null;unique"`
	CreatedAt int64  `gorm:"type:bigint unsigned;not null"`
	UpdatedAt int64  `gorm:"type:bigint unsigned;not null"`
	DeletedAt gorm.DeletedAt
}

// Message model
type Message struct {
	ID        uint64 `gorm:"type:bigint unsigned;primaryKey"`
	BucketID  string `gorm:"type:varchar(36);not null"`
	Title     string `gorm:"type:varchar(255);not null"`
	Content   string `gorm:"type:varchar(4096);not null"`
	Tag       string `gorm:"type:varchar(255);null"`
	CreatedAt int64  `gorm:"type:bigint unsigned;not null"`
}

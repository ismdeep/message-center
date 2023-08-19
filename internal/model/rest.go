package model

import "time"

// BucketCreateReq request model for create a bucket
type BucketCreateReq struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// BucketInfoResp bucket info
type BucketInfoResp struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// MessageAddReq add message request
type MessageAddReq struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

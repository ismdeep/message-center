package handler

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestMessageAdd(t *testing.T) {
	bucketID := gofakeit.UUID()
	assert.NoError(t, BucketCreate(context.TODO(), bucketID, gofakeit.Name()))

	type args struct {
		bucketID string
		title    string
		content  string
		tags     []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				bucketID: bucketID,
				title:    gofakeit.Name(),
				content:  gofakeit.BookTitle(),
				tags:     nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := MessageAdd(context.TODO(), tt.args.bucketID, tt.args.title, tt.args.content, tt.args.tags)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				t.Logf("id: %v", id)
			}
		})
	}
}

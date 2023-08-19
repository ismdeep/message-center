package handler

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestBucketCreate(t *testing.T) {
	type args struct {
		id   string
		name string
	}
	var tests = []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				id:   gofakeit.UUID(),
				name: gofakeit.Name(),
			},
			wantErr: false,
		},
		{
			name: "id is empty",
			args: args{
				id:   "",
				name: gofakeit.Name(),
			},
			wantErr: true,
		},
		{
			name: "id is too long",
			args: args{
				id:   "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
				name: gofakeit.Name(),
			},
			wantErr: true,
		},
		{
			name: "name is empty",
			args: args{
				id:   gofakeit.UUID(),
				name: "",
			},
			wantErr: true,
		},
		{
			name: "name too long",
			args: args{
				id:   gofakeit.UUID(),
				name: "name too long name too long name too long name too long name too long name too long name too longname too long name too long name too long name too long name too long name too long name too longname too long name too long name too long name too long name too long name too long name too long",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := BucketCreate(context.TODO(), tt.args.id, tt.args.name)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestBucketList(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "OK",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BucketList(context.TODO())
			assert.Equal(t, tt.wantErr, err != nil)
			t.Logf("bucket size: %v", len(got))
		})
	}
}

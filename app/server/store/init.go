package store

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ismdeep/message-center/internal/model"
)

var db *gorm.DB

func init() {
	tmpDB, err := gorm.Open(mysql.Open("root:message-center-123456@tcp(127.0.0.1:3306)/db_message_center?charset=utf8"))
	if err != nil {
		panic(err) // init
	}

	db = tmpDB

	if err := db.AutoMigrate(&model.Bucket{}, &model.Message{}); err != nil {
		panic(err)
	}
}

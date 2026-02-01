package main

import (
	"time"

	"github.com/sword-demon/gin-chat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const mysqlConfig = "root:admin888@tcp(127.0.0.1:3306)/gin-chat?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(mysql.Open(mysqlConfig), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 迁移 schame
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{
		Model: gorm.Model{
			ID: 1,
		},
		Name:          "test",
		Password:      "123456",
		Phone:         "12345678901",
		Email:         "test@example.com",
		LoginTime:     time.Now(),
		HeartbeatTime: time.Now(),
		LogoutTime:    time.Now(),
		IsLogout:      false,
		DeviceInfo:    "test",
		Identity:      "test",
		ClientIp:      "127.0.0.1",
		ClientPort:    "8081",
	}
	db.Create(user)
}

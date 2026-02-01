package models

import (
	"time"

	"github.com/sword-demon/gin-chat/utils"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LogoutTime    time.Time `gorm:"column:logout_time" json:"logout_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func NewUserBasic() *UserBasic {
	return &UserBasic{}
}

func (table *UserBasic) GetUserList() []*UserBasic {
	userList := make([]*UserBasic, 10)
	utils.DB.Find(&userList).Limit(10).Offset(0).Order("created_at desc")

	return userList
}

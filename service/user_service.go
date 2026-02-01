package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sword-demon/gin-chat/models"
)

func GetUserList(c *gin.Context) {
	userDao := models.NewUserBasic()
	userList := userDao.GetUserList()
	c.JSON(200, userList)
}

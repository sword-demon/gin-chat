package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sword-demon/gin-chat/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/index", service.GetIndex)
	r.GET("/user/list", service.GetUserList)

	return r
}

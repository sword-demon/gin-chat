package main

import (
	"github.com/sword-demon/gin-chat/router"
	"github.com/sword-demon/gin-chat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()

	r.Run(":8081")
}

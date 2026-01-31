package main

import (
	"github.com/sword-demon/gin-chat/router"
)

func main() {
	r := router.Router()

	r.Run(":8081")
}

package main

import (
	"gin/src/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()
	router.SetTrustedProxies(nil)

	handler.Init(router)

	router.Run(":3000")
}

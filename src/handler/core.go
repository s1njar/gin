package handler

import (
	"github.com/gin-gonic/gin"
)

type Context struct{}

func Init(router *gin.Engine) {
	ctx := &Context{}

	router.GET("/", ctx.health)
}

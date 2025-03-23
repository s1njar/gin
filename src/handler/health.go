package handler

import "github.com/gin-gonic/gin"

func (ctx *Context) health(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok"})
}

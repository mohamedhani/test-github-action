package router

import (
	"github.com/gin-gonic/gin"
)

func ping_route(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "pong"})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", ping_route)
	return r
}

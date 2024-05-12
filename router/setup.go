package router

import (
	"github.com/gin-gonic/gin"
)

func ping_route(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "pong"})
}
func health_check(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "healthy"})
}
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", ping_route)
	r.GET("/healthz", health_check)
	return r
}

package handlers

import (
	"github.com/gin-gonic/gin"
)

func Healthz(ctx *gin.Context) {
	ctx.String(200, "ok")
}

func Ping(ctx *gin.Context) {
	ctx.String(200, "pong")
}

func Greet(name string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"greetings": "Hello, " + name})
	}
}

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

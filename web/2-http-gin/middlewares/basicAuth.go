package middleware // the name of the package and the name of the immediate directory must be same, if different need to use alias while calling
// I purposely gave middleware and middlewares
import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth(ctx *gin.Context) {
	username := ctx.Request.Header.Get("username")
	password := ctx.Request.Header.Get("password")

	if username == "admin" && password == "admin123" {
		ctx.Next()
	} else {
		slog.Error("unauthorized user")
		ctx.String(http.StatusUnauthorized, "unauthorized user")
		ctx.Abort()
	}
}

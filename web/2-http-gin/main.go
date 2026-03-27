package main

import (
	"flag"
	"http-demo/handlers"
	middleware "http-demo/middlewares"
	"log/slog"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

var (
	PORT     string
	FILENAME string
)

const (
	DEFAULT_PORT          = "8089"
	DEFAULT_USER_FILENAME = "users.db"
)

func init() {
	PORT = os.Getenv("APP_PORT")
	FILENAME = os.Getenv("FILE_NAME")
}

func main() {

	if PORT == "" {
		flag.StringVar(&PORT, "port", DEFAULT_PORT, "--port=8089")
	}
	if FILENAME == "" {
		flag.StringVar(&FILENAME, "filename", DEFAULT_USER_FILENAME, "--filename=8089")
	}
	flag.Parse()

	slog.Info("Server started and listening on:" + PORT)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Gin based http applcation")
	})

	r.GET("/ping", handlers.Ping)
	r.GET("/greet", handlers.Greet("Jiten"))
	r.GET("healthz", handlers.Healthz)

	userhandler := handlers.NewUser(FILENAME)

	userGroup := r.Group("/private/v1", middleware.BasicAuth) // group of routers..

	userGroup.POST("/users", userhandler.Create)
	userGroup.GET("/users/:id", userhandler.GetBy)

	if err := r.Run(":" + PORT); err != nil {
		slog.Error(err.Error())
		runtime.Goexit() //it would ensure that all other goroutines would complete their executuion
	} // starts the server
}

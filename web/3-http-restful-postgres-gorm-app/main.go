package main

import (
	"demo-gorm-app/database"
	"demo-gorm-app/handlers"
	"log/slog"
	"os"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
	DSN  string
)

const (
	DEFAULT_PORT = "8089"
	DEFAULT_DSN  = `host=localhost user=postgres password=postgres dbname=user-profiles-db port=5432 sslmode=disable TimeZone=Asia/Shanghai`
)

func init() {
	PORT = os.Getenv("APP_PORT")
	DSN = os.Getenv("DB_URL")

	if PORT == "" {
		PORT = DEFAULT_PORT
	}
	if DSN == "" {
		DSN = DEFAULT_DSN
	}
}

func shortFile(file string) string {
	parts := strings.Split(file, "/")
	return parts[len(parts)-1]
}

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	logger.Info("hello world")
	r := gin.Default()

	r.GET("/ping", handlers.Ping)
	r.GET("/healthz", handlers.Healthz)

	db, err := database.Connect(DSN)

	if err != nil {
		slog.Error(err.Error())
		panic(err.Error())
	}

	iuserdb := database.NewUserDB(db, logger)
	userhandler := handlers.NewUser(iuserdb)

	userGroup := r.Group("/v1/private") // yet to implement the auth part for private

	userGroup.POST("/users", userhandler.CreateUser)
	userGroup.GET("/users/profiles/:id", userhandler.GetUserByID)
	userGroup.GET("/users/:id", userhandler.GetOnlyUserByID)

	if err := r.Run(":" + PORT); err != nil {
		slog.Error(err.Error())
		runtime.Goexit()
	}

}

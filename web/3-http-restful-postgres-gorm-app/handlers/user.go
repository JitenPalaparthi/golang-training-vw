package handlers

import (
	"demo-gorm-app/contracts"
	"demo-gorm-app/models"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	contracts.IUserDB // promoted field
}

func NewUser(iuserDb contracts.IUserDB) *User {
	return &User{iuserDb}
}

func (u *User) CreateUser(ctx *gin.Context) {
	// take bytes.. convert them into User struct
	user := new(models.User)

	err := ctx.Bind(user) // automatically bind the data from r.body to the object, as long as data is jaon,xml orhtml

	err = user.Validate()
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}

	user.Status = "active"
	//user.Id = uint(rand.IntN(10000))

	user.LastModified = time.Now().Unix()

	fmt.Println(user)

	user, err = u.Insert(user)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}
	ctx.JSON(201, user)
}

func (u *User) GetUserByID(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		slog.Error("id not provided or invalid id")
		ctx.String(http.StatusBadRequest, "id not prodided or invalid")
		ctx.Abort()
	}

	_id, err := strconv.Atoi(id)

	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, "id not prodided or invalid")
		ctx.Abort()
	}

	user, err := u.GetByID(uint(_id))

	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, "no data is available or something went wrong")
		ctx.Abort()
	}

	ctx.JSON(200, user)

}

func (u *User) GetOnlyUserByID(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		slog.Error("id not provided or invalid id")
		ctx.String(http.StatusBadRequest, "id not prodided or invalid")
		ctx.Abort()
	}

	_id, err := strconv.Atoi(id)

	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, "id not prodided or invalid")
		ctx.Abort()
	}

	user, err := u.GetOnlyUserID(uint(_id))

	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, "no data is available or something went wrong")
		ctx.Abort()
	}

	ctx.JSON(200, user)

}

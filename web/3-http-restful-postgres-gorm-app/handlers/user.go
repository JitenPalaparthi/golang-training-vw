package handlers

import (
	"demo-gorm-app/contracts"
	"demo-gorm-app/models"
	"net/http"
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

	user, err = u.Insert(user)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}
	ctx.JSON(201, user)
}

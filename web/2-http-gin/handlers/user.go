package handlers

import (
	"http-demo/fileops"
	"http-demo/models"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	FileName string
}

func NewUser(fileName string) *User {
	go fileops.SaveToFileCh(fileName)
	return &User{fileName}
}

func (u *User) Create(ctx *gin.Context) {
	// take bytes.. convert them into User struct
	user := new(models.User)

	err := ctx.Bind(user) // automatically bind the data from r.body to the object, as long as data is jaon,xml orhtml

	err = user.Validate()
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}

	user.Status = "active"
	user.Id = uint(rand.IntN(10000))

	user.LastModified = time.Now().Unix()

	fileops.UserCh <- user // send the value to channel

	ctx.JSON(201, user)

}

func (u *User) GetBy(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.String(400, "invalid path parameter")
		ctx.Abort()
	}
	user, err := fileops.ReadFileByID(u.FileName, id)
	if err != nil {
		ctx.String(400, err.Error())
		ctx.Abort()
	}
	ctx.JSON(200, user)
}

package handlers

import (
	"encoding/json"
	"http-demo/fileops"
	"http-demo/models"
	"io"
	"math/rand/v2"
	"net/http"
	"time"
)

type User struct {
	FileName string
}

func NewUser(fileName string) *User {
	go fileops.SaveToFileCh(fileName)
	return &User{fileName}
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// take bytes.. convert them into User struct

		user := new(models.User)

		err = json.Unmarshal(bytes, user) // deserialization
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		err = user.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		user.Status = "active"
		user.Id = uint(rand.IntN(10000))

		user.LastModified = time.Now().Unix()

		fileops.UserCh <- user // send the value to channel

		//err = fileops.SaveToFile(u.FileName, user)

		// if err != nil {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }

		w.WriteHeader(201)
		w.Write([]byte("User successfully stored in the file"))

	} else {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
}

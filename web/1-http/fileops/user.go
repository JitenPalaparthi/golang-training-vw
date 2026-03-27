package fileops

import (
	"encoding/json"
	"http-demo/models"
	"log/slog"
	"os"
)

var UserCh chan *models.User

func init() {
	UserCh = make(chan *models.User, 10)
}

func SaveToFileCh(filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		slog.Error(err.Error())
	}
	defer f.Close()

	slog.Info("The SavetoFileCh has started running and processing to save users to file")
	for user := range UserCh {
		bytes, err := json.Marshal(user)
		if err != nil {
			slog.Error(err.Error())
		}
		bytes = append(bytes, 10)
		_, err = f.Write(bytes)
		if err != nil {
			slog.Error(err.Error())
		}
	}
}

func SaveToFile(filename string, user *models.User) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	bytes = append(bytes, 10)
	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

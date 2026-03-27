package models

import "errors"

type User struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Status       string `json:"status"`
	LastModified int64  `json:"last_modified" gorm:"column:last_modified;index" `
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("invalid name")
	}
	if u.Email == "" {
		return errors.New("invalid email")
	}
	return nil
}

// I do serialize and deserialize
// `` --> tags.. Go uses reflection for this internally
// there are many tags , based on the format ,even tags can be used for kind of validations etc..

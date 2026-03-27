package models

import "errors"

type User struct {
	Id           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Profiles     []Profile `json:"profiles,omitempty" "foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status       string    `json:"status"`
	LastModified int64     `json:"last_modified" gorm:"column:last_modified;index" `
}

type Profile struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	UserID       uint   `json:"user_id" gorm:"not null;index"`
	ProfileName  string `json:"profile_name"`
	Bio          string `json:"bio"`
	Status       string `json:"status"`
	LastModified int64  `json:"last_modified" gorm:"column:last_modified;index" `
}

// one min

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

package database

import (
	"demo-gorm-app/contracts"
	"demo-gorm-app/models"

	"gorm.io/gorm"
)

type UserDB struct {
	*gorm.DB // promoted field
}

func NewUserDB(db *gorm.DB) contracts.IUserDB {
	return &UserDB{db}
}

func (udb *UserDB) Insert(user *models.User) (*models.User, error) {
	udb.AutoMigrate(&models.User{}) // creates a table
	tx := udb.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

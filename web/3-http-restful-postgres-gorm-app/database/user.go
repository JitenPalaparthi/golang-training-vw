package database

import (
	"demo-gorm-app/contracts"
	"demo-gorm-app/models"
	"log/slog"

	"gorm.io/gorm"
)

type UserDB struct {
	*gorm.DB // promoted field
	Logger   *slog.Logger
}

func NewUserDB(db *gorm.DB, logger *slog.Logger) contracts.IUserDB {
	return &UserDB{db, logger}
}

func (udb *UserDB) Insert(user *models.User) (*models.User, error) {
	udb.AutoMigrate(&models.User{}, &models.Profile{}) // creates a table
	tx := udb.Create(user)
	if tx.Error != nil {
		udb.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return user, nil
}

func (udb *UserDB) GetByID(id uint) (*models.User, error) {
	var user *models.User = new(models.User)
	tx := udb.Preload("Profiles").First(user, id)
	if tx.Error != nil {
		udb.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return user, nil
}

func (udb *UserDB) GetOnlyUserID(id uint) (*models.User, error) {
	var user *models.User = new(models.User)
	tx := udb.Omit("Profiles").First(user, id)
	if tx.Error != nil {
		udb.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return user, nil
}

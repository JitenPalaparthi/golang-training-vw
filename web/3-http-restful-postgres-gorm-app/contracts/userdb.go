package contracts

import "demo-gorm-app/models"

type IUserDB interface {
	Insert(user *models.User) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	GetOnlyUserID(id uint) (*models.User, error)
}

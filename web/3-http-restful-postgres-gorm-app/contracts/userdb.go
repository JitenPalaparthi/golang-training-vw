package contracts

import "demo-gorm-app/models"

type IUserDB interface {
	Insert(user *models.User) (*models.User, error)
}

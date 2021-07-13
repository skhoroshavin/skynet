package spi

import (
	"skynet/domain/models"
)

type UsersRepository interface {
	Insert(id string, password string) error

	UpdatePassword(id string, password string) error
	UpdateUserData(id string, data *models.UserData) error

	Password(id string) (string, error)
	UserData(id string) (*models.UserData, error)
}

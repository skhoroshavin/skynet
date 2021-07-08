package api

import (
	"skynet/domain/models"
)

type Users interface {
	CreateUser(id string, password string) error
	UpdatePassword(id string, oldPassword string, newPassword string) error
	UpdateUserData(id string, data *models.UserData) error

	UserData(id string) (*models.UserData, error)
}

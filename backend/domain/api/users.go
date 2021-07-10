package api

import (
	"skynet/domain/models"
)

type Users interface {
	UpdateUserData(id string, data *models.UserData) error

	UserData(id string) (*models.UserData, error)
}

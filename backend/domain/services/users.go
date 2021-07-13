package services

import (
	"skynet/domain/models"
	"skynet/domain/spi"
)

type UserService struct {
	storage spi.Storage
}

func NewUserService(storage spi.Storage) UserService {
	return UserService{
		storage: storage,
	}
}

func (u UserService) UpdateUserData(id string, data *models.UserData) error {
	return u.storage.Transaction(func(sd spi.StorageData) error {
		return sd.Users.UpdateUserData(id, data)
	})
}

func (u UserService) UserData(id string) (*models.UserData, error) {
	var result *models.UserData
	err := u.storage.Transaction(func(sd spi.StorageData) error {
		res, err := sd.Users.UserData(id)
		if err == nil {
			result = res
		}
		return err
	})
	return result, err
}

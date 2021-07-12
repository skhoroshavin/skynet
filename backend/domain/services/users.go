package services

import (
	"skynet/domain/models"
	"skynet/domain/spi"
)

type UserService struct {
	storage spi.UsersStorage
}

func NewUserService(users spi.UsersStorage) UserService {
	return UserService{
		storage: users,
	}
}

func (u UserService) UpdateUserData(id string, data *models.UserData) error {
	return spi.Transactional(u.storage, func(users spi.UsersRepository) error {
		return users.UpdateUserData(id, data)
	})
}

func (u UserService) UserData(id string) (*models.UserData, error) {
	var result *models.UserData
	err := spi.Transactional(u.storage, func(users spi.UsersRepository) error {
		res, err := users.UserData(id)
		if err == nil {
			result = res
		}
		return err
	})
	return result, err
}

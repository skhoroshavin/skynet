package services

import (
	"errors"
	"skynet/domain/spi"
)

type AuthService struct {
	users spi.UsersStorage
}

func NewAuthService(users spi.UsersStorage) AuthService {
	return AuthService{
		users: users,
	}
}

func (a AuthService) CreateUser(id string, password string) error {
	if len(password) < 1 {
		return errors.New("password cannot be empty")
	}

	return spi.Transactional(a.users, func(users spi.UsersRepository) error {
		return users.Insert(id, password)
	})
}

func (a AuthService) UpdatePassword(id string, oldPassword string, newPassword string) error {
	if len(newPassword) < 1 {
		return errors.New("password cannot be empty")
	}

	return spi.Transactional(a.users, func(users spi.UsersRepository) error {
		password, err := users.Password(id)
		if err != nil {
			return err
		}

		if password != oldPassword {
			return errors.New("invalid old password")
		}

		return users.UpdatePassword(id, newPassword)
	})
}

func (a AuthService) CreateSession(id string) (string, error) {
	panic("not implemented")
}

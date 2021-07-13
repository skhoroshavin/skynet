package services

import (
	"errors"
	"skynet/domain/spi"
)

type AuthService struct {
	storage spi.Storage
}

func NewAuthService(storage spi.Storage) AuthService {
	return AuthService{
		storage: storage,
	}
}

func (a AuthService) SignUp(id string, password string) (string, error) {
	if len(password) < 1 {
		return "", errors.New("password cannot be empty")
	}

	var session string
	err := a.storage.Transaction(func(r spi.Repositories) error {
		err := r.Users().Insert(id, password)
		if err != nil {
			return err
		}

		// TODO: Implement session creation
		//session, err = data.Sessions.Create()
		return err
	})
	if err != nil {
		return "", err
	}

	return session, nil
}

func (a AuthService) UserID(session string) (string, error) {
	panic("implement me")
}

func (a AuthService) UpdatePassword(id string, oldPassword string, newPassword string) error {
	if len(newPassword) < 1 {
		return errors.New("password cannot be empty")
	}

	return a.storage.Transaction(func(r spi.Repositories) error {
		password, err := r.Users().Password(id)
		if err != nil {
			return err
		}

		if password != oldPassword {
			return errors.New("invalid old password")
		}

		return r.Users().UpdatePassword(id, newPassword)
	})
}

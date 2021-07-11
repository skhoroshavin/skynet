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

	err := a.users.Insert(id, password)
	if err != nil {
		return err
	}

	return nil
}

func (a AuthService) UpdatePassword(id string, oldPassword string, newPassword string) error {
	if len(newPassword) < 1 {
		return errors.New("password cannot be empty")
	}

	txn, err := a.users.Transaction(id)
	if err != nil {
		return errors.New("failed to start transaction")
	}
	defer txn.Rollback()

	if !txn.MatchPassword(oldPassword) {
		return errors.New("invalid old password")
	}

	txn.UpdatePassword(newPassword)

	if txn.Commit() != nil {
		return errors.New("failed to change password")
	}

	return nil
}

func (a AuthService) CreateSession(id string) (string, error) {
	panic("not implemented")
}

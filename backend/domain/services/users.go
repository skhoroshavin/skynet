package services

import (
	"errors"
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

func (u UserService) CreateUser(id string, password string) error {
	if len(password) < 1 {
		return errors.New("password cannot be empty")
	}

	err := u.storage.Insert(id, password)
	if err != nil {
		return err
	}

	return nil
}

func (u UserService) UpdatePassword(id string, oldPassword string, newPassword string) error {
	if len(newPassword) < 1 {
		return errors.New("password cannot be empty")
	}

	txn, err := u.storage.Transaction(id)
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

func (u UserService) UpdateUserData(id string, data *models.UserData) error {
	txn, err := u.storage.Transaction(id)
	if err != nil {
		return errors.New("failed to start transaction")
	}
	defer txn.Rollback()

	txn.UpdateUserData(data)

	if txn.Commit() != nil {
		return errors.New("failed to update user data")
	}

	return nil
}

func (u UserService) UserData(id string) (*models.UserData, error) {
	txn, err := u.storage.Transaction(id)
	if err != nil {
		return nil, errors.New("failed to start transaction")
	}
	defer txn.Rollback()

	return txn.UserData(), nil
}

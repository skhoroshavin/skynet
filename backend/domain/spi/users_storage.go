package spi

import (
	"skynet/domain/models"
)

type UsersStorage interface {
	Begin() (UsersRepository, error)
}

type UsersRepository interface {
	Save() error
	Cancel()

	Insert(id string, password string) error

	UpdatePassword(id string, password string) error
	UpdateUserData(id string, data *models.UserData) error

	Password(id string) (string, error)
	UserData(id string) (*models.UserData, error)
}

func Transactional(storage UsersStorage, txn func(users UsersRepository) error) error {
	users, err := storage.Begin()
	if err != nil {
		return err
	}

	err = txn(users)
	if err != nil {
		users.Cancel()
		return err
	}

	return users.Save()
}

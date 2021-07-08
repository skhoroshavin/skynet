package spi

import (
	"skynet/domain/models"
)

type UsersStorage interface {
	Insert(id string, password string) error
	Transaction(id string) (UserTransaction, error)
}

type UserTransaction interface {
	Commit() error
	Rollback()

	UpdatePassword(password string)
	UpdateUserData(data *models.UserData)

	MatchPassword(password string) bool
	UserData() *models.UserData
}

package services

import (
	"errors"
	"skynet/domain/models"
	"skynet/domain/spi"
	"skynet/domain/spi_testing"
	"testing"
)

type TestUsersStorage struct {
	passwords map[string]string
	data      map[string]models.UserData
}

func newTestUsersStorage() *TestUsersStorage {
	r := &TestUsersStorage{}
	r.reset()
	return r
}

func (u *TestUsersStorage) reset() {
	u.passwords = map[string]string{}
	u.data = map[string]models.UserData{}
}

func (u *TestUsersStorage) Begin() (spi.UsersRepository, error) {
	return TestUsersRepository{
		storage: u,
		passwords: u.passwords,
		data: u.data,
	}, nil
}

type TestUsersRepository struct {
	storage *TestUsersStorage

	passwords map[string]string
	data      map[string]models.UserData
}

func (u TestUsersRepository) Save() error {
	u.storage.passwords = u.passwords
	u.storage.data = u.data
	return nil
}

func (u TestUsersRepository) Cancel() {

}

func (u TestUsersRepository) Insert(id string, password string) error {
	_, ok := u.passwords[id]
	if ok {
		return errors.New("user already exists")
	}

	u.passwords[id] = password
	u.data[id] = models.UserData{}
	return nil
}

func (u TestUsersRepository) UpdatePassword(id string, password string) error {
	_, ok := u.passwords[id]
	if !ok {
		return errors.New("user not found")
	}

	u.passwords[id] = password
	return nil
}

func (u TestUsersRepository) UpdateUserData(id string, data *models.UserData) error {
	_, ok := u.passwords[id]
	if !ok {
		return errors.New("user not found")
	}

	u.data[id] = *data
	return nil
}

func (u TestUsersRepository) Password(id string) (string, error) {
	pass, ok := u.passwords[id]
	if !ok {
		return "", errors.New("user not found")
	}
	return pass, nil
}

func (u TestUsersRepository) UserData(id string) (*models.UserData, error) {
	data, ok := u.data[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return &data, nil
}

func Test_TestUserStorage(t *testing.T) {
	spi_testing.UsersStorageTestSuite(t, newTestUsersStorage())
}

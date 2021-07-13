package services

import (
	"errors"
	"skynet/domain/models"
	"skynet/domain/spi"
	"skynet/domain/spi_testing"
	"testing"
)

type TestUsersRepository struct {
	passwords map[string]string
	data      map[string]models.UserData
}

func (u *TestUsersRepository) reset() {
	u.passwords = map[string]string{}
	u.data = map[string]models.UserData{}
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


type TestStorage struct {
	users TestUsersRepository
}

func testStorage() *TestStorage {
	r := &TestStorage{}
	r.reset()
	return r
}

func (s *TestStorage) reset() {
	s.users.reset()
}

func (s *TestStorage) Transaction(wrk func(sd spi.StorageData) error) error {
	sd := spi.StorageData{
		Users: s.users,
	}

	err := wrk(sd)
	if err != nil {
		return err
	}

	s.users = sd.Users.(TestUsersRepository)
	return nil
}

func Test_TestUserStorage(t *testing.T) {
	spi_testing.UsersStorageTestSuite(t, testStorage())
}

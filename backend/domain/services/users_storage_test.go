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

type TestUsersStorageTxn struct {
	storage *TestUsersStorage

	id       string
	password string
	data     models.UserData
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

func (u *TestUsersStorage) Insert(id string, password string) error {
	_, ok := u.passwords[id]
	if ok {
		return errors.New("user already exists")
	}

	u.passwords[id] = password
	return nil
}

func (u *TestUsersStorage) Transaction(id string) (spi.UserTransaction, error) {
	pwd, ok := u.passwords[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &TestUsersStorageTxn{
		storage:  u,
		id:       id,
		password: pwd,
		data:     u.data[id],
	}, nil
}

func (t *TestUsersStorageTxn) Commit() error {
	t.storage.passwords[t.id] = t.password
	t.storage.data[t.id] = t.data
	return nil
}

func (t *TestUsersStorageTxn) Rollback() {

}

func (t *TestUsersStorageTxn) UpdatePassword(password string) {
	t.password = password
}

func (t *TestUsersStorageTxn) UpdateUserData(data *models.UserData) {
	t.data = *data
}

func (t *TestUsersStorageTxn) MatchPassword(password string) bool {
	return t.password == password
}

func (t *TestUsersStorageTxn) UserData() *models.UserData {
	return &t.data
}

func Test_TestUserStorage(t *testing.T) {
	spi_testing.UsersStorageTestSuite(t, newTestUsersStorage())
}

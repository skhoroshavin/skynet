package services

import (
	"github.com/stretchr/testify/mock"
	"skynet/domain/models"
	"skynet/domain/spi"
)

type StorageMock struct {
	users    UsersRepositoryMock
	sessions SessionsRepositoryMock
}

func testStorage() *StorageMock {
	return &StorageMock{}
}

func (s *StorageMock) Transaction(wrk func(r spi.Repositories) error) error {
	return wrk(s)
}

func (s *StorageMock) Users() spi.UsersRepository {
	return &s.users
}

func (s *StorageMock) Sessions() spi.SessionsRepository {
	return &s.sessions
}

type UsersRepositoryMock struct {
	mock.Mock
}

func (u *UsersRepositoryMock) Insert(id string, password string) error {
	args := u.Called(id, password)
	return args.Error(0)
}

func (u *UsersRepositoryMock) UpdatePassword(id string, password string) error {
	args := u.Called(id, password)
	return args.Error(0)
}

func (u *UsersRepositoryMock) UpdateUserData(id string, data *models.UserData) error {
	panic("implement me")
}

func (u *UsersRepositoryMock) Password(id string) (string, error) {
	args := u.Called(id)
	return args.String(0), args.Error(1)
}

func (u *UsersRepositoryMock) UserData(id string) (*models.UserData, error) {
	panic("implement me")
}

type SessionsRepositoryMock struct {
	mock.Mock
}

func (s *SessionsRepositoryMock) Insert(session string, id string) error {
	args := s.Called(session, id)
	return args.Error(0)
}

func (s *SessionsRepositoryMock) Delete(session string) error {
	panic("implement me")
}

func (s *SessionsRepositoryMock) UserID(session string) (string, error) {
	panic("implement me")
}


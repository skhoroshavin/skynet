package services

import "skynet/domain/spi"

type TestRepositories struct {
	users TestUsersRepository
}

func (r TestRepositories) Users() spi.UsersRepository {
	return &r.users
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

func (s *TestStorage) Transaction(wrk func(r spi.Repositories) error) error {
	r := TestRepositories{
		s.users,
	}

	err := wrk(r)
	if err != nil {
		return err
	}

	s.users = r.users
	return nil
}

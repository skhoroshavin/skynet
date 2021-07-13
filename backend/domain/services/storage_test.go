package services

import "skynet/domain/spi"

type TestSessionsRepository struct {
}

func (t TestSessionsRepository) CreateSession(id string) string {
	panic("implement me")
}

type TestRepositories struct {
	users    TestUsersRepository
	sessions TestSessionsRepository
}

func (r TestRepositories) Users() spi.UsersRepository {
	return &r.users
}

func (r TestRepositories) Sessions() spi.SessionsRepository {
	return &r.sessions
}

type TestStorage struct {
	users    TestUsersRepository
	sessions TestSessionsRepository
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
		s.sessions,
	}

	err := wrk(r)
	if err != nil {
		return err
	}

	s.users = r.users
	return nil
}

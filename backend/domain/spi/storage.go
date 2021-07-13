package spi

type Storage interface {
	Transaction(wrk func(r Repositories) error) error
}

type Repositories interface {
	Users() UsersRepository
}

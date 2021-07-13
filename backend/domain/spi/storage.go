package spi

type StorageData struct {
	Users UsersRepository
}

type Storage interface {
	Transaction(wrk func(sd StorageData) error) error
}

package spi

type SessionsRepository interface {
	Insert(session string, id string) error
	Delete(session string) error

	UserID(session string) (string, error)
}

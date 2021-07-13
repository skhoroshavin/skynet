package spi

type SessionsRepository interface {
	CreateSession(id string) string
}

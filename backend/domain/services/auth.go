package services

import (
	"errors"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"skynet/domain/spi"
)

type AuthService struct {
	storage spi.Storage
}

func NewAuthService(storage spi.Storage) AuthService {
	return AuthService{
		storage: storage,
	}
}

func (a AuthService) SignUp(id string, password string) (string, error) {
	if len(password) < 1 {
		return "", errors.New("password cannot be empty")
	}

	var session string
	err := a.storage.Transaction(func(r spi.Repositories) error {
		err := r.Users().Insert(id, encryptPassword(password))
		if err != nil {
			return err
		}

		uid, err := ksuid.NewRandom()
		if err != nil {
			return err
		}

		session = uid.String()
		return r.Sessions().Insert(session, id)
	})
	if err != nil {
		return "", err
	}

	return session, nil
}

func (a AuthService) UserID(session string) (string, error) {
	panic("implement me")
}

func (a AuthService) UpdatePassword(id string, oldPassword string, newPassword string) error {
	if len(newPassword) < 1 {
		return errors.New("password cannot be empty")
	}

	return a.storage.Transaction(func(r spi.Repositories) error {
		encryptedPassword, err := r.Users().Password(id)
		if err != nil {
			return err
		}

		if !checkPassword(encryptedPassword, oldPassword) {
			return errors.New("invalid old password")
		}

		return r.Users().UpdatePassword(id, encryptPassword(newPassword))
	})
}

func encryptPassword(password string) string {
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("failed to encrypt password")
	}
	return string(res)
}

func checkPassword(encryptedPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password)) == nil
}

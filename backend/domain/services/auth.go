package services

import (
	"errors"
	"fmt"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"skynet/domain/spi"
	"unicode"
)

type AuthService struct {
	storage spi.Storage
}

func NewAuthService(storage spi.Storage) AuthService {
	return AuthService{
		storage: storage,
	}
}

func validateUserId(id string) error {
	if len(id) < 1 {
		return errors.New("user id cannot be empty")
	}

	if len(id) > 30 {
		return errors.New("user id cannot be longer than 24 symbols")
	}

	for _, c := range id {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) && c != '.' && c != '_' {
			return fmt.Errorf("user id must be either letter or digit, got %v", id)
		}
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 1 {
		return errors.New("password cannot be empty")
	}

	return nil
}

func (a AuthService) SignUp(id string, password string) (string, error) {
	if err := validateUserId(id); err != nil {
		return "", err
	}

	if err := validatePassword(password); err != nil {
		return "", err
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

// TODO: Test me
func (a AuthService) UserID(session string) (string, error) {
	var userID string
	err := a.storage.Transaction(func(r spi.Repositories) error {
		var err error
		userID, err = r.Sessions().UserID(session)
		return err
	})
	if err != nil {
		return "", err
	}
	return userID, nil
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
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic("failed to encrypt password")
	}
	return string(res)
}

func checkPassword(encryptedPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password)) == nil
}

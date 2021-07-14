package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAuthService_SignUp(t *testing.T) {
	t.Run("signing up new user stores password in encrypted way and creates session", func(t *testing.T) {
		storage := testStorage()
		svc := NewAuthService(storage)

		storage.users.On("Insert", "john", mock.Anything).Return(nil)
		storage.sessions.On("Insert", mock.Anything, "john").Return(nil)

		session, err := svc.SignUp("john", "easy")

		assert.Nil(t, err)
		storage.users.AssertNotCalled(t, "Insert", "john", "easy")
		storage.sessions.AssertCalled(t, "Insert", session, "john")

		encryptedPass := storage.users.Calls[0].Arguments.String(1)
		assert.True(t, checkPassword(encryptedPass, "easy"))
	})

	t.Run("cannot create user with empty password", func(t *testing.T) {
		storage := testStorage()
		svc := NewAuthService(storage)

		_, err := svc.SignUp("john", "")

		assert.Error(t, err)
	})

	t.Run("cannot create duplicate user", func(t *testing.T) {
		storage := testStorage()
		svc := NewAuthService(storage)

		storage.users.On("Insert", "john", mock.Anything).Return(fmt.Errorf("user john already exists"))

		_, err := svc.SignUp("john", "easy")

		assert.Error(t, err)
	})
}

func TestAuthService_UpdatePassword(t *testing.T) {
	t.Run("update password checks old password for validity and stores new in encrypted way", func(t *testing.T) {
		storage := testStorage()
		svc := NewAuthService(storage)

		storage.users.On("Password", "john").Return(encryptPassword("easy"), nil)
		storage.users.On("UpdatePassword", "john", mock.Anything).Return(nil)

		err := svc.UpdatePassword("john", "easy", "peasy")

		assert.Nil(t, err)
		storage.users.AssertNotCalled(t, "UpdatePassword", "john", "peasy")

		encryptedPass := storage.users.Calls[1].Arguments.String(1)
		assert.True(t, checkPassword(encryptedPass, "peasy"))
	})

	t.Run("user must exist", func(t *testing.T) {
		storage := testStorage()
		svc := NewAuthService(storage)

		storage.users.On("Password", "john").Return("", fmt.Errorf("user john not found"))

		err := svc.UpdatePassword("john", "easy", "peasy")

		assert.Error(t, err)
	})

	t.Run("cannot update password to empty", func(t *testing.T) {
		storage := testStorage()
		svc := NewAuthService(storage)

		err := svc.UpdatePassword("john", "easy", "")

		assert.Error(t, err)
	})
}

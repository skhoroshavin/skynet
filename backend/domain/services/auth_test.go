package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthService_CreateUser(t *testing.T) {
	storage := newTestUsersStorage()
	svc := NewAuthService(storage)

	t.Run("can create new user", func(t *testing.T) {
		storage.reset()

		err := svc.CreateUser("john", "easy")

		assert.Nil(t, err)
		assert.Contains(t, storage.passwords, "john")
		assert.Equal(t, "easy", storage.passwords["john"])
	})

	t.Run("cannot create user with empty password", func(t *testing.T) {
		storage.reset()

		err := svc.CreateUser("john", "")

		assert.Error(t, err)
		assert.NotContains(t, "john", storage.passwords)
	})

	t.Run("cannot create duplicate user", func(t *testing.T) {
		storage.reset()
		storage.passwords["john"] = "easy"

		err := svc.CreateUser("john", "peasy")

		assert.Error(t, err)
		assert.Contains(t, storage.passwords, "john")
		assert.Equal(t, "easy", storage.passwords["john"])
	})
}

func TestAuthService_UpdatePassword(t *testing.T) {
	storage := newTestUsersStorage()
	svc := NewAuthService(storage)

	t.Run("can update password", func(t *testing.T) {
		storage.reset()
		storage.passwords["john"] = "easy"

		err := svc.UpdatePassword("john", "easy", "peasy")

		assert.Nil(t, err)
		assert.Contains(t, storage.passwords, "john")
		assert.Equal(t, "peasy", storage.passwords["john"])
	})

	t.Run("user must exist", func(t *testing.T) {
		storage.reset()

		err := svc.UpdatePassword("john", "easy", "peasy")

		assert.Error(t, err)
		assert.NotContains(t, "john", storage.passwords)
	})

	t.Run("cannot update password to empty", func(t *testing.T) {
		storage.reset()
		storage.passwords["john"] = "easy"

		err := svc.UpdatePassword("john", "easy", "")

		assert.Error(t, err)
		assert.Contains(t, storage.passwords, "john")
		assert.Equal(t, "easy", storage.passwords["john"])
	})

	t.Run("must provide correct old password", func(t *testing.T) {
		storage.reset()
		storage.passwords["john"] = "easy"

		err := svc.UpdatePassword("john", "hard", "peasy")

		assert.Error(t, err)
		assert.Contains(t, storage.passwords, "john")
		assert.Equal(t, "easy", storage.passwords["john"])
	})
}

package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthService_SignUp(t *testing.T) {
	storage := testStorage()
	svc := NewAuthService(storage)

	t.Run("can sign up new user", func(t *testing.T) {
		storage.reset()

		_, err := svc.SignUp("john", "easy")

		assert.Nil(t, err)
		assert.Contains(t, storage.users.passwords, "john")
		assert.Equal(t, "easy", storage.users.passwords["john"])
	})

	t.Run("cannot create user with empty password", func(t *testing.T) {
		storage.reset()

		_, err := svc.SignUp("john", "")

		assert.Error(t, err)
		assert.NotContains(t, "john", storage.users.passwords)
	})

	t.Run("cannot create duplicate user", func(t *testing.T) {
		storage.reset()
		storage.users.passwords["john"] = "easy"

		_, err := svc.SignUp("john", "peasy")

		assert.Error(t, err)
		assert.Contains(t, storage.users.passwords, "john")
		assert.Equal(t, "easy", storage.users.passwords["john"])
	})
}

func TestAuthService_UpdatePassword(t *testing.T) {
	storage := testStorage()
	svc := NewAuthService(storage)

	t.Run("can update password", func(t *testing.T) {
		storage.reset()
		storage.users.passwords["john"] = "easy"

		err := svc.UpdatePassword("john", "easy", "peasy")

		assert.Nil(t, err)
		assert.Contains(t, storage.users.passwords, "john")
		assert.Equal(t, "peasy", storage.users.passwords["john"])
	})

	t.Run("user must exist", func(t *testing.T) {
		storage.reset()

		err := svc.UpdatePassword("john", "easy", "peasy")

		assert.Error(t, err)
		assert.NotContains(t, "john", storage.users.passwords)
	})

	t.Run("cannot update password to empty", func(t *testing.T) {
		storage.reset()
		storage.users.passwords["john"] = "easy"

		err := svc.UpdatePassword("john", "easy", "")

		assert.Error(t, err)
		assert.Contains(t, storage.users.passwords, "john")
		assert.Equal(t, "easy", storage.users.passwords["john"])
	})

	t.Run("must provide correct old password", func(t *testing.T) {
		storage.reset()
		storage.users.passwords["john"] = "easy"

		err := svc.UpdatePassword("john", "hard", "peasy")

		assert.Error(t, err)
		assert.Contains(t, storage.users.passwords, "john")
		assert.Equal(t, "easy", storage.users.passwords["john"])
	})
}

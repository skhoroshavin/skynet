package services

import (
	"github.com/stretchr/testify/assert"
	"skynet/domain/models"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	storage := newTestUsersStorage()
	svc := NewUserService(storage)

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

func TestUserService_UpdatePassword(t *testing.T) {
	storage := newTestUsersStorage()
	svc := NewUserService(storage)

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

func TestUserService_UserData(t *testing.T) {
	storage := newTestUsersStorage()
	svc := NewUserService(storage)

	t.Run("new user has empty user data", func(t *testing.T) {
		storage.reset()
		svc.CreateUser("john", "easy")

		data, err := svc.UserData("john")

		assert.Nil(t, err)
		assert.Empty(t, data.FirstName)
		assert.Empty(t, data.LastName)
		assert.Nil(t, data.Birthday)
		assert.Equal(t, models.GenderUndefined, data.Gender)
		assert.Empty(t, data.City)
		assert.Empty(t, data.Interests)
	})

	t.Run("cannot get data from non-existent user", func(t *testing.T) {
		storage.reset()

		_, err := svc.UserData("john")

		assert.Error(t, err)
	})

	t.Run("user can update user data", func(t *testing.T) {
		storage.reset()
		svc.CreateUser("john", "easy")

		userData := &models.UserData{
			FirstName: "",
			LastName:  "",
			Birthday:  nil,
			Gender:    models.GenderMale,
			City:      "",
			Interests: "",
		}
		err := svc.UpdateUserData("john", userData)
		assert.Nil(t, err)

	})
}

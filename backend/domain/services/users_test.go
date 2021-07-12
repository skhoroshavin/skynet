package services

import (
	"github.com/stretchr/testify/assert"
	"skynet/domain/models"
	"testing"
)

func TestUserService_UserData(t *testing.T) {
	storage := newTestUsersStorage()
	auth := NewAuthService(storage)
	svc := NewUserService(storage)

	t.Run("new user has empty user data", func(t *testing.T) {
		storage.reset()
		auth.CreateUser("john", "easy")

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
		auth.CreateUser("john", "easy")

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

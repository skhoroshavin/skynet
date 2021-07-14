package mysql_storage

import (
	"github.com/stretchr/testify/assert"
	"skynet/domain/models"
	"skynet/domain/spi"
	"testing"
	"time"
)

func TestUsers(t *testing.T) {
	config, err := createTestDatabase()
	assert.Nil(t, err)
	defer dropTestDatabase(config)

	storage, err := NewStorage(config)
	assert.Nil(t, err)

	if !assert.Nil(t, storage.CreateSchema()) {
		return
	}

	t.Run("insert user", func(t *testing.T) {
		t.Run("can insert new user", func(t *testing.T) {
			err := storage.Transaction(func(r spi.Repositories) error {
				return r.Users().Insert("john", "easy")
			})
			assert.Nil(t, err)
		})

		t.Run("cannot insert duplicate user even with same password", func(t *testing.T) {
			err := storage.Transaction(func(r spi.Repositories) error {
				return r.Users().Insert("john", "easy")
			})
			assert.Error(t, err)
		})

		t.Run("cannot insert duplicate user", func(t *testing.T) {
			err := storage.Transaction(func(r spi.Repositories) error {
				return r.Users().Insert("john", "peasy")
			})
			assert.Error(t, err)
		})
	})

	t.Run("password", func(t *testing.T) {
		t.Run("password returns user password", func(t *testing.T) {
			err := storage.Transaction(func(r spi.Repositories) error {
				password, err := r.Users().Password("john")
				assert.Equal(t, "easy", password)
				return err
			})
			assert.Nil(t, err)
		})

		t.Run("update password updates user password", func(t *testing.T) {
			err := storage.Transaction(func(r spi.Repositories) error {
				err := r.Users().UpdatePassword("john", "peasy")
				if !assert.Nil(t, err) {
					return err
				}

				password, err := r.Users().Password("john")
				if !assert.Nil(t, err) {
					return err
				}

				assert.Equal(t, "peasy", password)
				return nil
			})
			assert.Nil(t, err)
		})
	})

	t.Run("user data", func(t *testing.T) {
		t.Run("user data for new user is empty", func(t *testing.T) {
			err := storage.Transaction(func(r spi.Repositories) error {
				data, err := r.Users().UserData("john")
				if !assert.Nil(t, err) {
					return err
				}

				assert.Empty(t, data.FirstName)
				assert.Empty(t, data.LastName)
				assert.Nil(t, data.Birthday)
				assert.Equal(t, models.GenderUndefined, data.Gender)
				assert.Empty(t, data.City)
				assert.Empty(t, data.Interests)
				return nil
			})
			assert.Nil(t, err)
		})

		t.Run("user data can be fully updated", func(t *testing.T) {
			err := storage.Transaction(func(r spi.Repositories) error {
				time := time.Date(1983, 11, 18, 0, 0, 0, 0, time.UTC)
				storedData := &models.UserData{
					FirstName: "John",
					LastName:  "Doe",
					Birthday:  &time,
					Gender:    models.GenderMale,
					City:      "New Vegas",
					Interests: "Dismantling cyborgs",
				}

				err := r.Users().UpdateUserData("john", storedData)
				if !assert.Nil(t, err) {
					return err
				}

				retrievedData, err := r.Users().UserData("john")
				if !assert.Nil(t, err) {
					return err
				}

				assert.EqualValues(t, storedData, retrievedData)
				return nil
			})
			assert.Nil(t, err)
		})

		t.Run("user data can be partially updated", func(t *testing.T) {
			err := storage.Transaction(func(r spi.Repositories) error {
				data, err := r.Users().UserData("john")
				if !assert.Nil(t, err) {
					return err
				}

				data.FirstName = "Jonny"
				data.LastName = "B"
				err = r.Users().UpdateUserData("john", data)
				if !assert.Nil(t, err) {
					return err
				}

				updatedData, err := r.Users().UserData("john")
				if !assert.Nil(t, err) {
					return err
				}
				assert.EqualValues(t, data, updatedData)

				return nil
			})
			assert.Nil(t, err)
		})

		t.Run("user data fields can be reset", func(t *testing.T) {
			err := storage.Transaction(func(r spi.Repositories) error {
				data, err := r.Users().UserData("john")
				if !assert.Nil(t, err) {
					return err
				}

				data.Birthday = nil
				data.Gender = models.GenderUndefined
				err = r.Users().UpdateUserData("john", data)
				if !assert.Nil(t, err) {
					return err
				}

				updatedData, err := r.Users().UserData("john")
				if !assert.Nil(t, err) {
					return err
				}
				assert.Equal(t, *data, *updatedData)

				return nil
			})
			assert.Nil(t, err)
		})
	})
}

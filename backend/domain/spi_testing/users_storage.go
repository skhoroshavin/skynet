package spi_testing

import (
	"github.com/stretchr/testify/assert"
	"skynet/domain/models"
	"skynet/domain/spi"
	"testing"
	"time"
)

func UsersStorageTestSuite(t *testing.T, storage spi.UsersStorage) {
	t.Run("insert user", func(t *testing.T) {
		t.Run("can insert new user", func(t *testing.T) {
			err := storage.Insert("john", "easy")
			assert.Nil(t, err)
		})

		t.Run("cannot insert duplicate user even with same password", func(t *testing.T) {
			err := storage.Insert("john", "easy")
			assert.Error(t, err)
		})

		t.Run("cannot insert duplicate user", func(t *testing.T) {
			err := storage.Insert("john", "peasy")
			assert.Error(t, err)
		})
	})

	t.Run("transaction", func(t *testing.T) {
		t.Run("can start transaction for existing user", func(t *testing.T) {
			txn, err := storage.Transaction("john")
			assert.Nil(t, err)
			txn.Rollback()
		})

		t.Run("cannot start transaction for non-existent user", func(t *testing.T) {
			_, err := storage.Transaction("ripper")
			assert.Error(t, err)
		})

		t.Run("password", func(t *testing.T) {
			t.Run("match password returns whether password is correct", func(t *testing.T) {
				txn, _ := storage.Transaction("john")
				defer txn.Rollback()

				assert.True(t, txn.MatchPassword("easy"))
				assert.False(t, txn.MatchPassword("peasy"))
			})

			t.Run("update password effects are visible within transaction", func(t *testing.T) {
				txn, _ := storage.Transaction("john")
				defer txn.Rollback()

				txn.UpdatePassword("peasy")

				assert.False(t, txn.MatchPassword("easy"))
				assert.True(t, txn.MatchPassword("peasy"))
			})

			t.Run("effects of transaction without commit are not visible", func(t *testing.T) {
				txn, _ := storage.Transaction("john")
				defer txn.Rollback()

				assert.True(t, txn.MatchPassword("easy"))
				assert.False(t, txn.MatchPassword("peasy"))
			})

			t.Run("updates are visible to subsequent transactions after commit", func(t *testing.T) {
				txn, _ := storage.Transaction("john")
				txn.UpdatePassword("peasy")
				assert.Nil(t, txn.Commit())

				txn, _ = storage.Transaction("john")
				defer txn.Rollback()

				assert.False(t, txn.MatchPassword("easy"))
				assert.True(t, txn.MatchPassword("peasy"))
			})
		})

		t.Run("user data", func(t *testing.T) {
			t.Run("user data for new user is empty", func(t *testing.T) {
				txn, _ := storage.Transaction("john")
				defer txn.Rollback()

				data := txn.UserData()

				assert.NotNil(t, data)
				assert.Empty(t, data.FirstName)
				assert.Empty(t, data.LastName)
				assert.Nil(t, data.Birthday)
				assert.Equal(t, models.GenderUndefined, data.Gender)
				assert.Empty(t, data.City)
				assert.Empty(t, data.Interests)
			})

			t.Run("user data can be fully updated", func(t *testing.T) {
				txn, _ := storage.Transaction("john")
				time := time.Date(1983, 11, 18, 0, 0, 0, 0, time.UTC)
				data := &models.UserData{
					FirstName: "John",
					LastName:  "Doe",
					Birthday:  &time,
					Gender:    3,
					City:      "New Vegas",
					Interests: "Dismantling cyborgs",
				}
				txn.UpdateUserData(data)
				assert.Nil(t, txn.Commit())

				txn, _ = storage.Transaction("john")
				defer txn.Rollback()
				updatedData := txn.UserData()

				assert.NotNil(t, updatedData)
				assert.Equal(t, *data, *updatedData)
			})

			t.Run("user data can be partially updated", func(t *testing.T) {
				txn, _ := storage.Transaction("john")
				data := *txn.UserData()
				data.FirstName = "Jonny"
				data.LastName = "B"
				txn.UpdateUserData(&data)
				assert.Nil(t, txn.Commit())

				txn, _ = storage.Transaction("john")
				defer txn.Rollback()
				updatedData := txn.UserData()

				assert.NotNil(t, updatedData)
				assert.Equal(t, data, *updatedData)
			})

			t.Run("user data fields can be reset", func(t *testing.T) {
				txn, _ := storage.Transaction("john")
				data := *txn.UserData()
				data.Birthday = nil
				data.Gender = models.GenderUndefined
				txn.UpdateUserData(&data)
				assert.Nil(t, txn.Commit())

				txn, _ = storage.Transaction("john")
				defer txn.Rollback()
				updatedData := txn.UserData()

				assert.NotNil(t, updatedData)
				assert.Equal(t, data, *updatedData)
			})
		})
	})
}

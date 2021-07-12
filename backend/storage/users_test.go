package storage

import (
	"github.com/stretchr/testify/assert"
	"skynet/domain/spi_testing"
	"testing"
)

func TestUsers(t *testing.T) {
	config, err := createTestDatabase()
	assert.Nil(t, err)
	defer dropTestDatabase(config)

	db, err := NewDatabase(config)
	assert.Nil(t, err)

	if !assert.Nil(t, db.CreateSchema()) {
		return
	}

	spi_testing.UsersStorageTestSuite(t, db.Users())
}

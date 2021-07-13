package mysql_storage

import (
	"github.com/stretchr/testify/assert"
	"skynet/domain/spi_testing"
	"testing"
)

func TestUsers(t *testing.T) {
	config, err := createTestDatabase()
	assert.Nil(t, err)
	defer dropTestDatabase(config)

	storage, err := NewMySqlStorage(config)
	assert.Nil(t, err)

	if !assert.Nil(t, storage.CreateSchema()) {
		return
	}

	spi_testing.UsersStorageTestSuite(t, storage)
}

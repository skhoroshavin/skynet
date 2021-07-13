package mysql_storage

import (
	"database/sql"
	"fmt"
)

func createTestDatabase() (*DBConfig, error) {
	config := EnvConfig()
	config.DBName = ""

	db, err := sql.Open("mysql", config.mysqlDsn())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %s", err)
	}

	for i := 0; i < 20; i++ {
		config.DBName = fmt.Sprintf("skynet_test_%d", i)
		_, err := db.Exec(fmt.Sprintf("create database %s", config.DBName))
		if err == nil {
			return config, nil
		}
	}

	return nil, fmt.Errorf("failed to create database in 20 attempts")
}

func dropTestDatabase(config *DBConfig) {
	db, err := sql.Open("mysql", config.mysqlDsn())
	if err != nil {
		panic(fmt.Sprintf("failed to open database: %s", err))
	}

	_, err = db.Exec(fmt.Sprintf("drop database %s", config.DBName))
	if err != nil {
		panic(fmt.Sprintf("failed to drop database: %s", err))
	}
}

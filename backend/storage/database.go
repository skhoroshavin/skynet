package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	users *Users
}

func NewDatabase(config *DBConfig) (*Database, error) {
	db, err := sql.Open("mysql", config.mysqlDsn())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %s", err)
	}

	return &Database{
		users: newUsers(db),
	}, nil
}

func (d Database) Users() *Users {
	return d.users
}

func (d Database) CreateSchema() error {
	if err := d.users.createSchema(); err != nil {
		return err
	}
	return nil
}

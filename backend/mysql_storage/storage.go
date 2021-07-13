package mysql_storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"skynet/domain/spi"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(config *DBConfig) (*Storage, error) {
	db, err := sql.Open("mysql", config.mysqlDsn())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %s", err)
	}

	return &Storage{
		db: db,
	}, nil
}

func (s Storage) Transaction(wrk func(sd spi.StorageData) error) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	sd := spi.StorageData{
		Users: newUsers(tx),
	}

	err = wrk(sd)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s Storage) CreateSchema() error {
	if err := createUsersSchema(s.db); err != nil {
		return err
	}
	return nil
}

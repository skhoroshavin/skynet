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

func NewMySqlStorage(config *DBConfig) (*Storage, error) {
	db, err := sql.Open("mysql", config.mysqlDsn())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %s", err)
	}

	return &Storage{
		db: db,
	}, nil
}

func (d Storage) Transaction(wrk func(sd spi.StorageData) error) error {
	tx, err := d.db.Begin()
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

func (d Storage) CreateSchema() error {
	if err := createUsersSchema(d.db); err != nil {
		return err
	}
	return nil
}

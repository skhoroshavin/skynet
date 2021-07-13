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

type Repositories struct {
	tx *sql.Tx
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

func (s Storage) Transaction(wrk func(r spi.Repositories) error) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	err = wrk(Repositories{tx})
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

func (r Repositories) Users() spi.UsersRepository {
	return Users{r.tx}
}
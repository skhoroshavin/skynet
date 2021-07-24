package mysql_storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rubenv/sql-migrate"
	"skynet/domain/spi"
)

var migrations = migrate.MemoryMigrationSource{
	Migrations: []*migrate.Migration{
		{
			Id: "1",
			Up: []string{
				MigrateUsersV1,
				MigrateSessionsV1,
			},
			Down: []string{
				MigrateUsersV1Down,
				MigrateSessionsV1Down,
			},
		},
	},
}

type Storage struct {
	db *sql.DB
}

type Repositories struct {
	tx *sql.Tx
}

func NewStorage(config *Config) (*Storage, error) {
	db, err := sql.Open("mysql", config.mysqlDsn())
	if err != nil {
		return nil, err
	}

	s := &Storage{db}
	if err := s.Setup(); err != nil {
		return nil, err
	}

	return s, nil
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

func (s Storage) Setup() error {
	migrate.SetTable("migrations")
	_, err := migrate.Exec(s.db, "mysql", migrations, migrate.Up)
	if err != nil {
		return err
	}
	return nil
}

func (s Storage) CleanUp() error {
	migrate.SetTable("migrations")
	_, err := migrate.Exec(s.db, "mysql", migrations, migrate.Down)
	if err != nil {
		return err
	}
	return nil
}

func (r Repositories) Users() spi.UsersRepository {
	return Users{r.tx}
}

func (r Repositories) Sessions() spi.SessionsRepository {
	return Sessions{r.tx}
}

package mysql_storage

import "database/sql"

type Sessions struct {
	tx *sql.Tx
}

func (s Sessions) CreateSession(id string) string {
	panic("implement me")
}


package mysql_storage

import (
	"database/sql"
)

type Sessions struct {
	tx *sql.Tx
}

func createSessionsSchema(db *sql.DB) error {
	_, err := db.Exec(`
        create table sessions (
            session_id varchar(64) primary key, 
            user_id    varchar(64)
        );
    `)
	return err
}

func (s Sessions) Insert(sessionId string, userId string) error {
	_, err := s.tx.Exec("insert into sessions(session_id, user_id) values (?, ?)", sessionId, userId)
	return err
}

func (s Sessions) Delete(session string) error {
	panic("implement me")
}

// TODO: Test me
func (s Sessions) UserID(session string) (string, error) {
	var userID string

	err := s.tx.QueryRow("select user_id from sessions where session_id = ?", session).Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

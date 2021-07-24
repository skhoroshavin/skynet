package mysql_storage

import (
	"database/sql"
	"skynet/domain/models"
)

const (
	MigrateUsersV1 = `
create table users (
	id         varchar(64) primary key,
	password   varchar(64),
	first_name varchar(64) character set utf8mb4,
	last_name  varchar(64) character set utf8mb4,
	birthday   date,
	gender     enum('undefined', 'male', 'female'),
	city       varchar(64) character set utf8mb4,
	interests  text character set utf8mb4
);`
	MigrateUsersV1Down = `
drop table users;`
)

type Users struct {
	tx *sql.Tx
}

func (u Users) Insert(id string, password string) error {
	_, err := u.tx.Exec("insert into users(id, password) values (?, ?)", id, password)
	return err
}

func (u Users) UpdatePassword(id string, password string) error {
	_, err := u.tx.Exec("update users set password = ? where id = ?", password, id)
	return err
}

func (u Users) UpdateUserData(id string, data *models.UserData) error {
	_, err := u.tx.Exec("update users set "+
		"first_name = ?, "+
		"last_name = ?, "+
		"birthday = ?, "+
		"gender = ?, "+
		"city = ?, "+
		"interests = ? "+
		"where id = ?",
		data.FirstName,
		data.LastName,
		data.Birthday,
		data.Gender.String(),
		data.City,
		data.Interests,
		id)
	return err
}

func (u Users) Password(id string) (string, error) {
	var password string
	err := u.tx.QueryRow("select password from users where id = ?", id).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (u Users) UserData(id string) (*models.UserData, error) {
	var firstNameN, lastNameN, genderN, cityN, interestsN sql.NullString
	var birthdayN sql.NullTime
	err := u.tx.QueryRow("select "+
		"first_name, "+
		"last_name, "+
		"birthday, "+
		"gender, "+
		"city, "+
		"interests "+
		"from users where id = ?", id).Scan(
		&firstNameN,
		&lastNameN,
		&birthdayN,
		&genderN,
		&cityN,
		&interestsN,
	)
	if err != nil {
		return nil, err
	}

	result := &models.UserData{}

	if firstNameN.Valid {
		result.FirstName = firstNameN.String
	}

	if lastNameN.Valid {
		result.LastName = lastNameN.String
	}

	if birthdayN.Valid {
		result.Birthday = &birthdayN.Time
	}

	if genderN.Valid {
		result.Gender, _ = models.GenderFromString(genderN.String)
	}

	if cityN.Valid {
		result.City = cityN.String
	}

	if interestsN.Valid {
		result.Interests = interestsN.String
	}

	return result, nil
}

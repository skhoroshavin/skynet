package mysql_storage

import (
	"database/sql"
	"skynet/domain/models"
)

type Users struct {
	tx *sql.Tx
}

func newUsers(tx *sql.Tx) Users {
	return Users{
		tx: tx,
	}
}

func createUsersSchema(db *sql.DB) error {
	_, err := db.Exec(`
        create table users (
            id         varchar(64) primary key, 
            password   varchar(64),
            first_name varchar(64) not null default "",
            last_name  varchar(64) not null default "",
            birthday   date,
            gender     enum('undefined', 'male', 'female') not null default "undefined",
            city       varchar(64) not null default "",
            interests  text
        );
    `)
	return err
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
	var firstName, lastName, genderS, city string
	var birthdayN sql.NullTime
	var interestsN sql.NullString

	err := u.tx.QueryRow("select "+
		"first_name, "+
		"last_name, "+
		"birthday, "+
		"gender, "+
		"city, "+
		"interests "+
		"from users where id = ?", id).Scan(
		&firstName,
		&lastName,
		&birthdayN,
		&genderS,
		&city,
		&interestsN,
	)
	if err != nil {
		return nil, err
	}

	birthday := &birthdayN.Time
	if !birthdayN.Valid {
		birthday = nil
	}

	var gender models.Gender
	switch genderS {
	case "male":
		gender = models.GenderMale
	case "female":
		gender = models.GenderMale
	default:
		gender = models.GenderUndefined
	}

	interests := interestsN.String
	if !interestsN.Valid {
		interests = ""
	}

	return &models.UserData{
		FirstName: firstName,
		LastName:  lastName,
		Birthday:  birthday,
		Gender:    gender,
		City:      city,
		Interests: interests,
	}, nil
}

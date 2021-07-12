package storage

import (
	"database/sql"
	"log"
	"skynet/domain/models"
	"skynet/domain/spi"
)

type Users struct {
	db *sql.DB
}

type UsersRepository struct {
	tx *sql.Tx
}

func newUsers(db *sql.DB) *Users {
	return &Users{
		db: db,
	}
}

func (u Users) createSchema() error {
	_, err := u.db.Exec(`
        create table users (
            id         varchar(30) primary key, 
            password   varchar(30),
            first_name varchar(30) not null default "",
            last_name  varchar(30) not null default "",
            birthday   date,
            gender     enum('undefined', 'male', 'female') not null default "undefined",
            city       varchar(30) not null default "",
            interests  text
        );
    `)
	return err
}

func (u Users) Begin() (spi.UsersRepository, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return nil, err
	}
	return &UsersRepository{tx}, nil
}

func (u UsersRepository) Save() error {
	return u.tx.Commit()
}

func (u UsersRepository) Cancel() {
	err := u.tx.Rollback()
	if err != nil {
		log.Printf("Failed to rollback transaction: %s", err)
	}
}

func (u UsersRepository) Insert(id string, password string) error {
	_, err := u.tx.Exec("insert into users(id, password) values (?, ?)", id, password)
	return err
}

func (u UsersRepository) UpdatePassword(id string, password string) error {
	_, err := u.tx.Exec("update users set password = ? where id = ?", password, id)
	return err
}

func (u UsersRepository) UpdateUserData(id string, data *models.UserData) error {
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

func (u UsersRepository) Password(id string) (string, error) {
	var password string
	err := u.tx.QueryRow("select password from users where id = ?", id).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (u UsersRepository) UserData(id string) (*models.UserData, error) {
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

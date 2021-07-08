package services

import "time"

type User struct {
	FirstName string
	LastName string
	BirthDate time.Time
}

type Users interface {
	GetUser(id int) (*User, error)
}

type SvcUsers struct {

}

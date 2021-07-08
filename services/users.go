package services

import (
	"errors"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Birthday  time.Time `json:"birthday"`
}

type Users interface {
	GetUser(id int) (*User, error)
}

type UsersSvc struct {
}

func (u UsersSvc) GetUser(id int) (*User, error) {
	return nil, errors.New("user not found")
}

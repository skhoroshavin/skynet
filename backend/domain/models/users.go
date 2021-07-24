package models

import (
	"fmt"
	"strings"
	"time"
)

type Gender int

const (
	GenderUndefined Gender = iota
	GenderMale
	GenderFemale
)

func (g Gender) String() string {
	switch g {
	case GenderMale:
		return "male"
	case GenderFemale:
		return "female"
	default:
		return "undefined"
	}
}

func GenderFromString(s string) (Gender, error) {
	switch strings.ToLower(s) {
	case "undefined": return GenderUndefined, nil
	case "male": return GenderMale, nil
	case "female": return GenderFemale, nil
	default: return GenderUndefined, fmt.Errorf("invalid gender %v", s)
	}
}

type UserData struct {
	FirstName string     `json:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	Birthday  *time.Time `json:"birthday,omitempty"`
	Gender    Gender     `json:"gender,omitempty"`
	City      string     `json:"city,omitempty"`
	Interests string     `json:"interests,omitempty"`
}

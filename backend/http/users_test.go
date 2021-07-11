package http

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"skynet/domain/models"
	"testing"
	"time"
)

func TestGetExistingUser(t *testing.T) {
	s := newTestServer()

	birthday := time.Date(1983, time.November, 18, 0, 0, 0, 0, time.UTC)
	exampleUser := models.UserData{
		FirstName: "John",
		LastName:  "Doe",
		Birthday:  &birthday,
		Gender:    models.GenderFemale,
		City:      "New Vegas",
		Interests: "Dismantling cyborgs",
	}
	s.users.On("UserData", "john").Return(&exampleUser, nil)

	req := httptest.NewRequest("GET", "/users/john", nil)
	res := s.serve(req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.JSONEq(t,
		`{
			"first_name": "John",
			"last_name": "Doe", 
			"birthday": "1983-11-18T00:00:00Z",
            "gender": 2,
            "city": "New Vegas",
            "interests": "Dismantling cyborgs"
		}`,
		res.Body.String())
}

func TestGetNonExistentUser(t *testing.T) {
	s := newTestServer()
	s.users.On("UserData", "jonny").
		Return(nil, errors.New("getUser jonny not found"))

	req := httptest.NewRequest("GET", "/users/jonny", nil)
	res := s.serve(req)

	assert.Equal(t, http.StatusNotFound, res.Code)
	assert.JSONEq(t, `{"error": "getUser jonny not found"}`, res.Body.String())
}

package http

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"skynet/domain/models"
	"testing"
	"time"
)

type UsersServiceMock struct {
	mock.Mock
}

func (m *UsersServiceMock) CreateUser(id string, password string) error {
	args := m.Called(id, password)
	return args.Error(0)
}

func (m *UsersServiceMock) UpdatePassword(id string, oldPassword string, newPassword string) error {
	args := m.Called(id, oldPassword, newPassword)
	return args.Error(0)
}

func (m *UsersServiceMock) UpdateUserData(id string, data *models.UserData) error {
	args := m.Called(id, data)
	return args.Error(0)
}

func (m *UsersServiceMock) UserData(id string) (*models.UserData, error) {
	args := m.Called(id)
	res, _ := args.Get(0).(*models.UserData)
	return res, args.Error(1)
}

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
		Return(nil, errors.New("user jonny not found"))

	req := httptest.NewRequest("GET", "/users/jonny", nil)
	res := s.serve(req)

	assert.Equal(t, http.StatusNotFound, res.Code)
	assert.JSONEq(t, `{"error": "user jonny not found"}`, res.Body.String())
}

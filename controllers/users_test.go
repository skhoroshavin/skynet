package controllers

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"otus_social/services"
	"testing"
	"time"
)

type UsersSvcMock struct {
	mock.Mock
}

func (m UsersSvcMock) GetUser(id int) (*services.User, error) {
	args := m.Called(id)

	user, _ := args.Get(0).(*services.User)
	return user, args.Error(1)
}

func TestOnlyNumericalUsersAreValid(t *testing.T) {
	s := CreateServer(nil)
	req := httptest.NewRequest("GET", "/users/john", nil)
	res := httptest.NewRecorder()
	s.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestGetExistingUser(t *testing.T) {
	exampleUser := services.User{
		FirstName: "John",
		LastName: "Doe",
		BirthDate: time.Date(1983, time.November, 18, 0, 0, 0, 0, time.UTC),
	}

	usersSvc := new(UsersSvcMock)
	usersSvc.On("GetUser", mock.Anything).Return(exampleUser, nil)

	s := CreateServer(usersSvc)
	req := httptest.NewRequest("GET", "/users/3", nil)
	res := httptest.NewRecorder()
	s.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var body map[string]string
	err := json.Unmarshal(res.Body.Bytes(), &body)
	assert.Nil(t, err)

}

func TestGetNonExistentUser(t *testing.T) {
	usersSvc := new(UsersSvcMock)
	usersSvc.On("GetUser", mock.Anything).Return(nil, errors.New("user not found"))

	s := CreateServer(usersSvc)
	req := httptest.NewRequest("GET", "/users/3", nil)
	res := httptest.NewRecorder()
	s.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNotFound, res.Code)
}

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"skynet/domain/models"
	"testing"
)

type AuthServiceMock struct {
	mock.Mock
}

func (m *AuthServiceMock) CreateUser(id string, password string) error {
	args := m.Called(id, password)
	return args.Error(0)
}

func (m *AuthServiceMock) UpdatePassword(id string, oldPassword string, newPassword string) error {
	args := m.Called(id, oldPassword, newPassword)
	return args.Error(0)
}

func (m *AuthServiceMock) CreateSession(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

type UsersServiceMock struct {
	mock.Mock
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

type TestServer struct {
	s     *gin.Engine
	auth  *AuthServiceMock
	users *UsersServiceMock
}

func newTestServer() *TestServer {
	auth := new(AuthServiceMock)
	users := new(UsersServiceMock)
	s := NewServer(auth, users)

	return &TestServer{
		s:     s,
		auth:  auth,
		users: users,
	}
}

func (t TestServer) serve(req *http.Request) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	t.s.ServeHTTP(res, req)
	return res
}

func TestServer_NotFound(t *testing.T) {
	s := newTestServer()

	req := httptest.NewRequest("GET", "/", nil)
	res := s.serve(req)

	assert.Equal(t, http.StatusNotFound, res.Code)
	assert.JSONEq(t, `{ "error": "not found" }`, res.Body.String())
}

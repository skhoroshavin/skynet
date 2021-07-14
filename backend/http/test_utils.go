package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"skynet/domain/models"
)

type AuthServiceMock struct {
	mock.Mock
}

func (m *AuthServiceMock) SignUp(id string, password string) (string, error) {
	args := m.Called(id, password)
	return args.String(0), args.Error(1)
}

func (m *AuthServiceMock) UserID(session string) (string, error) {
	args := m.Called(session)
	return args.String(0), args.Error(1)
}

func (m *AuthServiceMock) UpdatePassword(id string, oldPassword string, newPassword string) error {
	args := m.Called(id, oldPassword, newPassword)
	return args.Error(0)
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
	s := NewServer(EnvConfig(), auth, users)

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

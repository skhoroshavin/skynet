package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestServer struct {
	s *gin.Engine
	users *UsersServiceMock
}

func newTestServer() *TestServer {
	users := new(UsersServiceMock)
	s := NewServer(users)
	return &TestServer{
		s: s,
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

package http

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func sessionID(res *httptest.ResponseRecorder) (string, error) {
	var sessionID string
	secure, httpOnly := false, false
	for _, item := range res.Header().Values("Set-Cookie") {
		if strings.Contains(item, "sessionid") {
			for _, item := range strings.Split(item, ";") {
				item := strings.Trim(item, " ")
				switch {
				case strings.Contains(item, "sessionid"):
					sessionID = strings.Split(item, "=")[1]
				case item == "Secure":
					secure = true
				case item == "HttpOnly":
					httpOnly = true
				}
			}
		}
	}

	if len(sessionID) == 0 {
		return "", errors.New("session id not found")
	}
	if !secure {
		return "", errors.New("cookie is not secure")
	}
	if !httpOnly {
		return "", errors.New("cookie is not http only")
	}
	return sessionID, nil
}

func TestAuthSignUp(t *testing.T) {
	s := newTestServer()
	s.auth.On("SignUp", "john", "easy").Return("some_session_id", nil)

	body := `{"id": "john", "password": "easy"}`
	req := httptest.NewRequest("POST", "/auth/signup", strings.NewReader(body))
	res := s.serve(req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.JSONEq(t, `{}`, res.Body.String())

	sessionId, err := sessionID(res)
	assert.Nil(t, err)
	assert.Equal(t, "some_session_id", sessionId)

	s.auth.AssertExpectations(t)
}

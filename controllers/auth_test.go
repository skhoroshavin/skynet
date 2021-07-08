package controllers

import (
	"net/http/httptest"
	"testing"
)

func TestSignUp(t *testing.T) {
	s := CreateServer(nil)
	req := httptest.NewRequest("POST", "/signup", nil)
	res := httptest.NewRecorder()
	s.ServeHTTP(res, req)

	//assert.Equal(t, http.StatusOK, res.Code)
}

func TestSignIn(t *testing.T) {
	s := CreateServer(nil)
	req := httptest.NewRequest("POST", "/signin", nil)
	res := httptest.NewRecorder()
	s.ServeHTTP(res, req)

	//assert.Equal(t, http.StatusOk, res.Code)
}

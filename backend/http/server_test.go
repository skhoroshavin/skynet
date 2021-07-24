package http

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_NotFound(t *testing.T) {
	s := newTestServer()

	req := httptest.NewRequest("GET", "/", nil)
	res := s.serve(req)

	assert.Equal(t, http.StatusNotFound, res.Code)
	assert.JSONEq(t, `{ "message": "Not Found" }`, res.Body.String())
}

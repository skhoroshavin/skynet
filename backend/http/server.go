package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"skynet/domain/api"
)

func NewServer(auth api.Auth, users api.Users) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.Default()

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	})

	attachAuth(e, auth)
	attachUsers(e, users)

	return e
}

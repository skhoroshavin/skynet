package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"skynet/domain/api"
)

func NewServer(users api.Users) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.Default()

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	})

	usersController(users).attach(e)

	return e
}

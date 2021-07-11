package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"skynet/domain/api"
)

type Auth struct {
	auth api.Auth
}

func attachAuth(e *gin.Engine, auth api.Auth) {
	a := Auth{
		auth: auth,
	}
	e.POST("/auth/signup", a.signup)
}

type UserCredentials struct {
	ID       string   `json:"id" binding:"required"`
	Password string   `json:"password" binding:"required"`
}

func (a Auth) signup(c *gin.Context) {
	var credentials UserCredentials
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := a.auth.CreateUser(credentials.ID, credentials.Password); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	sessionID, err := a.auth.CreateSession(credentials.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": "john"})
	c.SetCookie("sessionid", sessionID, 0, "/", "", true, true)
}

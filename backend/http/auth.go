package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"skynet/domain/api"
)

type Auth struct {
	config *Config
	auth   api.Auth
}

func attachAuth(e *gin.Engine, config *Config, auth api.Auth) {
	a := Auth{config, auth}
	e.POST("/auth/signup", a.signup)
}

type UserCredentials struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a Auth) signup(c *gin.Context) {
	var credentials UserCredentials
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	sessionID, err := a.auth.SignUp(credentials.ID, credentials.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"err": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("sessionid", sessionID, 0, "/", a.config.DomainName, true, true)
	c.JSON(200, gin.H{"id": credentials.ID})
}

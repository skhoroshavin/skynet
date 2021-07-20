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

type UserCredentials struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func attachAuth(e *gin.Engine, config *Config, auth api.Auth) {
	a := Auth{config, auth}
	e.POST("/auth/signup", a.signup)
	e.GET("/auth/me", a.me)
}

func (a Auth) signup(c *gin.Context) {
	var credentials UserCredentials
	if err := c.BindJSON(&credentials); err != nil {
		Error(c, http.StatusBadRequest, err)
		return
	}

	sessionID, err := a.auth.SignUp(credentials.ID, credentials.Password)
	if err != nil {
		Error(c, http.StatusConflict, err)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("sessionid", sessionID, 0, "/", a.config.DomainName, true, true)
	c.JSON(200, gin.H{})
}

// TODO: Test me
func (a Auth) me(c *gin.Context) {
	sessionID, err := c.Cookie("sessionid")
	if err != nil {
		Error(c, http.StatusUnauthorized, err)
		return
	}

	userID, err := a.auth.UserID(sessionID)
	if err != nil {
		Error(c, http.StatusUnauthorized, err)
		return
	}

	c.JSON(200, gin.H{"id": userID})
}

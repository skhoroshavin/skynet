package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
}

func authController() Auth {
	return Auth{}
}

func (a Auth) attach(e *gin.Engine) {
	e.POST("/signup", a.signup)
	e.POST("/signin", a.signin)
}

func (a Auth) signup(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

func (a Auth) signin(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

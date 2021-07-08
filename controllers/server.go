package controllers

import (
	"github.com/gin-gonic/gin"
	"otus_social/services"
)

func CreateServer(users services.Users) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.Default()

	authController().attach(e)
	usersController(users).attach(e)

	return e
}

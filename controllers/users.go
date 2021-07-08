package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"otus_social/services"
	"strconv"
)

type Users struct {
	svcUsers services.Users
}

func usersController(u services.Users) Users {
	return Users{
		svcUsers: u,
	}
}

func (u Users) attach(e *gin.Engine) {
	e.GET("/users/:id", u.getUser)
}

func (u Users) getUser(c *gin.Context) {
	sid := c.Param("id")

	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"error": err})
		return
	}

	user, err := u.svcUsers.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H {"error": err})
		return
	}

	c.JSON(http.StatusOK, user)
}

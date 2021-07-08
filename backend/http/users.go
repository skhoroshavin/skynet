package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"skynet/domain/api"
)

type Users struct {
	svcUsers api.Users
}

func usersController(u api.Users) Users {
	return Users{
		svcUsers: u,
	}
}

func (u Users) attach(e *gin.Engine) {
	e.PUT("/users/:id", u.updateUser)
	e.GET("/users/:id", u.user)
}

func (u Users) updateUser(c *gin.Context) {
	id := c.Param("id")

	user, err := u.svcUsers.UserData(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user %s not found", id)})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u Users) user(c *gin.Context) {
	id := c.Param("id")

	user, err := u.svcUsers.UserData(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user %s not found", id)})
		return
	}

	c.JSON(http.StatusOK, user)
}

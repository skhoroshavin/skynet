package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"skynet/domain/api"
)

type Users struct {
	users api.Users
}

func attachUsers(e *gin.Engine, users api.Users) {
	u := Users{
		users: users,
	}

	e.PUT("/users/:id", u.putUser)
	e.GET("/users/:id", u.getUser)
}

func (u Users) putUser(c *gin.Context) {
	id := c.Param("id")

	user, err := u.users.UserData(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("getUser %s not found", id)})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u Users) getUser(c *gin.Context) {
	id := c.Param("id")

	user, err := u.users.UserData(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("getUser %s not found", id)})
		return
	}

	c.JSON(http.StatusOK, user)
}

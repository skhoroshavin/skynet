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

	e.GET("/users/:id", u.getUser)
}

func (u Users) getUser(c *gin.Context) {
	id := c.Param("id")

	user, err := u.users.UserData(id)
	if err != nil {
		Error(c, http.StatusNotFound, fmt.Errorf(`user "%s" not found`, id))
		return
	}

	c.JSON(http.StatusOK, user)
}

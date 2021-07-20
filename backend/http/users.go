package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"skynet/domain/api"
	"skynet/domain/models"
	"time"
)

type Users struct {
	users api.Users
}

type UserData struct {
	FirstName string     `json:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	Birthday  *time.Time `json:"birthday,omitempty"`
	Gender    string     `json:"gender,omitempty"`
	City      string     `json:"city,omitempty"`
	Interests string     `json:"interests,omitempty"`
}

func attachUsers(e *gin.Engine, users api.Users) {
	u := Users{
		users: users,
	}

	e.GET("/users/:id", u.getUser)
	e.PUT("/users/:id", u.putUser)
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

func (u Users) putUser(c *gin.Context) {
	var data UserData

	id := c.Param("id")
	if err := c.BindJSON(&data); err != nil {
		Error(c, http.StatusBadRequest, err)
		return
	}

	userData := models.UserData{
		data.FirstName,
		data.LastName,
		data.Birthday,
		models.GenderUndefined,
		data.City,
		data.Interests,
	}

	if data.Gender == "male" {
		userData.Gender = models.GenderMale
	}
	if data.Gender == "female" {
		userData.Gender = models.GenderFemale
	}

	if err := u.users.UpdateUserData(id, &userData); err != nil {
		Error(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

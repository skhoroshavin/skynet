package http

import (
	"github.com/labstack/echo/v4"
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

func attachUsers(e *echo.Echo, users api.Users) {
	u := Users{
		users: users,
	}

	e.GET("/users/:id", u.getUser)
	e.PUT("/users/:id", u.putUser)
}

func (u Users) getUser(c echo.Context) error {
	id := c.Param("id")

	user, err := u.users.UserData(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (u Users) putUser(c echo.Context) error {
	var data UserData

	id := c.Param("id")
	if err := c.Bind(&data); err != nil {
		return err
	}

	userData := models.UserData{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Birthday:  data.Birthday,
		Gender:    models.GenderUndefined,
		City:      data.City,
		Interests: data.Interests,
	}

	userData.Gender, _ = models.GenderFromString(data.Gender)

	if err := u.users.UpdateUserData(id, &userData); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"skynet/domain/api"
	"skynet/domain/models"
)

type Users struct {
	users api.Users
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
	var userData models.UserData
	id := c.Param("id")
	if err := c.Bind(&userData); err != nil {
		return err
	}

	if err := u.users.UpdateUserData(id, &userData); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

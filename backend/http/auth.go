package http

import (
	"github.com/labstack/echo/v4"
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

func attachAuth(e *echo.Echo, config *Config, auth api.Auth) {
	a := Auth{config, auth}
	e.POST("/auth/signup", a.signup)
	e.GET("/auth/me", a.me)
}

func (a Auth) signup(c echo.Context) error {
	var credentials UserCredentials
	if err := c.Bind(&credentials); err != nil {
		return err
	}

	sessionID, err := a.auth.SignUp(credentials.ID, credentials.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:       "sessionid",
		Value:      sessionID,
		Path:       "/",
		Domain:     a.config.DomainName,
		Secure:     true,
		HttpOnly:   true,
		SameSite:   http.SameSiteLaxMode,
	})
	return c.JSON(http.StatusOK, nil)
}

// TODO: Test me
func (a Auth) me(c echo.Context) error {
	sessionID, err := c.Cookie("sessionid")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	userID, err := a.auth.UserID(sessionID.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"id": userID})
}

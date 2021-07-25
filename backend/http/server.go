package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"runtime"
	"skynet/domain/api"
	"time"
)

func NewServer(config *Config, auth api.Auth, users api.Users) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	middleware.Logger()
	e.Use(middleware.RequestID())
	e.Use(Logger)
	e.Use(ForceJSON)
	e.Use(middleware.Secure())
	e.Use(Recover)

	attachAuth(e, config, auth)
	attachUsers(e, users)

	return e
}

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		start := time.Now()

		id := req.Header.Get(echo.HeaderXRequestID)
		if id == "" {
			id = c.Response().Header().Get(echo.HeaderXRequestID)
		}

		l := log.With().
			Str("id", id).
			Str("method", c.Request().Method).
			Str("uri", c.Request().RequestURI).
			Str("client_ip", c.RealIP()).
			Logger()
		c.SetRequest(req.WithContext((&l).WithContext(req.Context())))

		err := next(c)
		if err != nil {
			c.Error(err)
		}

		level := zerolog.InfoLevel
		if c.Response().Status >= 400 {
			level = zerolog.ErrorLevel
		}

		l.WithLevel(level).
			Err(err).
			Int("status", c.Response().Status).
			Int64("latency", time.Now().Sub(start).Milliseconds()).
			Msg("access")
		return nil
	}
}

func ForceJSON(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		return next(c)
	}
}

func Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				l := log.Ctx(c.Request().Context())

				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				stack := make([]byte, 4<<10)
				length := runtime.Stack(stack, true)

				l.Error().
					Err(err).
					Bytes("stack", stack[:length]).
					Msg("PANIC RECOVER")
				c.Error(err)
			}
		}()
		return next(c)
	}
}

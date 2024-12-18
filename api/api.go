package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RegisterRouter() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}

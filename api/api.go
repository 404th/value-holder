package api

import (
	"github.com/404th/value-holder/internal/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
)

func Run(cfg *config.Config, sugar *zap.SugaredLogger) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}

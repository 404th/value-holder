package main

import (
	"net/http"
	"os"

	"github.com/404th/value-holder/api"
	"github.com/404th/value-holder/internal/config"
	"github.com/404th/value-holder/internal/server"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func main() {
	// 0. Init cfg
	cfg, err := config.NewConfig()
	if err != nil {
		zap.L().Error("Error happened while initializing .env file")
	}

	// 1. Init log
	logger := new(zap.Logger)
	switch os.Getenv(cfg.ValueHolderProjectMode) {
	case config.ValueHolderProjectModeDevelopment:
		logger = zap.Must(zap.NewDevelopment())
	case config.ValueHolderProjectModeProduction:
		logger = zap.Must(zap.NewProduction())
	default:
		logger = zap.Must(zap.NewDevelopment())
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	e := api.Run(cfg, sugar)

	server.Run(cfg, sugar, e)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

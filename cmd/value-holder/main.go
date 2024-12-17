package main

import (
	"os"

	"github.com/404th/value-holder/internal/config"
	"go.uber.org/zap"
)

func main() {
	// 1. Init cfg
	cfg, err := config.NewConfig()
	if err != nil {
		zap.L().Error("Error happened while initializing .env file")
	}

	// 0. Init log
	logger := new(zap.Logger)
	switch os.Getenv(cfg.ProjectMode) {
	case config.ProjectModeDevelopment:
		logger = zap.Must(zap.NewDevelopment())
	case config.ProjectModeProduction:
		logger = zap.Must(zap.NewProduction())
	default:
		logger = zap.Must(zap.NewDevelopment())
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infow("User logged in",
		"username", "johndoe",
		"userid", 123456,
		zap.String("provider", "google"),
	)
}

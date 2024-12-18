package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/404th/value-holder/internal/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", hello, func(hf echo.HandlerFunc) echo.HandlerFunc {
		sugar.Infoln("Successfully catched from middleware")

		return hello
	})

	// 2. Run server
	shutdownChan := make(chan bool, 1)

	go func() {
		// Start server
		if err := e.Start(fmt.Sprintf("%s:%d", cfg.ValueHolderProjectHost, cfg.ValueHolderProjectPort)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			sugar.Errorln("failed to start server", "error", err)
		}

		// simulate time to close connections
		time.Sleep(1 * time.Millisecond)

		sugar.Infoln("Stopped serving new connections.")
		shutdownChan <- true
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := e.Shutdown(shutdownCtx); err != nil {
		sugar.Fatalf("HTTP shutdown error: %v", err)
	}

	<-shutdownChan
	sugar.Infoln("Graceful shutdown complete.")
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

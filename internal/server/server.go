package server

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
	"go.uber.org/zap"
)

func Run(cfg *config.Config, sugar *zap.SugaredLogger, e *echo.Echo) {
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

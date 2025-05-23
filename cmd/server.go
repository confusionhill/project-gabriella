package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"com.github/confusionhill/df/private/server/internal/config"
	"github.com/labstack/echo/v4"
)

var (
	serverRunning = false
)

func startServer(cfg *config.Config, e *echo.Echo) {
	go func() {
		serverRunning = true
		if err := e.Start(cfg.Server.Port); err != nil && err != http.ErrServerClosed {
			// logger.LoggerInstance.Fatal(fmt.Sprintf("shutting down the server: %s", err))
			e.Logger.Fatal("shutting down the server:", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")
	killServer(cfg, e)
}

func killServer(cfg *config.Config, e *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	// logger.LoggerInstance.Info("Shutting down server...")
	log.Println("Shutting down server...")
	serverRunning = false
}

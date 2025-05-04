package main

import (
	"context"
	"log"
	"net/http"
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
			e.Logger.Fatal("shutting down the server:", err)
		}
	}()
}

func killServer(cfg *config.Config, e *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("Shutting down server...")
	serverRunning = false
}

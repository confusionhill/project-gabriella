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

func startServer(cfg *config.Config, e *echo.Echo) {
	go func() {
		if err := e.Start(cfg.Server.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server:", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // listen for SIGINT (Ctrl+C)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

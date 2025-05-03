package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/domain/session"
	"com.github/confusionhill/df/private/server/internal/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.LoadConfigFromJson()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	startApp(cfg)
}

func test() {
	e := echo.New()
	cfg, err := config.LoadConfigFromJson()
	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	// Middleware (optional but good for logging, recovery, etc.)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve static assets under /cdn (e.g. .swf files)
	e.Static("/cdn", "public/assets")
	e.Static("/assets", "public/pages/assets")

	//MARK: Instiate Routers
	pagesRouter, err := router.NewPagesRouter()
	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	sessionUsecase, err := session.NewUsecase(cfg)
	if err != nil {
		e.Logger.Fatal(err)
		return
	}
	sessionHandler, err := session.NewHandler(cfg, sessionUsecase)
	if err != nil {
		e.Logger.Fatal(err)
		return
	}
	sessionRouter, err := router.NewSessionRouter(cfg, sessionHandler)
	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	//MARK: SETUP ROUTERS
	pagesRouter.Setup(e)
	sessionRouter.Setup(e)

	fmt.Println(cfg)
	// Start server in a goroutine
	go func() {
		if err := e.Start(":8081"); err != nil && err != http.ErrServerClosed {
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

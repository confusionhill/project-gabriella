package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"com.github/confusionhill/df/private/server/application"
	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func startApp(cfg *config.Config) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve static assets under /cdn (e.g. .swf files)
	e.GET("/cdn/*", func(c echo.Context) error {
		reqPath := c.Param("*")
		localPath := filepath.Join("public/assets", reqPath)

		// Serve local file if it exists and is not a directory
		if fileInfo, err := os.Stat(localPath); err == nil && !fileInfo.IsDir() {
			return c.File(localPath)
		}

		// File doesn't exist locally, proxy the request
		remoteURL := "https://play.dragonfable.com/game/" + reqPath
		resp, err := http.Get(remoteURL)
		if err != nil {
			return c.String(http.StatusBadGateway, "Failed to fetch from remote server")
		}
		defer resp.Body.Close()

		// Set headers (optional: copy all headers if needed)
		for k, v := range resp.Header {
			for _, vv := range v {
				c.Response().Header().Add(k, vv)
			}
		}

		c.Response().WriteHeader(resp.StatusCode)
		_, err = io.Copy(c.Response(), resp.Body)
		return err
	})

	e.Static("/assets", "public/pages/assets")

	rsc, err := application.LoadResources(cfg)
	if err != nil {
		e.Logger.Fatal(err)
		return
	}
	repo, err := application.LoadRepositories(cfg, rsc)
	if err != nil {
		e.Logger.Fatal(err)
		return
	}
	usecases, err := application.LoadUsecases(cfg, repo)
	if err != nil {
		e.Logger.Fatal(err)
		return
	}
	handlers, err := application.LoadHandlers(cfg, usecases)
	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	//MARK: Instiate Routers
	pagesRouter, err := router.NewPagesRouter()
	if err != nil {
		e.Logger.Fatal(err)
		return
	}
	sessionRouter, err := router.NewSessionRouter(cfg, handlers.SessionHandler)
	if err != nil {
		e.Logger.Fatal(err)
		return
	}
	authRouter, err := router.NewAuthRouter(cfg, handlers.AuthHandler)
	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	pagesRouter.Setup(e)
	sessionRouter.Setup(e)
	authRouter.Setup(e)

	startGui(cfg, e)
}

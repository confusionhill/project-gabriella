package router

import (
	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/domain/game"
	"github.com/labstack/echo/v4"
)

type GameRouter struct {
	cfg     *config.Config
	handler *game.Handler
}

func NewGameRouter(cfg *config.Config, handler *game.Handler) (*GameRouter, error) {
	return &GameRouter{
		cfg:     cfg,
		handler: handler,
	}, nil
}

func (r *GameRouter) Setup(e *echo.Echo) {
	serverGroup := e.Group("/server")
	serverGroup.POST("/cf-characterload.asp", r.handler.LoadCharacterHandler)
}

package router

import (
	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/domain/session"
	"github.com/labstack/echo/v4"
)

type SessionRouter struct {
	cfg     *config.Config
	handler *session.Handler
}

func NewSessionRouter(cfg *config.Config, handler *session.Handler) (*SessionRouter, error) {
	return &SessionRouter{
		cfg:     cfg,
		handler: handler,
	}, nil
}

func (r *SessionRouter) Setup(e *echo.Echo) {
	serverEmulatorGroup := e.Group("/server-emulator/server.php")
	serverEmulatorGroup.GET("/DFversion.asp", r.handler.GameSettingsHandler)
}

package router

import (
	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/domain/authentication"
	"github.com/labstack/echo/v4"
)

type AuthRouter struct {
	cfg     *config.Config
	handler *authentication.Handler
}

func NewAuthRouter(cfg *config.Config, handler *authentication.Handler) (*AuthRouter, error) {
	return &AuthRouter{
		cfg:     cfg,
		handler: handler,
	}, nil
}

func (r *AuthRouter) Setup(e *echo.Echo) {
	serverGroup := e.Group("/server")
	serverGroup.POST("/cf-usersignup.asp", r.handler.RegisterAccountHandler)
	serverGroup.POST("/cf-userlogin.asp", r.handler.AuthenticateAccountHandler)
}

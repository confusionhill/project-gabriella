package application

import (
	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/domain/authentication"
	"com.github/confusionhill/df/private/server/internal/domain/game"
	"com.github/confusionhill/df/private/server/internal/domain/session"
)

type Handlers struct {
	AuthHandler    *authentication.Handler
	SessionHandler *session.Handler
	GameHandler    *game.Handler
}

func LoadHandlers(cfg *config.Config, usecase *Usecases) (*Handlers, error) {
	authHandler, err := authentication.NewHandler(cfg, usecase.authUsecase)
	if err != nil {
		return nil, err
	}
	sessionHandler, err := session.NewHandler(cfg, usecase.sessionUsecase)
	if err != nil {
		return nil, err
	}
	gameHandler, err := game.NewHandler(cfg)
	if err != nil {
		return nil, err
	}
	return &Handlers{
		AuthHandler:    authHandler,
		SessionHandler: sessionHandler,
		GameHandler:    gameHandler,
	}, nil
}

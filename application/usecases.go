package application

import (
	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/domain/authentication"
	"com.github/confusionhill/df/private/server/internal/domain/session"
)

type Usecases struct {
	sessionUsecase *session.Usecase
	authUsecase    *authentication.Usecase
}

func LoadUsecases(cfg *config.Config, repo *Repositories) (*Usecases, error) {
	sessionUsecase, err := session.NewUsecase(cfg)
	if err != nil {
		return nil, err
	}
	authUsecase, err := authentication.NewUsecase(cfg, repo.auth)
	if err != nil {
		return nil, err
	}
	return &Usecases{
		sessionUsecase: sessionUsecase,
		authUsecase:    authUsecase,
	}, nil
}

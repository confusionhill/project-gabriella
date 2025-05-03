package application

import (
	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/domain/session"
)

type Usecases struct {
	sessionUsecase *session.Usecase
}

func LoadUsecases(cfg *config.Config, repo *Repositories) (*Usecases, error) {
	sessionUsecase, err := session.NewUsecase(cfg)
	if err != nil {
		return nil, err
	}
	return &Usecases{
		sessionUsecase: sessionUsecase,
	}, nil
}

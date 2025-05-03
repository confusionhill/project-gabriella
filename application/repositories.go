package application

import (
	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/domain/authentication"
)

type Repositories struct {
	auth *authentication.Repository
}

func LoadRepositories(cfg *config.Config, rsc *Resources) (*Repositories, error) {
	authRepo, err := authentication.NewRepository(cfg, rsc.db)
	if err != nil {
		return nil, err
	}
	return &Repositories{
		auth: authRepo,
	}, nil
}

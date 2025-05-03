package application

import "com.github/confusionhill/df/private/server/internal/config"

type Repositories struct {
}

func LoadRepositories(cfg *config.Config, rsc *Resources) (*Repositories, error) {
	return &Repositories{}, nil
}

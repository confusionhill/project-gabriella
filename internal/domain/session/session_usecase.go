package session

import (
	"context"

	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/data/dto/settings"
)

type Usecase struct {
	cfg *config.Config
}

func NewUsecase(cfg *config.Config) (*Usecase, error) {
	return &Usecase{
		cfg: cfg,
	}, nil
}

func (u *Usecase) GameSettingsUsecase(ctx context.Context) settings.GameSettingsDTO {
	return settings.NewGameSettings(u.cfg.Settings)
}

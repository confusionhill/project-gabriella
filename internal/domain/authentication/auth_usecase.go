package authentication

import (
	"context"

	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/data/entity/game"
	"github.com/google/uuid"
)

type Usecase struct {
	cfg        *config.Config
	repository *Repository
}

func NewUsecase(cfg *config.Config, repo *Repository) (*Usecase, error) {
	return &Usecase{
		cfg:        cfg,
		repository: repo,
	}, nil
}

func (u *Usecase) RegisterUser(ctx context.Context, user *game.User) error {
	user.SessionToken = u.generateRandomSessionToken()
	return u.repository.CreateUser(ctx, user)
}

func (u *Usecase) generateRandomSessionToken() string {
	return uuid.New().String()
}

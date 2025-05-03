package authentication

import (
	"context"
	"errors"

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

func (u *Usecase) RegisterUserUsecase(ctx context.Context, user *game.User) error {
	user.SessionToken = u.generateRandomSessionToken()
	return u.repository.CreateUser(ctx, user)
}

func (u *Usecase) generateRandomSessionToken() string {
	return uuid.New().String()
}

func (u *Usecase) AuthenticateUserUsecase(ctx context.Context, username string, password string) (*game.User, error) {
	user, err := u.repository.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("invalid password")
	}
	return user, nil
}

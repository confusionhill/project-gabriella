package authentication

import (
	"context"

	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/data/entity/game"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	cfg *config.Config
	db  *sqlx.DB
}

func NewRepository(cfg *config.Config, db *sqlx.DB) (*Repository, error) {
	return &Repository{
		cfg: cfg,
		db:  db,
	}, nil
}

func (r *Repository) CreateUser(ctx context.Context, user *game.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (username, password, email, birth_date, session_token) VALUES (:username, :password, :email, :birth_date, :session_token)", user.Username, user.Password, user.Email, user.BirthDate, user.SessionToken)
	return err
}

package authentication

import (
	"context"
	"time"

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

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*game.User, error) {
	var user game.User
	err := r.db.GetContext(ctx, &user, "SELECT id, username, password, email, birth_date, session_token FROM users WHERE username = ?", username)
	return &user, err
}

func (r *Repository) CreateCharacter(ctx context.Context, character *game.Character) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO characters (user_id, name, level, experience, strength, dexterity, intelligence, luck, charisma, endurance, wisdom, race_id, class_id, base_class_id, created_at, gender, pronoun, hair_id, color_hair, color_skin, color_base, color_trim) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		character.UserID,
		character.Name,
		character.Level,
		character.Experience,
		character.Strength,
		character.Dexterity,
		character.Intelligence,
		character.Luck,
		character.Charisma,
		character.Endurance,
		character.Wisdom,
		character.RaceID,
		character.ClassID,
		character.BaseClassID,
		time.Now(),
		character.Gender,
		character.Pronoun,
		character.HairID,
		character.ColorHair,
		character.ColorSkin,
		character.ColorBase,
		character.ColorTrim,
		character.QuestID,
	)
	return err
}

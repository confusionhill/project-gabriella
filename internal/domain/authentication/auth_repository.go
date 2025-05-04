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
	_, err := r.db.ExecContext(ctx, "INSERT INTO df_user (username, password, email, birthdate, sessionToken) VALUES (:username, :password, :email, :birth_date, :session_token)", user.Username, user.Password, user.Email, user.Birthdate, user.SessionToken)
	return err
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) ([]game.User, error) {
	var users []game.User
	query := `SELECT
	 u.id as user_id, u.password, u.email,
	  u.birth_date, u.session_token, 
	  c.id as id, c.name, c.level, c.dragon_amulet, c.base_class_id, c.class_id, c.race_id
	FROM users as u LEFT JOIN characters
	 as c ON u.id = c.user_id WHERE u.username = ?`
	err := r.db.SelectContext(ctx, &users, query, username)
	return users, err
}

func (r *Repository) CreateCharacter(ctx context.Context, character *game.Character) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO characters (user_id, name, level, experience, strength, dexterity, intelligence, luck, charisma, endurance, wisdom, race_id, class_id, base_class_id, created_at, gender, pronoun, hair_id, color_hair, color_skin, color_base, color_trim) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		character.UserId,
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
		character.RaceId,
		character.ClassId,
		character.BaseClassId,
		time.Now(),
		character.Gender,
		character.Pronoun,
		character.HairId,
		character.ColorHair,
		character.ColorSkin,
		character.ColorBase,
		character.ColorTrim,
		character.QuestId,
	)
	return err
}

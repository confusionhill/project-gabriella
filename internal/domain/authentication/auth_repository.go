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
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	tx.Commit()
	defer tx.Rollback()
	query := `SELECT
	u.id as id, u.password, u.email,
	u.birthdate, u.sessionToken, 
	c.id as charId, c.name as charName, c.level as charLevel, 
	c.dragonAmulet as charDragonAmulet, c.baseClassId as charBaseClassId, 
	c.classId as charClassId, c.raceId as charRaceId,
	cl.name as charClassName, r.name as charRaceName
FROM df_user as u 
LEFT JOIN df_char as c ON u.id = c.userId
LEFT JOIN df_class as cl ON c.classId = cl.id
LEFT JOIN df_race as r ON c.raceId = r.id
WHERE u.username = ?`
	err = r.db.SelectContext(ctx, &users, query, username)
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

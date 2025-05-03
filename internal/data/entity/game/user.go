package game

import (
	"database/sql"
	"time"
)

type User struct {
	Id               int64        `db:"id"`
	Username         string       `db:"username"`
	Password         string       `db:"password"`
	Email            string       `db:"email"`
	BirthDate        string       `db:"birth_date"`
	SessionToken     string       `db:"session_token"`
	CreatedAt        time.Time    `db:"created_at"`
	IsUpgraded       bool         `db:"is_upgraded"`
	IsSpecial        bool         `db:"is_special"`
	IsEmailActivated bool         `db:"is_email_activated"`
	IsBanned         bool         `db:"is_banned"`
	LastLogin        sql.NullTime `db:"last_login"`
	Character
}

type Character struct {
	Id        int64        `db:"id"`
	UserID    int          `db:"user_id"`
	CreatedAt time.Time    `db:"created_at"` // Consider using time.Time
	UpdatedAt sql.NullTime `db:"updated_at"` // Consider using time.Time

	Name string `db:"name"`

	Level      int `db:"level"`
	Experience int `db:"experience"`

	Silver int `db:"silver"`
	Gold   int `db:"gold"`
	Gems   int `db:"gems"`
	Coins  int `db:"coins"`

	DragonAmulet bool `db:"dragon_amulet"`
	PVPStatus    bool `db:"pvp_status"`

	Gender  string `db:"gender"`
	Pronoun string `db:"pronoun"`

	HairID    int    `db:"hair_id"`
	ColorHair string `db:"color_hair"` // hex string
	ColorSkin string `db:"color_skin"` // hex string
	ColorBase string `db:"color_base"` // hex string
	ColorTrim string `db:"color_trim"` // hex string

	QuestID int `db:"quest_id"`

	Strength     int `db:"strength"`
	Dexterity    int `db:"dexterity"`
	Intelligence int `db:"intelligence"`
	Luck         int `db:"luck"`
	Charisma     int `db:"charisma"`
	Endurance    int `db:"endurance"`
	Wisdom       int `db:"wisdom"`

	LastDailyQuestDone string `db:"last_daily_quest_done"`

	Armor  string `db:"armor"`
	Skills string `db:"skills"`
	Quests string `db:"quests"`

	RaceID      int `db:"race_id"`
	ClassID     int `db:"class_id"`
	BaseClassID int `db:"base_class_id"`

	LastTimeSeen sql.NullTime `db:"last_time_seen"` // Consider using time.Time
}

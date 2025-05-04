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
	UserID    int64        `db:"user_id"`
	CreatedAt time.Time    `db:"created_at"` // Consider using time.Time
	UpdatedAt sql.NullTime `db:"updated_at"` // Consider using time.Time

	Name string `db:"name"`

	Level      int64 `db:"level"`
	Experience int64 `db:"experience"`

	Silver int64 `db:"silver"`
	Gold   int64 `db:"gold"`
	Gems   int64 `db:"gems"`
	Coins  int64 `db:"coins"`

	DragonAmulet bool `db:"dragon_amulet"`
	PVPStatus    bool `db:"pvp_status"`

	Gender  string `db:"gender"`
	Pronoun string `db:"pronoun"`

	HairID    int64  `db:"hair_id"`
	ColorHair string `db:"color_hair"` // hex string
	ColorSkin string `db:"color_skin"` // hex string
	ColorBase string `db:"color_base"` // hex string
	ColorTrim string `db:"color_trim"` // hex string

	QuestID int `db:"quest_id"`

	Strength     int64 `db:"strength"`
	Dexterity    int64 `db:"dexterity"`
	Intelligence int64 `db:"intelligence"`
	Luck         int64 `db:"luck"`
	Charisma     int64 `db:"charisma"`
	Endurance    int64 `db:"endurance"`
	Wisdom       int64 `db:"wisdom"`

	LastDailyQuestDone string `db:"last_daily_quest_done"`

	Armor  string `db:"armor"`
	Skills string `db:"skills"`
	Quests string `db:"quests"`

	RaceID      int64 `db:"race_id"`
	ClassID     int64 `db:"class_id"`
	BaseClassID int64 `db:"base_class_id"`

	LastTimeSeen sql.NullTime `db:"last_time_seen"` // Consider using time.Time
}

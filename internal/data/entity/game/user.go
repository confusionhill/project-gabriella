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
}

package game

import (
	"database/sql"
	"time"
)

type User struct {
	ID              int64        `db:"id"`
	CreatedAt       time.Time    `db:"createdAt"`
	UpdatedAt       sql.NullTime `db:"updatedAt"`
	Username        string       `db:"username"`
	Password        string       `db:"password"`
	Email           string       `db:"email"`
	Birthdate       time.Time    `db:"birthdate"`
	SessionToken    string       `db:"sessionToken"`
	Upgraded        bool         `db:"upgraded"`
	Activated       bool         `db:"activated"`
	OptIn           bool         `db:"optIn"`
	Special         bool         `db:"special"`
	Banned          bool         `db:"banned"`
	LastLogin       sql.NullTime `db:"lastLogin"`
	RecoveryCode    string       `db:"recoveryCode"`
	RecoveryExpires sql.NullTime `db:"recoveryExpires"`
	IsDeleted       int64        `db:"_isDeleted"`
	Character
}

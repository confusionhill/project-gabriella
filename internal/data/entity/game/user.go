package game

import "time"

type User struct {
	Id           int64     `db:"id"`
	Username     string    `db:"username"`
	Password     string    `db:"password"`
	Email        string    `db:"email"`
	BirthDate    string    `db:"birth_date"`
	SessionToken string    `db:"session_token"`
	CreatedAt    time.Time `db:"created_at"`
}

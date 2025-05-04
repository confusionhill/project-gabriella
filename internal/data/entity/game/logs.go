package game

import "time"

type Logs struct {
	ID             int64     `db:"id"`
	CreatedAt      time.Time `db:"createdAt"`
	UserId         int64     `db:"userId"`
	CharId         int64     `db:"charId"`
	Service        string    `db:"service"`
	Method         string    `db:"method"`
	Action         string    `db:"action"`
	Description    string    `db:"description"`
	ReferenceClass string    `db:"referenceClass"`
	ReferenceId    int64     `db:"referenceId"`
	AdditionalData string    `db:"additionalData"`
	Severity       string    `db:"severity"`
	Ip             string    `db:"ip"`
	UserAgent      string    `db:"userAgent"`
	IsDeleted      int64     `db:"_isDeleted"`
}

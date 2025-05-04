package game

import "time"

type CharacterItem struct {
	ID        int64     `db:"id"`
	CharId    int64     `db:"charId"`
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
	ItemId    int64     `db:"itemId"`
	Equipped  bool      `db:"equipped"`
	Count     int64     `db:"count"`
	IsDeleted int64     `db:"_isDeleted"`
}

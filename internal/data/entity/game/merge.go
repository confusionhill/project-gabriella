package game

type Merge struct {
	ID        int64 `db:"id"`
	ItemId1   int64 `db:"itemId1"`
	Amount1   int64 `db:"amount1"`
	ItemId2   int64 `db:"itemId2"`
	Amount2   int64 `db:"amount2"`
	ItemId    int64 `db:"itemId"`
	String    int64 `db:"string"`
	Index     int64 `db:"index"`
	Value     int64 `db:"value"`
	Level     int64 `db:"level"`
	IsDeleted int64 `db:"_isDeleted"`
}

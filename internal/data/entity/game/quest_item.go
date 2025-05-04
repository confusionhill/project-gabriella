package game

type QuestItem struct {
	ID        int64 `db:"id"`
	QuestId   int64 `db:"questId"`
	ItemId    int64 `db:"itemId"`
	IsDeleted int64 `db:"_isDeleted"`
}

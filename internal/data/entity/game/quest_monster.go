package game

type QuestMonster struct {
	ID        int64 `db:"id"`
	QuestId   int64 `db:"questId"`
	MonsterId int64 `db:"monsterId"`
	IsDeleted int64 `db:"_isDeleted"`
}

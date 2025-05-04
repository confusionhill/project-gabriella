package game

type ItemShopItem struct {
	ID         int64 `db:"id"`
	ItemShopId int64 `db:"itemShopId"`
	ItemId     int64 `db:"itemId"`
	IsDeleted  int64 `db:"_isDeleted"`
}

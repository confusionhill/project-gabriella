package game

type HairShopHair struct {
	ID         int64 `db:"id"`
	HairShopId int64 `db:"hairShopId"`
	HairId     int64 `db:"hairId"`
	IsDeleted  int64 `db:"_isDeleted"`
}

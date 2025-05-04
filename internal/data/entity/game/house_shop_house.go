package game

type HouseShopHouse struct {
	ID          int64 `db:"id"`
	HouseShopId int64 `db:"houseShopId"`
	HouseId     int64 `db:"houseId"`
	IsDeleted   int64 `db:"_isDeleted"`
}

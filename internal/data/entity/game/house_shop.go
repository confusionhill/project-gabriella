package game

type HouseShop struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	IsDeleted int64  `db:"_isDeleted"`
}

package game

type ItemShop struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	Count     int64  `db:"count"`
	IsDeleted int64  `db:"_isDeleted"`
}

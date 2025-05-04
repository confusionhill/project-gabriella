package game

type Race struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	Resists   string `db:"resists"`
	IsDeleted int64  `db:"_isDeleted"`
}

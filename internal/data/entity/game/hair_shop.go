package game

type HairShop struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	Swf       string `db:"swf"`
	IsDeleted int64  `db:"_isDeleted"`
}

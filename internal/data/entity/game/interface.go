package game

type Interface struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	Swf       string `db:"swf"`
	LoadUnder int64  `db:"loadUnder"`
	IsDeleted int64  `db:"_isDeleted"`
}

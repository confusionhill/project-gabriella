package game

type Hair struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	Swf        string `db:"swf"`
	Frame      bool   `db:"frame"`
	Price      int64  `db:"price"`
	Gender     string `db:"gender"`
	RaceId     int64  `db:"raceId"`
	EarVisible bool   `db:"earVisible"`
	IsDeleted  int64  `db:"_isDeleted"`
}

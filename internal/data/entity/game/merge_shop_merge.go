package game

type MergeShopMerge struct {
	ID          int64 `db:"id"`
	MergeShopId int64 `db:"mergeShopId"`
	MergeId     int64 `db:"mergeId"`
	IsDeleted   int64 `db:"_isDeleted"`
}

package game

type Class struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	Element    string `db:"element"`
	Equippable string `db:"equippable"`
	Swf        string `db:"swf"`
	ArmorId    int64  `db:"armorId"`
	WeaponId   int64  `db:"weaponId"`
	Savable    int64  `db:"savable"`
	IsDeleted  int64  `db:"_isDeleted"`
}

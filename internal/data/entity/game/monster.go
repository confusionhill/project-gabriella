package game

type Monster struct {
	ID           int64  `db:"id"`
	Name         string `db:"name"`
	Level        int64  `db:"level"`
	Experience   int64  `db:"experience"`
	HitPoints    int64  `db:"hitPoints"`
	ManaPoints   int64  `db:"manaPoints"`
	Silver       int64  `db:"silver"`
	Gold         int64  `db:"gold"`
	Gems         int64  `db:"gems"`
	Coins        int64  `db:"coins"`
	Gender       string `db:"gender"`
	HairStyle    int64  `db:"hairStyle"`
	ColorHair    string `db:"colorHair"`
	ColorSkin    string `db:"colorSkin"`
	ColorBase    string `db:"colorBase"`
	ColorTrim    string `db:"colorTrim"`
	Strength     int64  `db:"strength"`
	Dexterity    int64  `db:"dexterity"`
	Intelligence int64  `db:"intelligence"`
	Luck         int64  `db:"luck"`
	Charisma     int64  `db:"charisma"`
	Endurance    int64  `db:"endurance"`
	Wisdom       int64  `db:"wisdom"`
	Element      string `db:"element"`
	RaceId       int64  `db:"raceId"`
	ArmorId      int64  `db:"armorId"`
	WeaponId     int64  `db:"weaponId"`
	MovName      string `db:"movName"`
	Swf          string `db:"swf"`
	IsDeleted    int64  `db:"_isDeleted"`
}

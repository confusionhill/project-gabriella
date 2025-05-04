package game

type Item struct {
	ID            int64  `db:"id"`
	Name          string `db:"name"`
	Description   string `db:"description"`
	DesignInfo    string `db:"designInfo"`
	Visible       int64  `db:"visible"`
	Destroyable   int64  `db:"destroyable"`
	Sellable      int64  `db:"sellable"`
	DragonAmulet  int64  `db:"dragonAmulet"`
	Currency      int64  `db:"currency"`
	Cost          int64  `db:"cost"`
	MaxStackSize  int64  `db:"maxStackSize"`
	Bonus         int64  `db:"bonus"`
	Rarity        int64  `db:"rarity"`
	Level         int64  `db:"level"`
	Type          string `db:"type"`
	Element       string `db:"element"`
	CategoryId    int64  `db:"categoryId"`
	EquipSpot     string `db:"equipSpot"`
	ItemType      string `db:"itemType"`
	Swf           string `db:"swf"`
	Icon          string `db:"icon"`
	Strength      int64  `db:"strength"`
	Dexterity     int64  `db:"dexterity"`
	Intelligence  int64  `db:"intelligence"`
	Luck          int64  `db:"luck"`
	Charisma      int64  `db:"charisma"`
	Endurance     int64  `db:"endurance"`
	Wisdom        int64  `db:"wisdom"`
	DamageMin     int64  `db:"damageMin"`
	DamageMax     int64  `db:"damageMax"`
	DefenseMelee  int64  `db:"defenseMelee"`
	DefensePierce int64  `db:"defensePierce"`
	DefenseMagic  int64  `db:"defenseMagic"`
	Critical      int64  `db:"critical"`
	Parry         int64  `db:"parry"`
	Dodge         int64  `db:"dodge"`
	Block         int64  `db:"block"`
	Resists       string `db:"resists"`
	IsDeleted     int64  `db:"_isDeleted"`
}

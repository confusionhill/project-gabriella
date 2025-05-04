package game

type House struct {
	ID           int64  `db:"id"`
	Name         string `db:"name"`
	Description  string `db:"description"`
	Visible      int64  `db:"visible"`
	Destroyable  bool   `db:"destroyable"`
	Equippable   bool   `db:"equippable"`
	RandomDrop   bool   `db:"randomDrop"`
	Sellable     bool   `db:"sellable"`
	DragonAmulet bool   `db:"dragonAmulet"`
	Enc          bool   `db:"enc"`
	Cost         int64  `db:"cost"`
	Currency     int64  `db:"currency"`
	Rarity       bool   `db:"rarity"`
	Level        int64  `db:"level"`
	Category     int64  `db:"category"`
	EquipSpot    int64  `db:"equipSpot"`
	Type         string `db:"type"`
	Random       int64  `db:"random"`
	Element      int64  `db:"element"`
	Icon         string `db:"icon"`
	DesignInfo   string `db:"designInfo"`
	Swf          string `db:"swf"`
	Region       int64  `db:"region"`
	Theme        int64  `db:"theme"`
	Size         int64  `db:"size"`
	BaseHP       int64  `db:"baseHP"`
	StorageSize  int64  `db:"storageSize"`
	MaxGuards    int64  `db:"maxGuards"`
	MaxRooms     int64  `db:"maxRooms"`
	MaxExtItems  int64  `db:"maxExtItems"`
	IsDeleted    int64  `db:"_isDeleted"`
}

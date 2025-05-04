package game

type Quest struct {
	ID              int64  `db:"id"`
	Name            string `db:"name"`
	Description     string `db:"description"`
	Complete        string `db:"complete"`
	Swf             string `db:"swf"`
	SwfX            string `db:"swfX"`
	MaxSilver       int64  `db:"maxSilver"`
	MaxGold         int64  `db:"maxGold"`
	MaxGems         int64  `db:"maxGems"`
	MaxExp          int64  `db:"maxExp"`
	Mint64ime       int64  `db:"mint64ime"`
	Counter         int64  `db:"counter"`
	Extra           string `db:"extra"`
	DailyIndex      int64  `db:"dailyIndex"`
	DailyReward     int64  `db:"dailyReward"`
	MonsterMinLevel int64  `db:"monsterMinLevel"`
	MonsterMaxLevel int64  `db:"monsterMaxLevel"`
	MonsterType     string `db:"monsterType"`
	MonsterGroupSwf string `db:"monsterGroupSwf"`
	IsDeleted       int64  `db:"_isDeleted"`
}

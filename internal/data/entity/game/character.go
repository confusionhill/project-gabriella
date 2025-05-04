package game

import (
	"database/sql"
	"time"
)

type Character struct {
	ID                 int64        `db:"id"`
	UserId             int64        `db:"userId"`
	CreatedAt          time.Time    `db:"createdAt"`
	UpdatedAt          sql.NullTime `db:"updatedAt"`
	Name               string       `db:"name"`
	Level              int64        `db:"level"`
	Experience         int64        `db:"experience"`
	Silver             int64        `db:"silver"`
	Gold               int64        `db:"gold"`
	Gems               int64        `db:"gems"`
	Coins              int64        `db:"coins"`
	DragonAmulet       int64        `db:"dragonAmulet"`
	PvpStatus          int64        `db:"pvpStatus"`
	Gender             string       `db:"gender"`
	Pronoun            string       `db:"pronoun"`
	HairId             int64        `db:"hairId"`
	ColorHair          string       `db:"colorHair"`
	ColorSkin          string       `db:"colorSkin"`
	ColorBase          string       `db:"colorBase"`
	ColorTrim          string       `db:"colorTrim"`
	Strength           int64        `db:"strength"`
	Dexterity          int64        `db:"dexterity"`
	Intelligence       int64        `db:"intelligence"`
	Luck               int64        `db:"luck"`
	Charisma           int64        `db:"charisma"`
	Endurance          int64        `db:"endurance"`
	Wisdom             int64        `db:"wisdom"`
	LastDailyQuestDone sql.NullTime `db:"lastDailyQuestDone"`
	Armor              string       `db:"armor"`
	Skills             string       `db:"skills"`
	Quests             string       `db:"quests"`
	RaceId             int64        `db:"raceId"`
	ClassId            int64        `db:"classId"`
	BaseClassId        int64        `db:"baseClassId"`
	QuestId            int64        `db:"questId"`
	LastTimeSeen       sql.NullTime `db:"lastTimeSeen"`
	IsDeleted          int64        `db:"_isDeleted"`
}

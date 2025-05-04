package game

import (
	"database/sql"
	"time"
)

type Settings struct {
	ID                           int64        `db:"id"`
	CreatedAt                    time.Time    `db:"createdAt"`
	UpdatedAt                    sql.NullTime `db:"updatedAt"`
	ServerName                   string       `db:"serverName"`
	GameSwf                      string       `db:"gameSwf"`
	ServerVersion                string       `db:"serverVersion"`
	ServerLocation               string       `db:"serverLocation"`
	GamefilesPath                string       `db:"gamefilesPath"`
	HomeUrl                      string       `db:"homeUrl"`
	PlayUrl                      string       `db:"playUrl"`
	SignUpUrl                    string       `db:"signUpUrl"`
	LostPasswordUrl              string       `db:"lostPasswordUrl"`
	TosUrl                       string       `db:"tosUrl"`
	CharDetailUrl                string       `db:"charDetailUrl"`
	SignUpMessage                string       `db:"signUpMessage"`
	News                         string       `db:"news"`
	EnableAdvertising            bool         `db:"enableAdvertising"`
	DailyQuestCoinsReward        int64        `db:"dailyQuestCoinsReward"`
	DragonAmuletForAll           bool         `db:"dragonAmuletForAll"`
	RevalidateClientValues       bool         `db:"revalidateClientValues"`
	BanInvalidClientValues       bool         `db:"banInvalidClientValues"`
	CanDeleteUpgradedChar        bool         `db:"canDeleteUpgradedChar"`
	NonUpgradedChars             int64        `db:"nonUpgradedChars"`
	UpgradedChars                int64        `db:"upgradedChars"`
	NonUpgradedMaxBagSlots       int64        `db:"nonUpgradedMaxBagSlots"`
	UpgradedMaxBagSlots          int64        `db:"upgradedMaxBagSlots"`
	NonUpgradedMaxBankSlots      int64        `db:"nonUpgradedMaxBankSlots"`
	UpgradedMaxBankSlots         int64        `db:"upgradedMaxBankSlots"`
	NonUpgradedMaxHouseSlots     int64        `db:"nonUpgradedMaxHouseSlots"`
	UpgradedMaxHouseSlots        int64        `db:"upgradedMaxHouseSlots"`
	NonUpgradedMaxHouseItemSlots int64        `db:"nonUpgradedMaxHouseItemSlots"`
	UpgradedMaxHouseItemSlots    int64        `db:"upgradedMaxHouseItemSlots"`
	ExperienceMultiplier         float64      `db:"experienceMultiplier"`
	GemsMultiplier               float64      `db:"gemsMultiplier"`
	GoldMultiplier               float64      `db:"goldMultiplier"`
	SilverMultiplier             float64      `db:"silverMultiplier"`
	OnlineThreshold              int64        `db:"onlineThreshold"`
	Detailed404ClientError       bool         `db:"detailed404ClientError"`
	SendEmails                   bool         `db:"sendEmails"`
	EmailApiUrl                  string       `db:"emailApiUrl"`
	EmailApiToken                string       `db:"emailApiToken"`
	EmailAddress                 string       `db:"emailAddress"`
	IsDeleted                    int64        `db:"_isDeleted"`
}

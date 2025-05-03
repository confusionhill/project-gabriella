package game

import (
	"encoding/xml"
	"time"

	utils "com.github/confusionhill/df/private/server/internal/Utils"
	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/data/entity/game"
)

type AuthUserDTO struct {
	XMLName xml.Name `xml:"characters"`
	User    UserDTO  `xml:"user"`
}

type CharacterDTO struct {
	XMLName      xml.Name `xml:"characters"`
	CharID       int64    `xml:"CharID,attr"`
	Name         string   `xml:"strCharacterName,attr"`
	Level        int64    `xml:"intLevel,attr"`
	AccessLevel  int64    `xml:"intAccessLevel,attr"`
	DragonAmulet int64    `xml:"intDragonAmulet,attr"`
	BaseClassID  int64    `xml:"orgClassID,attr"`
	ClassName    string   `xml:"strClassName,attr"`
	RaceName     string   `xml:"strRaceName,attr"`
}

type UserDTO struct {
	XMLName        xml.Name       `xml:"user"`
	UserID         int64          `xml:"UserID,attr"`
	CharsAllowed   int64          `xml:"intCharsAllowed,attr"`
	AccessLevel    int64          `xml:"intAccessLevel,attr"`
	Upgrade        int64          `xml:"intUpgrade,attr"`
	ActivationFlag int64          `xml:"intActivationFlag,attr"`
	OptIn          int64          `xml:"bitOptin,attr"`
	Token          string         `xml:"strToken,attr"`
	News           string         `xml:"strNews,attr"`
	AdFlag         int64          `xml:"bitAdFlag,attr"`
	DateToday      string         `xml:"dateToday,attr"`
	CustomUsername string         `xml:"customParam_username,attr"`
	Characters     []CharacterDTO `xml:"characters"`
}

func NewAuthUserDTO(user *game.User) *AuthUserDTO {
	userdto := UserDTO{
		UserID:         user.Id,
		CharsAllowed:   6,
		AccessLevel:    1,
		Upgrade:        6,
		ActivationFlag: 0,
		OptIn:          1,
		Token:          user.SessionToken,
		News:           "Test News, i have no idea what to put here, just enjoy the game!",
		AdFlag:         1,
		DateToday:      time.Now().Format("2006-01-02"),
		CustomUsername: user.Username,
		Characters:     []CharacterDTO{},
	}

	if user.IsEmailActivated {
		userdto.ActivationFlag = 1
	}

	return &AuthUserDTO{
		User: userdto,
	}
}

func (u *AuthUserDTO) ToString(cfg *config.Config) (string, error) {
	xmlBytes, err := xml.Marshal(u)
	if err != nil {
		return "", err
	}
	toEncrypt := string(xmlBytes)
	return utils.EncryptNinja(&cfg.Server.DragonFable, toEncrypt), nil
}

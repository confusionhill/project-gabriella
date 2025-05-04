package game

import (
	"encoding/xml"
	"fmt"
	"time"

	utils "com.github/confusionhill/df/private/server/internal/Utils"
	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/data/entity/game"
)

type CreateCharacterRequestDTO struct {
	UserID        int64  `form:"intUserID"`
	Username      string `form:"strUsername"`
	Password      string `form:"strPassword"`
	Token         string `form:"strToken"`
	CharacterName string `form:"strCharacterName"`
	Gender        string `form:"strGender"`
	Pronoun       string `form:"strPronoun"`
	HairID        int64  `form:"intHairID"`
	ColorHair     int64  `form:"intColorHair"`
	ColorSkin     int64  `form:"intColorSkin"`
	ColorBase     int64  `form:"intColorBase"`
	ColorTrim     int64  `form:"intColorTrim"`
	ClassID       int64  `form:"intClassID"`
	RaceID        int64  `form:"intRaceID"`
	ClassName     string `form:"strClass"`
	HairFrame     int64  `form:"intHairFrame"`
}

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

func NewAuthUserDTO(users []game.User) *AuthUserDTO {
	user := users[0]
	characters := []CharacterDTO{}
	for _, character := range users {
		fmt.Println(character)
		char := CharacterDTO{
			CharID:      character.ID,
			Name:        character.Name,
			Level:       character.Level,
			AccessLevel: 1,
			BaseClassID: character.BaseClassId,
			ClassName:   "warior",
			RaceName:    "human",
		}
		char.DragonAmulet = character.DragonAmulet
		characters = append(characters, char)
	}
	userdto := UserDTO{
		UserID:         user.UserId,
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
		Characters:     characters,
	}

	if user.Activated {
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

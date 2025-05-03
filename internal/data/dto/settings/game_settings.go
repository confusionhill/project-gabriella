package settings

import (
	"net/url"

	"com.github/confusionhill/df/private/server/internal/config"
)

type GameSettingsDTO struct {
	GameMovie     string
	SignUpMessage string
	Server        string
	GamefilesPath string
	GameVersion   string
	Online        bool
	End           string
}

func NewGameSettings(entity config.GameSettings) GameSettingsDTO {
	return GameSettingsDTO{
		GameMovie:     entity.GameMovie,
		SignUpMessage: entity.SignUpMessage,
		Server:        entity.Server,
		GamefilesPath: entity.GamefilesPath,
		GameVersion:   entity.GameVersion,
		Online:        entity.Online,
		End:           entity.End,
	}
}

func (settings *GameSettingsDTO) ToStringValues() string {
	values := url.Values{}
	values.Set("gamemovie", settings.GameMovie)
	values.Set("signUpMessage", settings.SignUpMessage)
	values.Set("server", settings.Server)
	values.Set("gamefilesPath", settings.GamefilesPath)
	values.Set("gameVersion", settings.GameVersion)
	values.Set("online", map[bool]string{true: "true", false: "false"}[settings.Online])
	values.Set("end", settings.End)
	return "&" + values.Encode()
}

package config

type Config struct {
	Server   ServerConfig `json:"server`
	Settings GameSettings `json:"settings"`
}

type ServerConfig struct {
	Database    DatabaseConfig    `json:"database"`
	DragonFable DragonFableConfig `json:"dragonFable"`
	Name        string            `json:"name"`
	Port        string            `json:"port"`
}

type DatabaseConfig struct {
	Type string `json:"type"`
	Host string `json:"host"`
}

type DragonFableConfig struct {
	DfSettingsId int64  `json:"fSettingsId`
	Ninja2Key    string `json:"ninja2Key`
}

type GameSettings struct {
	GameMovie     string `json:"gameMovie"`
	SignUpMessage string `json:"signUpMessage"`
	Server        string `json:"server"`
	GamefilesPath string `json:"gameFilesPath"`
	GameVersion   string `json:"gameVersion"`
	Online        bool   `json:"online"`
	End           string `json:"end"`
}

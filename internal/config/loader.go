package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadConfigFromJson() (*Config, error) {
	file, err := os.Open("private/df.secrets.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config JSON: %w", err)
	}

	return &cfg, nil
}

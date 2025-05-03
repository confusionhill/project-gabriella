package main

import (
	"fmt"

	"com.github/confusionhill/df/private/server/internal/config"
)

func main() {
	cfg, err := config.LoadConfigFromJson()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	startApp(cfg)
}

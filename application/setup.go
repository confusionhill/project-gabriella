package application

import (
	"fmt"
	"log"
	"os"

	"com.github/confusionhill/df/private/server/internal/config"
)

func Setup(cfg *config.Config, rsc *Resources) {
	// db. execute sql query from folder private/db_template.sql
	go func() {
		sqlBytes, err := os.ReadFile("private/db_template.sql")
		if err != nil {
			log.Fatalf("failed to read SQL file: %v", err)
		}

		sqlQuery := string(sqlBytes)

		// Execute the SQL query
		_, err = rsc.db.Exec(sqlQuery)
		if err != nil {
			log.Fatalf("failed to execute SQL query: %v", err)
		}

		fmt.Println("SQL query executed successfully")
	}()
}

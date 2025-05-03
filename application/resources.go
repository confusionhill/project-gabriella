package application

import (
	"com.github/confusionhill/df/private/server/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Resources struct {
	db *sqlx.DB
}

func LoadResources(cfg *config.Config) (*Resources, error) {
	db, err := sqlx.Connect(cfg.Server.Database.Type, cfg.Server.Database.Host)
	if err != nil {
		return nil, err
	}
	schema := `
	CREATE TABLE users IF NOT EXISTS (
    id INT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    birth_date DATE NOT NULL,
    session_token VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);`
	db.MustExec(schema)
	return &Resources{
		db: db,
	}, nil
}

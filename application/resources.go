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
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		author TEXT
	);`
	db.MustExec(schema)
	return &Resources{
		db: db,
	}, nil
}

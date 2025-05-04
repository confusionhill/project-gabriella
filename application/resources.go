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
	CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    birth_date DATE NOT NULL,
    session_token VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`
	db.MustExec(schema)
	charSchema := `CREATE TABLE IF NOT EXISTS characters (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,

    name TEXT NOT NULL,

    level INTEGER NOT NULL default 1,
    experience INTEGER NOT NULL default 0,

    silver INTEGER NOT NULL default 0,
    gold INTEGER NOT NULL default 0,
    gems INTEGER NOT NULL default 0,
    coins INTEGER NOT NULL default 0,

    dragon_amulet BOOLEAN NOT NULL default true,
    pvp_status BOOLEAN NOT NULL default true,

    gender TEXT NOT NULL,
    pronoun TEXT NOT NULL,

    hair_id INTEGER NOT NULL,
    color_hair TEXT NOT NULL,
    color_skin TEXT NOT NULL,
    color_base TEXT NOT NULL,
    color_trim TEXT NOT NULL,

    quest_id INTEGER,

    strength INTEGER NOT NULL default 10,
    dexterity INTEGER NOT NULL default 10,
    intelligence INTEGER NOT NULL default 10,
    luck INTEGER NOT NULL default 10,
    charisma INTEGER NOT NULL default 10,
    endurance INTEGER NOT NULL default 10,
    wisdom INTEGER NOT NULL default 10,

    last_daily_quest_done TEXT,

    armor TEXT,
    skills TEXT,
    quests TEXT,

    race_id INTEGER NOT NULL default 1,
    class_id INTEGER NOT NULL default 1,
    base_class_id INTEGER NOT NULL default 1,

    last_time_seen TIMESTAMP default CURRENT_TIMESTAMP
);`
	db.MustExec(charSchema)

	return &Resources{
		db: db,
	}, nil
}

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    birth_date DATE NOT NULL,
    session_token VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS characters (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,

    name TEXT NOT NULL,

    level INTEGER NOT NULL,
    experience INTEGER NOT NULL,

    silver INTEGER NOT NULL,
    gold INTEGER NOT NULL,
    gems INTEGER NOT NULL,
    coins INTEGER NOT NULL,

    dragon_amulet BOOLEAN NOT NULL,
    pvp_status BOOLEAN NOT NULL,

    gender TEXT,
    pronoun TEXT,

    hair_id INTEGER,
    color_hair TEXT,
    color_skin TEXT,
    color_base TEXT,
    color_trim TEXT,

    quest_id INTEGER,

    strength INTEGER,
    dexterity INTEGER,
    intelligence INTEGER,
    luck INTEGER,
    charisma INTEGER,
    endurance INTEGER,
    wisdom INTEGER,

    last_daily_quest_done TEXT,

    armor TEXT,
    skills TEXT,
    quests TEXT,

    race_id INTEGER,
    class_id INTEGER,
    base_class_id INTEGER,

    last_time_seen TIMESTAMP
);

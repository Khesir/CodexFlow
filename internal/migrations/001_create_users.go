package migrations

import "github.com/jmoiron/sqlx"

func MigrationCreateUsers(db sqlx.Ext) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username TEXT NOT NULL UNIQUE,
            email TEXT NOT NULL UNIQUE,
            created_at TIMESTAMP DEFAULT NOW()
        );
    `)
	return err
}

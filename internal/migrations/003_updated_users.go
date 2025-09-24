package migrations

import "github.com/jmoiron/sqlx"

func MigrationUpdateUsers(db sqlx.Ext) error {
	_, err := db.Exec(`
		ALTER TABLE users
		ADD COLUMN IF NOT EXISTS password TEXT NOT NULL DEFAULT '',
		ADD COLUMN IF NOT EXISTS is_admin BOOLEAN DEFAULT FALSE,
		ADD COLUMN IF NOT EXISTS is_google_user BOOLEAN DEFAULT FALSE,
		ADD COLUMN IF NOT EXISTS last_login_at TIMESTAMP,
		ADD COLUMN IF NOT EXISTS refresh_token TEXT,
		ADD COLUMN IF NOT EXISTS modified_at TIMESTAMP DEFAULT NOW(),
		ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;
	`)
	return err
}

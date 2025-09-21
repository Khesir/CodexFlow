package migrations

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var Migrations = []struct {
	Version string
	Up      func(e sqlx.Ext) error
}{
	{"001_create_users", MigrationCreateUsers},
	{"002_create_posts", MigrationCreatePosts},
}

func RunMigrations(db *sqlx.DB) {
	// ensure migrations table exists
	db.MustExec(`CREATE TABLE IF NOT EXISTS schema_migrations (
		version TEXT PRIMARY KEY
	)`)

	applied := 0
	for _, m := range Migrations {
		var exists bool
		err := db.Get(&exists,
			`SELECT EXISTS (
            SELECT 1 FROM schema_migrations WHERE version = $1
        )`, m.Version)
		if err != nil {
			panic("failed to check migrations: " + err.Error())
		}

		if exists {
			continue
		}

		tx, err := db.Beginx()
		if err != nil {
			panic("failed to start transaction: " + err.Error())
		}

		if err := m.Up(tx); err != nil {
			tx.Rollback()
			panic("migration " + m.Version + " failed: " + err.Error())
		}

		_, err = tx.Exec(`INSERT INTO schema_migrations (version) VALUES ($1)`, m.Version)
		if err != nil {
			tx.Rollback()
			panic("failed to record migration " + m.Version + ": " + err.Error())
		}

		if err := tx.Commit(); err != nil {
			panic("failed to commit migration " + m.Version + ": " + err.Error())
		}

		applied++
	}

	if applied == 0 {
		fmt.Println("No new migrations to apply.")
	} else {
		fmt.Printf("Applied %d new migration(s).\n", applied)
	}
}

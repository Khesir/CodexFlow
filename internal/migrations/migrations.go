package migrations

import "github.com/jmoiron/sqlx"

var Migrations = map[string]func(db *sqlx.DB) error{
	"001_create_users": MigrationCreateUsers,
	"002_create_posts": MigrationCreatePosts,
}

func RunMigrations(db *sqlx.DB) {
	for version, m := range Migrations {
		if err := m(db); err != nil {
			panic("migration " + version + " failed: " + err.Error())
		}
	}
}

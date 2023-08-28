package bootstrap

import databasemigrations "todo/internal/databaseMigrations"

func Seed() {
	Migrate()
	databasemigrations.Seed()
}

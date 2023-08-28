package bootstrap

import (
	"log"
	databasemigrations "todo/internal/databaseMigrations"

	"github.com/joho/godotenv"
)

func Migrate() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	databasemigrations.Migrate()

}

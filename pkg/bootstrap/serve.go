package bootstrap

import (
	"log"
	"todo/pkg/routing"

	"github.com/joho/godotenv"
)

func Serve() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Migrate()
	routing.RouteBuilder()
}

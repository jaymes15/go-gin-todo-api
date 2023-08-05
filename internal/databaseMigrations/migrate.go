package databasemigrations

import (
	"fmt"
	"log"
	"todo/internal/apps/tasks"
	"todo/internal/apps/users"
	"todo/pkg/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&tasks.TaskModel{}, &users.UserModel{})

	if err != nil {
		log.Fatalf(":::::Error Reading Configs::::: %s", err)
	}

	fmt.Println("Migration done")
}

package databasemigrations

import (
	"fmt"
	"log"
	"strconv"
	"todo/internal/apps/tasks"
	"todo/internal/apps/users"
	"todo/pkg/database"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()

	password := "test123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		log.Fatalf("unable to decode into struct %s", err)
	}
	user := users.UserModel{UserName: "Jinzhu", Password: string(hashedPassword)}

	db.Create(&user)

	user2 := users.UserModel{UserName: "Jinzhu2", Password: string(hashedPassword)}
	db.Create(&user2)

	for i := 1; i <= 10; i++ {
		title := fmt.Sprintf("new task title %s", strconv.Itoa(i))
		content := fmt.Sprintf("new task content %s", strconv.Itoa(i))
		task := tasks.TaskModel{Title: title, Content: content, UserId: user.ID}
		db.Create(&task)

		log.Printf("Task created successfully with title %s", task.Title)
	}
	log.Println("Db successfully seeded")

}

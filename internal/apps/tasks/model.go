package tasks

import (
	"todo/internal/apps/users"

	"gorm.io/gorm"
)

type TaskModel struct {
	gorm.Model
	Title     string `gorm:"varchar:300;not null"`
	Content   string `gorm:"text"`
	Completed bool   `gorm:"bool;default:false"`
	UserId    uint
	User      users.UserModel
}

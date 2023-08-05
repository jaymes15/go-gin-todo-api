package users

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	UserName string `gorm:"varchar:191;unique;not null"`
	Password string `gorm:"varchar:191;not null"`
}

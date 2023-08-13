package users

import (
	"errors"
	"todo/pkg/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRespository struct {
	DB *gorm.DB
}

func NewUserRespository() *UserRespository {
	return &UserRespository{
		DB: database.Connection(),
	}

}

func (userRespository *UserRespository) Create(user UserModel) (UserModel, error) {
	var newUser UserModel

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	if err != nil {
		return newUser, errors.New("error hashing password")
	}
	user.Password = string(hashedPassword)

	userRespository.DB.Create(&user).Scan(&newUser)

	return newUser, nil
}

func (userRepository *UserRespository) FindByID(id uint) UserModel {
	var user UserModel

	userRepository.DB.First(&user, "id = ?", id)

	return user
}

func (userRepository *UserRespository) FindByUserName(username string) UserModel {
	var user UserModel

	userRepository.DB.First(&user, "user_name = ?", username)

	return user
}

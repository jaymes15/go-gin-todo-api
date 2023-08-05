package users

import (
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRespository UserRespositoryInterface
}

func NewUserService() *UserService {
	return &UserService{
		userRespository: NewUserRespository(),
	}
}

func (userService *UserService) CreateUser(c *gin.Context, request Register) (UserRes, error) {
	var response UserRes

	newUser := UserModel{
		UserName: request.UserName,
		Password: request.Password,
	}
	user, err := userService.userRespository.Create(newUser)

	if err != nil {
		return response, err
	}

	if user.ID == 0 {
		return response, errors.New("error on creating user")
	}

	return ToUser(user), nil
}

func (userService *UserService) LoginUser(request Login) (UserRes, error) {
	var response UserRes

	user := userService.userRespository.FindByUserName(request.UserName)

	if user.ID == 0 {
		return response, errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return response, errors.New("invalid credentials")
	}

	return ToUser(user), nil
}

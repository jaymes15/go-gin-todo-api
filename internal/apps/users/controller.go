package users

import (
	"log"
	"net/http"
	"strconv"
	"todo/pkg/sessions"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService UserServiceInterface
}

func NewAuthController() *AuthController {
	return &AuthController{
		userService: NewUserService(),
	}
}

func (authController *AuthController) Register(c *gin.Context) {
	var request Register

	if err := c.ShouldBind(&request); err != nil {
		log.Printf("Could not Bind:::::: %s", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := authController.userService.CreateUser(c, request)

	if err != nil {
		log.Printf("Error:::::: %s", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (authController *AuthController) Login(c *gin.Context) {
	var request Login

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Could not Bind:::::: %s", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := authController.userService.LoginUser(request)

	if err != nil {
		log.Printf("Error:::::: %s", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (authController *AuthController) HandleLogout(c *gin.Context) {
	sessions.Remove(c, "auth")

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

package users

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	userController := NewAuthController()

	router.POST("/users", userController.Register)
	router.POST("/users/token", userController.Login)

}

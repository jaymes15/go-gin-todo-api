package users

import "github.com/gin-gonic/gin"

type UserServiceInterface interface {
	CreateUser(c *gin.Context, request Register) (UserRes, error)
	LoginUser(request Login) (UserRes, error)
}

package middlewares

import (
	"log"
	"net/http"
	"strconv"
	UserRepository "todo/internal/apps/users"
	"todo/pkg/auth/custom_jwt"
	"todo/pkg/sessions"

	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {
	var userRepo = UserRepository.NewUserRespository()

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.ParseUint(authID, 10, 64)

		user := userRepo.FindByID(uint(userID))

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		// before request

		c.Next()
	}
}

func ValidateAuth() gin.HandlerFunc {
	var userRepo = UserRepository.NewUserRespository()

	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")

		if err != nil {
			tokenString = c.GetHeader("Authorization")
			if tokenString == "" {
				log.Println("Error in ValidateAuth function: ", err.Error())
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}
		}
		userId, err := custom_jwt.ValidateJWT(tokenString)

		if err != nil {
			log.Println("Error in ValidateAuth function: ", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Auth token expired"})
			c.Abort()
			return
		}

		userID, _ := strconv.ParseUint(userId, 10, 64)

		user := userRepo.FindByID(uint(userID))

		if user.ID == 0 {
			log.Println("Error in ValidateAuth function: User not found")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}

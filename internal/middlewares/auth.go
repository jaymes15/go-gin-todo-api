package middlewares

import (
	"net/http"
	"strconv"
	UserRepository "todo/internal/apps/users"
	"todo/pkg/sessions"

	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {
	var userRepo = UserRepository.NewUserRespository()

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// before request

		c.Next()
	}
}

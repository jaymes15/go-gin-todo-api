package middlewares

import (
	"net/http"
	"strconv"
	"todo/pkg/sessions"

	"github.com/gin-gonic/gin"
)

func IsGuest() gin.HandlerFunc {

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		if userID != 0 {
			c.JSON(http.StatusFound, gin.H{"message": "OK"})
			return
		}
		// before request

		c.Next()
	}
}

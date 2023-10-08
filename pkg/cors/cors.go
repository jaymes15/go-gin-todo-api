package cors

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	return config
}

func corsMiddleware(config cors.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(config)(c)
		if c.Request.Method == http.MethodOptions {
			c.JSON(http.StatusOK, struct{}{})
			return
		}
		c.Next()
	}
}

func UseCors(route *gin.Engine) {
	config := corsConfig()
	route.Use(corsMiddleware(config))
}

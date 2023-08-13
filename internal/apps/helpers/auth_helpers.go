package helpers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"todo/pkg/auth/custom_jwt"

	"github.com/gin-gonic/gin"
)

func parseBoolEnv(key string) bool {
	val := os.Getenv(key)
	log.Println("parseBoolEnv:::::: ", key, val)
	return strings.ToLower(val) == "true"
}

func SetAuthtoken(c *gin.Context, userId uint) (string, error) {
	HTTP_ONLY := parseBoolEnv("HTTP_ONLY")
	SECURE := parseBoolEnv("SECURE")
	DOMAIN := os.Getenv("DOMAIN")
	jwtToken, err := custom_jwt.CreateJWT(userId)

	if err != nil {
		log.Printf("Failed to Created token Error:::::: %s", err.Error())
		return "", errors.New(err.Error())

	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", jwtToken, 3600*24*30, "", DOMAIN, SECURE, HTTP_ONLY)
	return jwtToken, nil
}

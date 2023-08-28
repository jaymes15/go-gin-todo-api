package custom_jwt

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func CreateJWT(userId uint) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}

	SECRET_KEY := os.Getenv("SECRET_KEY")

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(userId)),
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(SECRET_KEY))

	fmt.Println("CreateJWT:::::", tokenString, err)
	if err != nil {
		log.Println("Failed to Created token Error:::::: ", err.Error())
		return "", errors.New(err.Error())

	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}

	SECRET_KEY := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SECRET_KEY), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["sub"])
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return "", errors.New("Expired auth token")

		} else {
			return claims["sub"].(string), nil
		}

	} else {
		fmt.Println(err)
		return "", errors.New(err.Error())
	}

}

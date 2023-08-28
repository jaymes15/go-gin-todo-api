package routing

import (
	"fmt"
	"log"
	"todo/pkg/config"

	"github.com/gin-gonic/gin"
)

func Serve(router *gin.Engine) {
	config.Set()

	configs := config.Get()

	serverConfig := fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)
	err := router.Run(serverConfig)

	if err != nil {
		log.Fatalf("Error in routing: %s", err)
	}
}

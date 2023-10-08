package chat

import (
	"todo/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	chatController := NewChatController()
	authGroup := router.Group("")
	authGroup.Use(middlewares.ValidateAuth())
	{
		authGroup.GET("/ws/chat", chatController.Chat)

	}

}

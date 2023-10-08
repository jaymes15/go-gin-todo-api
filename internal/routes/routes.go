package routes

import (
	"todo/internal/apps/chat"
	"todo/internal/apps/tasks"
	"todo/internal/apps/users"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	tasks.Routes(router)
	users.Routes(router)
	chat.Routes(router)

}

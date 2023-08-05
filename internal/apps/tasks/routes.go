package tasks

import (
	"todo/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	taskController := NewTaskController()

	authGroup := router.Group("/tasks")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/", taskController.AllTask)
		authGroup.POST("/", taskController.Create)
		authGroup.PATCH("/:id", taskController.Update)
	}
}

package tasks

import "github.com/gin-gonic/gin"

type TaskServiceInterface interface {
	GetAllTask(c *gin.Context) TasksRes
	CreateTask(c *gin.Context, request CreateTask) (TaskRes, error)
	UpdateTask(taskId uint, request UpdateTask, c *gin.Context) (TaskRes, error)
}

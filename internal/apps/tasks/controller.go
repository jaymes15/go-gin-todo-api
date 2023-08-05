package tasks

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type TaskController struct {
	taskService TaskServiceInterface
}

func NewTaskController() *TaskController {
	return &TaskController{
		taskService: NewTaskService(),
	}
}

func (taskControllers *TaskController) AllTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"allTask":  taskControllers.taskService.GetAllTask(c),
		"app name": viper.Get("App.Name"),
	})
}

func (taskControllers *TaskController) Create(c *gin.Context) {

	var request CreateTask

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Could not Bind:::::: %s", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := taskControllers.taskService.CreateTask(c, request)

	if err != nil {
		log.Printf("Error:::::: %s", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"task":     task,
		"app name": viper.Get("App.Name"),
	})

}

func (taskControllers *TaskController) Update(c *gin.Context) {

	var request UpdateTask

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Could not Bind:::::: %s", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	task, err := taskControllers.taskService.UpdateTask(uint(id), request, c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"title": "Entity not found", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task":     task,
		"app name": viper.Get("App.Name"),
	})
}

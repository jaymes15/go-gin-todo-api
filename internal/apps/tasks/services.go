package tasks

import (
	"errors"
	"strconv"
	"todo/internal/apps/users"
	"todo/pkg/sessions"

	"github.com/gin-gonic/gin"
)

type TaskService struct {
	taskRespository TaskRespositoryInterface
}

func NewTaskService() *TaskService {
	return &TaskService{
		taskRespository: NewTaskRespository(),
	}
}

func (taskService *TaskService) GetAllTask(c *gin.Context) TasksRes {
	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)
	user := users.NewUserRespository().FindByID(userID)

	if user.ID == 0 {

		return TasksRes{
			Data: []TaskRes{},
		}
	}
	allTask := taskService.taskRespository.List(user.ID)
	return ToTasks(allTask)
}

func (taskService *TaskService) CreateTask(c *gin.Context, request CreateTask) (TaskRes, error) {
	var response TaskRes

	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)
	user := users.NewUserRespository().FindByID(userID)

	if user.ID == 0 {
		return response, errors.New("task could not be created, user not found")
	}

	newTask := TaskModel{
		Title:     request.Title,
		Content:   request.Content,
		Completed: request.Completed,
		UserId:    user.ID,
	}

	task := taskService.taskRespository.Create(newTask)

	if task.ID == 0 {
		return response, errors.New("task could not be created")
	}

	return ToTask(task), nil
}

func (taskService *TaskService) UpdateTask(taskId uint, request UpdateTask, c *gin.Context) (TaskRes, error) {
	var response TaskRes

	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)
	user := users.NewUserRespository().FindByID(userID)

	if user.ID == 0 {
		return response, errors.New("user not found")
	}
	retrievedTask := taskService.taskRespository.GetTaskById(taskId)

	if userID != int(retrievedTask.UserId) {
		return response, errors.New("UnAuthorized")
	}

	task := TaskModel{
		Title:     request.Title,
		Content:   request.Content,
		Completed: request.Completed,
	}

	task, err := taskService.taskRespository.Update(taskId, task)

	if err != nil {
		return response, err
	}

	return ToTask(task), nil
}

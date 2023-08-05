package tasks

import (
	"fmt"
	"log"
	"todo/pkg/database"

	"gorm.io/gorm"
)

type TaskRespository struct {
	DB *gorm.DB
}

func NewTaskRespository() *TaskRespository {
	return &TaskRespository{
		DB: database.Connection(),
	}

}

func (taskRespository *TaskRespository) List(userID uint) []TaskModel {
	var allTask []TaskModel

	fmt.Println("userID:::::", userID)

	taskRespository.DB.Where("user_id = ?", userID).Order("created_at").Find(&allTask)

	return allTask
}

func (taskRespository *TaskRespository) Create(task TaskModel) TaskModel {
	var newTask TaskModel

	taskRespository.DB.Create(&task).Scan(&newTask)

	return newTask
}

func (taskRespository *TaskRespository) Update(taskId uint, task TaskModel) (TaskModel, error) {

	updates := make(map[string]interface{})

	updates["Completed"] = task.Completed

	if task.Content != "" {
		updates["Content"] = task.Content
	}
	if task.Title != "" {
		updates["Title"] = task.Title
	}

	result := taskRespository.DB.Model(&TaskModel{}).Where("id = ?", taskId).Updates(updates)
	if result.Error != nil {
		log.Println(result.Error)
		return TaskModel{}, result.Error
	}

	var updatedTask TaskModel
	if err := taskRespository.DB.First(&updatedTask, taskId).Error; err != nil {
		log.Println(err)
		return TaskModel{}, result.Error
	}

	return updatedTask, nil
}

func (taskRespository *TaskRespository) GetTaskById(taskID uint) TaskModel {
	var task TaskModel

	taskRespository.DB.First(&task, "id = ?", taskID)

	return task
}

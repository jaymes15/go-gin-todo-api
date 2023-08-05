package tasks

import (
	"fmt"
)

type TaskRes struct {
	ID        uint
	Title     string
	Content   string
	CreatedAt string
	Completed bool
}

type TasksRes struct {
	Data []TaskRes
}

func ToTask(task TaskModel) TaskRes {
	return TaskRes{
		ID:        task.ID,
		Title:     task.Title,
		Content:   task.Content,
		Completed: task.Completed,
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", task.CreatedAt.Year(), task.CreatedAt.Month(), task.CreatedAt.Day()),
	}
}

func ToTasks(tasks []TaskModel) TasksRes {
	var response TasksRes

	for _, task := range tasks {
		response.Data = append(response.Data, ToTask(task))
	}

	return response
}

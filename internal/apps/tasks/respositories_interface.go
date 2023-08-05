package tasks

type TaskRespositoryInterface interface {
	List(userID uint) []TaskModel
	Create(task TaskModel) TaskModel
	Update(taskId uint, task TaskModel) (TaskModel, error)
	GetTaskById(taskID uint) TaskModel
}

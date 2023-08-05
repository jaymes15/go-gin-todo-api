package tasks

// Binding from JSON
type CreateTask struct {
	Title     string `json:"title" binding:"required,min=2"`
	Content   string `json:"content" binding:"required"`
	Completed bool   `json:"completed"`
}

type UpdateTask struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

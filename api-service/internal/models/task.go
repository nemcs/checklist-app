package models

// структуры данных (Task, AddTaskRequest)
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Desctiption string `json:"desctiption"`
	Done        bool   `json:"done"`
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"desctiption"`
}

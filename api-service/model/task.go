package model

//├── model/                  # структуры данных
//│   └── task.go             # Task, CreateTaskRequest, и т.д.

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

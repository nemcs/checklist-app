package models

import "time"

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

type NewTaskToDB struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type NewTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

type TasksListResponse struct {
	Data []Task `json:"data"`
}
type TasksSingleResponse struct {
	Data Task `json:"data"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type ID struct {
	ID string `json:"id"`
}

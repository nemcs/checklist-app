package models

import "time"

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

type ID struct {
	ID string `json:"id"`
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
}

type MarkDoneRequest struct {
}

type DeleteTaskRequest struct {
}

type DataResponse struct {
	Data any `json:"data,omitempty"`
}
type MessageResponse struct {
	Message string `json:"message"`
}

type APIResponse struct {
	Message *MessageResponse
	Data    *DataResponse
}

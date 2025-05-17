package service

import (
	"github.com/google/uuid"
	"github.com/nemcs/checklist-app/api-service/model"
	"github.com/nemcs/checklist-app/api-service/storage"
)

//├── service/                # бизнес-логика
//│   └── task.go             # логика добавления, удаления, и т.д.

func AddTask(req model.CreateTaskRequest) model.Task {
	id := uuid.New().String()
	task := model.Task{
		ID:          id,
		Title:       req.Title,
		Desctiption: req.Description,
		Done:        false,
	}
	storage.Tasks[id] = task
	return task
}

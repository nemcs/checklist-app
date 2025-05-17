package services

// бизнес-логика (addTask, getTasks)
import (
	"github.com/google/uuid"
	"github.com/nemcs/checklist-app/api-service/internal/models"
)

var tasks = make(map[string]models.Task)

func AddTask(req models.CreateTaskRequest) models.Task {
	id := uuid.New().String()
	task := models.Task{
		ID:          id,
		Title:       req.Title,
		Desctiption: req.Description,
		Done:        false,
	}
	tasks[id] = task
	return task
}

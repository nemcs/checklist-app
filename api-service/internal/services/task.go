package services

// бизнес-логика (addTask, getTasks)
import (
	"github.com/google/uuid"
	"github.com/nemcs/checklist-app/api-service/internal/models"
)

var storage = make(map[string]models.Task)

func AddTask(req models.CreateTaskRequest) models.Task {
	id := uuid.New().String()
	task := models.Task{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Done:        false,
	}
	storage[id] = task
	return task
}

func GetAllTasks() []models.Task {
	var tasks []models.Task
	for _, v := range storage {
		tasks = append(tasks, v)
	}
	return tasks
}

package services

// бизнес-логика (addTask, getTasks)
import (
	"errors"
	"github.com/google/uuid"
	"github.com/nemcs/checklist-app/api-service/internal/models"
)

var ErrTaskNotFound = errors.New("Задача не найдена")

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

func DeleteTask(id string) error {
	_, found := getTask(id)
	if !found {
		return ErrTaskNotFound
	}
	delete(storage, id)
	return nil
}

func ChangeStatusDone(id string) error {
	_, found := getTask(id)
	if !found {
		return ErrTaskNotFound
	}
	task := storage[id]
	task.Done = true
	storage[id] = task
	return nil
}

func getTask(id string) (models.Task, bool) {
	task, found := storage[id]
	return task, found

}

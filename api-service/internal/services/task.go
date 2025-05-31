package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/nemcs/checklist-app/api-service/internal/errors"
	"github.com/nemcs/checklist-app/api-service/internal/models"
	"github.com/nemcs/checklist-app/api-service/internal/utils"
	"net/http"
	"net/url"
)

func (c *DBClient) GetAllTasks() (*models.TasksListResponse, error) {
	buildURL, err := url.JoinPath(c.baseURL, "/tasks")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrURLBuildFailed, err)
	}

	r, err := c.Client.Get(buildURL)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrDBRequestFailed, err)
	}

	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %d", errors.ErrDBResponseBadStatus, r.StatusCode)
	}

	var tasks models.TasksListResponse
	if err = json.NewDecoder(r.Body).Decode(&tasks); err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrJSONDecodeFailed, err)
	}

	return &tasks, nil
}

func (c *DBClient) GetTask(id string) (*models.TasksSingleResponse, error) {
	body, err := json.Marshal(models.ID{ID: id})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrJSONEncodeFailed, err)
	}

	buildURL, err := url.JoinPath(c.baseURL, "/task")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrURLBuildFailed, err)
	}
	r, err := c.Client.Post(buildURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrDBRequestFailed, err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %d", errors.ErrDBResponseBadStatus, r.StatusCode)
	}

	var tasks *models.TasksSingleResponse
	if err = json.NewDecoder(r.Body).Decode(&tasks); err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrJSONDecodeFailed, err)
	}

	return tasks, nil
}

func (c *DBClient) ChangeStatusDone(id string) error {
	body, err := json.Marshal(models.ID{ID: id})
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrJSONEncodeFailed, err)
	}

	buildURL, err := url.JoinPath(c.baseURL, "/task/done")
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrURLBuildFailed, err)
	}

	return utils.SendJSONRequest(http.MethodPatch, buildURL, body, c.Client)
}

func (c *DBClient) DeleteTask(id string) error {
	body, err := json.Marshal(models.ID{ID: id})
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrJSONEncodeFailed, err)
	}

	buildURL, err := url.JoinPath(c.baseURL, "/task/delete")
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrURLBuildFailed, err)
	}

	return utils.SendJSONRequest(http.MethodDelete, buildURL, body, c.Client)
}

func (c *DBClient) AddTask(req *http.Request) error {
	//в body получили title / description
	//парсим в структуру
	var task models.NewTaskRequest
	if err := utils.DecodeJSON(req, &task); err != nil {
		return fmt.Errorf("%w: %v", errors.ErrJSONDecodeFailed, err)
	}
	//формируем структуру с айдишником и добавляем title / description

	newTask := models.NewTaskToDB{
		ID:          uuid.New().String(),
		Title:       task.Title,
		Description: task.Description,
	}
	body, err := json.Marshal(newTask)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrJSONEncodeFailed, err)
	}
	// Отправляем запрос в бд
	// самое главное здесь сформировать айдишник или в хендлере db-service лучше делать

	buildURL, err := url.JoinPath(c.baseURL, "/task/create")
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrURLBuildFailed, err)
	}
	resp, err := c.Client.Post(buildURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrDBRequestFailed, err)
	}
	defer req.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %d", errors.ErrDBResponseBadStatus, resp.StatusCode)
	}

	return nil
}

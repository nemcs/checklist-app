package utils

// хелперы (парсинг, генерация ID, etc)
import (
	"encoding/json"
	"github.com/nemcs/checklist-app/api-service/internal/models"
	"net/http"
)

func SendJSON(w http.ResponseWriter, status int, task models.Task) {
	resp := map[string]string{"id": task.ID, "message": "created"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

func DecodeJsonBodyCreateTask(r *http.Request) (*models.CreateTaskRequest, error) {
	var req models.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

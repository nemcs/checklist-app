package util

import (
	"encoding/json"
	"github.com/nemcs/checklist-app/api-service/model"
	"net/http"
)

//├── util/                   # вспомогательные функции, если появятся
//│   └── response.go         # функции для возврата JSON, ошибок и т.д.

func SendJSON(w http.ResponseWriter, status int, task model.Task) {
	resp := map[string]string{"id": task.ID, "message": "created"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

func DecodeJsonBodyCreateTask(r *http.Request) (*model.CreateTaskRequest, error) {
	var req model.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

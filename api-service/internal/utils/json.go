package utils

import (
	"encoding/json"
	"github.com/nemcs/checklist-app/api-service/internal/models"
	"log"
	"net/http"
)

// TODO норм ли юзать any?
func SendJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("SendJSON error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func DecodeJsonBodyCreateTask(r *http.Request) (*models.CreateTaskRequest, error) {
	var req models.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

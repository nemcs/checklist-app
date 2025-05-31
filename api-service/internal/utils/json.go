package utils

// хелперы (парсинг, генерация ID, etc)
import (
	"encoding/json"
	"fmt"
	"github.com/nemcs/checklist-app/api-service/internal/models"
	"log"
	"net/http"
)

func SendJSON(w http.ResponseWriter, status int, data any) {
	SetHeaderJSON(w)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("SendJSON error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func DecodeJSON(r *http.Request, dst any) error {
	if err := json.NewDecoder(r.Body).Decode(&dst); err != nil {
		return fmt.Errorf("[DecodeJSON]: %w", err)
	}
	return nil

}
func EncodeJSON(w http.ResponseWriter, src any) error {
	SetHeaderJSON(w)
	if err := json.NewEncoder(w).Encode(src); err != nil {
		return fmt.Errorf("[EncodeJSON]: %w", err)
	}
	return nil
}

func DecodeJsonBodyCreateTask(r *http.Request) (*models.NewTaskRequest, error) {
	var req models.NewTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

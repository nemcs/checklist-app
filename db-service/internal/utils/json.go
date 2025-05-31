package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("SendJSON error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func SendSuccessJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("SendJSON error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
func SendErrorJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("SendJSON error: %v", err)
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
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(src); err != nil {
		return fmt.Errorf("[EncodeJSON]: %w", err)
	}
	return nil
}

//func DecodeJsonBodyCreateTask(r *http.Request) (*models.CreateTaskRequest, error) {
//	var req models.CreateTaskRequest
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		return nil, err
//	}
//	return &req, nil
//}

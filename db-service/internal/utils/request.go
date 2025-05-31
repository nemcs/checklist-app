package utils

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/nemcs/checklist-app/db-service/internal/models"
	"net/http"
)

func DecodeAndValidateIDFromBody(r *http.Request) (string, error) {
	var id models.ID
	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		return "", fmt.Errorf("неверный json")
	}
	if err := uuid.Validate(id.ID); err != nil {
		return "", fmt.Errorf("неверный id, необходим UUID")
	}
	return id.ID, nil
}

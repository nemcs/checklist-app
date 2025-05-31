package utils

import (
	"bytes"
	"fmt"
	"github.com/nemcs/checklist-app/api-service/internal/errors"
	"net/http"
)

func SetHeaderJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func SendJSONRequest(method string, url string, body []byte, client *http.Client) error {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrDBRequestFailed, err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrDBRequestFailed, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %d", errors.ErrDBResponseBadStatus, resp.StatusCode)
	}
	return nil
}

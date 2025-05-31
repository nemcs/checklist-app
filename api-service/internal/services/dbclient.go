package services

import "net/http"

type DBClient struct {
	baseURL string
	Client  *http.Client
}

func NewDBClient(url string) *DBClient {
	return &DBClient{
		baseURL: url,
		Client:  &http.Client{},
	}
}

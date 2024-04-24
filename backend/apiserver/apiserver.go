package apiserver

import (
	"net/http"
)

// API server is the HTTP API server
type APIServer interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// NewAPIServer creates a new API server instance
func NewAPIServer() (APIServer, error) {
	return nil, nil // TODO
}

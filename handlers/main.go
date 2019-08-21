package handlers

import (
	"net/http"
)

// Health Check Endpoint
func Health(w http.ResponseWriter, r *http.Request) {
	message := "Service is healthy"
	w.Write([]byte(message))
}

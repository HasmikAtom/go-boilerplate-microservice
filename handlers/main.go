package handlers

import (
	"net/http"
	"strings"
)

// Health Endpoint
func Health(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

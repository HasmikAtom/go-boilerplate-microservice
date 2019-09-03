package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/hasmikatom/go-boilerplate-service/handlers"
)

// SetRoutes
func SetRoutes(router *mux.Router) {

	router.HandleFunc("/", redirectToSwaggerUI)
	router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", swaggerFileServer))

	router.HandleFunc("/health", handlers.Health).Methods("GET")
}

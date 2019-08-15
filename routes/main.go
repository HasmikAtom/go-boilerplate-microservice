package routes

import (
	"github.com/gorilla/mux"

	"github.com/hasmikatom/go-boilerplate-service/handlers"
)

// Routes
func SetRoutes(router *mux.Router) {

	router.HandleFunc("/health", handlers.Health).Methods("GET")

}

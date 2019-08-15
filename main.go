package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hasmikatom/go-boilerplate-service/routes"

	"github.com/gorilla/mux"
)

func main() {
	//** Setting the connection port **//
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := mux.NewRouter()

	routes.SetRoutes(router)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

# Go-Swagger boilerplate microservice

This a boilerplate Go microservice with Swagger implementation created for Mutable but can also be used independently.
It is meant to provide the basic functionality for a Swagger-ready Go server.


# Dependencies

## Swagger
This service is using Swagger 2.0 package for API documentation

Swagger Ui assets are stored in the ``/swaggerui`` directory and are served on ``/swaggerui/`` endpoint. ``/swaggerui/swagger.json`` contains the swagger configuration for the current boilerplate. 

Root endpoint ``/`` is set to redirect to ``/swaggerui/`` endpoint for convenience.

## Gorilla/Mux
This service is using Mux router package. 
```
package main

import (
	"github.com/hasmikatom/go-boilerplate-service/routes"
	"github.com/gorilla/mux"
)
func main () {
	router := mux.NewRouter()

	routes.SetRoutes(router)
}
```
The Mux package is imported and the router is initialized in our main function.

# Infrastructure
**Routes** and **Handlers** logic is separated into their own packages which makes adding new endpoints and handlers easier.

Initialized Mux router is passed to ``router.SetRoutes(router *mux.Router)`` function, where it is used to set and manage endpoints and their handler functions which are imported from **Handler** package.

```
package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/hasmikatom/go-boilerplate-service/handlers"
)

func SetRoutes (router *mux.Router) {
	router.HandleFunc("/", redirectToSwaggerUI)
	router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", swaggerFileServer))

	router.HandleFunc("/health", handlers.Health).Methods("GET")
}
```

Mutable app checks the service's health through the ``/health`` endpoint.

## References

- [GoDoc]([https://godoc.org/](https://godoc.org/))
- [Swagger 2.0](https://swagger.io/docs/specification/2-0/basic-structure/) 
- [Swagger UI](https://swagger.io/tools/swagger-ui/)
- [gorilla/mux](https://github.com/gorilla/mux)


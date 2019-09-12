# Go-Swagger boilerplate microservice

  

This a boilerplate Go microservice with Swagger implementation created for Mutable but can also be used independently.

It is meant to provide the basic functionality for a Swagger-ready Go server.

  
  

# Dependencies

  
## Mutable Meta

This package is a wrapper to the background service `Mutable Meta`. Mutable Meta is responsible for service discovery and service configuration.

### Service Discovery
Returns all the currently known services.



### Config Manager
Fetches the JSON configuration for the current service.
```
//** Structs for Mutable's config.json **//
type  config  struct {
	Prod envConfig
	Dev envConfig
}

//** Example Environment Configs //
type  envConfig  struct {
	// Postgres postgresConfig
}
```
```
import meta "github.com/HasmikAtom/meta-go"

func  getConfig() config {
	var  conf config
	if  err  := meta.UnmarshalConfig(&conf); err !=  nil {
		panic(err)
	}
	return conf
}

// Conf stores the database configuration
var  Conf  =  getConfig()
```

## Swagger

Swagger 2.0 package is used for API documentation

  

Swagger Ui assets are stored in the ``/swaggerui`` directory and are served on ``/swaggerui/`` endpoint. ``/swaggerui/swagger.json`` contains the swagger configuration for the current boilerplate.

  

Root endpoint ``/`` is set to redirect to ``/swaggerui/`` endpoint for convenience.

  

## Gorilla/Mux

Mux router package is used for routing.

```

package main

import (
	"github.com/gorilla/mux"
)

func main () {
	router := mux.NewRouter()
	setRoutes(router)
}

```

The Mux package is imported and the router is initialized in our main function.

  

# Infrastructure

**Routes** are set in the **main package** while **Handlers** and **Database** logic is separated into its own package.

  

Initialized Mux router is passed to ``setRoutes(router *mux.Router)`` function, where it is used to set and manage endpoints and their handler functions which are imported from **Handler** package.
```

package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/hasmikatom/go-boilerplate-service/handlers"
)

func setRoutes (router *mux.Router) {
	router.HandleFunc("/", redirectToSwaggerUI)
	router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", swaggerFileServer))
	router.HandleFunc("/health", handlers.Health).Methods("GET")
}

```

**Database** package is set to handle all the external connections through configs.json, and **Mutable meta** .

  

Mutable app checks the service's health through the ``/health`` endpoint.

  

## References

  

-  [GoDoc](https://godoc.org/)
-  [Mutable Meta](https://github.com/HasmikAtom/meta-go)
-  [Swagger 2.0](https://swagger.io/docs/specification/2-0/basic-structure/)
-  [Swagger UI](https://swagger.io/tools/swagger-ui/)
-  [gorilla/mux](https://github.com/gorilla/mux)
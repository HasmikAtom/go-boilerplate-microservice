package main

import "net/http"

//** Swagger_UI Utils **//
var swaggerFileServer = http.FileServer(http.Dir("./swaggerui/"))

func redirectToSwaggerUI(writer http.ResponseWriter, req *http.Request) {
	http.Redirect(writer, req, "/swaggerui/", http.StatusMovedPermanently)
}

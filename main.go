package main

import "os"

func main() {
	//** Setting the connection port **//
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

}

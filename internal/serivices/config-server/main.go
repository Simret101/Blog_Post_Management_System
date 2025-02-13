package main

import (
	"config-server/handlers"
	"log"
	"net/http"
)

func main() {
	// Define the HTTP routes
	http.HandleFunc("/config", handlers.GetConfig)

	// Start the server
	port := ":8085"
	log.Printf("Starting config server on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

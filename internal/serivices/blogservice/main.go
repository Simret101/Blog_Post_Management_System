package main

import (
	"log"
	"net/http"
)

func blogServiceHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("BlogService is healthy!"))
}

func main() {
	http.HandleFunc("/health", blogServiceHealth)
	log.Println("BlogService running on :8081")
	http.ListenAndServe(":8081", nil)
}

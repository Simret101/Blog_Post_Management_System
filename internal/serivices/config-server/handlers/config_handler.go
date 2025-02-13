package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var configFile = "./config/config.json"

// GetConfig reads the configuration file and returns it as JSON
func GetConfig(w http.ResponseWriter, r *http.Request) {
	// Ensure only GET requests are allowed
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Open the config file
	file, err := os.Open(configFile)
	if err != nil {
		log.Printf("Error opening config file: %v", err)
		http.Error(w, "Could not read config", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Read the file contents
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error reading config file: %v", err)
		http.Error(w, "Could not read config", http.StatusInternalServerError)
		return
	}

	// Set headers and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

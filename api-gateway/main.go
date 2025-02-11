package main

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Health check for the API Gateway
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "API Gateway is running"})
	})

	// Display all services in a single endpoint
	r.GET("/services", func(c *gin.Context) {
		services := gin.H{
			"services": []gin.H{
				{"name": "AuthService", "url": "http://localhost:8083", "route": "/auth/*path"},
				{"name": "BlogService", "url": "http://localhost:8081", "route": "/blog/*path"},
				{"name": "UserService", "url": "http://localhost:8082", "route": "/user/*path"},
			},
		}
		c.JSON(http.StatusOK, services)
	})

	// Run the API Gateway server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}


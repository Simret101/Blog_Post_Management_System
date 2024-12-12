package main

import (
	"practice/blog_service/route"

	"github.com/gin-gonic/gin"
)

func main() {
	// Enable debug mode for detailed logs
	gin.SetMode(gin.DebugMode)

	// Initialize Gin router
	router := gin.Default()

	// Set up routes
	route.SetUp(router)

	// Start the server on port 8080
	router.Run(":8080")
}

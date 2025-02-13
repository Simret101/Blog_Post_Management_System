package main

import (
	"log"
	_ "practice/docs"   // Import the docs package to initialize Swagger docs
	"practice/rabbitmq" // Import the rabbitmq package
	routes "practice/route"
	"user/controller"
	"user/repository"
	"user/usecase"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/gin-gonic/gin"
)

//	@title			Blog Post API
//	@version		1.0
//	@description	This is a simple blog post API
//	@termsOfService	http://swagger.io/terms/

func main() {

	// Step 1: Connect to RabbitMQ using the new package
	rabbitConn, rabbitCh, err := rabbitmq.NewRabbitMQConnection()
	if err != nil {
		log.Fatalf("Error setting up RabbitMQ: %s", err)
	}
	defer rabbitmq.CloseConnection(rabbitConn, rabbitCh) // Ensure the connection and channel are closed when done
	var x repository.UserRepository
	// Step 2: Create repositories and use cases for RabbitMQ integration
	userRepo := repository.NewUploadRepository(x) // Example: Pass a valid UserRepository
	uploadUsecase := usecase.NewUploadUsecase(*userRepo)

	// Step 3: Create the controller for image upload handling
	uploadController := controller.NewUploadController(*uploadUsecase)

	// Step 4: Set up the Gin router
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // Set up the Swagger docs route

	// Step 5: Register routes for image upload
	router.POST("/upload/:id", uploadController.UploadImg())

	// Step 6: Set up the rest of the routes
	routes.SetUp(router)

	// Step 7: Run the server
	router.Run("127.0.0.1:8080")

}

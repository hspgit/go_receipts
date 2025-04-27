package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures the gin router with all endpoints
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Add middleware
	router.Use(ErrorMiddleware())

	api := router.Group("/api")
	{
		api.GET("/users", GetAllUsers)
		api.POST("/users", CreateUser)

		api.GET("/items", GetAllItems)
		api.POST("/items", CreateItem)
	}

	return router
}

// StartServer starts the HTTP server
func StartServer(address string) {
	router := SetupRouter()
	log.Printf("Server starting on %s", address)
	if err := router.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

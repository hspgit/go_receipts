package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/users", GetAllUsers)
	}

	return router
}

func StartServer(address string) {
	router := SetupRouter()
	log.Printf("Server starting on %s", address)
	if err := router.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

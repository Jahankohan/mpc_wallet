package cmd

import (
	"log"

	"github.com/Jahankohan/mpc_wallet/handlers"
	"github.com/Jahankohan/mpc_wallet/utils"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	// Load configurations
	configuration := utils.LoadConfig()

	// Set up the database
	utils.SetupDatabase(configuration.Database)


	// Create a new contract handler
	contractHandler := handlers.NewContractHandler()


	// Create a new gin router
	router := gin.Default()

	// Set up the routes
	SetupRoutes(router, contractHandler)

	// Run the server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

package cmd

import (
	"fmt"
	"log"

	"github.com/Jahankohan/mpc_wallet/handlers"
	"github.com/Jahankohan/mpc_wallet/key_manager"
	"github.com/Jahankohan/mpc_wallet/models"
	"github.com/Jahankohan/mpc_wallet/utils"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	// Load configurations
	configuration := utils.LoadConfig()
	dbConfig := configuration.Database
	km := key_manager.KeyManager{}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DBName, dbConfig.Port)
	// Set up the database
	models.SetupDatabase(dsn)


	// Create a new contract handler
	contractHandler := handlers.NewContractHandler(configuration)

	userWalletHandler := handlers.NewUserWalletHandler(km, configuration)

	transactionHandler := handlers.NewTransactionHandler(km, configuration)

	// Create a new gin router
	router := gin.Default()

	// Set up the routes
	SetupRoutes(router, contractHandler, userWalletHandler, transactionHandler)

	// Run the server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

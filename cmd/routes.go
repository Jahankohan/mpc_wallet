package cmd

import (
	"github.com/Jahankohan/mpc_wallet/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the routes for the server
func SetupRoutes(router *gin.Engine, contractHandler *handlers.ContractHandler, 
	userWalletHandler *handlers.UserWalletHandler,
	transactionHandler *handlers.TransactionHandler) {
	// Contract routes
	contractRoutes := router.Group("/contracts")
	{
		contractRoutes.POST("/", contractHandler.CreateContract)
		contractRoutes.GET("/", contractHandler.GetAllContracts)
		contractRoutes.GET("/:id", contractHandler.GetContractByID)
		contractRoutes.PUT("/:id", contractHandler.UpdateContract)
		contractRoutes.DELETE("/:id", contractHandler.DeleteContract)
		contractRoutes.GET("/:id/endpoints", contractHandler.GetContractEndpoints)
		contractRoutes.GET("/:id/endpoints/:endpoint", contractHandler.GetEndpointInputVariables)
		contractRoutes.GET("/forwarder", contractHandler.GetContractForwarders)
		contractRoutes.GET("/keymanager", contractHandler.GetContractKeyManagers)
	}

	// User routes
	userWalletRoutes := router.Group("/users")
	{
		userWalletRoutes.POST("/", userWalletHandler.CreateUser)
		userWalletRoutes.GET("/", userWalletHandler.GetAllUserWallets)
		userWalletRoutes.GET("/:id", userWalletHandler.GetUserWalletByID)
	}

	// registers the transaction handlers to the given router group.
	transactionRoutes := router.Group("/transactions")
	{
		transactionRoutes.POST("/", transactionHandler.CreateRegularTransaction)
		transactionRoutes.POST("/meta", transactionHandler.CreateMetaTransaction)
		transactionRoutes.GET("/:id", transactionHandler.GetTransactionHandler)
		transactionRoutes.GET("/", transactionHandler.GetAllTransactionsHandler)
		transactionRoutes.DELETE("/:id", transactionHandler.DeleteTransactionHandler)
	}

}

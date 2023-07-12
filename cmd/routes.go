package cmd

import (
	"net/http"

	"github.com/Jahankohan/mpc_wallet/handlers"
	"github.com/Jahankohan/mpc_wallet/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the routes for the server
func SetupProtectedRoutes(router *gin.Engine, contractHandler *handlers.ContractHandler, 
	userWalletHandler *handlers.UserWalletHandler,
	transactionHandler *handlers.TransactionHandler, userHandler *handlers.UserHandler) {

	protectedRouter := router.Group("/api")
	protectedRouter.Use(middlewares.JWTMiddleware())
	protectedRouter.Use(middlewares.RoleBasedAuthorizationMiddleware())
	
	// User Routes
	userRoutes := protectedRouter.Group("/users")
	{
		userRoutes.GET("", userHandler.GetAllUsers)
		userRoutes.GET("/:id", userHandler.GetUserByID)
	}
	
	// Contract routes
	contractRoutes := protectedRouter.Group("/contracts")
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

	// UserWallet routes
	userWalletRoutes := protectedRouter.Group("/wallets")
	{
		userWalletRoutes.POST("/", userWalletHandler.CreateUser)
		userWalletRoutes.GET("/", userWalletHandler.GetAllUserWallets)
		userWalletRoutes.GET("/:id", userWalletHandler.GetUserWalletByID)
	}

	// registers the transaction handlers to the given router group.
	transactionRoutes := protectedRouter.Group("/transactions")
	{
		transactionRoutes.POST("/", transactionHandler.CreateRegularTransaction)
		transactionRoutes.POST("/meta", transactionHandler.CreateMetaTransaction)
		transactionRoutes.GET("/:id", transactionHandler.GetTransactionHandler)
		transactionRoutes.GET("/", transactionHandler.GetAllTransactionsHandler)
		transactionRoutes.DELETE("/:id", transactionHandler.DeleteTransactionHandler)
	}

}

func SetupPublicRoutes(router *gin.Engine, contractHandler *handlers.ContractHandler, 
	userWalletHandler *handlers.UserWalletHandler,
	transactionHandler *handlers.TransactionHandler, userHandler *handlers.UserHandler) {
	
	publicRouter := router.Group("/api")

	// User routes
	userRoutes := publicRouter.Group("/users")
	{
		userRoutes.POST("/register", userHandler.RegisterUser)
		userRoutes.POST("/login", userHandler.AuthenticateUser)
	}
	
}

func SetupRoutes(router *gin.Engine, contractHandler *handlers.ContractHandler, 
	userWalletHandler *handlers.UserWalletHandler,
	transactionHandler *handlers.TransactionHandler, userHandler *handlers.UserHandler) {
	
	// Set up protected routes
	SetupProtectedRoutes(router, contractHandler, userWalletHandler, transactionHandler, userHandler)

	// Set up public routes
	SetupPublicRoutes(router, contractHandler, userWalletHandler, transactionHandler, userHandler)

	// Define a NoRoute handler to return JSON response for unknown routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
	})

	// Define a NoMethod handler to return JSON response for unknown methods
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	})
}

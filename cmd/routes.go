package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/Jahankohan/mpc_wallet/handlers"
)

// SetupRoutes configures the routes for the server
func SetupRoutes(router *gin.Engine, contractHandler *handlers.ContractHandler) {
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
	}

	// // User routes
	// userRoutes := router.Group("/users")
	// {
	// 	userRoutes.POST("/", userHandler.CreateUser)
	// 	userRoutes.GET("/", userHandler.GetAllUsers)
	// 	userRoutes.GET("/:id", userHandler.GetUserByID)
	// 	userRoutes.PUT("/:id", userHandler.UpdateUser)
	// 	userRoutes.DELETE("/:id", userHandler.DeleteUser)
	// }
}

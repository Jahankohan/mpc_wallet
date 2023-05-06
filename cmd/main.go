package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Jahankohan/mpc_wallet/handlers"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/register", handlers.RegisterUser)
		api.POST("/authenticate", handlers.AuthenticateUser)
	}

	router.Run(":8080")
}

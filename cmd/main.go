package main

import (
	"github.com/Jahankohan/mpc_wallet/handlers"
	"github.com/gin-gonic/gin"
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

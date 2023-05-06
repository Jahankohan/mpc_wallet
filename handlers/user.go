package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Jahankohan/mpc_wallet/models"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add user registration logic here

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func AuthenticateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add user authentication logic here

	c.JSON(http.StatusOK, gin.H{"message": "User authenticated successfully"})
}

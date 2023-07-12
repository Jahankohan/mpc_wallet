package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Jahankohan/mpc_wallet/models"
	"github.com/Jahankohan/mpc_wallet/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get JWT token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Parse JWT token
		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.GetJWTSecret()), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		fmt.Println("Token:", token)

		// Verify token claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    	// Retrieve user ID from token claims
    	userID := claims["user_id"].(float64) // Update the key to "user_id"

	    // Retrieve user from database based on userID
    	user, err := models.GetUserByID(uint(userID))
	    if err != nil {
        	c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
    	    c.Abort()
	        return
    	}

    	// Set user context
    	c.Set("user", user)

    	// Check role-based authorization
    	if endpoint := c.FullPath(); !hasPermission(user, endpoint) {
        	c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
    	    c.Abort()
	        return
    	}

    	// Proceed to the next handler
	    c.Next()
		} else {
	    	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	    	c.Abort()
		    return
		}
	}
}


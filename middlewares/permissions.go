package middlewares

import (
	"net/http"

	"github.com/Jahankohan/mpc_wallet/models"
	"github.com/gin-gonic/gin"
)

type Permission struct {
	Endpoint string   `json:"endpoint"`
	Roles    []string `json:"roles"`
}

var permissions []Permission

func LoadPermissions() {
	// Load permissions from a configuration file or database
	// and populate the `permissions` slice accordingly
	permissions = []Permission{
		{
			Endpoint: "/api/contracts/",
			Roles:    []string{"Admin", "Viewer"},
		},
		{
			Endpoint: "/api/transactions/",
			Roles:    []string{"Admin"},
		},
	}
}

func RoleBasedAuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user from the context
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Check if the user's role has permission to access the endpoint
		if endpoint := c.FullPath(); !hasPermission(user.(*models.User), endpoint) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden"})
			c.Abort()
			return
		}

		// Proceed to the next handler
		c.Next()
	}
}

func hasPermission(user *models.User, endpoint string) bool {
	flag := false
	for _, permission := range permissions {
		if permission.Endpoint == endpoint {
			for _, role := range permission.Roles {
				if user.Role == role {
					return true
				} else {
					flag = true
				}
			}
			break
		}
	}
	return !flag
}
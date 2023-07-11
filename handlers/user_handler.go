package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Jahankohan/mpc_wallet/models"
	"github.com/Jahankohan/mpc_wallet/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type UserHandler struct{
}


type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Convert User model to UserResponse
func ConvertToUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// RegisterUser handles user registration
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the username or email is already taken
	existingUser, err := models.GetUserByUsernameOrEmail(user.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username or email already taken"})
		return
	}

	// Encrypt the user's password
	encryptedPassword, err := utils.EncryptPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
		return
	}

	// Set the encrypted password
	user.Password = encryptedPassword

	// Create the user
	if err := models.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}


func (h *UserHandler) AuthenticateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the user from the database based on the provided username or email
	dbUser, err := models.GetUserByUsernameOrEmail(user.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Verify the password
	if err := utils.VerifyPassword(user.Password, dbUser.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(dbUser.ID, dbUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set the token in the response header
	c.Header("Authorization", "Bearer "+token)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Convert User models to UserResponse structs
	var usersResponse []*UserResponse
	for _, user := range users {
		userResponse := ConvertToUserResponse(&user)
		usersResponse = append(usersResponse, userResponse)
	}

	c.JSON(http.StatusOK, usersResponse)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := models.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Convert User model to UserResponse
	response := ConvertToUserResponse(user)

	c.JSON(http.StatusOK, response)
}


// Middleware to verify JWT token and set user context
func (h *UserHandler) JWTMiddleware() gin.HandlerFunc {
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

		// Verify token claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Retrieve user ID and role from token claims
			userID := claims["userId"].(float64)
			role := claims["role"].(string)

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
			if !checkAuthorization(role, c.Request.Method, c.Request.URL.Path) {
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

func checkAuthorization(role string, method string, path string) bool {
	// allowedRoles := map[string]map[string][]string{
	// 	"GET": {
	// 		"/api/users": {"admin", "viewer"},
	// 		"/api/posts": {"admin", "viewer"},
	// 	},
	// 	"POST": {
	// 		"/api/users": {"admin"},
	// 		"/api/posts": {"admin"},
	// 	},
	// 	"PUT": {
	// 		"/api/users": {"admin"},
	// 		"/api/posts": {"admin"},
	// 	},
	// 	"DELETE": {
	// 		"/api/users": {"admin"},
	// 		"/api/posts": {"admin"},
	// 	},
	// }

	// // Check if the path and method exist in the allowedRoles map
	// if methodRoles, ok := allowedRoles[method]; ok {
	// 	if allowedRoles, ok := methodRoles[path]; ok {
	// 		// Check if the user's role is in the list of allowed roles for the specific path and method
	// 		for _, allowedRole := range allowedRoles {
	// 			if role == allowedRole {
	// 				return true
	// 			}
	// 		}
	// 	}
	// }
	return true
}



// Authorization middleware to check user role
func (h *UserHandler) AuthorizeMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user from the request context
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Verify if the user has the required role
		if user.(*models.User).Role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		// Proceed to the next handler
		c.Next()
	}
}

// Generate JWT token
func generateJWTToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with secret key
	tokenString, err := token.SignedString([]byte(utils.GetJWTSecret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (h *UserHandler) RoleBasedAuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user from the context
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Check if the user's role is allowed for the requested path and method
		if userRole := user.(*models.User).Role; userRole != "" {
			if !checkAuthorization(userRole, c.Request.Method, c.Request.URL.Path) {
				c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden"})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user"})
			c.Abort()
			return
		}

		// Proceed to the next handler
		c.Next()
	}
}

// Example protected endpoint for admin role
func (h *UserHandler) AdminEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Admin endpoint"})
}

// Example protected endpoint for viewer role
func (h *UserHandler) ViewerEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Viewer endpoint"})
}

package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Jahankohan/mpc_wallet/models"
	"github.com/Jahankohan/mpc_wallet/utils"
	"github.com/gin-gonic/gin"
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


// Example protected endpoint for admin role
func (h *UserHandler) AdminEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Admin endpoint"})
}

// Example protected endpoint for viewer role
func (h *UserHandler) ViewerEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Viewer endpoint"})
}

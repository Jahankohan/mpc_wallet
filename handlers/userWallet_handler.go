package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/Jahankohan/mpc_wallet/key_manager"
	"github.com/Jahankohan/mpc_wallet/models"
	"github.com/Jahankohan/mpc_wallet/utils"
	"github.com/gin-gonic/gin"
)

type UserWalletHandler struct {
	km            key_manager.KeyManager
	configuration config.Configurations
}

func NewUserWalletHandler(km key_manager.KeyManager, configuration config.Configurations) *UserWalletHandler {
	return &UserWalletHandler{
		km:            km,
		configuration: configuration,
	}
}

func (h *UserWalletHandler) CreateUser(c *gin.Context) {
	// Parse the request body to get the userId
	var newUserWallet models.UserWallet
	if err := c.ShouldBindJSON(&newUserWallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	confs := utils.GetNetworkConfigurations(h.configuration, true)
	fmt.Println("Network Configurations:", confs)

	// Generate private key
	privateKey := h.km.CreatePrivateKey()

	// Split private key into shares and store them on different networks
	shares, err := h.km.SplitToShares(privateKey, 2, 3)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to split private key"})
		return
	}
	new_shares := h.km.ConvertByteSliceToStringSlice(shares)
	// Store each share on different networks
	h.km.StoreSharesToTheBlockchain(newUserWallet.UserID, new_shares, confs) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store shares on networks"})
		return
	}

	// Create user record in the database
	walletAddress := h.km.GenerateAddress(privateKey)
	newUserWallet.WalletAddress = walletAddress.Hex()
	err = models.CreateUserWallet(&newUserWallet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return the user ID and wallet address as the response
	c.JSON(http.StatusCreated, gin.H{"userId": newUserWallet.UserID, "walletAddress": newUserWallet.WalletAddress})
}

func (h *UserWalletHandler) GetAllUserWallets(c *gin.Context) {
	userWallets, err := models.GetAllUserWallets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user wallets"})
		return
	}

	c.JSON(http.StatusOK, userWallets)
}


func (h *UserWalletHandler) GetUserWalletByID(c *gin.Context) {
	// Get the user wallet ID from the URL parameters
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user wallet ID"})
		return
	}

	// Retrieve the user wallet from the database
	userWallet, err := models.GetUserWalletByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user wallet"})
		return
	}

	// Return the user wallet as the response
	c.JSON(http.StatusOK, userWallet)
}

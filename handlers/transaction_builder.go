package handlers

import (
	"math/big"
	"net/http"

	"github.com/Jahankohan/mpc_wallet/key_manager"
	"github.com/Jahankohan/mpc_wallet/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionBuilder *transaction.TransactionBuilder
	keyManager         *key_manager.KeyManager
}

func NewTransactionHandler(transactionBuilder *transaction.TransactionBuilder, keyManager *key_manager.KeyManager) *TransactionHandler {
	return &TransactionHandler{
		transactionBuilder: transactionBuilder,
		keyManager:         keyManager,
	}
}

func (th *TransactionHandler) SendRawTransaction(c *gin.Context) {
	// Parse the request body
	var requestBody struct {
		UserID  string `json:"userID"`
		To      string `json:"to"`
		Data    string `json:"data"`
		Value   string `json:"value"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// Retrieve shares and reconstruct the private key using KeyManager
	shares, err := th.keyManager.RetrieveShares(requestBody.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve key shares"})
		return
	}
	privateKey, err := th.keyManager.ReconstructPrivateKey(shares)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reconstruct private key"})
		return
	}

	// Convert the value string to a big integer
	value := new(big.Int)
	value.SetString(requestBody.Value, 10)

	// Sending the transaction using the transactionBuilder from middleware
	txHash, err := th.transactionBuilder.SendRawTransaction(privateKey, common.HexToAddress(requestBody.To), value, common.FromHex(requestBody.Data))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send transaction"})
		return
	}

	// Return the transaction hash
	c.JSON(http.StatusOK, gin.H{
		"status":          "Transaction sent",
		"transactionHash": txHash,
	})
}

package handlers

import (
	"log"
	"math/big"
	"net/http"

	"github.com/Jahankohan/mpc_wallet/transaction"
	"github.com/Jahankohan/mpc_wallet/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
    transactionBuilder *transaction.TransactionBuilder
}

type TransactionRequest struct {
	UserId  string `json:"userId"`
	To      string `json:"to"`
	Value   string `json:"value"`
	Data    string `json:"data"` // Additional field for transaction data
	Network string `json:"network"`
}

func NewTransactionHandler(tb *transaction.TransactionBuilder) *TransactionHandler {
    return &TransactionHandler{
        transactionBuilder: tb,
    }
}

func (th *TransactionHandler) HandleTransaction(c *gin.Context) {
	var requestData TransactionRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate network type
	if requestData.Network != "mainnet" && requestData.Network != "testnet" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid network type. Must be 'mainnet' or 'testnet'"})
		return
	}

	is_testnet := true
	if requestData.Network == "mainnet" {
		is_testnet = false
	}

	configurations := utils.LoadConfig()
	// Retrieve network configurations based on the specified network type
	networkConfigurations := utils.GetNetworkConfigurations(configurations, is_testnet)

	// Convert strings to appropriate types
	toAddress := common.HexToAddress(requestData.To)
	value, success := new(big.Int).SetString(requestData.Value, 10)
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value format"})
		return
	}
	txData := []byte(requestData.Data)

	// Send raw transaction
	txHash, err := th.transactionBuilder.SendRawTransaction(requestData.UserId, networkConfigurations, toAddress, value, txData)
	if err != nil {
		log.Fatalf("Error sending transaction: %v", err)
	}

	// Response
	response := struct {
		TransactionHash string `json:"transactionHash"`
	}{
		TransactionHash: txHash,
	}

	c.JSON(http.StatusOK, response)
}

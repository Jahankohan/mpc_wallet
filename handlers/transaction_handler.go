// handlers/transaction_handler.go
package handlers

import (
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/Jahankohan/mpc_wallet/key_manager"
	"github.com/Jahankohan/mpc_wallet/models"
	"github.com/Jahankohan/mpc_wallet/transaction"
	"github.com/Jahankohan/mpc_wallet/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type CreateTransactionRequest struct {
	UserID          string   `json:"userID"`
	IsTestnet       bool     `json:"isTestnet"`
	Network         string   `json:"network"`
	TargetContract  string   `json:"targetContract"`
	FunctionName    string   `json:"functionName"`
	Args            []string `json:"args"`
}

type TransactionHandler struct {
	km					key_manager.KeyManager
	configuration 		config.Configurations
}


func NewTransactionHandler(km key_manager.KeyManager, configuration config.Configurations ) *TransactionHandler {
	return &TransactionHandler{
		km:            		km,
		configuration: 		configuration,
	}
}

func (h *TransactionHandler) CreateRegularTransaction(c *gin.Context) {
	var req CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	networkType := "mainnet"
	if req.IsTestnet {
		networkType = "testnet"
	}
	networkConfig, err := utils.GetSpecificNetworkConfiguration(h.configuration, networkType, req.Network)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(networkConfig.Network)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum network: %s", err)
	}

	networks := utils.GetNetworkConfigurations(h.configuration, req.IsTestnet)

	// Initialize Transaction Broadcaster and Handler
	broadcaster := transaction.NewTransactionBroadcaster(client)
	metaProcessor, err := transaction.NewMetaTransactionProcessor(networkConfig.Network)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the TransactionHandler with the broadcaster and metaProcessor
	handler := transaction.NewTransactionHandler(broadcaster, *metaProcessor, &h.km)

	contractID, err := strconv.ParseUint(req.TargetContract, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contract ID"})
		return
	}
	contract, err := models.GetContractByID(uint(contractID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve contract"})
		return
	}

	abiStr := contract.ABI
	contractAddress := common.HexToAddress(contract.Address)

	value := new(big.Int)
	value.SetString(req.Args[0], 10)

	args := []interface{}{value}

	txHash, err := handler.HandleTransaction(networkConfig, req.UserID, networks, abiStr, contractAddress, req.FunctionName, args...)
	if err != nil {
		log.Fatalf("Failed to handle transaction: %s", err)
	}
	argsStr := strings.Join(req.Args, ",")
	// Create a transaction object
	transactionObj := models.Transaction{
		HashorResult:       txHash.(string),
		UserID:             req.UserID,
		Network:            req.Network,
		TargetContract:     contractAddress.Hex(),
		TargetFunction:     req.FunctionName,
		Args:               argsStr,
		IsMetaTransaction:  false,
		TransactionType:    models.MethodTypeWrite,
		IsSuccessful:       true,                        // Set the success status based on the actual result
		IsTestnet:          req.IsTestnet,
	}

	// Save the transaction object into the database
	err = models.CreateTransaction(&transactionObj)
	if err != nil {
		log.Fatalf("Failed to create transaction: %s", err)
	}

	c.JSON(http.StatusOK, gin.H{"result": txHash})
}




// GetTransactionHandler handles the retrieval of a transaction by ID.
func (h *TransactionHandler) GetTransactionHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transaction, err := models.GetTransactionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// GetAllTransactionsHandler handles the retrieval of all transactions.
func (h *TransactionHandler) GetAllTransactionsHandler(c *gin.Context) {
	transactions, err := models.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// DeleteTransactionHandler handles the deletion of a transaction.
func (h *TransactionHandler) DeleteTransactionHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transaction, err := models.GetTransactionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = models.DeleteTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
package handlers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jahankohan/mpc_wallet/transaction"
	"github.com/Jahankohan/mpc_wallet/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type TransactionHandler struct {
	transactionCreator  transaction.TransactionCreator
	transactionSigner   transaction.TransactionSigner
	transactionBroadcaster transaction.TransactionBroadcaster
}

type TransactionRequest struct {
	UserId 	string	`json:"userId"`
	To     	string	`json:"to"`
	Value  	string	`json:"value"`
	Data   	string	`json:"data"` // Additional field for transaction data
	Network	string	`json:"network"` // Network type (e.g., "mainnet" or "testnet")
}

func NewTransactionHandler(tc transaction.TransactionCreator, ts transaction.TransactionSigner, tb transaction.TransactionBroadcaster) *TransactionHandler {
	return &TransactionHandler{
		transactionCreator:  tc,
		transactionSigner:   ts,
		transactionBroadcaster: tb,
	}
}

func (th *TransactionHandler) HandleTransaction(w http.ResponseWriter, r *http.Request) {
	var requestData TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate network type
	if requestData.Network != "mainnet" && requestData.Network != "testnet" {
		http.Error(w, "Invalid network type. Must be 'mainnet' or 'testnet'", http.StatusBadRequest)
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
	value, err := strconv.ParseInt(requestData.Value, 10, 64)
	if err != nil {
		http.Error(w, "Invalid value format", http.StatusBadRequest)
		return
	}
	txData := []byte(requestData.Data)

	// Create transaction
	unsignedTx, err := th.transactionCreator.CreateTransaction(toAddress, value, txData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create transaction: %v", err), http.StatusInternalServerError)
		return
	}

	// Sign transaction
	signedTx, err := th.transactionSigner.SignTransaction(requestData.UserId, networkConfigurations, unsignedTx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to sign transaction: %v", err), http.StatusInternalServerError)
		return
	}

	// Encode transaction to raw bytes and convert to hex string
	rawTxBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode transaction: %v", err), http.StatusInternalServerError)
		return
	}
	rawTxHex := hex.EncodeToString(rawTxBytes)

	// Response
	response := struct {
		SignedTransaction string `json:"signedTransaction"`
	}{
		SignedTransaction: rawTxHex,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

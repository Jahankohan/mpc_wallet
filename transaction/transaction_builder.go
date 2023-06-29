package transaction

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/Jahankohan/mpc_wallet/key_manager"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionBuilder struct {
	client     *ethclient.Client
	keyManager *key_manager.KeyManager
}

func NewTransactionBuilder(client *ethclient.Client, keyManager *key_manager.KeyManager) *TransactionBuilder {
	return &TransactionBuilder{
		client:     client,
		keyManager: keyManager,
	}
}

func (tb *TransactionBuilder) ReadContract(abiStr string, contractAddress common.Address, functionName string, args ...interface{}) (string, error) {
	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return "", fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	// Pack the function call
	data, err := parsedABI.Pack(functionName, args...)
	if err != nil {
		return "", fmt.Errorf("failed to pack data for function %q: %w", functionName, err)
	}

	// Define the call message
	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	// Call the contract
	result, err := tb.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %w", err)
	}

	// The result is the raw bytes returned by the contract function. 
	// You may want to further process this result depending on your application's needs.
	return string(result), nil
}

func (tb *TransactionBuilder) WriteContract(userID string, confs []config.NetworkConfiguration, abiStr string, contractAddress common.Address, functionName string, args ...interface{}) (string, error) {
	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return "", fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	// Pack the function call
	data, err := parsedABI.Pack(functionName, args...)
	if err != nil {
		return "", fmt.Errorf("failed to pack data for function %q: %w", functionName, err)
	}

	// Retrieve shares and reconstruct the private key using KeyManager
	shares, err := tb.keyManager.RetrieveAllShares(confs, userID)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve key shares: %v", err)
	}
	privateKey, err := tb.keyManager.ReconstructPrivateKey(shares)
	if err != nil {
		return "", fmt.Errorf("failed to reconstruct private key: %v", err)
	}
	log.Println("Reconstructed Key:", privateKey)
	log.Println("Public Key:", privateKey.PublicKey)

	// Get the nonce for the account
	nonce, err := tb.client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		return "", err
	}

	// Get the gas price
	gasPrice, err := tb.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	// Estimate the gas needed for the transaction
	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}
	gasLimit, err := tb.client.EstimateGas(context.Background(), msg)
	if err != nil {
		return "", err
	}

	// Create a new transaction
	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// Sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1)), privateKey)
	if err != nil {
		return "", err
	}

	// Log and return the transaction hash
	txHash := signedTx.Hash().Hex()
	log.Printf("Transaction sent: %s", txHash)
	err = tb.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return txHash, nil
}

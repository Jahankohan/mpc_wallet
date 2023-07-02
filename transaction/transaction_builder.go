package transaction

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
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
	keyManager    *key_manager.KeyManager
	networkConfig config.NetworkConfiguration
}

func NewTransactionBuilder(keyManager *key_manager.KeyManager, networkConfig config.NetworkConfiguration) *TransactionBuilder {
	return &TransactionBuilder{
		keyManager:    keyManager,
		networkConfig: networkConfig,
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

	// Create an Ethereum client dynamically
	client, err := ethclient.Dial(tb.networkConfig.Network)
	if err != nil {
		return "", fmt.Errorf("failed to connect to Ethereum RPC: %w", err)
	}
	defer client.Close()

	// Call the contract
	result, err := client.CallContract(context.Background(), msg, nil)
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

	// Create an Ethereum client dynamically
	client, err := ethclient.Dial(tb.networkConfig.Network)
	if err != nil {
		return "", fmt.Errorf("failed to connect to Ethereum RPC: %w", err)
	}
	defer client.Close()

	// Get the nonce for the account
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		return "", err
	}

	// Get the gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	// Estimate the gas needed for the transaction
	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return "", err
	}

	// Create a new transaction
	fmt.Println("Data:", data)
	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// Convert ChainId from string to *big.Int
	chainID, success := new(big.Int).SetString(tb.networkConfig.ChainId, 10)
	if !success {
		return "", fmt.Errorf("invalid chain ID")
	}

	// Sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	// Log and return the transaction hash
	txHash := signedTx.Hash().Hex()
	log.Printf("Transaction sent: %s", txHash)
	return txHash, nil
}

func (tb *TransactionBuilder) WriteMetaContract(userID string, abiStr string, contractAddress common.Address, functionName string, args ...interface{}) (string, error) {
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

	// Create an Ethereum client dynamically
	client, err := ethclient.Dial(tb.networkConfig.Network)
	if err != nil {
		return "", fmt.Errorf("failed to connect to Ethereum RPC: %w", err)
	}
	defer client.Close()

	// Get the relayer's address
	relayerAddress := common.HexToAddress(os.Getenv("RELAYER_ADDRESS"))

	// Get the nonce for the relayer account
	nonce, err := client.PendingNonceAt(context.Background(), relayerAddress)
	if err != nil {
		return "", err
	}

	// Get the gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	// Estimate the gas needed for the transaction
	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return "", err
	}

	// Create a new transaction
	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// Convert ChainId from string to *big.Int
	chainID, success := new(big.Int).SetString(tb.networkConfig.ChainId, 10)
	if !success {
		return "", fmt.Errorf("invalid chain ID")
	}

	// Get relayer's private key from environment variable
	privateKeyStr := os.Getenv("RELAYER_PRIVATE_KEY")
	if privateKeyStr == "" {
		return "", fmt.Errorf("RELAYER_PRIVATE_KEY environment variable is not set")
	}
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return "", fmt.Errorf("failed to convert private key: %s", err)
	}

	// Sign the transaction with relayer's private key
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	// Log and return the transaction hash
	txHash := signedTx.Hash().Hex()
	log.Printf("Meta-transaction sent: %s", txHash)
	return txHash, nil
}


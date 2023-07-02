package transaction

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionBroadcaster interface {
	BroadcastRegularTransaction(ctx context.Context, privateKey string, to common.Address, value *big.Int, data []byte, gasLimit uint64, gasPrice *big.Int, chainID *big.Int) (string, error)
	BroadcastMetaTransaction(ctx context.Context, fromAddress common.Address, to common.Address, value *big.Int, data []byte, gasLimit uint64, gasPrice *big.Int, chainID *big.Int) (string, error)
	EstimateGasLimit(ctx context.Context, to common.Address, data []byte) (uint64, error)
	EstimateGasPrice(ctx context.Context) (*big.Int, error)
	SignAndBroadcastTransaction(ctx context.Context, privateKey *ecdsa.PrivateKey, fromAddress common.Address, to common.Address, data []byte, gasPrice *big.Int, gasLimit uint64, chainID *big.Int) (string, error)
}

type transactionBroadcaster struct {
	client *ethclient.Client
}

func NewTransactionBroadcaster(client *ethclient.Client) TransactionBroadcaster {
	return &transactionBroadcaster{client: client}
}

func (tb *transactionBroadcaster) BroadcastRegularTransaction(ctx context.Context, privateKey string, to common.Address, value *big.Int, data []byte, gasLimit uint64, gasPrice *big.Int, chainID *big.Int) (string, error) {
	// Convert the private key string to a private key object
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to convert private key: %s", err)
	}
	fromAddress := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)

	return tb.SignAndBroadcastTransaction(ctx, privateKeyECDSA, fromAddress, to, data, gasPrice, gasLimit, chainID)
}

func (tb *transactionBroadcaster) BroadcastMetaTransaction(ctx context.Context, fromAddress common.Address, to common.Address, value *big.Int, data []byte, gasLimit uint64, gasPrice *big.Int, chainID *big.Int) (string, error) {
	// Get relayer's private key from environment variable
	privateKeyStr := os.Getenv("RELAYER_PRIVATE_KEY")
	if privateKeyStr == "" {
		return "", fmt.Errorf("RELAYER_PRIVATE_KEY environment variable is not set")
	}
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return "", fmt.Errorf("failed to convert private key: %s", err)
	}

	return tb.SignAndBroadcastTransaction(ctx, privateKey, fromAddress, to, data, gasPrice, gasLimit, chainID)
}

func (tb *transactionBroadcaster) EstimateGasLimit(ctx context.Context, to common.Address, data []byte) (uint64, error) {
	msg := ethereum.CallMsg{To: &to, Data: data}
	gasLimit, err := tb.client.EstimateGas(ctx, msg)
	if err != nil {
		return 0, err
	}
	return gasLimit, nil
}

func (tb *transactionBroadcaster) EstimateGasPrice(ctx context.Context) (*big.Int, error) {
	gasPrice, err := tb.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	return gasPrice, nil
}

func (tb *transactionBroadcaster) SignAndBroadcastTransaction(ctx context.Context, privateKey *ecdsa.PrivateKey, fromAddress common.Address, to common.Address, data []byte, gasPrice *big.Int, gasLimit uint64, chainID *big.Int) (string, error) {
	// Retrieve the nonce for the account
	nonce, err := tb.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Printf("Error retrieving account nonce: %s", err)
		return "", err
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, to, new(big.Int), gasLimit, gasPrice, data)

	// Sign the transaction with provided private key
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %s", err)
	}

	// Broadcast the transaction
	if err := tb.client.SendTransaction(ctx, signedTx); err != nil {
		return "", fmt.Errorf("failed to broadcast transaction: %s", err)
	}

	// Return transaction hash
	return signedTx.Hash().Hex(), nil
}

package transaction

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/Jahankohan/mpc_wallet/key_manager"
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

func (tb *TransactionBuilder) SendRawTransaction(userID string, toAddress common.Address, value *big.Int, data []byte) (string, error) {
	// Retrieve shares and reconstruct the private key using KeyManager
	shares, err := tb.keyManager.RetrieveShares(userID)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve key shares: %v", err)
	}
	privateKey, err := tb.keyManager.ReconstructPrivateKey(shares)
	if err != nil {
		return "", fmt.Errorf("failed to reconstruct private key: %v", err)
	}

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

	// Create a new transaction
	tx := types.NewTransaction(nonce, toAddress, value, 21000, gasPrice, data)

	// Sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1)), privateKey)
	if err != nil {
		return "", err
	}

	// Broadcast the transaction
	err = tb.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	// Log and return the transaction hash
	txHash := signedTx.Hash().Hex()
	log.Printf("Transaction sent: %s", txHash)
	return txHash, nil
}

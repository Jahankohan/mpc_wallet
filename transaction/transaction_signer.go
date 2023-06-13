package transaction

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/Jahankohan/mpc_wallet/key_manager"
)

type TransactionSigner interface {
	SignTransaction(userId string, conf []config.NetworkConfiguration, unsignedTx *types.Transaction) (*types.Transaction, error)
}

type transactionSigner struct {
	keyManager key_manager.KeyManager
}

func NewTransactionSigner(keyManager key_manager.KeyManager) TransactionSigner {
	return &transactionSigner{
		keyManager: keyManager,
	}
}

func (ts *transactionSigner) SignTransaction(userId string, conf []config.NetworkConfiguration, unsignedTx *types.Transaction) (*types.Transaction, error) {
	// Retrieve private key of the user using KeyManager
	shares, err := ts.keyManager.RetrieveAllShares(conf, userId)
	if err != nil {
		return nil, err
	}

	keyManager := key_manager.KeyManager{}
	privateKey, err := keyManager.ReconstructPrivateKey(shares)
	if err != nil {
		return nil, err
	}
	
	// Sign the transaction
	signer := types.NewEIP155Signer(big.NewInt(1)) // example chainID
	signedTx, err := types.SignTx(unsignedTx, signer, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	return signedTx, nil
}

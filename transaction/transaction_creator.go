package transaction

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

type TransactionCreator interface {
	CreateTransaction(to common.Address, value int64, data []byte) (*types.Transaction, error)
}

type transactionCreator struct{}

func NewTransactionCreator() TransactionCreator {
	return &transactionCreator{}
}

func (tc *transactionCreator) CreateTransaction(to common.Address, value int64, data []byte) (*types.Transaction, error) {
	// Here you would create an unsigned transaction
	gasLimit := uint64(21000) // hardcoded gas limit
	gasPrice := big.NewInt(18000000000) // hardcoded gas price
	chainID := big.NewInt(params.MainnetChainConfig.ChainID.Int64()) // Mainnet chain ID
	nonce := uint64(0) // You should fetch the actual nonce

	tx := types.NewTransaction(nonce, to, big.NewInt(value), gasLimit, gasPrice, data)

	signer := types.NewEIP155Signer(chainID)
	unsignedTx, _ := types.SignTx(tx, signer, nil) // nil private key for unsigned tx

	return unsignedTx, nil
}

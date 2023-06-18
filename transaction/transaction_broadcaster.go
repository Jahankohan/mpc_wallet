package transaction

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionBroadcaster interface {
	BroadcastTransaction(signedTx *types.Transaction)
}

type transactionBroadcaster struct {
	client *ethclient.Client
}

func NewTransactionBroadcaster(client *ethclient.Client) TransactionBroadcaster {
	return &transactionBroadcaster{client: client}
}

func (tb *transactionBroadcaster) BroadcastTransaction(signedTx *types.Transaction) {
	err := tb.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

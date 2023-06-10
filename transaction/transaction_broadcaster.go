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

type transactionBroadcaster struct{}

func NewTransactionBroadcaster() TransactionBroadcaster {
	return &transactionBroadcaster{}
}

func (tb *transactionBroadcaster) BroadcastTransaction(signedTx *types.Transaction) {
	client, err := ethclient.Dial("https://mainnet.infura.io") // example endpoint
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

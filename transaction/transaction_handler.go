package transaction

import (
	"log"
	"math/big"

	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/Jahankohan/mpc_wallet/key_manager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionHandler struct {
	builder     *TransactionBuilder
	broadcaster TransactionBroadcaster
}

func NewTransactionHandler(client *ethclient.Client, keyManager *key_manager.KeyManager, broadcaster TransactionBroadcaster) *TransactionHandler {
	return &TransactionHandler{
		builder:     NewTransactionBuilder(client, keyManager),
		broadcaster: broadcaster,
	}
}

func (th *TransactionHandler) HandleTransaction(userId string, conf []config.NetworkConfiguration, to common.Address, value int64, data []byte) {
	txHash, err := th.builder.SendRawTransaction(userId, conf, to, big.NewInt(value), data)
	if err != nil {
		log.Fatalf("Error creating and signing transaction: %v", err)
		return
	}

	log.Printf("Transaction successfully sent. Hash: %s", txHash)
}

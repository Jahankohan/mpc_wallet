package transaction

import (
	"log"

	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/ethereum/go-ethereum/common"
)

type TransactionHandler struct {
	creator    TransactionCreator
	signer     TransactionSigner
	broadcaster TransactionBroadcaster
}

func NewTransactionHandler(creator TransactionCreator, signer TransactionSigner, broadcaster TransactionBroadcaster) *TransactionHandler {
	return &TransactionHandler{
		creator:    creator,
		signer:     signer,
		broadcaster: broadcaster,
	}
}

func (th *TransactionHandler) HandleTransaction(userId string, conf []config.NetworkConfiguration, to common.Address, value int64, data []byte) {
	unsignedTx, err := th.creator.CreateTransaction(to, value, data)
	if err != nil {
		log.Fatalf("Error creating transaction: %v", err)
		return
	}

	signedTx, err := th.signer.SignTransaction(userId, conf, unsignedTx)
	if err != nil {
		log.Fatalf("Error signing transaction: %v", err)
		return
	}

	th.broadcaster.BroadcastTransaction(signedTx)
}

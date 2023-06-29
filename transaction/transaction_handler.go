package transaction

import (
	"fmt"
	"strings"

	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/Jahankohan/mpc_wallet/key_manager"
	"github.com/ethereum/go-ethereum/accounts/abi"
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

func (th *TransactionHandler) HandleTransaction(userId string, conf []config.NetworkConfiguration, abiStr string, contractAddress common.Address, functionName string, args ...interface{}) (interface{}, error) {
	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	// Check if the function is constant
	method, exist := parsedABI.Methods[functionName]
	if !exist {
		return nil, fmt.Errorf("function %s does not exist in contract ABI", functionName)
	}


	// If the function is constant, call ReadContract. Otherwise, call WriteContract.
	if (method.StateMutability == "pure" || method.StateMutability == "view")  {
		return th.builder.ReadContract(abiStr, contractAddress, functionName, args...)
	} else {
		return th.builder.WriteContract(userId, conf, abiStr, contractAddress, functionName, args...)
	}
}
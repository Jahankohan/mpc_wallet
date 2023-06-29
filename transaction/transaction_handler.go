package transaction

import (
	"fmt"
	"strings"

	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/Jahankohan/mpc_wallet/key_manager"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type TransactionHandler struct {
	broadcaster TransactionBroadcaster
}

func NewTransactionHandler(broadcaster TransactionBroadcaster) *TransactionHandler {
	return &TransactionHandler{
		broadcaster: broadcaster,
	}
}

func (th *TransactionHandler) HandleTransaction(networkConfig config.NetworkConfiguration, keyManager *key_manager.KeyManager, userId string, confs []config.NetworkConfiguration, abiStr string, contractAddress common.Address, functionName string, args ...interface{}) (interface{}, error) {
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

	// Create TransactionBuilder with the provided NetworkConfiguration
	builder := NewTransactionBuilder(keyManager, networkConfig)

	// If the function is constant, call ReadContract. Otherwise, call WriteContract.
	if method.StateMutability == "pure" || method.StateMutability == "view" {
		return builder.ReadContract(abiStr, contractAddress, functionName, args...)
	} else {
		// Pass network configurations for key retrieval, this might need to be adjusted based on your setup.
		return builder.WriteContract(userId, confs, abiStr, contractAddress, functionName, args...)
	}
}

package transaction

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/Jahankohan/mpc_wallet/key_manager"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type TransactionHandler struct {
	broadcaster           TransactionBroadcaster
	metaProcessor         MetaTransactionProcessor
	transactionBuilder    *TransactionBuilder
}

func NewTransactionHandler(broadcaster TransactionBroadcaster, metaProcessor MetaTransactionProcessor, keyManager *key_manager.KeyManager, networkConfig config.NetworkConfiguration) *TransactionHandler {
	return &TransactionHandler{
		broadcaster:        broadcaster,
		metaProcessor:      metaProcessor,
		transactionBuilder: NewTransactionBuilder(keyManager, networkConfig),
	}
}

func (th *TransactionHandler) HandleTransaction(networkConfig config.NetworkConfiguration, userId string, confs []config.NetworkConfiguration, abiStr string, contractAddress common.Address, functionName string, args ...interface{}) (interface{}, error) {
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
	if method.StateMutability == "pure" || method.StateMutability == "view" {
		return th.transactionBuilder.ReadContract(abiStr, contractAddress, functionName, args...)
	} else {
		// Pass network configurations for key retrieval, this might need to be adjusted based on your setup.
		return th.transactionBuilder.WriteContract(userId, confs, abiStr, contractAddress, functionName, args...)
	}
}

func (th *TransactionHandler) HandleMetaTransaction(ctx context.Context, networkConfig config.NetworkConfiguration, keyManager *key_manager.KeyManager, userId string, confs []config.NetworkConfiguration, forwarderAbiStr string, forwarderAddress common.Address, forwarderFunctionName string, nonce *big.Int, contractAddress common.Address, contractAbiStr string, contractFunctionName string, args ...interface{}) (string, error) {
	// Call the forwarder contract with the encoded meta-transaction as a parameter
	txHash, err := th.metaProcessor.ProcessMetaTransaction(ctx, networkConfig, keyManager, userId, confs, forwarderAbiStr, forwarderAddress, forwarderFunctionName, nonce, contractAddress, contractAbiStr, contractFunctionName, args...)
	if err != nil {
	    return "", fmt.Errorf("failed to send meta-transaction through forwarder: %w", err)
	}

	// Return the transaction hash
	return txHash, nil
}

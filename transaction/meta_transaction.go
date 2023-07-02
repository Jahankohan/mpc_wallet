package transaction

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/Jahankohan/mpc_wallet/key_manager"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MetaTransactionProcessor struct {
	broadcaster TransactionBroadcaster
	ethClient   *ethclient.Client
}

func NewMetaTransactionProcessor(broadcaster TransactionBroadcaster, rpcURL string) (*MetaTransactionProcessor, error) {
	// Create Ethereum client
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create Ethereum client: %v", err)
	}

	return &MetaTransactionProcessor{
		broadcaster: broadcaster,
		ethClient:   ethClient,
	}, nil
}

func (mtp *MetaTransactionProcessor) ProcessMetaTransaction(ctx context.Context, networkConfig config.NetworkConfiguration, keyManager *key_manager.KeyManager, userId string, forwarderAbiStr string, forwarderAddress common.Address, forwarderFunctionName string, nonce *big.Int, contractAddress common.Address, contractAbiStr string, contractFunctionName string, args ...interface{}) (string, error) {

	// Parse the Forwarder Contract ABI
	forwarderABI, err := abi.JSON(strings.NewReader(forwarderAbiStr))
	if err != nil {
		return "", fmt.Errorf("failed to parse forwarder contract ABI: %w", err)
	}

	// Parse the Target Contract ABI
	contractABI, err := abi.JSON(strings.NewReader(contractAbiStr))
	if err != nil {
		return "", fmt.Errorf("failed to parse target contract ABI: %w", err)
	}

	// Pack the data for the function we want to call in the target contract
	functionData, err := contractABI.Pack(contractFunctionName, args...)
	if err != nil {
		return "", fmt.Errorf("failed to pack data for target contract: %w", err)
	}

	// Pack the data for the function in the forwarder contract (processMetaTransaction) along with the required meta-transaction parameters
	data, err := forwarderABI.Pack(
        forwarderFunctionName, 
        nonce, 
        contractAddress, 
        big.NewInt(0), 
        functionData,
    )
	if err != nil {
		return "", fmt.Errorf("failed to pack data for forwarder contract: %w", err)
	}

	// Get relayer's private key from environment variable
	privateKey, err := crypto.HexToECDSA(os.Getenv("RELAYER_PRIVATE_KEY"))
	if err != nil {
		return "", fmt.Errorf("failed to load relayer's private key: %w", err)
	}

	// Estimate gas limit
	gasLimit, err := mtp.estimateGasLimit(ctx, forwarderAddress, data)
	if err != nil {
		return "", fmt.Errorf("failed to estimate gas limit: %w", err)
	}

	// Get gas price
	gasPrice, err := mtp.estimateGasPrice(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to estimate gas price: %w", err)
	}

    fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
    chainIdInt := new(big.Int)
    _, success := chainIdInt.SetString(networkConfig.ChainId, 10)
    if !success {
        log.Fatal("Error Converting ChainId")
    }

	// Sign and broadcast the meta-transaction via the forwarder contract
	txHash, err := mtp.broadcaster.BroadcastMetaTransaction(
        ctx, 
        fromAddress,
        forwarderAddress,
        big.NewInt(0), // Value is 0 for meta transactions
        data,
        gasLimit,
        gasPrice,
        chainIdInt,
    )
	if err != nil {
		return "", fmt.Errorf("failed to broadcast meta-transaction: %w", err)
	}

	// Return the transaction hash
	return txHash, nil
}

func (mtp *MetaTransactionProcessor) estimateGasLimit(ctx context.Context, toAddress common.Address, data []byte) (uint64, error) {
	// Estimate the gas required for the transaction
	msg := ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	}
	gasLimit, err := mtp.ethClient.EstimateGas(ctx, msg)
	if err != nil {
		return 0, fmt.Errorf("failed to estimate gas limit: %v", err)
	}

	return gasLimit, nil
}

func (mtp *MetaTransactionProcessor) estimateGasPrice(ctx context.Context) (*big.Int, error) {
	// Suggest gas price
	gasPrice, err := mtp.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}

	return gasPrice, nil
}

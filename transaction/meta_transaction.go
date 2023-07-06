package transaction

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
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
	"golang.org/x/crypto/sha3"
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

func hexStringToBytes32(s string) [32]byte {
    bytes, err := hex.DecodeString(s)
    if err != nil {
        log.Fatalf("Failed to decode hex string: %v", err)
    }
    
    var bytes32 [32]byte
    copy(bytes32[:], bytes)
    return bytes32
}

func (mtp *MetaTransactionProcessor) ProcessMetaTransaction(ctx context.Context, networkConfig config.NetworkConfiguration, keyManager *key_manager.KeyManager, userId string, confs []config.NetworkConfiguration, forwarderAbiStr string, forwarderAddress common.Address, forwarderFunctionName string, nonce *big.Int, contractAddress common.Address, contractAbiStr string, contractFunctionName string, args ...interface{}) (string, error) {

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

	// Retrieve shares and reconstruct the private key using KeyManager
	shares, err := keyManager.RetrieveAllShares(confs, userId)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve key shares: %v", err)
	}
	rawUserPrivateKey, err := keyManager.ReconstructPrivateKey(shares)
	
	privateKeyBytes := rawUserPrivateKey.D.Bytes()
	userPrivateKey := fmt.Sprintf("%x", privateKeyBytes)
	fmt.Println("PK:", userPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to reconstruct private key: %v", err)
	}
	userAddress := crypto.PubkeyToAddress(rawUserPrivateKey.PublicKey)

	// Retrieve nonce from the forwarder contract
	forwarderNonce, err := mtp.getNonce(ctx, userAddress, forwarderAddress, forwarderABI)
	if err != nil {
    	return "", fmt.Errorf("failed to retrieve nonce: %v", err)
	}

	fmt.Println("Forwarder Nonce:", forwarderNonce)

	dataToSign := crypto.Keccak256(
		userAddress.Bytes(),
		forwarderAddress.Bytes(),
		forwarderNonce.Bytes(),
		functionData,
	)

	prefix := []byte("\x19Ethereum Signed Message:\n32")
	prefixedData := append(prefix, dataToSign...)
	prefixedHash := crypto.Keccak256(prefixedData)

	userPrivateKeyBytes, err := hex.DecodeString(userPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Convert bytes to ecdsa.PrivateKey
	privateKeyECDSA, err := crypto.ToECDSA(userPrivateKeyBytes)
	if err != nil {
		log.Fatal(err)
	}
	// sign data with user's private key
	signature, err := crypto.Sign(prefixedHash, privateKeyECDSA)
	if err != nil {
        return "", fmt.Errorf("failed to sign data: %w", err)
    }

	chainIdInt := new(big.Int)
    _, success := chainIdInt.SetString(networkConfig.ChainId, 10)
    if !success {
        log.Fatal("Error Converting ChainId")
    }
	fmt.Println("chainIdInt:", chainIdInt)

    sigR := signature[:32]
    sigS := signature[32:64]
    sigV := signature[64]

	var sigRFixed [32]byte
	var sigSFixed [32]byte
	copy(sigRFixed[:], sigR)
	copy(sigSFixed[:], sigS)

	// Convert byte to uint8
	sigVUint8 := uint8(sigV)
	fmt.Println(sigVUint8)
	// if sigVUint8 < 27 {
	// 	sigVUint8 += byte(chainIdInt.Uint64()*2 + 35)
	// }
	// fmt.Println("New SigV value:", sigVUint8)

	// Calculate the new v value based on the chain ID
	chainID := big.NewInt(43113) // Your chainID here
	v := big.NewInt(int64(sigV)) // Your recovery id (0 or 1) as big.Int

	// Calculate sigV according to EIP-155
	v.Add(v, chainID.Mul(chainID, big.NewInt(2)))
	v.Add(v, big.NewInt(35))

	// Convert to uint8
	sigVUint8 = byte(v.Uint64())
	fmt.Println("New Sigv:", sigVUint8)
	fmt.Println("SigR:", sigRFixed)
	fmt.Println("SigS:", sigSFixed)
	fmt.Println("function:", functionData)
	fmt.Println("target:", forwarderAddress)
	fmt.Println("user address:", userAddress)

	Sv, sR, sS, err := CreateAndSignMetaTransaction(privateKeyECDSA, userAddress, 
		forwarderAddress, forwarderNonce, functionData, chainIdInt)
	if err != nil {
		fmt.Println("Unable to create and sign.")
	}
	fmt.Println("New Calculated Vares")
	fmt.Println("New Sigv:", Sv)
	fmt.Println("SigR:", sR)
	fmt.Println("SigS:", sS)

	verifyResult, err := VerifyMetaTransaction(userAddress, forwarderAddress, forwarderNonce, 
	functionData, 
	Sv,
	sR, 
	sS)
	if err != nil {
		return "", fmt.Errorf("failed to verify the hash: %w", err)
	}
	fmt.Println("Hash Verified?", verifyResult)

	// Pack the data for the function in the forwarder contract (processMetaTransaction) along with the required meta-transaction parameters
	data, err := forwarderABI.Pack(
        forwarderFunctionName, 
        userAddress, 
        functionData,
		sigRFixed,
		sigSFixed,
		sigVUint8,
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

func GetFunctionSignature(functionText string) []byte {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(functionText))
	return hasher.Sum(nil)[:4]
}

func CreateAndSignMetaTransaction(userPrivateKey *ecdsa.PrivateKey, userAddress common.Address, targetContract common.Address, nonce *big.Int, functionSignature []byte, chainID *big.Int) (uint8, [32]byte, [32]byte, error) {

	// Hash the data
	hash := crypto.Keccak256Hash(
		userAddress.Bytes(),
		targetContract.Bytes(),
		nonce.Bytes(),
		functionSignature,
	)

	// Prefix
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(hash.Bytes()))
	prefixedHash := crypto.Keccak256Hash(append([]byte(prefix), hash.Bytes()...))

	// Sign the prefixed hash with the user's private key
	signature, err := crypto.Sign(prefixedHash.Bytes(), userPrivateKey)
	if err != nil {
		return 0, [32]byte{}, [32]byte{}, fmt.Errorf("failed to sign data: %v", err)
	}

	// Extract V, R, S from the signature
	sigV := signature[64] // V
	if chainID.Sign() != 0 {
		sigV += byte(chainID.Uint64()*2 + 35)
	}
	var sigR, sigS [32]byte
	copy(sigR[:], signature[:32]) // R
	copy(sigS[:], signature[32:64]) // S

	return sigV, sigR, sigS, nil
}

func VerifyMetaTransaction(userAddress common.Address, targetContract common.Address, nonce *big.Int, functionSignature []byte, sigV uint8, sigR [32]byte, sigS [32]byte) (bool, error) {

	// Hash the data
	hash := crypto.Keccak256Hash(
		userAddress.Bytes(),
		targetContract.Bytes(),
		nonce.Bytes(),
		functionSignature,
	)

	fmt.Println("Hash Computed:", hash)

	// Prefix
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(hash.Bytes()))
	prefixedHash := crypto.Keccak256Hash(append([]byte(prefix), hash.Bytes()...))

	fmt.Println("Prefixhash computed:", prefixedHash)

	// Reconstruct the sender's public key using the signature
	sig := append(sigR[:], sigS[:]...)
	sig = append(sig, sigV-27)

	pubKey, err := crypto.SigToPub(prefixedHash.Bytes(), sig)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %v", err)
	}
	fmt.Println("PubKey Created:", pubKey)

	recoveredAddress := crypto.PubkeyToAddress(*pubKey)
	fmt.Println("Recovered Address:", recoveredAddress)
	return recoveredAddress == userAddress, nil
}

func (mtp *MetaTransactionProcessor) getNonce(ctx context.Context, userAddress common.Address, forwarderAddress common.Address, forwarderABI abi.ABI) (*big.Int, error) {
    callData, err := forwarderABI.Pack("getNonce", userAddress)
    if err != nil {
        return nil, fmt.Errorf("failed to pack data for getNonce: %v", err)
    }

    msg := ethereum.CallMsg{
        To:   &forwarderAddress,
        Data: callData,
    }
    output, err := mtp.ethClient.CallContract(ctx, msg, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to call getNonce: %v", err)
    }

    // Unpack result
    var nonce *big.Int
    err = forwarderABI.UnpackIntoInterface(&nonce, "getNonce", output)
    if err != nil {
        return nil, fmt.Errorf("failed to unpack getNonce output: %v", err)
    }

    return nonce, nil
}


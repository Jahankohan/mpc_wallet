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

	"github.com/Jahankohan/mpc_wallet/Forwarder"
	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/Jahankohan/mpc_wallet/key_manager"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func (mtp *MetaTransactionProcessor) padTo32Bytes(input []byte) []byte {
	padded := make([]byte, 32)
	copy(padded[32-len(input):], input)
	return padded
}

func (mtp *MetaTransactionProcessor) getNonce(userAddress common.Address, forwarder *Forwarder.Forwarder) (*big.Int, error){
	nonce, err := forwarder.GetNonce(&bind.CallOpts{}, userAddress)
	if err != nil {
		fmt.Printf("Failed to call getNonce: %v\n", err)
		return big.NewInt(-1), err
	}
	return nonce, nil
}

func (mtp *MetaTransactionProcessor) convertNonceToBytes(nonce string) ([]byte, error) {
	nonceInBytes := mtp.padTo32Bytes([]byte(nonce))
	fmt.Println("nonce:", hex.EncodeToString(nonceInBytes[:]))
	return nonceInBytes, nil
}

func (mtp *MetaTransactionProcessor) calculateDigest(domainSeparator []byte, 
	messageTypeHash []byte, 
	functionName []byte, 
	paddedUserNonce string,
	userAddress string) []byte {

	functionSignaure := crypto.Keccak256(functionName)
	firstLayer := hex.EncodeToString(messageTypeHash[:]) + paddedUserNonce + userAddress + hex.EncodeToString(functionSignaure[:])
	bytes, err := hex.DecodeString(firstLayer)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
	}
	firstLayerDigest := crypto.Keccak256(bytes)
	fmt.Println("Calculate Digest ---  Type Hash:", hex.EncodeToString(messageTypeHash[:]))
	fmt.Println("Calculate Digest ---  User Nonce:", paddedUserNonce)
	fmt.Println("Calculate Digest ---  User Address:", userAddress)
	fmt.Println("Calculate Digest ---  Function Signature:", hex.EncodeToString(functionSignaure[:]))
	fmt.Println("Calculate Digest ---  First Layer ABI:", firstLayer)
	fmt.Println("Calculate Digest ---  First Layer Keccak:", hex.EncodeToString(firstLayerDigest[:]))

	prefix := []byte{0x19, 0x01}
    encodedPackedData := append(prefix, domainSeparator[:]...)
    encodedPackedData = append(encodedPackedData, firstLayerDigest...)
	fmt.Println("Calculate Digest ---  Second Layer", hex.EncodeToString(encodedPackedData[:]))

	digest := crypto.Keccak256(encodedPackedData)
	fmt.Println("Calculate Digest ---  Digest", hex.EncodeToString(digest[:]))

	return digest

} 

func (mtp *MetaTransactionProcessor) calculateDomainSeparator(EIP712_DOMAIN_TYPEHASH string, nameHash string, 
								versionHash string, paddedChainId string, contractAddress string) []byte{

	fmt.Println("Forwarder Address:", contractAddress)
	fmt.Println("Padded ChainId:", (paddedChainId))

	// Concatenate all bytes
	encoded := EIP712_DOMAIN_TYPEHASH + nameHash + versionHash + paddedChainId + contractAddress

	// Display the concatenated bytes
	fmt.Println("Concatenated Bytes:", encoded)

	// Get bytes from the hexadecimal string
	bytes, err := hex.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
	}

	// Now, you can hash the concatenated bytes with Keccak256
	domainSeparator := crypto.Keccak256(bytes) // Assuming you have a Keccak256 function defined elsewhere

	fmt.Printf("DOMAIN SEPARATOR: %x\n", domainSeparator)
	return domainSeparator
}

func (mtp *MetaTransactionProcessor) SignMetaTransaction(
	privateKey *ecdsa.PrivateKey, 
	userAddress common.Address, 
	functionSignature []byte, 
	chainID *big.Int, 
	forwarder common.Address, 
	nonce *big.Int,
	forwarderInstance *Forwarder.Forwarder,
	) ([]byte, error) {
    // Define EIP-712 types
    EIP712Domain := "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
    META_TRANSACTION_TYPE := "MetaTransaction(uint256 nonce,address from,bytes functionSignature)"
	name := "MyApp"
	version := "1"
	nameHash := crypto.Keccak256([]byte(name))
	versionHash := crypto.Keccak256([]byte(version))

    // Domain separator and type hashes
    domainSeparatorTypeHash := crypto.Keccak256([]byte(EIP712Domain))
    metaTransactionTypeHash := crypto.Keccak256([]byte(META_TRANSACTION_TYPE))


	paddedChainId := mtp.padTo32Bytes(chainID.Bytes())
	paddedAddress := mtp.padTo32Bytes(forwarder.Bytes())
	domainSeparator := mtp.calculateDomainSeparator(hex.EncodeToString(domainSeparatorTypeHash[:]), 
												hex.EncodeToString(nameHash[:]), 
												hex.EncodeToString(versionHash[:]), 
												hex.EncodeToString(paddedChainId[:]), 
												hex.EncodeToString(paddedAddress[:]))
	
	userNonce, err := mtp.getNonce(userAddress, forwarderInstance)
	if err != nil {
		log.Fatal(err)
	}
	paddedUserNonce := mtp.padTo32Bytes(userNonce.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	paddedUserAddress := mtp.padTo32Bytes(userAddress.Bytes())
												
	digest := mtp.calculateDigest(domainSeparator, metaTransactionTypeHash, functionSignature, 
								hex.EncodeToString(paddedUserNonce), hex.EncodeToString(paddedUserAddress[:]))


    // Sign the digest
    signature, err := crypto.Sign(digest, privateKey)
    if err != nil {
        return nil, err
    }

    return signature, nil
}


func (mtp *MetaTransactionProcessor) ProcessMetaTransaction(ctx context.Context, 
	networkConfig config.NetworkConfiguration, keyManager *key_manager.KeyManager, userId string, 
	confs []config.NetworkConfiguration, forwarderAddress common.Address, 
	nonce *big.Int, contractAddress common.Address, contractAbiStr string, 
	contractFunctionName string, args ...interface{}) (string, error) {
	
	// Get relayer's private key from environment variable
	privateKey, err := crypto.HexToECDSA(os.Getenv("RELAYER_PRIVATE_KEY"))
	if err != nil {
		return "", fmt.Errorf("failed to load relayer's private key: %w", err)
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


	// Load the Forwarder contract
	forwarderInstance, err := Forwarder.NewForwarder(forwarderAddress, mtp.ethClient)
	if err != nil {
		log.Fatal(err)
	}

	chainId := new(big.Int)

	// Convert the string to big.Int
	
	_, ok := chainId.SetString(networkConfig.ChainID, 10) // 10 is the base (decimal in this case)
	if !ok {
		fmt.Println("Error in conversion")
	}

	// Create an authorized transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		fmt.Println("Error Binding the auth, %w", err)
	}
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	gasPrice, err := mtp.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
    	log.Fatal(err)
	}
	auth.GasPrice = gasPrice

	// Parse the Target Contract ABI
	contractABI, err := abi.JSON(strings.NewReader(contractAbiStr))
	if err != nil {
		return "", fmt.Errorf("failed to parse target contract ABI: %w", err)
	}

	functionSignature, err := contractABI.Pack(contractFunctionName, args...)
	if err != nil {
		log.Fatal(err)
	}
	
	// Sign the meta transaction
	signature, err := mtp.SignMetaTransaction(rawUserPrivateKey, userAddress, 
		functionSignature, chainId, forwarderAddress, big.NewInt(0), forwarderInstance)
	if err != nil {
		log.Fatal(err)
	}

	// Extract r, s, v from the signature
	r := [32]byte{}
	s := [32]byte{}
	copy(r[:], signature[:32])
	copy(s[:], signature[32:64])
	v := uint8(signature[64]) + 27

	// Pack the data for the function we want to call in the target contract
	functionData, err := contractABI.Pack(contractFunctionName, args...)
    if err != nil {
        return "", fmt.Errorf("failed to pack data for target contract: %w", err)
    }

	tx, err := forwarderInstance.ExecuteMetaTransaction(auth, userAddress, contractAddress, functionData, r, s, v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Meta Transaction sent: %s\n", tx.Hash().Hex())

	return tx.Hash().Hex(), nil
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

package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/Jahankohan/mpc_wallet/Forwarder"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func padTo32Bytes(input []byte) []byte {
	padded := make([]byte, 32)
	copy(padded[32-len(input):], input)
	return padded
}

func getNonce(userAddress common.Address, forwarder *Forwarder.Forwarder) (*big.Int, error){
	nonce, err := forwarder.GetNonce(&bind.CallOpts{}, userAddress)
	if err != nil {
		fmt.Printf("Failed to call getNonce: %v\n", err)
		return big.NewInt(-1), err
	}
	return nonce, nil
}

func convertNonceToBytes(nonce string) ([]byte, error) {
	nonceInBytes := padTo32Bytes([]byte(nonce))
	fmt.Println("nonce:", hex.EncodeToString(nonceInBytes[:]))
	return nonceInBytes, nil
}

func calculateDigest(domainSeparator []byte, 
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

func calculateDomainSeparator(EIP712_DOMAIN_TYPEHASH string, nameHash string, 
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

func signMetaTransaction(
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


	paddedChainId := padTo32Bytes(chainID.Bytes())
	paddedAddress := padTo32Bytes(forwarder.Bytes())
	domainSeparator := calculateDomainSeparator(hex.EncodeToString(domainSeparatorTypeHash[:]), 
												hex.EncodeToString(nameHash[:]), 
												hex.EncodeToString(versionHash[:]), 
												hex.EncodeToString(paddedChainId[:]), 
												hex.EncodeToString(paddedAddress[:]))
	
	userNonce, err := getNonce(userAddress, forwarderInstance)
	if err != nil {
		log.Fatal(err)
	}
	paddedUserNonce := padTo32Bytes(userNonce.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	paddedUserAddress := padTo32Bytes(userAddress.Bytes())
												
	digest := calculateDigest(domainSeparator, metaTransactionTypeHash, functionSignature, 
								hex.EncodeToString(paddedUserNonce), hex.EncodeToString(paddedUserAddress[:]))


    // Sign the digest
    signature, err := crypto.Sign(digest, privateKey)
    if err != nil {
        return nil, err
    }

    return signature, nil
}


func relayMetaTx(
	clientURL string, 
	privateKeyHex, 
	userPrivateKey string,
	forwarderAddressHex string, 
	functionSignature []byte, 
	chainId *big.Int,
	) {
	// Setup Ethereum client
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		log.Fatal(err)
	}

	// Private key of the relayer
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	// Get user's address from the private key
	userPrivateKeyECDSA, err := crypto.HexToECDSA(userPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	userAddress := crypto.PubkeyToAddress(userPrivateKeyECDSA.PublicKey)

	// Address of the Forwarder contract
	forwarderAddress := common.HexToAddress(forwarderAddressHex)

	// Load the Forwarder contract
	forwarderInstance, err := Forwarder.NewForwarder(forwarderAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Create an authorized transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		fmt.Println("Error Binding the auth, %w", err)
	}
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
    	log.Fatal(err)
	}
	auth.GasPrice = gasPrice

	// Sign the meta transaction
	signature, err := signMetaTransaction(userPrivateKeyECDSA, userAddress, 
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

	// Print sigR, sigS and sigV in hex format
	fmt.Println("sigR:", hex.EncodeToString(r[:]))
	fmt.Println("sigS:", hex.EncodeToString(s[:]))
	fmt.Println("sigV:", v)

	// Print function signature in bytes
	fmt.Println("Function Signature (in bytes):", hex.EncodeToString(functionSignature))

	targetContract := common.HexToAddress("0xA99dCd104F08d01fD7A142cB7451c9E64Be8f6e3")

	tx, err := forwarderInstance.ExecuteMetaTransaction(auth, userAddress, targetContract, functionSignature, r, s, v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Meta Transaction sent: %s\n", tx.Hash().Hex())
}
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

func calculateDigest(domainSeparator []byte, 
	messageTypeHash []byte, 
	functionName []byte, 
	userNonce *big.Int) []byte {
	userAddress := "8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199"
	
	nonce := "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	functionSignaure := crypto.Keccak256(functionName)
	firstLayer := hex.EncodeToString(messageTypeHash[:]) + nonce + userAddress + hex.EncodeToString(functionSignaure[:])
	bytes, err := hex.DecodeString(firstLayer)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
	}
	firstLayerDigest := crypto.Keccak256(bytes)
	fmt.Println("Calculate Digest ---  Type Hash:", hex.EncodeToString(messageTypeHash[:]))
	fmt.Println("Calculate Digest ---  User Nonce:", nonce)
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

func calculateDomainSeparator() []byte{
	EIP712_DOMAIN_TYPEHASH := "8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f"
	nameHash := "1db874843672c48c2b3f8b32abc9aedf3c4132bc5c1094d62797110a8f762d54"
	versionHash := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
	chainID := "0000000000000000000000000000000000000000000000000000000000013881000000000000000000000000"
	contractAddress := "59EdA1Cde4E71168d119b8089281a88491d7b58b"

	chainIDNew := big.NewInt(80001)
	chainIDBytes := chainIDNew.Bytes()
	chainIDPadded := make([]byte, 32)
	copy(chainIDPadded[32-len(chainIDBytes):], chainIDBytes)
	fmt.Println("ChainID Padded:", hex.EncodeToString(chainIDPadded[:]))
	fmt.Println("ChainID Bytes:", hex.EncodeToString(chainIDBytes[:]))


	// Concatenate all bytes
	encoded := EIP712_DOMAIN_TYPEHASH + nameHash + versionHash + chainID + contractAddress

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

func signMetaTransaction(privateKey *ecdsa.PrivateKey, userAddress common.Address, functionSignature []byte, chainID *big.Int, forwarder common.Address, nonce *big.Int) ([]byte, error) {
    // Define EIP-712 types
    EIP712Domain := "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
    META_TRANSACTION_TYPE := "MetaTransaction(uint256 nonce,address from,bytes functionSignature)"
	// name := "MyApp"
	// version := "1"
	// nameHash := crypto.Keccak256([]byte(name))
	// versionHash := crypto.Keccak256([]byte(version))

    // Domain separator and type hashes
    domainSeparatorTypeHash := crypto.Keccak256([]byte(EIP712Domain))
    metaTransactionTypeHash := crypto.Keccak256([]byte(META_TRANSACTION_TYPE))

	fmt.Println("Domain Type Hash:", hex.EncodeToString(domainSeparatorTypeHash[:]))
	fmt.Println("meta type hash:", hex.EncodeToString(metaTransactionTypeHash[:]))

    // domainData := append(append(append([]byte("MyApp"), []byte("1")...), common.LeftPadBytes(chainID.Bytes(), 32)...), forwarder.Bytes()...)

    // domainSeparator := crypto.Keccak256(
    //     domainSeparatorTypeHash,
    //     crypto.Keccak256(domainData),
    // )

	chainIDBytes := make([]byte, 32)
	chainID.FillBytes(chainIDBytes)
	// paddedChainIDBytes := append(make([]byte, 32-len(chainIDBytes)), chainIDBytes...)

	// domainSeparators := crypto.Keccak256(
	// 	domainSeparatorTypeHash,
	// 	crypto.Keccak256([]byte("MyApp")),
	// 	crypto.Keccak256([]byte("1")),
	// 	chainIDBytes,
	// 	forwarder.Bytes(),
	// )
	domainSeparator := calculateDomainSeparator()
	
	digest := calculateDigest(domainSeparator, metaTransactionTypeHash, functionSignature, big.NewInt(0))


    // Sign the digest
    signature, err := crypto.Sign(digest, privateKey)
    if err != nil {
        return nil, err
    }

    return signature, nil
}


func verifySignature(
	userAddress common.Address,
	functionSignature []byte,
	chainID *big.Int,
	forwarder common.Address,
	nonce *big.Int,
	sigR, sigS [32]byte,
	sigV uint8,
) bool {

	// Define EIP-712 types
	EIP712Domain := "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
	META_TRANSACTION_TYPE := "MetaTransaction(uint256 nonce,address from,bytes functionSignature)"

	// Domain separator and type hashes
	domainSeparatorTypeHash := crypto.Keccak256([]byte(EIP712Domain))
	metaTransactionTypeHash := crypto.Keccak256([]byte(META_TRANSACTION_TYPE))

	domainData := []byte("MyApp")
	domainData = append(domainData, []byte("1")...)
	domainData = append(domainData, common.LeftPadBytes(chainID.Bytes(), 32)...) // Pad to 32 bytes
	domainData = append(domainData, forwarder.Bytes()...)

	domainSeparator := crypto.Keccak256(
		domainSeparatorTypeHash,
		crypto.Keccak256(domainData),
	)

	functionSignatureHash := crypto.Keccak256(functionSignature)
	nonceBytes := common.LeftPadBytes(nonce.Bytes(), 32) // Pad to 32 bytes

	typedDataHash := crypto.Keccak256(
		metaTransactionTypeHash,
		crypto.Keccak256(nonceBytes, userAddress.Bytes(), functionSignatureHash),
	)

	// The final digest that will be used for ecrecover
	digest := crypto.Keccak256([]byte("\x19\x01"), domainSeparator, typedDataHash)

	pubKey, err := crypto.SigToPub(digest, append(append(sigR[:], sigS[:]...), sigV-27))
	if err != nil {
		log.Fatal(err)
	}

	recoveredAddress := crypto.PubkeyToAddress(*pubKey)

	fmt.Println("User Address:", userAddress.Hex())
	fmt.Println("Recovered Address:", recoveredAddress.Hex())

	return recoveredAddress == userAddress
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
	signature, err := signMetaTransaction(userPrivateKeyECDSA, userAddress, functionSignature, chainId, forwarderAddress, big.NewInt(0))
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

	// res := verifySignature(userAddress, functionSignature, chainId, forwarderAddress, big.NewInt(0), r, s, v)
	// if res {
	// 	// Execute the meta transaction
	

	// } else {
	// 	fmt.Println("Signnature is not valid.")
	// }
	targetContract := common.HexToAddress("0xA99dCd104F08d01fD7A142cB7451c9E64Be8f6e3")

	tx, err := forwarderInstance.ExecuteMetaTransaction(auth, userAddress, targetContract, functionSignature, r, s, v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Meta Transaction sent: %s\n", tx.Hash().Hex())
}
package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

type EIP712Domain struct {
	Name              string
	Version           string
	ChainID           *big.Int
	VerifyingContract common.Address
}

func createMetaTransaction(userPrivateKey *ecdsa.PrivateKey, functionSignature []byte, forwarder string) ([]byte, common.Address) {
	// User's public key
	userPublicKey := userPrivateKey.Public()
	userPublicKeyECDSA, ok := userPublicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	userAddress := crypto.PubkeyToAddress(*userPublicKeyECDSA)

	domain := EIP712Domain{
		Name:              "MyApp",
		Version:           "1",
		ChainID:           big.NewInt(1), // Mainnet
		VerifyingContract: common.HexToAddress(forwarder),
	}

	// EIP-712 encoding
	typeHash := sha3.Sum256([]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))
	domainHash := sha3.Sum256([]byte(domain.Name + domain.Version + string(domain.ChainID.Bytes()) + string(domain.VerifyingContract.Bytes())))

	domainSeparator := sha3.Sum256(append(typeHash[:], domainHash[:]...))

	// Construct the EIP712 digest
	digest := sha3.NewLegacyKeccak256()
	digest.Write([]byte("\x19\x01"))
	digest.Write(domainSeparator[:])

	funcSigHashArray := sha3.Sum256(functionSignature)
	digest.Write(funcSigHashArray[:])
	fmt.Println("Digest:", digest)

	// Sign the EIP712 digest
	signature, err := crypto.Sign(digest.Sum(nil), userPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	return signature, userAddress
}

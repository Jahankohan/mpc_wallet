package key_manager

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/hashicorp/vault/shamir"

	d "github.com/Jahankohan/mpc_wallet/middleware"
	u "github.com/Jahankohan/mpc_wallet/utils"
)

type KeyManager struct{}

func (km *KeyManager) CreatePrivateKey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	return privateKey
}

func (km *KeyManager) SplitToShares(privateKey *ecdsa.PrivateKey, minimumShares int, totalShares int) ([][]byte, error) {
	if privateKey == nil {
		return nil, fmt.Errorf("private key cannot be nil")
	}
	if minimumShares < 1 || minimumShares > totalShares {
		return nil, fmt.Errorf("invalid share count: minimum %d, total %d", minimumShares, totalShares)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	
	shares, err := shamir.Split(privateKeyBytes, totalShares, minimumShares)
	if err != nil {
		return nil, fmt.Errorf("failed to split private key: %v", err)
	}

	return shares, nil
}


func createShareID(userID string, shareIndex int) [32]byte {
	shareIDString := fmt.Sprintf("%s_share%d", userID, shareIndex)
	var shareID [32]byte
	copy(shareID[:], []byte(shareIDString)[:32])
	return shareID
}



func (km *KeyManager) StoreSharesToTheBlockchain(userID string, shares []string) {
	// Store each share into KeyShareStorage contracts on different networks
	// Replace with actual contract addresses and private keys for each network

	for i, share := range shares {
		shareID := createShareID(userID, i+1)
		d.StoreShares(u.LoadConfig().Local, shareID, share)
		d.StoreShares(u.LoadConfig().POLYTestnet, shareID, share)
		d.StoreShares(u.LoadConfig().AVATestnet, shareID, share)
	}
}


func (km *KeyManager) ReconstructPrivateKey(shares [][]byte) (*ecdsa.PrivateKey, error) {
	if len(shares) < 2 {
		return nil, fmt.Errorf("at least two shares are required to reconstruct the private key")
	}

	privateKeyBytes, err := shamir.Combine(shares)
	if err != nil {
		return nil, fmt.Errorf("failed to reconstruct private key: %v", err)
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to convert byte slice to ECDSA private key: %v", err)
	}

	return privateKey, nil
}

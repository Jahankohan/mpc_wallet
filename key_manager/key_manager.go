package key_manager

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/tyler-smith/go-bip39"

	d "github.com/Jahankohan/mpc_wallet/deploy"
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

func (km *KeyManager) SplitToShares(privateKey *ecdsa.PrivateKey) []string {
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Printf("Generated private key: %s\n", hexutil.Encode(privateKeyBytes))

	// Generate 24-word BIP39 mnemonic for secret sharing
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Split mnemonic into 3 shares
	shares := make([]string, 3)
	words := strings.Split(mnemonic, " ")
	for i := 0; i < 3; i++ {
		start := i * 8
		end := start + 8
		shareWords := words[start:end]
		shares[i] = strings.Join(shareWords, " ")
		fmt.Printf("Share %d: %s\n", i+1, shares[i])
	}

	return shares
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

func (km *KeyManager) retrieveShares() []string {
	// Implement the logic to retrieve the shares from the KeyShareStorage contracts on each network
	return nil
}

func (km *KeyManager) reconstructThePrivateKey(shares []string) *ecdsa.PrivateKey {
	// Implement the logic to reconstruct the private key from the shares
	return nil
}
package key_manager

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/hashicorp/vault/shamir"

	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/Jahankohan/mpc_wallet/middleware"
)

type KeyManager struct{}

func (km *KeyManager) CreatePrivateKey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	return privateKey
}

func (km *KeyManager) GenerateAddress(privateKey *ecdsa.PrivateKey) common.Address{
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	return address
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

func (km *KeyManager) ConvertByteSliceToStringSlice(byteSlice [][]byte) []string {
    var stringSlice []string
    for _, byteArr := range byteSlice {
        stringSlice = append(stringSlice, string(byteArr))
    }
    return stringSlice
}


func createShareID(userID string, shareIndex int) [32]byte {
	shareIDString := fmt.Sprintf("%s_share%d", userID, shareIndex)
	var shareID [32]byte
	copy(shareID[:], []byte(shareIDString)[:32])
	return shareID
}

func stringToBytes32(input string) [32]byte {
	var output [32]byte
	copy(output[:], []byte(input)[:32])
	return output
}


func (km *KeyManager) StoreSharesToTheBlockchain(userID string, shares []string, networks []config.NetworkConfiguration) {
	for i, share := range shares {
		fmt.Println("Share to be stored:", share, ", on chain:", networks[i].Network)
		middleware.StoreShares(networks[i], stringToBytes32(userID), share)
	}
}


// RetrieveAllShares retrieves shares from different networks based on the configurations passed
func (km *KeyManager) RetrieveAllShares(confs []config.NetworkConfiguration, userId string) ([][]byte, error) {
	shares := make([][]byte, len(confs))
	
	for i, conf := range confs {
		share, err := middleware.RetrieveShares(conf, userId)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve share from network %s: %v", conf.Network, err)
		}
		shares[i] = []byte(share)
	}
	
	return shares, nil
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

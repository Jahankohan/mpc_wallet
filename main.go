package main

import (
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	// Parameters
	clientURL := "https://polygon-mumbai.g.alchemy.com/v2/3EBQaaPHYFJv7HXoFcHCkxgpitirS2H5"
	relayerPrivateKey := "8eb2e13f92e850fb487aa6ff5aa786818d440395115ba91baf34e33d6722ac24"
	// forwarderAddress := "0x5F2fB290798715afc2A8a2af221b883A243Da47D"
	forwarderAddress := "0x59EdA1Cde4E71168d119b8089281a88491d7b58b"

	userPrivateKey := "df57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e"
	
	abiDefinition := `[
		{
			"inputs": [],
			"name": "retrieve",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "uint256",
					"name": "num",
					"type": "uint256"
				}
			],
			"name": "store",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		}
	]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDefinition))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}
	// Pack the function signature and data for the SimpleStorage's store function
	number := big.NewInt(42) // The number to be stored
	packedData, err := parsedABI.Pack("store", number)
	if err != nil {
		log.Fatal(err)
	}
	
	chainId := big.NewInt(80001)
	// Relay the meta transaction
	relayMetaTx(clientURL, relayerPrivateKey, userPrivateKey, forwarderAddress, packedData, chainId)
}

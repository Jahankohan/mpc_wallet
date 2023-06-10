package middleware

import (
	"log"
	"strings"

	config "github.com/Jahankohan/mpc_wallet/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)


func stringToBytes32(input string) [32]byte {
	var output [32]byte
	copy(output[:], []byte(input)[:32])
	return output
}

func bytes32ToString(input [32]byte) string {
	return strings.TrimRight(string(input[:]), "\x00")
}

func StoreShares(conf config.NetworkConfiguration, userId [32]byte, share string) {
    client := Connect(conf)
    instance, err := LoadContract(client, conf.DeployedAddress)
    auth := GetOwnerAuth(conf)
    if err != nil {
        log.Panic()
    }
    instance.StoreShare(auth, userId, stringToBytes32(share))
}

func RetrieveShares(conf config.NetworkConfiguration, userId string) (string) {
    client := Connect(conf)
    instance, err := LoadContract(client, conf.DeployedAddress)
    if err != nil {
        log.Panic()
    }
    auth := GetOwnerAuth(conf)
    bind := bind.CallOpts{
        Pending: true,
        From:    auth.From,
        Context: auth.Context,
    }
    share, err := instance.GetShare(&bind, stringToBytes32(userId))
    return bytes32ToString(share)
}
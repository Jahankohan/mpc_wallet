package middleware

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	keyShare "github.com/Jahankohan/mpc_wallet/artifacts"
	config "github.com/Jahankohan/mpc_wallet/config"
	u "github.com/Jahankohan/mpc_wallet/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type NetworkConnections struct {
    network     config.NetworkConfiguration
    instance    keyShare.MpcWallet
}

func Connect(conf config.NetworkConfiguration) *ethclient.Client {
	
	client, err := ethclient.Dial(conf.Network)
	if err != nil {
		log.Fatal(err)
	}
	// Checking Connection
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil{
		log.Fatal(err)
	}
    _ = header
    fmt.Println("client created successfully!")
	return client
}

func GetOwnerAuth(conf config.NetworkConfiguration) *bind.TransactOpts {
	pK := conf.PrivateKey
    auth := GetAuthOwnerKeys(pK, conf)
    return auth
}

func DeployAllContracts(is_testnet bool) ([]NetworkConnections, error){
    
    var mainnet = []config.NetworkConfiguration {
        u.LoadConfig().ETHMainnet,
        u.LoadConfig().POLYMainnet,
        u.LoadConfig().AVAMainnet,
    }

    var testnet = []config.NetworkConfiguration {
        u.LoadConfig().Local,
        u.LoadConfig().AVATestnet,
        u.LoadConfig().POLYTestnet,
    }

    var deployedConfigurations = []NetworkConnections {}
	
	
	if (is_testnet) {
        for _, net := range testnet {
            client := Connect(net)
            instance, err := DeployContract(client, net)
            deployedConfigurations = append(deployedConfigurations, NetworkConnections{net, *instance})
            if err != nil {
                log.Fatal(err)
                return deployedConfigurations, err
            }
            client.Close()
        }
    } else {
        for _, net := range mainnet {
            client := Connect(net)
            instance, err := DeployContract(client, net)
            deployedConfigurations = append(deployedConfigurations, NetworkConnections{net, *instance})
            if err != nil {
                log.Fatal(err)
                return deployedConfigurations, err
            }

            client.Close()
        }
    }

    return deployedConfigurations, nil
}

func DeployContract(client *ethclient.Client, conf config.NetworkConfiguration) (*keyShare.MpcWallet, error){
	// This Method Deploy Contract 
    fmt.Println("Preparing to Deploy on: ", conf.Network)

    auth := GetOwnerAuth(conf)
    fmt.Println("Auth created Successfully!!!")
    
    address, tx, instance, err := keyShare.DeployMpcWallet(auth, client)
    if err != nil {
        log.Fatal(err)
        return instance, err
    }


    fmt.Println("Contract Deployed to: ", address.Hex(), ", On Network: ", conf.Network)
    fmt.Println("TX Hash: ", tx.Hash().Hex())
	os.Setenv("CONTRACT_ADDRESS", address.Hex())
    return instance, err
}

func LoadContract(client *ethclient.Client, _address string) (*keyShare.MpcWallet, error) {
    address := common.HexToAddress(_address)
    instance, err := keyShare.NewMpcWallet(address, client)
    if err != nil {
        log.Fatal(err)
    }
    return instance, nil
}

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

func GetAuthOwnerKeys(pK string, conf config.NetworkConfiguration) *bind.TransactOpts{
    client := Connect(conf)
    
    privateKey, err := crypto.HexToECDSA(pK)
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    chainID, err := client.ChainID(context.Background())
    if err != nil {
        panic(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    if err != nil {
        panic(err)
    }
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(3000000) // in units
    auth.GasPrice = gasPrice
    return auth;
}

func UpdateNonce(fromAddress common.Address, conf config.NetworkConfiguration) *big.Int{
    client := Connect(conf)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }
    return big.NewInt(int64(nonce));
}


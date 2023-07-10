// models/user_wallet.go
package models

import (
	"context"
	"log"
	"math/big"

	"github.com/Jahankohan/mpc_wallet/balance"
	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/Jahankohan/mpc_wallet/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type UserWallet struct {
	gorm.Model
	UserID        	string 	`gorm:"unique;not null"`
	WalletAddress 	string 	`gorm:"unique;not null"`
	EthereumBalance float64	`gorm:"-"`
	MaticBalance   	float64	`gorm:"-"`
	AvaxBalance    	float64	`gorm:"-"`
}

func createNetworkClients(configuration config.Configurations, isTestnet bool, chainName string) (*balance.BalanceFetcher){
	networkType := "testnet"
	if !isTestnet {
		networkType = "mainnet"
	}
	networkConfig, err := utils.GetSpecificNetworkConfiguration(configuration, networkType, chainName)
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(networkConfig.Network)
	if err != nil {
		log.Fatal(err)
	}
	balanceChecker := balance.NewBalanceFetcher(client)
	return balanceChecker
}

func CreateUserWallet(userWallet *UserWallet) error {
	result := DB.Create(userWallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserWalletByID(id uint) (*UserWallet, error) {
	var userWallet UserWallet
	result := DB.First(&userWallet, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userWallet, nil
}

func ConvertToFloat64(value *big.Float) float64 {
	f, _ := value.Float64()
	return f
}

func GetBalancesForAllChains(walletAddress string) (float64, float64, float64){
	configuration := utils.LoadConfig()
	ethBalanceChecker := createNetworkClients(configuration, true, "sepoila")
	mumbaiBalanceChecker := createNetworkClients(configuration, true, "mumbai")
	fujiBalanceChecker := createNetworkClients(configuration, true, "fuji")

	EthereumBalance, _ := ethBalanceChecker.GetNativeTokenBalance(context.Background(), common.HexToAddress(walletAddress))
	MaticBalance, _ := mumbaiBalanceChecker.GetNativeTokenBalance(context.Background(), common.HexToAddress(walletAddress))
	AvaxBalance, _ := fujiBalanceChecker.GetNativeTokenBalance(context.Background(), common.HexToAddress(walletAddress))

	return ConvertToFloat64(EthereumBalance), ConvertToFloat64(MaticBalance), ConvertToFloat64(AvaxBalance)
}

func GetUserWalletByUserID(userID string) (*UserWallet, error) {
	var userWallet UserWallet
	result := DB.Where("user_id = ?", userID).First(&userWallet)
	if result.Error != nil {
		return nil, result.Error
	}
	
	userWallet.EthereumBalance,	userWallet.MaticBalance, userWallet.AvaxBalance = GetBalancesForAllChains(userWallet.WalletAddress)

	return &userWallet, nil
}

func GetAllUserWallets() ([]UserWallet, error) {
	var userWallets []UserWallet
	result := DB.Find(&userWallets)
	if result.Error != nil {
		return nil, result.Error
	}

	for i := range userWallets {
		userWallets[i].EthereumBalance, userWallets[i].MaticBalance, userWallets[i].AvaxBalance = GetBalancesForAllChains(userWallets[i].WalletAddress)
	}

	return userWallets, nil
}

func UpdateUserWallet(userWallet *UserWallet) error {
	result := DB.Save(userWallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUserWallet(userWallet *UserWallet) error {
	result := DB.Delete(userWallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

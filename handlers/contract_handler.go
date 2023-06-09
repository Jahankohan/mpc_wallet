package handlers

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/Jahankohan/mpc_wallet/balance"
	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/Jahankohan/mpc_wallet/models"
	"github.com/Jahankohan/mpc_wallet/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type ContractHandler struct{
	configuration config.Configurations
}

type ContractInfo struct {
	ContractAddress string `json:"contract_address"`
	Network         string `json:"network"`
	ContractBalance string `json:"contract_balance"`
}


func NewContractHandler(configuration config.Configurations) *ContractHandler {
	return &ContractHandler{
		configuration:	configuration,
	}
}

func (h *ContractHandler) CreateContract(c *gin.Context) {
	var contract models.Contract
	if err := c.ShouldBindJSON(&contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := contract.CreateContract(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, contract)
}

func (h *ContractHandler) GetContractByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	contract, err := models.GetContractByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contract)
}

func (h *ContractHandler) GetContractByAddress(c *gin.Context) {
	address := c.Param("address")

	contract, err := models.GetContractByAddress(address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contract)
}

func (h *ContractHandler) GetAllContracts(c *gin.Context) {
	contracts, err := models.GetAllContracts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func (h *ContractHandler) UpdateContract(c *gin.Context) {
	var contract models.Contract
	if err := c.ShouldBindJSON(&contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := contract.UpdateContract()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contract)
}

func (h *ContractHandler) DeleteContract(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	contract, err := models.GetContractByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := contract.DeleteContract(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contract deleted successfully"})
}


func (ch *ContractHandler) GetContractEndpoints(c *gin.Context) {
	// Get the contract ID from the URL parameter
	contractID := c.Param("id")

	id, err := strconv.ParseUint(contractID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Retrieve the contract from the database
	contract, err := models.GetContractByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract not found"})
		return
	}

	// Parse the ABI JSON
	contractABI, err := abi.JSON(strings.NewReader(contract.ABI))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse contract ABI"})
		return
	}

	// Get the list of methods
	methods := make([]string, 0)
	for _, method := range contractABI.Methods {
		methods = append(methods, method.Name)
	}

	// Return the list of methods as the response
	c.JSON(http.StatusOK, gin.H{"endpoints": methods})
}

func (ch *ContractHandler) GetEndpointInputVariables(c *gin.Context) {
	// Get the contract ID and endpoint name from the URL parameters
	contractID := c.Param("id")
	endpoint := c.Param("endpoint")

	id, err := strconv.ParseUint(contractID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Retrieve the contract from the database
	contract, err := models.GetContractByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract not found"})
		return
	}

	// Parse the ABI JSON
	contractABI, err := abi.JSON(strings.NewReader(contract.ABI))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse contract ABI"})
		return
	}

	// Find the specified endpoint in the ABI
	var endpointABI abi.Method
	for _, method := range contractABI.Methods {
		if method.Name == endpoint {
			endpointABI = method
			break
		}
	}

	if endpointABI.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endpoint not found"})
		return
	}

	// Prepare the input variables and their types
	inputVariables := make(map[string]string)
	for _, input := range endpointABI.Inputs {
		inputVariables[input.Name] = input.Type.String()
	}

	// Return the input variables as the response
	c.JSON(http.StatusOK, gin.H{"inputVariables": inputVariables})
}


func (h *ContractHandler) GetContractKeyManagers(c *gin.Context) {
	var contracts []ContractInfo

	// Retrieve the contract information from the configuration
	testConfiguration := utils.GetNetworkConfigurations(h.configuration, true)
	for _, config := range testConfiguration {
		contractAddress := config.DeployedAddress
		contracts = append(contracts, ContractInfo{
			ContractAddress: contractAddress,
			Network:         config.Network,
		})
	}

	c.JSON(http.StatusOK, contracts)
}


func (h *ContractHandler) GetContractForwarders(c *gin.Context) {
    var contracts []ContractInfo

    // Retrieve the contract and balance information from the configuration
    testConfiguration := utils.GetNetworkConfigurationsWithName(h.configuration, true)
    for network, config := range testConfiguration {
        contractAddress := config.ForwarderAddress
        contractBalance := GetBalancesForAllChains(contractAddress, h.configuration, network)
        contracts = append(contracts, ContractInfo{
            ContractAddress: contractAddress,
            Network:         config.Network,
            ContractBalance: fmt.Sprintf("%f", contractBalance),
        })
    }

    c.JSON(http.StatusOK, contracts)
}


func createNetworkClients(configuration config.Configurations, isTestnet bool, chainName string) (*balance.BalanceFetcher){
	networkType := "mainnet"
	if isTestnet {
		networkType = "testnet"
	}
	networkConfig, err := utils.GetSpecificNetworkConfiguration(configuration, networkType, chainName)
	if err != nil {
		log.Fatal("Network Config Error:",err)
	}
	client, err := ethclient.Dial(networkConfig.Network)
	if err != nil {
		log.Fatal("Failed to create the client", err)
	}
	balanceChecker := balance.NewBalanceFetcher(client)
	return balanceChecker
}

func ConvertToFloat64(value *big.Float) float64 {
	f, _ := value.Float64()
	return f
}

func GetBalancesForAllChains(walletAddress string, configuration config.Configurations, networkName string) (float64){
	balanceChecker := createNetworkClients(configuration, true, networkName)

	balance, _ := balanceChecker.GetNativeTokenBalance(context.Background(), common.HexToAddress(walletAddress))


	return ConvertToFloat64(balance)
}
package utils

import (
	"fmt"
	"os"

	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() config.Configurations {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Read private key from environment variable
	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		fmt.Println("PRIVATE_KEY environment variable is not set")
	} else {
		// Update the private key in the configurations for each network
		for networkType, networkConfig := range configuration.Networks {
			for key, config := range networkConfig {
				config.PrivateKey = privateKey
				networkConfig[key] = config
			}
			configuration.Networks[networkType] = networkConfig
		}
	}

	return configuration
}

func GetNetworkConfigurations(configuration config.Configurations, isTestnet bool) []config.NetworkConfiguration {
	var networkConfigs []config.NetworkConfiguration

	for networkType, networkConfig := range configuration.Networks {
		if isTestnet && isTestnetNetwork(networkType) {
			for _, config := range networkConfig {
				networkConfigs = append(networkConfigs, config)
			}
		} else if !isTestnet && isMainnetNetwork(networkType) {
			for _, config := range networkConfig {
				networkConfigs = append(networkConfigs, config)
			}
		}
	}
	return networkConfigs
}

func GetSpecificNetworkConfiguration(configuration config.Configurations, networkType string, networkName string) (config.NetworkConfiguration, error) {
	networkConfig, ok := configuration.Networks[networkType]
	if !ok {
		return config.NetworkConfiguration{}, fmt.Errorf("network type %s not found in configuration", networkType)
	}

	specificConfig, ok := networkConfig[networkName]
	if !ok {
		return config.NetworkConfiguration{}, fmt.Errorf("network name %s not found in configuration for network type %s", networkName, networkType)
	}

	return specificConfig, nil
}


func GetNetworkConfigurationsWithName(configuration config.Configurations, isTestnet bool) map[string]config.NetworkConfiguration {
    networkConfigs := make(map[string]config.NetworkConfiguration)

    for networkType, networkConfig := range configuration.Networks {
        if isTestnet && isTestnetNetwork(networkType) {
            for networkName, config := range networkConfig {
                networkConfigs[networkName] = config
            }
        } else if !isTestnet && isMainnetNetwork(networkType) {
            for networkName, config := range networkConfig {
                networkConfigs[networkName] = config
            }
        }
    }
    return networkConfigs
}


func isTestnetNetwork(networkType string) bool {
	// You can extend this list with other testnet identifiers
	return networkType == "testnet"
}

func isMainnetNetwork(networkType string) bool {
	// You can extend this list with other mainnet identifiers
	return networkType == "mainnet"
}

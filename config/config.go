package config

import "math/big"

type Configurations struct {
	Networks map[string]map[string]NetworkConfiguration `mapstructure:"networks"`
}

type NetworkConfiguration struct {
	PrivateKey      string
	Network         string
	DeployedAddress string
	ChainId         *big.Int
}

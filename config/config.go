package config

import "math/big"

type Configurations struct {
	Local			NetworkConfiguration
	ETHTestnet		NetworkConfiguration
	ETHMainnet		NetworkConfiguration
	POLYTestnet		NetworkConfiguration
	POLYMainnet		NetworkConfiguration
	AVATestnet		NetworkConfiguration
	AVAMainnet		NetworkConfiguration
}

type NetworkConfiguration struct {
	PrivateKey		string
	Network			string
	DeployedAddress	string
	ChainId			*big.Int
}

package config

type Configurations struct {
	Networks map[string]map[string]NetworkConfiguration `mapstructure:"networks"`
	Database DatabaseConfig                             `mapstructure:"database"`
}

type NetworkConfiguration struct {
	PrivateKey      	string `mapstructure:"privatekey"`
	Network         	string `mapstructure:"network"`
	DeployedAddress 	string `mapstructure:"deployedAddress"`
	ForwarderAddress	string `mapstructure:"forwarderAddress"`
	ChainID         	string `mapstructure:"chainid"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}
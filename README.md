# MPC Custodial Wallet

MPC Custodial Wallet is an open-source custodial wallet service that utilizes Multi-Party Computation (MPC) to securely store and manage cryptocurrency private keys. The wallet service is designed for high flexibility and can interact with various blockchain networks, such as Ethereum, Polygon, and Avalanche.

## Features

- **Key Management**: Split a user's private key into multiple shares using Shamir's Secret Sharing algorithm and store each share on different blockchain networks.
- **Transaction Building and Signing**: Create, sign, and broadcast transactions to the supported blockchain networks.
- **OAuth Integration**: Allow users to register and log in using their social media accounts (e.g., Google).
- **Modular Architecture**: Comprising of modular components, the service is easy to extend and adapt to various blockchain networks.
- **Smart Contract Interaction**: Flexible interaction with smart contracts by parsing contract ABIs.

## Getting Started

### Prerequisites

- Golang (version 1.16 or higher)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- Ethereum and other supported blockchain networks' testnet/private keys and addresses.

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/Jahankohan/mpc_wallet.git
    ```

2. Change to the project directory:

    ```sh
    cd mpc_wallet
    ```

3. Install the required Go packages:

    ```sh
    go mod download
    ```

4. Set up the configuration in `config.yml`:

    ```yaml
    local:
      privatekey: "YOUR_ETHEREUM_PRIVATE_KEY"
      network: "http://localhost:8545"
      deployedAddress: "YOUR_SMART_CONTRACT_ADDRESS"
      chainid: 1337

    ethTestnet:
      privatekey: "YOUR_ETHEREUM_TESTNET_PRIVATE_KEY"
      network: "YOUR_ETHEREUM_TESTNET_RPC_URL"
      deployedAddress: "YOUR_SMART_CONTRACT_ADDRESS"
      chainid: "11155111"

    ethMainnet:
      privatekey: "YOUR_ETHEREUM_MAINNET_PRIVATE_KEY"
      network: "YOUR_ETHEREUM_MAINNET_RPC_URL"
      deployedAddress: "YOUR_SMART_CONTRACT_ADDRESS"
      chainid: "1"

    polyTestnet:
      privatekey: "YOUR_POLYGON_TESTNET_PRIVATE_KEY"
      network: "YOUR_POLYGON_TESTNET_RPC_URL"
      deployedAddress: "YOUR_SMART_CONTRACT_ADDRESS"
      chainid: "80001"

    polyMainnet:
      privatekey: "YOUR_POLYGON_MAINNET_PRIVATE_KEY"
      network: "YOUR_POLYGON_MAINNET_RPC_URL"
      deployedAddress: "YOUR_SMART_CONTRACT_ADDRESS"
      chainid: "137"

    avaTestnet:
      privatekey: "YOUR_AVALANCHE_TESTNET_PRIVATE_KEY"
      network: "YOUR_AVALANCHE_TESTNET_RPC_URL"
      deployedAddress: "YOUR_SMART_CONTRACT_ADDRESS"
      chainid: "43113"

    avaMainnet:
      privatekey: "YOUR_AVALANCHE_MAINNET_PRIVATE_KEY"
      network: "YOUR_AVALANCHE_MAINNET_RPC_URL"
      deployedAddress: "YOUR_SMART_CONTRACT_ADDRESS"
      chainid: "43114"
    ```


5. Run the server:

    ```sh
    go run main.go
    ```

Your server should now be running at `http://localhost:8080`.

## Usage

- Register/Login through OAuth.
- Store a private key by splitting it into shares and storing them on different blockchains.
- Build, sign, and broadcast transactions to interact with blockchain networks and smart contracts.
- Retrieve and reconstruct private keys through stored shares.

## Contributing

We welcome contributions from the community. Please submit your pull requests to this GitHub repository.

## License

This project is licensed under the MIT License.

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Go Ethereum](https://github.com/ethereum/go-ethereum)
- [Hashicorp Vault's Shamir's Secret Sharing Library](https://github.com/hashicorp/vault/tree/main/shamir)

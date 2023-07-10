package balance

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BalanceFetcher struct {
	client *ethclient.Client
}

func NewBalanceFetcher(client *ethclient.Client) *BalanceFetcher {
	return &BalanceFetcher{
		client: client,
	}
}

func (bf *BalanceFetcher) GetNativeTokenBalance(ctx context.Context, address common.Address) (*big.Float, error) {
	balance, err := bf.client.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch balance: %w", err)
	}
	ethValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))
	return ethValue, nil
}

func (bf *BalanceFetcher) GetERC20Balance(ctx context.Context, tokenAddress common.Address, tokenABIString string, 
	userAddress common.Address) (*big.Float, error) {
	tokenABI, err := abi.JSON(strings.NewReader(tokenABIString))
	if err != nil {
		return nil, fmt.Errorf("failed to parse token ABI: %w", err)
	}

	data, err := tokenABI.Pack("balanceOf", userAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to pack data for balanceOf function: %w", err)
	}

	result, err := bf.client.CallContract(ctx, ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch token balance: %w", err)
	}

	var balance big.Int
	err = tokenABI.UnpackIntoInterface(&balance, "balanceOf", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack token balance: %w", err)
	}

	ethValue := new(big.Float).Quo(new(big.Float).SetInt(&balance), big.NewFloat(math.Pow10(18)))
	return ethValue, nil
}

func (bf *BalanceFetcher) GetERC721Balance(ctx context.Context, tokenAddress common.Address, tokenABIString string, userAddress common.Address) (uint64, error) {
	tokenABI, err := abi.JSON(strings.NewReader(tokenABIString))
	if err != nil {
		return 0, fmt.Errorf("failed to parse token ABI: %w", err)
	}

	data, err := tokenABI.Pack("balanceOf", userAddress)
	if err != nil {
		return 0, fmt.Errorf("failed to pack data for balanceOf function: %w", err)
	}

	result, err := bf.client.CallContract(ctx, ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch token balance: %w", err)
	}

	var balance *big.Int
	err = tokenABI.UnpackIntoInterface(&balance, "balanceOf", result)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack token balance: %w", err)
	}

	return balance.Uint64(), nil
}

func (bf *BalanceFetcher) GetERC721TokenList(ctx context.Context, tokenAddress common.Address, tokenABIString string, userAddress common.Address) ([]*big.Int, error) {
	tokenABI, err := abi.JSON(strings.NewReader(tokenABIString))
	if err != nil {
		return nil, fmt.Errorf("failed to parse token ABI: %w", err)
	}

	// Get the total supply of tokens
	totalSupplyData, err := tokenABI.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("failed to pack data for totalSupply function: %w", err)
	}

	totalSupplyResult, err := bf.client.CallContract(ctx, ethereum.CallMsg{
		To:   &tokenAddress,
		Data: totalSupplyData,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch total supply of tokens: %w", err)
	}

	var totalSupply *big.Int
	err = tokenABI.UnpackIntoInterface(&totalSupply, "totalSupply", totalSupplyResult)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack total supply of tokens: %w", err)
	}

	tokenIDs := make([]*big.Int, 0)

	// Iterate through token IDs and check the owner
	for i := uint64(0); i < totalSupply.Uint64(); i++ {
		ownerData, err := tokenABI.Pack("ownerOf", big.NewInt(int64(i)))
		if err != nil {
			return nil, fmt.Errorf("failed to pack data for ownerOf function: %w", err)
		}

		ownerResult, err := bf.client.CallContract(ctx, ethereum.CallMsg{
			To:   &tokenAddress,
			Data: ownerData,
		}, nil)
		if err != nil {
			// Handle the case where the token does not exist
			if strings.Contains(err.Error(), "ERC721: owner query for nonexistent token") {
				continue
			}
			return nil, fmt.Errorf("failed to unpack owner of token: %w", err)
		}

		var owner common.Address
		err = tokenABI.UnpackIntoInterface(&owner, "ownerOf", ownerResult)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack owner of token %s: %w", i, err)
		}

		if owner == userAddress {
			tokenIDs = append(tokenIDs, big.NewInt(int64(i)))
			fmt.Println("TokenIDS:", tokenIDs)
		}
		owner = common.HexToAddress("")
	}

	return tokenIDs, nil
}



func (bf *BalanceFetcher) GetERC721TokenURI(ctx context.Context, tokenAddress common.Address, tokenABIString string,
	 tokenID *big.Int) (string, error) {
	tokenABI, err := abi.JSON(strings.NewReader(tokenABIString))
	if err != nil {
		return "", fmt.Errorf("failed to parse token ABI: %w", err)
	}

	data, err := tokenABI.Pack("tokenURI", tokenID)
	if err != nil {
		return "", fmt.Errorf("failed to pack data for tokenURI function: %w", err)
	}

	result, err := bf.client.CallContract(ctx, ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}, nil)
	if err != nil {
		return "", fmt.Errorf("failed to fetch token URI: %w", err)
	}

	var tokenURI string
	err = tokenABI.UnpackIntoInterface(&tokenURI, "tokenURI", result)
	if err != nil {
		return "", fmt.Errorf("failed to unpack token URI: %w", err)
	}

	return tokenURI, nil
}

func (bf *BalanceFetcher) GetERC1155Balance(ctx context.Context, tokenAddress common.Address, tokenABIString string,
	 tokenID *big.Int, userAddress common.Address) (*big.Int, error) {
	tokenABI, err := abi.JSON(strings.NewReader(tokenABIString))
	if err != nil {
		return nil, fmt.Errorf("failed to parse token ABI: %w", err)
	}

	data, err := tokenABI.Pack("balanceOf", userAddress, tokenID)
	if err != nil {
		return nil, fmt.Errorf("failed to pack data for balanceOf function: %w", err)
	}

	result, err := bf.client.CallContract(ctx, ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch token balance: %w", err)
	}

	var balance big.Int
	err = tokenABI.UnpackIntoInterface(&balance, "balanceOf", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack token balance: %w", err)
	}

	return &balance, nil
}

func (bf *BalanceFetcher) GetERC1155TokenCount(ctx context.Context, tokenAddress common.Address, tokenABIString string) (*big.Int, error) {
	tokenABI, err := abi.JSON(strings.NewReader(tokenABIString))
	if err != nil {
		return nil, fmt.Errorf("failed to parse token ABI: %w", err)
	}

	data, err := tokenABI.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("failed to pack data for totalSupply function: %w", err)
	}

	result, err := bf.client.CallContract(ctx, ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch token count: %w", err)
	}

	var count *big.Int
	err = tokenABI.UnpackIntoInterface(&count, "totalSupply", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack token count: %w", err)
	}

	return count, nil
}

func (bf *BalanceFetcher) GetERC1155TokenURI(ctx context.Context, tokenAddress common.Address, tokenABIString string,
	 tokenID *big.Int) (string, error) {
	tokenABI, err := abi.JSON(strings.NewReader(tokenABIString))
	if err != nil {
		return "", fmt.Errorf("failed to parse token ABI: %w", err)
	}

	data, err := tokenABI.Pack("uri", tokenID)
	if err != nil {
		return "", fmt.Errorf("failed to pack data for uri function: %w", err)
	}

	result, err := bf.client.CallContract(ctx, ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}, nil)
	if err != nil {
		return "", fmt.Errorf("failed to fetch token URI: %w", err)
	}

	var tokenURI string
	err = tokenABI.UnpackIntoInterface(&tokenURI, "uri", result)
	if err != nil {
		return "", fmt.Errorf("failed to unpack token URI: %w", err)
	}

	return tokenURI, nil
}


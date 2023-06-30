package transaction

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Jahankohan/mpc_wallet/config"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type MetaTransaction struct {
	From        common.Address
	To          common.Address
	FunctionSig []byte
	Nonce       *big.Int
	ChainId     *big.Int
	Params      []byte
}

func NewMetaTransaction(from common.Address, to common.Address, functionSig []byte, nonce *big.Int, chainId *big.Int, params []byte) *MetaTransaction {
	return &MetaTransaction{
		From:        from,
		To:          to,
		FunctionSig: functionSig,
		Nonce:       nonce,
		ChainId:     chainId,
		Params:      params,
	}
}

func (mt *MetaTransaction) Encode() []byte {
	encodedData, err := rlp.EncodeToBytes([]interface{}{
		mt.From,
		mt.To,
		mt.FunctionSig,
		mt.Nonce,
		mt.ChainId,
		mt.Params,
	})

	if err != nil {
		fmt.Printf("Error encoding meta-transaction: %s", err)
		return nil
	}

	return encodedData
}

func BuildMetaTransaction(networkConfig config.NetworkConfiguration, contractAddress common.Address, abiStr string, functionName string, nonce *big.Int, params ...interface{}) (*MetaTransaction, error) {
	parsedABI, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	method, exist := parsedABI.Methods[functionName]
	if !exist {
		return nil, fmt.Errorf("function %s does not exist in contract ABI", functionName)
	}

	inputs, err := method.Inputs.Pack(params...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack data for inputs: %w", err)
	}

	functionSignature := crypto.Keccak256([]byte(method.Sig))
	functionSignature = functionSignature[:4]

	// Build the meta-transaction
	chainId := big.NewInt(0)
	chainId.SetString(networkConfig.ChainId, 10)

	mt := NewMetaTransaction(
		common.HexToAddress(networkConfig.DeployedAddress),
		contractAddress,
		functionSignature,
		nonce,
		chainId,
		inputs,
	)

	return mt, nil
}

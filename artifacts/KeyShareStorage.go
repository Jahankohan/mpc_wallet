// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KeyShareStorage

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// KeyShareStorageMetaData contains all meta data concerning the KeyShareStorage contract.
var KeyShareStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"shareId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"ShareDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"shareId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"ShareStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"shareId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"ShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAuthorizedAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"shareId\",\"type\":\"bytes32\"}],\"name\":\"deleteShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"shareId\",\"type\":\"bytes32\"}],\"name\":\"getShare\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAuthorizedAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"shareId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"share\",\"type\":\"bytes\"}],\"name\":\"storeShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"shareId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"newShare\",\"type\":\"bytes\"}],\"name\":\"updateShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506113d8806100d86000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063776bb9631161005b578063776bb963146100eb5780637b4ea770146101075780637dbf117b146101375780638456cb591461015357610088565b80631ca5655f1461008d5780633f4ba83a146100a957806342f1181e146100b357806370712939146100cf575b600080fd5b6100a760048036038101906100a29190610d44565b61015d565b005b6100b1610304565b005b6100cd60048036038101906100c89190610dfe565b610440565b005b6100e960048036038101906100e49190610dfe565b610528565b005b61010560048036038101906101009190610e2b565b610611565b005b610121600480360381019061011c9190610e2b565b610753565b60405161012e9190610ed7565b60405180910390f35b610151600480360381019061014c9190610d44565b6108d4565b005b61015b610a19565b005b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166101e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e090610f56565b60405180910390fd5b600060149054906101000a900460ff1615610239576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023090610fc2565b60405180910390fd5b600060026000848152602001908152602001600020805461025990611011565b90501461029b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102929061108e565b60405180910390fd5b806002600084815260200190815260200160002090816102bb9190611264565b503373ffffffffffffffffffffffffffffffffffffffff16827f498bef06e06c917324187e12264cb64a2ccde5e394d2f5ea9562c685f447f0c860405160405180910390a35050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610392576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038990610f56565b60405180910390fd5b600060149054906101000a900460ff166103e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d890611382565b60405180910390fd5b60008060146101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa60405160405180910390a2565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104c590610f56565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ad90610f56565b60405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661069d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069490610f56565b60405180910390fd5b600060149054906101000a900460ff16156106ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e490610fc2565b60405180910390fd5b60026000828152602001908152602001600020600061070c9190610b57565b3373ffffffffffffffffffffffffffffffffffffffff16817f8e9e05770a48baa919b86724933a643079c39db4ffda1e456b3403a892e6f7d860405160405180910390a350565b6060600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166107e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d890610f56565b60405180910390fd5b600060149054906101000a900460ff1615610831576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082890610fc2565b60405180910390fd5b60026000838152602001908152602001600020805461084f90611011565b80601f016020809104026020016040519081016040528092919081815260200182805461087b90611011565b80156108c85780601f1061089d576101008083540402835291602001916108c8565b820191906000526020600020905b8154815290600101906020018083116108ab57829003601f168201915b50505050509050919050565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610960576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095790610f56565b60405180910390fd5b600060149054906101000a900460ff16156109b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a790610fc2565b60405180910390fd5b806002600084815260200190815260200160002090816109d09190611264565b503373ffffffffffffffffffffffffffffffffffffffff16827f6d7c51f2ccbdd28127dc175d6f77522fff8e9a7a38c3513b4c329bb19703677160405160405180910390a35050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610aa7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9e90610f56565b60405180910390fd5b600060149054906101000a900460ff1615610af7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aee90610fc2565b60405180910390fd5b6001600060146101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25860405160405180910390a2565b508054610b6390611011565b6000825580601f10610b755750610b94565b601f016020900490600052602060002090810190610b939190610b97565b5b50565b5b80821115610bb0576000816000905550600101610b98565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610bdb81610bc8565b8114610be657600080fd5b50565b600081359050610bf881610bd2565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610c5182610c08565b810181811067ffffffffffffffff82111715610c7057610c6f610c19565b5b80604052505050565b6000610c83610bb4565b9050610c8f8282610c48565b919050565b600067ffffffffffffffff821115610caf57610cae610c19565b5b610cb882610c08565b9050602081019050919050565b82818337600083830152505050565b6000610ce7610ce284610c94565b610c79565b905082815260208101848484011115610d0357610d02610c03565b5b610d0e848285610cc5565b509392505050565b600082601f830112610d2b57610d2a610bfe565b5b8135610d3b848260208601610cd4565b91505092915050565b60008060408385031215610d5b57610d5a610bbe565b5b6000610d6985828601610be9565b925050602083013567ffffffffffffffff811115610d8a57610d89610bc3565b5b610d9685828601610d16565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610dcb82610da0565b9050919050565b610ddb81610dc0565b8114610de657600080fd5b50565b600081359050610df881610dd2565b92915050565b600060208284031215610e1457610e13610bbe565b5b6000610e2284828501610de9565b91505092915050565b600060208284031215610e4157610e40610bbe565b5b6000610e4f84828501610be9565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610e92578082015181840152602081019050610e77565b60008484015250505050565b6000610ea982610e58565b610eb38185610e63565b9350610ec3818560208601610e74565b610ecc81610c08565b840191505092915050565b60006020820190508181036000830152610ef18184610e9e565b905092915050565b600082825260208201905092915050565b7f4e6f7420617574686f72697a6564000000000000000000000000000000000000600082015250565b6000610f40600e83610ef9565b9150610f4b82610f0a565b602082019050919050565b60006020820190508181036000830152610f6f81610f33565b9050919050565b7f436f6e7472616374206973207061757365640000000000000000000000000000600082015250565b6000610fac601283610ef9565b9150610fb782610f76565b602082019050919050565b60006020820190508181036000830152610fdb81610f9f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061102957607f821691505b60208210810361103c5761103b610fe2565b5b50919050565b7f536861726520494420616c726561647920657869737473000000000000000000600082015250565b6000611078601783610ef9565b915061108382611042565b602082019050919050565b600060208201905081810360008301526110a78161106b565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026111107fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826110d3565b61111a86836110d3565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600061116161115c61115784611132565b61113c565b611132565b9050919050565b6000819050919050565b61117b83611146565b61118f61118782611168565b8484546110e0565b825550505050565b600090565b6111a4611197565b6111af818484611172565b505050565b5b818110156111d3576111c860008261119c565b6001810190506111b5565b5050565b601f821115611218576111e9816110ae565b6111f2846110c3565b81016020851015611201578190505b61121561120d856110c3565b8301826111b4565b50505b505050565b600082821c905092915050565b600061123b6000198460080261121d565b1980831691505092915050565b6000611254838361122a565b9150826002028217905092915050565b61126d82610e58565b67ffffffffffffffff81111561128657611285610c19565b5b6112908254611011565b61129b8282856111d7565b600060209050601f8311600181146112ce57600084156112bc578287015190505b6112c68582611248565b86555061132e565b601f1984166112dc866110ae565b60005b82811015611304578489015182556001820191506020850194506020810190506112df565b86831015611321578489015161131d601f89168261122a565b8355505b6001600288020188555050505b505050505050565b7f436f6e7472616374206973206e6f742070617573656400000000000000000000600082015250565b600061136c601683610ef9565b915061137782611336565b602082019050919050565b6000602082019050818103600083015261139b8161135f565b905091905056fea264697066735822122067cb921b2e79dd7437f39352e944f97bbb075f5f99a66271fd46defb088f450f64736f6c63430008120033",
}

// KeyShareStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use KeyShareStorageMetaData.ABI instead.
var KeyShareStorageABI = KeyShareStorageMetaData.ABI

// KeyShareStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KeyShareStorageMetaData.Bin instead.
var KeyShareStorageBin = KeyShareStorageMetaData.Bin

// DeployKeyShareStorage deploys a new Ethereum contract, binding an instance of KeyShareStorage to it.
func DeployKeyShareStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KeyShareStorage, error) {
	parsed, err := KeyShareStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KeyShareStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KeyShareStorage{KeyShareStorageCaller: KeyShareStorageCaller{contract: contract}, KeyShareStorageTransactor: KeyShareStorageTransactor{contract: contract}, KeyShareStorageFilterer: KeyShareStorageFilterer{contract: contract}}, nil
}

// KeyShareStorage is an auto generated Go binding around an Ethereum contract.
type KeyShareStorage struct {
	KeyShareStorageCaller     // Read-only binding to the contract
	KeyShareStorageTransactor // Write-only binding to the contract
	KeyShareStorageFilterer   // Log filterer for contract events
}

// KeyShareStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyShareStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyShareStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyShareStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyShareStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyShareStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyShareStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyShareStorageSession struct {
	Contract     *KeyShareStorage  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyShareStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyShareStorageCallerSession struct {
	Contract *KeyShareStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// KeyShareStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyShareStorageTransactorSession struct {
	Contract     *KeyShareStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// KeyShareStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyShareStorageRaw struct {
	Contract *KeyShareStorage // Generic contract binding to access the raw methods on
}

// KeyShareStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyShareStorageCallerRaw struct {
	Contract *KeyShareStorageCaller // Generic read-only contract binding to access the raw methods on
}

// KeyShareStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyShareStorageTransactorRaw struct {
	Contract *KeyShareStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyShareStorage creates a new instance of KeyShareStorage, bound to a specific deployed contract.
func NewKeyShareStorage(address common.Address, backend bind.ContractBackend) (*KeyShareStorage, error) {
	contract, err := bindKeyShareStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyShareStorage{KeyShareStorageCaller: KeyShareStorageCaller{contract: contract}, KeyShareStorageTransactor: KeyShareStorageTransactor{contract: contract}, KeyShareStorageFilterer: KeyShareStorageFilterer{contract: contract}}, nil
}

// NewKeyShareStorageCaller creates a new read-only instance of KeyShareStorage, bound to a specific deployed contract.
func NewKeyShareStorageCaller(address common.Address, caller bind.ContractCaller) (*KeyShareStorageCaller, error) {
	contract, err := bindKeyShareStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyShareStorageCaller{contract: contract}, nil
}

// NewKeyShareStorageTransactor creates a new write-only instance of KeyShareStorage, bound to a specific deployed contract.
func NewKeyShareStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyShareStorageTransactor, error) {
	contract, err := bindKeyShareStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyShareStorageTransactor{contract: contract}, nil
}

// NewKeyShareStorageFilterer creates a new log filterer instance of KeyShareStorage, bound to a specific deployed contract.
func NewKeyShareStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyShareStorageFilterer, error) {
	contract, err := bindKeyShareStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyShareStorageFilterer{contract: contract}, nil
}

// bindKeyShareStorage binds a generic wrapper to an already deployed contract.
func bindKeyShareStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyShareStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyShareStorage *KeyShareStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyShareStorage.Contract.KeyShareStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyShareStorage *KeyShareStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyShareStorage.Contract.KeyShareStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyShareStorage *KeyShareStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyShareStorage.Contract.KeyShareStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyShareStorage *KeyShareStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyShareStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyShareStorage *KeyShareStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyShareStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyShareStorage *KeyShareStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyShareStorage.Contract.contract.Transact(opts, method, params...)
}

// GetShare is a free data retrieval call binding the contract method 0x7b4ea770.
//
// Solidity: function getShare(bytes32 shareId) view returns(bytes)
func (_KeyShareStorage *KeyShareStorageCaller) GetShare(opts *bind.CallOpts, shareId [32]byte) ([]byte, error) {
	var out []interface{}
	err := _KeyShareStorage.contract.Call(opts, &out, "getShare", shareId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetShare is a free data retrieval call binding the contract method 0x7b4ea770.
//
// Solidity: function getShare(bytes32 shareId) view returns(bytes)
func (_KeyShareStorage *KeyShareStorageSession) GetShare(shareId [32]byte) ([]byte, error) {
	return _KeyShareStorage.Contract.GetShare(&_KeyShareStorage.CallOpts, shareId)
}

// GetShare is a free data retrieval call binding the contract method 0x7b4ea770.
//
// Solidity: function getShare(bytes32 shareId) view returns(bytes)
func (_KeyShareStorage *KeyShareStorageCallerSession) GetShare(shareId [32]byte) ([]byte, error) {
	return _KeyShareStorage.Contract.GetShare(&_KeyShareStorage.CallOpts, shareId)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address addr) returns()
func (_KeyShareStorage *KeyShareStorageTransactor) AddAuthorizedAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _KeyShareStorage.contract.Transact(opts, "addAuthorizedAddress", addr)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address addr) returns()
func (_KeyShareStorage *KeyShareStorageSession) AddAuthorizedAddress(addr common.Address) (*types.Transaction, error) {
	return _KeyShareStorage.Contract.AddAuthorizedAddress(&_KeyShareStorage.TransactOpts, addr)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address addr) returns()
func (_KeyShareStorage *KeyShareStorageTransactorSession) AddAuthorizedAddress(addr common.Address) (*types.Transaction, error) {
	return _KeyShareStorage.Contract.AddAuthorizedAddress(&_KeyShareStorage.TransactOpts, addr)
}

// DeleteShare is a paid mutator transaction binding the contract method 0x776bb963.
//
// Solidity: function deleteShare(bytes32 shareId) returns()
func (_KeyShareStorage *KeyShareStorageTransactor) DeleteShare(opts *bind.TransactOpts, shareId [32]byte) (*types.Transaction, error) {
	return _KeyShareStorage.contract.Transact(opts, "deleteShare", shareId)
}

// DeleteShare is a paid mutator transaction binding the contract method 0x776bb963.
//
// Solidity: function deleteShare(bytes32 shareId) returns()
func (_KeyShareStorage *KeyShareStorageSession) DeleteShare(shareId [32]byte) (*types.Transaction, error) {
	return _KeyShareStorage.Contract.DeleteShare(&_KeyShareStorage.TransactOpts, shareId)
}

// DeleteShare is a paid mutator transaction binding the contract method 0x776bb963.
//
// Solidity: function deleteShare(bytes32 shareId) returns()
func (_KeyShareStorage *KeyShareStorageTransactorSession) DeleteShare(shareId [32]byte) (*types.Transaction, error) {
	return _KeyShareStorage.Contract.DeleteShare(&_KeyShareStorage.TransactOpts, shareId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KeyShareStorage *KeyShareStorageTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyShareStorage.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KeyShareStorage *KeyShareStorageSession) Pause() (*types.Transaction, error) {
	return _KeyShareStorage.Contract.Pause(&_KeyShareStorage.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KeyShareStorage *KeyShareStorageTransactorSession) Pause() (*types.Transaction, error) {
	return _KeyShareStorage.Contract.Pause(&_KeyShareStorage.TransactOpts)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address addr) returns()
func (_KeyShareStorage *KeyShareStorageTransactor) RemoveAuthorizedAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _KeyShareStorage.contract.Transact(opts, "removeAuthorizedAddress", addr)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address addr) returns()
func (_KeyShareStorage *KeyShareStorageSession) RemoveAuthorizedAddress(addr common.Address) (*types.Transaction, error) {
	return _KeyShareStorage.Contract.RemoveAuthorizedAddress(&_KeyShareStorage.TransactOpts, addr)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address addr) returns()
func (_KeyShareStorage *KeyShareStorageTransactorSession) RemoveAuthorizedAddress(addr common.Address) (*types.Transaction, error) {
	return _KeyShareStorage.Contract.RemoveAuthorizedAddress(&_KeyShareStorage.TransactOpts, addr)
}

// StoreShare is a paid mutator transaction binding the contract method 0x1ca5655f.
//
// Solidity: function storeShare(bytes32 shareId, bytes share) returns()
func (_KeyShareStorage *KeyShareStorageTransactor) StoreShare(opts *bind.TransactOpts, shareId [32]byte, share []byte) (*types.Transaction, error) {
	return _KeyShareStorage.contract.Transact(opts, "storeShare", shareId, share)
}

// StoreShare is a paid mutator transaction binding the contract method 0x1ca5655f.
//
// Solidity: function storeShare(bytes32 shareId, bytes share) returns()
func (_KeyShareStorage *KeyShareStorageSession) StoreShare(shareId [32]byte, share []byte) (*types.Transaction, error) {
	return _KeyShareStorage.Contract.StoreShare(&_KeyShareStorage.TransactOpts, shareId, share)
}

// StoreShare is a paid mutator transaction binding the contract method 0x1ca5655f.
//
// Solidity: function storeShare(bytes32 shareId, bytes share) returns()
func (_KeyShareStorage *KeyShareStorageTransactorSession) StoreShare(shareId [32]byte, share []byte) (*types.Transaction, error) {
	return _KeyShareStorage.Contract.StoreShare(&_KeyShareStorage.TransactOpts, shareId, share)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KeyShareStorage *KeyShareStorageTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyShareStorage.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KeyShareStorage *KeyShareStorageSession) Unpause() (*types.Transaction, error) {
	return _KeyShareStorage.Contract.Unpause(&_KeyShareStorage.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KeyShareStorage *KeyShareStorageTransactorSession) Unpause() (*types.Transaction, error) {
	return _KeyShareStorage.Contract.Unpause(&_KeyShareStorage.TransactOpts)
}

// UpdateShare is a paid mutator transaction binding the contract method 0x7dbf117b.
//
// Solidity: function updateShare(bytes32 shareId, bytes newShare) returns()
func (_KeyShareStorage *KeyShareStorageTransactor) UpdateShare(opts *bind.TransactOpts, shareId [32]byte, newShare []byte) (*types.Transaction, error) {
	return _KeyShareStorage.contract.Transact(opts, "updateShare", shareId, newShare)
}

// UpdateShare is a paid mutator transaction binding the contract method 0x7dbf117b.
//
// Solidity: function updateShare(bytes32 shareId, bytes newShare) returns()
func (_KeyShareStorage *KeyShareStorageSession) UpdateShare(shareId [32]byte, newShare []byte) (*types.Transaction, error) {
	return _KeyShareStorage.Contract.UpdateShare(&_KeyShareStorage.TransactOpts, shareId, newShare)
}

// UpdateShare is a paid mutator transaction binding the contract method 0x7dbf117b.
//
// Solidity: function updateShare(bytes32 shareId, bytes newShare) returns()
func (_KeyShareStorage *KeyShareStorageTransactorSession) UpdateShare(shareId [32]byte, newShare []byte) (*types.Transaction, error) {
	return _KeyShareStorage.Contract.UpdateShare(&_KeyShareStorage.TransactOpts, shareId, newShare)
}

// KeyShareStoragePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the KeyShareStorage contract.
type KeyShareStoragePausedIterator struct {
	Event *KeyShareStoragePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyShareStoragePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyShareStoragePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyShareStoragePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyShareStoragePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyShareStoragePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyShareStoragePaused represents a Paused event raised by the KeyShareStorage contract.
type KeyShareStoragePaused struct {
	By  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) FilterPaused(opts *bind.FilterOpts, by []common.Address) (*KeyShareStoragePausedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _KeyShareStorage.contract.FilterLogs(opts, "Paused", byRule)
	if err != nil {
		return nil, err
	}
	return &KeyShareStoragePausedIterator{contract: _KeyShareStorage.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *KeyShareStoragePaused, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _KeyShareStorage.contract.WatchLogs(opts, "Paused", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyShareStoragePaused)
				if err := _KeyShareStorage.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) ParsePaused(log types.Log) (*KeyShareStoragePaused, error) {
	event := new(KeyShareStoragePaused)
	if err := _KeyShareStorage.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyShareStorageShareDeletedIterator is returned from FilterShareDeleted and is used to iterate over the raw logs and unpacked data for ShareDeleted events raised by the KeyShareStorage contract.
type KeyShareStorageShareDeletedIterator struct {
	Event *KeyShareStorageShareDeleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyShareStorageShareDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyShareStorageShareDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyShareStorageShareDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyShareStorageShareDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyShareStorageShareDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyShareStorageShareDeleted represents a ShareDeleted event raised by the KeyShareStorage contract.
type KeyShareStorageShareDeleted struct {
	ShareId [32]byte
	By      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterShareDeleted is a free log retrieval operation binding the contract event 0x8e9e05770a48baa919b86724933a643079c39db4ffda1e456b3403a892e6f7d8.
//
// Solidity: event ShareDeleted(bytes32 indexed shareId, address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) FilterShareDeleted(opts *bind.FilterOpts, shareId [][32]byte, by []common.Address) (*KeyShareStorageShareDeletedIterator, error) {

	var shareIdRule []interface{}
	for _, shareIdItem := range shareId {
		shareIdRule = append(shareIdRule, shareIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _KeyShareStorage.contract.FilterLogs(opts, "ShareDeleted", shareIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return &KeyShareStorageShareDeletedIterator{contract: _KeyShareStorage.contract, event: "ShareDeleted", logs: logs, sub: sub}, nil
}

// WatchShareDeleted is a free log subscription operation binding the contract event 0x8e9e05770a48baa919b86724933a643079c39db4ffda1e456b3403a892e6f7d8.
//
// Solidity: event ShareDeleted(bytes32 indexed shareId, address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) WatchShareDeleted(opts *bind.WatchOpts, sink chan<- *KeyShareStorageShareDeleted, shareId [][32]byte, by []common.Address) (event.Subscription, error) {

	var shareIdRule []interface{}
	for _, shareIdItem := range shareId {
		shareIdRule = append(shareIdRule, shareIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _KeyShareStorage.contract.WatchLogs(opts, "ShareDeleted", shareIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyShareStorageShareDeleted)
				if err := _KeyShareStorage.contract.UnpackLog(event, "ShareDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseShareDeleted is a log parse operation binding the contract event 0x8e9e05770a48baa919b86724933a643079c39db4ffda1e456b3403a892e6f7d8.
//
// Solidity: event ShareDeleted(bytes32 indexed shareId, address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) ParseShareDeleted(log types.Log) (*KeyShareStorageShareDeleted, error) {
	event := new(KeyShareStorageShareDeleted)
	if err := _KeyShareStorage.contract.UnpackLog(event, "ShareDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyShareStorageShareStoredIterator is returned from FilterShareStored and is used to iterate over the raw logs and unpacked data for ShareStored events raised by the KeyShareStorage contract.
type KeyShareStorageShareStoredIterator struct {
	Event *KeyShareStorageShareStored // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyShareStorageShareStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyShareStorageShareStored)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyShareStorageShareStored)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyShareStorageShareStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyShareStorageShareStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyShareStorageShareStored represents a ShareStored event raised by the KeyShareStorage contract.
type KeyShareStorageShareStored struct {
	ShareId [32]byte
	By      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterShareStored is a free log retrieval operation binding the contract event 0x498bef06e06c917324187e12264cb64a2ccde5e394d2f5ea9562c685f447f0c8.
//
// Solidity: event ShareStored(bytes32 indexed shareId, address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) FilterShareStored(opts *bind.FilterOpts, shareId [][32]byte, by []common.Address) (*KeyShareStorageShareStoredIterator, error) {

	var shareIdRule []interface{}
	for _, shareIdItem := range shareId {
		shareIdRule = append(shareIdRule, shareIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _KeyShareStorage.contract.FilterLogs(opts, "ShareStored", shareIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return &KeyShareStorageShareStoredIterator{contract: _KeyShareStorage.contract, event: "ShareStored", logs: logs, sub: sub}, nil
}

// WatchShareStored is a free log subscription operation binding the contract event 0x498bef06e06c917324187e12264cb64a2ccde5e394d2f5ea9562c685f447f0c8.
//
// Solidity: event ShareStored(bytes32 indexed shareId, address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) WatchShareStored(opts *bind.WatchOpts, sink chan<- *KeyShareStorageShareStored, shareId [][32]byte, by []common.Address) (event.Subscription, error) {

	var shareIdRule []interface{}
	for _, shareIdItem := range shareId {
		shareIdRule = append(shareIdRule, shareIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _KeyShareStorage.contract.WatchLogs(opts, "ShareStored", shareIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyShareStorageShareStored)
				if err := _KeyShareStorage.contract.UnpackLog(event, "ShareStored", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseShareStored is a log parse operation binding the contract event 0x498bef06e06c917324187e12264cb64a2ccde5e394d2f5ea9562c685f447f0c8.
//
// Solidity: event ShareStored(bytes32 indexed shareId, address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) ParseShareStored(log types.Log) (*KeyShareStorageShareStored, error) {
	event := new(KeyShareStorageShareStored)
	if err := _KeyShareStorage.contract.UnpackLog(event, "ShareStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyShareStorageShareUpdatedIterator is returned from FilterShareUpdated and is used to iterate over the raw logs and unpacked data for ShareUpdated events raised by the KeyShareStorage contract.
type KeyShareStorageShareUpdatedIterator struct {
	Event *KeyShareStorageShareUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyShareStorageShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyShareStorageShareUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyShareStorageShareUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyShareStorageShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyShareStorageShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyShareStorageShareUpdated represents a ShareUpdated event raised by the KeyShareStorage contract.
type KeyShareStorageShareUpdated struct {
	ShareId [32]byte
	By      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterShareUpdated is a free log retrieval operation binding the contract event 0x6d7c51f2ccbdd28127dc175d6f77522fff8e9a7a38c3513b4c329bb197036771.
//
// Solidity: event ShareUpdated(bytes32 indexed shareId, address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) FilterShareUpdated(opts *bind.FilterOpts, shareId [][32]byte, by []common.Address) (*KeyShareStorageShareUpdatedIterator, error) {

	var shareIdRule []interface{}
	for _, shareIdItem := range shareId {
		shareIdRule = append(shareIdRule, shareIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _KeyShareStorage.contract.FilterLogs(opts, "ShareUpdated", shareIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return &KeyShareStorageShareUpdatedIterator{contract: _KeyShareStorage.contract, event: "ShareUpdated", logs: logs, sub: sub}, nil
}

// WatchShareUpdated is a free log subscription operation binding the contract event 0x6d7c51f2ccbdd28127dc175d6f77522fff8e9a7a38c3513b4c329bb197036771.
//
// Solidity: event ShareUpdated(bytes32 indexed shareId, address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) WatchShareUpdated(opts *bind.WatchOpts, sink chan<- *KeyShareStorageShareUpdated, shareId [][32]byte, by []common.Address) (event.Subscription, error) {

	var shareIdRule []interface{}
	for _, shareIdItem := range shareId {
		shareIdRule = append(shareIdRule, shareIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _KeyShareStorage.contract.WatchLogs(opts, "ShareUpdated", shareIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyShareStorageShareUpdated)
				if err := _KeyShareStorage.contract.UnpackLog(event, "ShareUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseShareUpdated is a log parse operation binding the contract event 0x6d7c51f2ccbdd28127dc175d6f77522fff8e9a7a38c3513b4c329bb197036771.
//
// Solidity: event ShareUpdated(bytes32 indexed shareId, address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) ParseShareUpdated(log types.Log) (*KeyShareStorageShareUpdated, error) {
	event := new(KeyShareStorageShareUpdated)
	if err := _KeyShareStorage.contract.UnpackLog(event, "ShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyShareStorageUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the KeyShareStorage contract.
type KeyShareStorageUnpausedIterator struct {
	Event *KeyShareStorageUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KeyShareStorageUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyShareStorageUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KeyShareStorageUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KeyShareStorageUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyShareStorageUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyShareStorageUnpaused represents a Unpaused event raised by the KeyShareStorage contract.
type KeyShareStorageUnpaused struct {
	By  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) FilterUnpaused(opts *bind.FilterOpts, by []common.Address) (*KeyShareStorageUnpausedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _KeyShareStorage.contract.FilterLogs(opts, "Unpaused", byRule)
	if err != nil {
		return nil, err
	}
	return &KeyShareStorageUnpausedIterator{contract: _KeyShareStorage.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *KeyShareStorageUnpaused, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _KeyShareStorage.contract.WatchLogs(opts, "Unpaused", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyShareStorageUnpaused)
				if err := _KeyShareStorage.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed by)
func (_KeyShareStorage *KeyShareStorageFilterer) ParseUnpaused(log types.Log) (*KeyShareStorageUnpaused, error) {
	event := new(KeyShareStorageUnpaused)
	if err := _KeyShareStorage.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mpc_wallet

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

// MpcWalletMetaData contains all meta data concerning the MpcWallet contract.
var MpcWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"shareId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"ShareDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"shareId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"ShareStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"shareId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"ShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAuthorizedAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"shareId\",\"type\":\"bytes32\"}],\"name\":\"deleteShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"shareId\",\"type\":\"bytes32\"}],\"name\":\"getShare\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAuthorizedAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"shareId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"share\",\"type\":\"bytes32\"}],\"name\":\"storeShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"shareId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newShare\",\"type\":\"bytes32\"}],\"name\":\"updateShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550610dee806100d86000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063707129391161005b57806370712939146100eb578063776bb963146101075780637b4ea770146101235780638456cb591461015357610088565b8063098bf6eb1461008d5780630d5e4b39146100a95780633f4ba83a146100c557806342f1181e146100cf575b600080fd5b6100a760048036038101906100a29190610b27565b61015d565b005b6100c360048036038101906100be9190610b27565b6102f1565b005b6100cd61042d565b005b6100e960048036038101906100e49190610ad5565b610569565b005b61010560048036038101906101009190610ad5565b610651565b005b610121600480360381019061011c9190610afe565b61073a565b005b61013d60048036038101906101389190610afe565b610874565b60405161014a9190610bfe565b60405180910390f35b61015b61096d565b005b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166101e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e090610c79565b60405180910390fd5b600060149054906101000a900460ff1615610239576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023090610c39565b60405180910390fd5b6000801b600260008481526020019081526020016000205414610291576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028890610c59565b60405180910390fd5b8060026000848152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff16827f498bef06e06c917324187e12264cb64a2ccde5e394d2f5ea9562c685f447f0c860405160405180910390a35050565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661037d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037490610c79565b60405180910390fd5b600060149054906101000a900460ff16156103cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c490610c39565b60405180910390fd5b8060026000848152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff16827f6d7c51f2ccbdd28127dc175d6f77522fff8e9a7a38c3513b4c329bb19703677160405160405180910390a35050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b290610c79565b60405180910390fd5b600060149054906101000a900460ff1661050a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050190610c19565b60405180910390fd5b60008060146101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa60405160405180910390a2565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ee90610c79565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d690610c79565b60405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166107c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107bd90610c79565b60405180910390fd5b600060149054906101000a900460ff1615610816576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080d90610c39565b60405180910390fd5b60026000828152602001908152602001600020600090553373ffffffffffffffffffffffffffffffffffffffff16817f8e9e05770a48baa919b86724933a643079c39db4ffda1e456b3403a892e6f7d860405160405180910390a350565b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610902576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f990610c79565b60405180910390fd5b600060149054906101000a900460ff1615610952576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094990610c39565b60405180910390fd5b60026000838152602001908152602001600020549050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109f290610c79565b60405180910390fd5b600060149054906101000a900460ff1615610a4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4290610c39565b60405180910390fd5b6001600060146101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25860405160405180910390a2565b600081359050610aba81610d8a565b92915050565b600081359050610acf81610da1565b92915050565b600060208284031215610ae757600080fd5b6000610af584828501610aab565b91505092915050565b600060208284031215610b1057600080fd5b6000610b1e84828501610ac0565b91505092915050565b60008060408385031215610b3a57600080fd5b6000610b4885828601610ac0565b9250506020610b5985828601610ac0565b9150509250929050565b610b6c81610cbc565b82525050565b6000610b7f601683610c99565b9150610b8a82610ce6565b602082019050919050565b6000610ba2601283610c99565b9150610bad82610d0f565b602082019050919050565b6000610bc5601783610c99565b9150610bd082610d38565b602082019050919050565b6000610be8600e83610c99565b9150610bf382610d61565b602082019050919050565b6000602082019050610c136000830184610b63565b92915050565b60006020820190508181036000830152610c3281610b72565b9050919050565b60006020820190508181036000830152610c5281610b95565b9050919050565b60006020820190508181036000830152610c7281610bb8565b9050919050565b60006020820190508181036000830152610c9281610bdb565b9050919050565b600082825260208201905092915050565b6000610cb582610cc6565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b7f436f6e7472616374206973206e6f742070617573656400000000000000000000600082015250565b7f436f6e7472616374206973207061757365640000000000000000000000000000600082015250565b7f536861726520494420616c726561647920657869737473000000000000000000600082015250565b7f4e6f7420617574686f72697a6564000000000000000000000000000000000000600082015250565b610d9381610caa565b8114610d9e57600080fd5b50565b610daa81610cbc565b8114610db557600080fd5b5056fea26469706673582212206ecb8d19d46bc4c3a0f8442153a68621292d60dede1bf883b16fa649f4396c4264736f6c63430008040033",
}

// MpcWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use MpcWalletMetaData.ABI instead.
var MpcWalletABI = MpcWalletMetaData.ABI

// MpcWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MpcWalletMetaData.Bin instead.
var MpcWalletBin = MpcWalletMetaData.Bin

// DeployMpcWallet deploys a new Ethereum contract, binding an instance of MpcWallet to it.
func DeployMpcWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MpcWallet, error) {
	parsed, err := MpcWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MpcWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MpcWallet{MpcWalletCaller: MpcWalletCaller{contract: contract}, MpcWalletTransactor: MpcWalletTransactor{contract: contract}, MpcWalletFilterer: MpcWalletFilterer{contract: contract}}, nil
}

// MpcWallet is an auto generated Go binding around an Ethereum contract.
type MpcWallet struct {
	MpcWalletCaller     // Read-only binding to the contract
	MpcWalletTransactor // Write-only binding to the contract
	MpcWalletFilterer   // Log filterer for contract events
}

// MpcWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type MpcWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MpcWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MpcWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MpcWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MpcWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MpcWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MpcWalletSession struct {
	Contract     *MpcWallet        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MpcWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MpcWalletCallerSession struct {
	Contract *MpcWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MpcWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MpcWalletTransactorSession struct {
	Contract     *MpcWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MpcWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type MpcWalletRaw struct {
	Contract *MpcWallet // Generic contract binding to access the raw methods on
}

// MpcWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MpcWalletCallerRaw struct {
	Contract *MpcWalletCaller // Generic read-only contract binding to access the raw methods on
}

// MpcWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MpcWalletTransactorRaw struct {
	Contract *MpcWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMpcWallet creates a new instance of MpcWallet, bound to a specific deployed contract.
func NewMpcWallet(address common.Address, backend bind.ContractBackend) (*MpcWallet, error) {
	contract, err := bindMpcWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MpcWallet{MpcWalletCaller: MpcWalletCaller{contract: contract}, MpcWalletTransactor: MpcWalletTransactor{contract: contract}, MpcWalletFilterer: MpcWalletFilterer{contract: contract}}, nil
}

// NewMpcWalletCaller creates a new read-only instance of MpcWallet, bound to a specific deployed contract.
func NewMpcWalletCaller(address common.Address, caller bind.ContractCaller) (*MpcWalletCaller, error) {
	contract, err := bindMpcWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MpcWalletCaller{contract: contract}, nil
}

// NewMpcWalletTransactor creates a new write-only instance of MpcWallet, bound to a specific deployed contract.
func NewMpcWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*MpcWalletTransactor, error) {
	contract, err := bindMpcWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MpcWalletTransactor{contract: contract}, nil
}

// NewMpcWalletFilterer creates a new log filterer instance of MpcWallet, bound to a specific deployed contract.
func NewMpcWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*MpcWalletFilterer, error) {
	contract, err := bindMpcWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MpcWalletFilterer{contract: contract}, nil
}

// bindMpcWallet binds a generic wrapper to an already deployed contract.
func bindMpcWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MpcWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MpcWallet *MpcWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MpcWallet.Contract.MpcWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MpcWallet *MpcWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MpcWallet.Contract.MpcWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MpcWallet *MpcWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MpcWallet.Contract.MpcWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MpcWallet *MpcWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MpcWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MpcWallet *MpcWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MpcWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MpcWallet *MpcWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MpcWallet.Contract.contract.Transact(opts, method, params...)
}

// GetShare is a free data retrieval call binding the contract method 0x7b4ea770.
//
// Solidity: function getShare(bytes32 shareId) view returns(bytes32)
func (_MpcWallet *MpcWalletCaller) GetShare(opts *bind.CallOpts, shareId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MpcWallet.contract.Call(opts, &out, "getShare", shareId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetShare is a free data retrieval call binding the contract method 0x7b4ea770.
//
// Solidity: function getShare(bytes32 shareId) view returns(bytes32)
func (_MpcWallet *MpcWalletSession) GetShare(shareId [32]byte) ([32]byte, error) {
	return _MpcWallet.Contract.GetShare(&_MpcWallet.CallOpts, shareId)
}

// GetShare is a free data retrieval call binding the contract method 0x7b4ea770.
//
// Solidity: function getShare(bytes32 shareId) view returns(bytes32)
func (_MpcWallet *MpcWalletCallerSession) GetShare(shareId [32]byte) ([32]byte, error) {
	return _MpcWallet.Contract.GetShare(&_MpcWallet.CallOpts, shareId)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address addr) returns()
func (_MpcWallet *MpcWalletTransactor) AddAuthorizedAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _MpcWallet.contract.Transact(opts, "addAuthorizedAddress", addr)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address addr) returns()
func (_MpcWallet *MpcWalletSession) AddAuthorizedAddress(addr common.Address) (*types.Transaction, error) {
	return _MpcWallet.Contract.AddAuthorizedAddress(&_MpcWallet.TransactOpts, addr)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address addr) returns()
func (_MpcWallet *MpcWalletTransactorSession) AddAuthorizedAddress(addr common.Address) (*types.Transaction, error) {
	return _MpcWallet.Contract.AddAuthorizedAddress(&_MpcWallet.TransactOpts, addr)
}

// DeleteShare is a paid mutator transaction binding the contract method 0x776bb963.
//
// Solidity: function deleteShare(bytes32 shareId) returns()
func (_MpcWallet *MpcWalletTransactor) DeleteShare(opts *bind.TransactOpts, shareId [32]byte) (*types.Transaction, error) {
	return _MpcWallet.contract.Transact(opts, "deleteShare", shareId)
}

// DeleteShare is a paid mutator transaction binding the contract method 0x776bb963.
//
// Solidity: function deleteShare(bytes32 shareId) returns()
func (_MpcWallet *MpcWalletSession) DeleteShare(shareId [32]byte) (*types.Transaction, error) {
	return _MpcWallet.Contract.DeleteShare(&_MpcWallet.TransactOpts, shareId)
}

// DeleteShare is a paid mutator transaction binding the contract method 0x776bb963.
//
// Solidity: function deleteShare(bytes32 shareId) returns()
func (_MpcWallet *MpcWalletTransactorSession) DeleteShare(shareId [32]byte) (*types.Transaction, error) {
	return _MpcWallet.Contract.DeleteShare(&_MpcWallet.TransactOpts, shareId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MpcWallet *MpcWalletTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MpcWallet.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MpcWallet *MpcWalletSession) Pause() (*types.Transaction, error) {
	return _MpcWallet.Contract.Pause(&_MpcWallet.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MpcWallet *MpcWalletTransactorSession) Pause() (*types.Transaction, error) {
	return _MpcWallet.Contract.Pause(&_MpcWallet.TransactOpts)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address addr) returns()
func (_MpcWallet *MpcWalletTransactor) RemoveAuthorizedAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _MpcWallet.contract.Transact(opts, "removeAuthorizedAddress", addr)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address addr) returns()
func (_MpcWallet *MpcWalletSession) RemoveAuthorizedAddress(addr common.Address) (*types.Transaction, error) {
	return _MpcWallet.Contract.RemoveAuthorizedAddress(&_MpcWallet.TransactOpts, addr)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address addr) returns()
func (_MpcWallet *MpcWalletTransactorSession) RemoveAuthorizedAddress(addr common.Address) (*types.Transaction, error) {
	return _MpcWallet.Contract.RemoveAuthorizedAddress(&_MpcWallet.TransactOpts, addr)
}

// StoreShare is a paid mutator transaction binding the contract method 0x098bf6eb.
//
// Solidity: function storeShare(bytes32 shareId, bytes32 share) returns()
func (_MpcWallet *MpcWalletTransactor) StoreShare(opts *bind.TransactOpts, shareId [32]byte, share [32]byte) (*types.Transaction, error) {
	return _MpcWallet.contract.Transact(opts, "storeShare", shareId, share)
}

// StoreShare is a paid mutator transaction binding the contract method 0x098bf6eb.
//
// Solidity: function storeShare(bytes32 shareId, bytes32 share) returns()
func (_MpcWallet *MpcWalletSession) StoreShare(shareId [32]byte, share [32]byte) (*types.Transaction, error) {
	return _MpcWallet.Contract.StoreShare(&_MpcWallet.TransactOpts, shareId, share)
}

// StoreShare is a paid mutator transaction binding the contract method 0x098bf6eb.
//
// Solidity: function storeShare(bytes32 shareId, bytes32 share) returns()
func (_MpcWallet *MpcWalletTransactorSession) StoreShare(shareId [32]byte, share [32]byte) (*types.Transaction, error) {
	return _MpcWallet.Contract.StoreShare(&_MpcWallet.TransactOpts, shareId, share)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MpcWallet *MpcWalletTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MpcWallet.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MpcWallet *MpcWalletSession) Unpause() (*types.Transaction, error) {
	return _MpcWallet.Contract.Unpause(&_MpcWallet.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MpcWallet *MpcWalletTransactorSession) Unpause() (*types.Transaction, error) {
	return _MpcWallet.Contract.Unpause(&_MpcWallet.TransactOpts)
}

// UpdateShare is a paid mutator transaction binding the contract method 0x0d5e4b39.
//
// Solidity: function updateShare(bytes32 shareId, bytes32 newShare) returns()
func (_MpcWallet *MpcWalletTransactor) UpdateShare(opts *bind.TransactOpts, shareId [32]byte, newShare [32]byte) (*types.Transaction, error) {
	return _MpcWallet.contract.Transact(opts, "updateShare", shareId, newShare)
}

// UpdateShare is a paid mutator transaction binding the contract method 0x0d5e4b39.
//
// Solidity: function updateShare(bytes32 shareId, bytes32 newShare) returns()
func (_MpcWallet *MpcWalletSession) UpdateShare(shareId [32]byte, newShare [32]byte) (*types.Transaction, error) {
	return _MpcWallet.Contract.UpdateShare(&_MpcWallet.TransactOpts, shareId, newShare)
}

// UpdateShare is a paid mutator transaction binding the contract method 0x0d5e4b39.
//
// Solidity: function updateShare(bytes32 shareId, bytes32 newShare) returns()
func (_MpcWallet *MpcWalletTransactorSession) UpdateShare(shareId [32]byte, newShare [32]byte) (*types.Transaction, error) {
	return _MpcWallet.Contract.UpdateShare(&_MpcWallet.TransactOpts, shareId, newShare)
}

// MpcWalletPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MpcWallet contract.
type MpcWalletPausedIterator struct {
	Event *MpcWalletPaused // Event containing the contract specifics and raw log

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
func (it *MpcWalletPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MpcWalletPaused)
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
		it.Event = new(MpcWalletPaused)
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
func (it *MpcWalletPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MpcWalletPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MpcWalletPaused represents a Paused event raised by the MpcWallet contract.
type MpcWalletPaused struct {
	By  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed by)
func (_MpcWallet *MpcWalletFilterer) FilterPaused(opts *bind.FilterOpts, by []common.Address) (*MpcWalletPausedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MpcWallet.contract.FilterLogs(opts, "Paused", byRule)
	if err != nil {
		return nil, err
	}
	return &MpcWalletPausedIterator{contract: _MpcWallet.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed by)
func (_MpcWallet *MpcWalletFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MpcWalletPaused, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MpcWallet.contract.WatchLogs(opts, "Paused", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MpcWalletPaused)
				if err := _MpcWallet.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_MpcWallet *MpcWalletFilterer) ParsePaused(log types.Log) (*MpcWalletPaused, error) {
	event := new(MpcWalletPaused)
	if err := _MpcWallet.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MpcWalletShareDeletedIterator is returned from FilterShareDeleted and is used to iterate over the raw logs and unpacked data for ShareDeleted events raised by the MpcWallet contract.
type MpcWalletShareDeletedIterator struct {
	Event *MpcWalletShareDeleted // Event containing the contract specifics and raw log

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
func (it *MpcWalletShareDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MpcWalletShareDeleted)
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
		it.Event = new(MpcWalletShareDeleted)
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
func (it *MpcWalletShareDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MpcWalletShareDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MpcWalletShareDeleted represents a ShareDeleted event raised by the MpcWallet contract.
type MpcWalletShareDeleted struct {
	ShareId [32]byte
	By      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterShareDeleted is a free log retrieval operation binding the contract event 0x8e9e05770a48baa919b86724933a643079c39db4ffda1e456b3403a892e6f7d8.
//
// Solidity: event ShareDeleted(bytes32 indexed shareId, address indexed by)
func (_MpcWallet *MpcWalletFilterer) FilterShareDeleted(opts *bind.FilterOpts, shareId [][32]byte, by []common.Address) (*MpcWalletShareDeletedIterator, error) {

	var shareIdRule []interface{}
	for _, shareIdItem := range shareId {
		shareIdRule = append(shareIdRule, shareIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MpcWallet.contract.FilterLogs(opts, "ShareDeleted", shareIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return &MpcWalletShareDeletedIterator{contract: _MpcWallet.contract, event: "ShareDeleted", logs: logs, sub: sub}, nil
}

// WatchShareDeleted is a free log subscription operation binding the contract event 0x8e9e05770a48baa919b86724933a643079c39db4ffda1e456b3403a892e6f7d8.
//
// Solidity: event ShareDeleted(bytes32 indexed shareId, address indexed by)
func (_MpcWallet *MpcWalletFilterer) WatchShareDeleted(opts *bind.WatchOpts, sink chan<- *MpcWalletShareDeleted, shareId [][32]byte, by []common.Address) (event.Subscription, error) {

	var shareIdRule []interface{}
	for _, shareIdItem := range shareId {
		shareIdRule = append(shareIdRule, shareIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MpcWallet.contract.WatchLogs(opts, "ShareDeleted", shareIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MpcWalletShareDeleted)
				if err := _MpcWallet.contract.UnpackLog(event, "ShareDeleted", log); err != nil {
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
func (_MpcWallet *MpcWalletFilterer) ParseShareDeleted(log types.Log) (*MpcWalletShareDeleted, error) {
	event := new(MpcWalletShareDeleted)
	if err := _MpcWallet.contract.UnpackLog(event, "ShareDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MpcWalletShareStoredIterator is returned from FilterShareStored and is used to iterate over the raw logs and unpacked data for ShareStored events raised by the MpcWallet contract.
type MpcWalletShareStoredIterator struct {
	Event *MpcWalletShareStored // Event containing the contract specifics and raw log

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
func (it *MpcWalletShareStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MpcWalletShareStored)
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
		it.Event = new(MpcWalletShareStored)
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
func (it *MpcWalletShareStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MpcWalletShareStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MpcWalletShareStored represents a ShareStored event raised by the MpcWallet contract.
type MpcWalletShareStored struct {
	ShareId [32]byte
	By      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterShareStored is a free log retrieval operation binding the contract event 0x498bef06e06c917324187e12264cb64a2ccde5e394d2f5ea9562c685f447f0c8.
//
// Solidity: event ShareStored(bytes32 indexed shareId, address indexed by)
func (_MpcWallet *MpcWalletFilterer) FilterShareStored(opts *bind.FilterOpts, shareId [][32]byte, by []common.Address) (*MpcWalletShareStoredIterator, error) {

	var shareIdRule []interface{}
	for _, shareIdItem := range shareId {
		shareIdRule = append(shareIdRule, shareIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MpcWallet.contract.FilterLogs(opts, "ShareStored", shareIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return &MpcWalletShareStoredIterator{contract: _MpcWallet.contract, event: "ShareStored", logs: logs, sub: sub}, nil
}

// WatchShareStored is a free log subscription operation binding the contract event 0x498bef06e06c917324187e12264cb64a2ccde5e394d2f5ea9562c685f447f0c8.
//
// Solidity: event ShareStored(bytes32 indexed shareId, address indexed by)
func (_MpcWallet *MpcWalletFilterer) WatchShareStored(opts *bind.WatchOpts, sink chan<- *MpcWalletShareStored, shareId [][32]byte, by []common.Address) (event.Subscription, error) {

	var shareIdRule []interface{}
	for _, shareIdItem := range shareId {
		shareIdRule = append(shareIdRule, shareIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MpcWallet.contract.WatchLogs(opts, "ShareStored", shareIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MpcWalletShareStored)
				if err := _MpcWallet.contract.UnpackLog(event, "ShareStored", log); err != nil {
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
func (_MpcWallet *MpcWalletFilterer) ParseShareStored(log types.Log) (*MpcWalletShareStored, error) {
	event := new(MpcWalletShareStored)
	if err := _MpcWallet.contract.UnpackLog(event, "ShareStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MpcWalletShareUpdatedIterator is returned from FilterShareUpdated and is used to iterate over the raw logs and unpacked data for ShareUpdated events raised by the MpcWallet contract.
type MpcWalletShareUpdatedIterator struct {
	Event *MpcWalletShareUpdated // Event containing the contract specifics and raw log

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
func (it *MpcWalletShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MpcWalletShareUpdated)
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
		it.Event = new(MpcWalletShareUpdated)
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
func (it *MpcWalletShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MpcWalletShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MpcWalletShareUpdated represents a ShareUpdated event raised by the MpcWallet contract.
type MpcWalletShareUpdated struct {
	ShareId [32]byte
	By      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterShareUpdated is a free log retrieval operation binding the contract event 0x6d7c51f2ccbdd28127dc175d6f77522fff8e9a7a38c3513b4c329bb197036771.
//
// Solidity: event ShareUpdated(bytes32 indexed shareId, address indexed by)
func (_MpcWallet *MpcWalletFilterer) FilterShareUpdated(opts *bind.FilterOpts, shareId [][32]byte, by []common.Address) (*MpcWalletShareUpdatedIterator, error) {

	var shareIdRule []interface{}
	for _, shareIdItem := range shareId {
		shareIdRule = append(shareIdRule, shareIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MpcWallet.contract.FilterLogs(opts, "ShareUpdated", shareIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return &MpcWalletShareUpdatedIterator{contract: _MpcWallet.contract, event: "ShareUpdated", logs: logs, sub: sub}, nil
}

// WatchShareUpdated is a free log subscription operation binding the contract event 0x6d7c51f2ccbdd28127dc175d6f77522fff8e9a7a38c3513b4c329bb197036771.
//
// Solidity: event ShareUpdated(bytes32 indexed shareId, address indexed by)
func (_MpcWallet *MpcWalletFilterer) WatchShareUpdated(opts *bind.WatchOpts, sink chan<- *MpcWalletShareUpdated, shareId [][32]byte, by []common.Address) (event.Subscription, error) {

	var shareIdRule []interface{}
	for _, shareIdItem := range shareId {
		shareIdRule = append(shareIdRule, shareIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MpcWallet.contract.WatchLogs(opts, "ShareUpdated", shareIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MpcWalletShareUpdated)
				if err := _MpcWallet.contract.UnpackLog(event, "ShareUpdated", log); err != nil {
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
func (_MpcWallet *MpcWalletFilterer) ParseShareUpdated(log types.Log) (*MpcWalletShareUpdated, error) {
	event := new(MpcWalletShareUpdated)
	if err := _MpcWallet.contract.UnpackLog(event, "ShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MpcWalletUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MpcWallet contract.
type MpcWalletUnpausedIterator struct {
	Event *MpcWalletUnpaused // Event containing the contract specifics and raw log

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
func (it *MpcWalletUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MpcWalletUnpaused)
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
		it.Event = new(MpcWalletUnpaused)
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
func (it *MpcWalletUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MpcWalletUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MpcWalletUnpaused represents a Unpaused event raised by the MpcWallet contract.
type MpcWalletUnpaused struct {
	By  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed by)
func (_MpcWallet *MpcWalletFilterer) FilterUnpaused(opts *bind.FilterOpts, by []common.Address) (*MpcWalletUnpausedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MpcWallet.contract.FilterLogs(opts, "Unpaused", byRule)
	if err != nil {
		return nil, err
	}
	return &MpcWalletUnpausedIterator{contract: _MpcWallet.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed by)
func (_MpcWallet *MpcWalletFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MpcWalletUnpaused, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MpcWallet.contract.WatchLogs(opts, "Unpaused", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MpcWalletUnpaused)
				if err := _MpcWallet.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_MpcWallet *MpcWalletFilterer) ParseUnpaused(log types.Log) (*MpcWalletUnpaused, error) {
	event := new(MpcWalletUnpaused)
	if err := _MpcWallet.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// UtilsFacetMetaData contains all meta data concerning the UtilsFacet contract.
var UtilsFacetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"node\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// UtilsFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use UtilsFacetMetaData.ABI instead.
var UtilsFacetABI = UtilsFacetMetaData.ABI

// UtilsFacet is an auto generated Go binding around an Ethereum contract.
type UtilsFacet struct {
	UtilsFacetCaller     // Read-only binding to the contract
	UtilsFacetTransactor // Write-only binding to the contract
	UtilsFacetFilterer   // Log filterer for contract events
}

// UtilsFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsFacetSession struct {
	Contract     *UtilsFacet       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsFacetCallerSession struct {
	Contract *UtilsFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// UtilsFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsFacetTransactorSession struct {
	Contract     *UtilsFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UtilsFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsFacetRaw struct {
	Contract *UtilsFacet // Generic contract binding to access the raw methods on
}

// UtilsFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsFacetCallerRaw struct {
	Contract *UtilsFacetCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsFacetTransactorRaw struct {
	Contract *UtilsFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtilsFacet creates a new instance of UtilsFacet, bound to a specific deployed contract.
func NewUtilsFacet(address common.Address, backend bind.ContractBackend) (*UtilsFacet, error) {
	contract, err := bindUtilsFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UtilsFacet{UtilsFacetCaller: UtilsFacetCaller{contract: contract}, UtilsFacetTransactor: UtilsFacetTransactor{contract: contract}, UtilsFacetFilterer: UtilsFacetFilterer{contract: contract}}, nil
}

// NewUtilsFacetCaller creates a new read-only instance of UtilsFacet, bound to a specific deployed contract.
func NewUtilsFacetCaller(address common.Address, caller bind.ContractCaller) (*UtilsFacetCaller, error) {
	contract, err := bindUtilsFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsFacetCaller{contract: contract}, nil
}

// NewUtilsFacetTransactor creates a new write-only instance of UtilsFacet, bound to a specific deployed contract.
func NewUtilsFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsFacetTransactor, error) {
	contract, err := bindUtilsFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsFacetTransactor{contract: contract}, nil
}

// NewUtilsFacetFilterer creates a new log filterer instance of UtilsFacet, bound to a specific deployed contract.
func NewUtilsFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFacetFilterer, error) {
	contract, err := bindUtilsFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFacetFilterer{contract: contract}, nil
}

// bindUtilsFacet binds a generic wrapper to an already deployed contract.
func bindUtilsFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsFacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UtilsFacet *UtilsFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UtilsFacet.Contract.UtilsFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UtilsFacet *UtilsFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UtilsFacet.Contract.UtilsFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UtilsFacet *UtilsFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UtilsFacet.Contract.UtilsFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UtilsFacet *UtilsFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UtilsFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UtilsFacet *UtilsFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UtilsFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UtilsFacet *UtilsFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UtilsFacet.Contract.contract.Transact(opts, method, params...)
}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(address addr) pure returns(bytes32)
func (_UtilsFacet *UtilsFacetCaller) Node(opts *bind.CallOpts, addr common.Address) ([32]byte, error) {
	var out []interface{}
	err := _UtilsFacet.contract.Call(opts, &out, "node", addr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(address addr) pure returns(bytes32)
func (_UtilsFacet *UtilsFacetSession) Node(addr common.Address) ([32]byte, error) {
	return _UtilsFacet.Contract.Node(&_UtilsFacet.CallOpts, addr)
}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(address addr) pure returns(bytes32)
func (_UtilsFacet *UtilsFacetCallerSession) Node(addr common.Address) ([32]byte, error) {
	return _UtilsFacet.Contract.Node(&_UtilsFacet.CallOpts, addr)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diamondloupeabi

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
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

// IDiamondLoupeFacet is an auto generated low-level Go binding around an user-defined struct.
type IDiamondLoupeFacet struct {
	FacetAddress      common.Address
	FunctionSelectors [][4]byte
}

// IDiamondLoupeMetaData contains all meta data concerning the IDiamondLoupe contract.
var IDiamondLoupeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_functionSelector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facetAddress_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facetAddresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"facetFunctionSelectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondLoupe.Facet[]\",\"name\":\"facets_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IDiamondLoupeABI is the input ABI used to generate the binding from.
// Deprecated: Use IDiamondLoupeMetaData.ABI instead.
var IDiamondLoupeABI = IDiamondLoupeMetaData.ABI

// IDiamondLoupe is an auto generated Go binding around an Ethereum contract.
type IDiamondLoupe struct {
	IDiamondLoupeCaller     // Read-only binding to the contract
	IDiamondLoupeTransactor // Write-only binding to the contract
	IDiamondLoupeFilterer   // Log filterer for contract events
}

// IDiamondLoupeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDiamondLoupeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondLoupeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDiamondLoupeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondLoupeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDiamondLoupeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondLoupeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDiamondLoupeSession struct {
	Contract     *IDiamondLoupe    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDiamondLoupeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDiamondLoupeCallerSession struct {
	Contract *IDiamondLoupeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IDiamondLoupeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDiamondLoupeTransactorSession struct {
	Contract     *IDiamondLoupeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IDiamondLoupeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDiamondLoupeRaw struct {
	Contract *IDiamondLoupe // Generic contract binding to access the raw methods on
}

// IDiamondLoupeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDiamondLoupeCallerRaw struct {
	Contract *IDiamondLoupeCaller // Generic read-only contract binding to access the raw methods on
}

// IDiamondLoupeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDiamondLoupeTransactorRaw struct {
	Contract *IDiamondLoupeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDiamondLoupe creates a new instance of IDiamondLoupe, bound to a specific deployed contract.
func NewIDiamondLoupe(address common.Address, backend bind.ContractBackend) (*IDiamondLoupe, error) {
	contract, err := bindIDiamondLoupe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDiamondLoupe{IDiamondLoupeCaller: IDiamondLoupeCaller{contract: contract}, IDiamondLoupeTransactor: IDiamondLoupeTransactor{contract: contract}, IDiamondLoupeFilterer: IDiamondLoupeFilterer{contract: contract}}, nil
}

// NewIDiamondLoupeCaller creates a new read-only instance of IDiamondLoupe, bound to a specific deployed contract.
func NewIDiamondLoupeCaller(address common.Address, caller bind.ContractCaller) (*IDiamondLoupeCaller, error) {
	contract, err := bindIDiamondLoupe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDiamondLoupeCaller{contract: contract}, nil
}

// NewIDiamondLoupeTransactor creates a new write-only instance of IDiamondLoupe, bound to a specific deployed contract.
func NewIDiamondLoupeTransactor(address common.Address, transactor bind.ContractTransactor) (*IDiamondLoupeTransactor, error) {
	contract, err := bindIDiamondLoupe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDiamondLoupeTransactor{contract: contract}, nil
}

// NewIDiamondLoupeFilterer creates a new log filterer instance of IDiamondLoupe, bound to a specific deployed contract.
func NewIDiamondLoupeFilterer(address common.Address, filterer bind.ContractFilterer) (*IDiamondLoupeFilterer, error) {
	contract, err := bindIDiamondLoupe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDiamondLoupeFilterer{contract: contract}, nil
}

// bindIDiamondLoupe binds a generic wrapper to an already deployed contract.
func bindIDiamondLoupe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDiamondLoupeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDiamondLoupe *IDiamondLoupeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDiamondLoupe.Contract.IDiamondLoupeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDiamondLoupe *IDiamondLoupeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDiamondLoupe.Contract.IDiamondLoupeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDiamondLoupe *IDiamondLoupeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDiamondLoupe.Contract.IDiamondLoupeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDiamondLoupe *IDiamondLoupeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDiamondLoupe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDiamondLoupe *IDiamondLoupeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDiamondLoupe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDiamondLoupe *IDiamondLoupeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDiamondLoupe.Contract.contract.Transact(opts, method, params...)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_IDiamondLoupe *IDiamondLoupeCaller) FacetAddress(opts *bind.CallOpts, _functionSelector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _IDiamondLoupe.contract.Call(opts, &out, "facetAddress", _functionSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_IDiamondLoupe *IDiamondLoupeSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _IDiamondLoupe.Contract.FacetAddress(&_IDiamondLoupe.CallOpts, _functionSelector)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_IDiamondLoupe *IDiamondLoupeCallerSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _IDiamondLoupe.Contract.FacetAddress(&_IDiamondLoupe.CallOpts, _functionSelector)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_IDiamondLoupe *IDiamondLoupeCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IDiamondLoupe.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_IDiamondLoupe *IDiamondLoupeSession) FacetAddresses() ([]common.Address, error) {
	return _IDiamondLoupe.Contract.FacetAddresses(&_IDiamondLoupe.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_IDiamondLoupe *IDiamondLoupeCallerSession) FacetAddresses() ([]common.Address, error) {
	return _IDiamondLoupe.Contract.FacetAddresses(&_IDiamondLoupe.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_IDiamondLoupe *IDiamondLoupeCaller) FacetFunctionSelectors(opts *bind.CallOpts, _facet common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _IDiamondLoupe.contract.Call(opts, &out, "facetFunctionSelectors", _facet)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_IDiamondLoupe *IDiamondLoupeSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _IDiamondLoupe.Contract.FacetFunctionSelectors(&_IDiamondLoupe.CallOpts, _facet)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_IDiamondLoupe *IDiamondLoupeCallerSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _IDiamondLoupe.Contract.FacetFunctionSelectors(&_IDiamondLoupe.CallOpts, _facet)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_IDiamondLoupe *IDiamondLoupeCaller) Facets(opts *bind.CallOpts) ([]IDiamondLoupeFacet, error) {
	var out []interface{}
	err := _IDiamondLoupe.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new([]IDiamondLoupeFacet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDiamondLoupeFacet)).(*[]IDiamondLoupeFacet)

	return out0, err

}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_IDiamondLoupe *IDiamondLoupeSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _IDiamondLoupe.Contract.Facets(&_IDiamondLoupe.CallOpts)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_IDiamondLoupe *IDiamondLoupeCallerSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _IDiamondLoupe.Contract.Facets(&_IDiamondLoupe.CallOpts)
}

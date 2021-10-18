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

// NameResolverFacetMetaData contains all meta data concerning the NameResolverFacet contract.
var NameResolverFacetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NameResolverFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use NameResolverFacetMetaData.ABI instead.
var NameResolverFacetABI = NameResolverFacetMetaData.ABI

// NameResolverFacet is an auto generated Go binding around an Ethereum contract.
type NameResolverFacet struct {
	NameResolverFacetCaller     // Read-only binding to the contract
	NameResolverFacetTransactor // Write-only binding to the contract
	NameResolverFacetFilterer   // Log filterer for contract events
}

// NameResolverFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type NameResolverFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameResolverFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NameResolverFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameResolverFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NameResolverFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameResolverFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NameResolverFacetSession struct {
	Contract     *NameResolverFacet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NameResolverFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NameResolverFacetCallerSession struct {
	Contract *NameResolverFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NameResolverFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NameResolverFacetTransactorSession struct {
	Contract     *NameResolverFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NameResolverFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type NameResolverFacetRaw struct {
	Contract *NameResolverFacet // Generic contract binding to access the raw methods on
}

// NameResolverFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NameResolverFacetCallerRaw struct {
	Contract *NameResolverFacetCaller // Generic read-only contract binding to access the raw methods on
}

// NameResolverFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NameResolverFacetTransactorRaw struct {
	Contract *NameResolverFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNameResolverFacet creates a new instance of NameResolverFacet, bound to a specific deployed contract.
func NewNameResolverFacet(address common.Address, backend bind.ContractBackend) (*NameResolverFacet, error) {
	contract, err := bindNameResolverFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NameResolverFacet{NameResolverFacetCaller: NameResolverFacetCaller{contract: contract}, NameResolverFacetTransactor: NameResolverFacetTransactor{contract: contract}, NameResolverFacetFilterer: NameResolverFacetFilterer{contract: contract}}, nil
}

// NewNameResolverFacetCaller creates a new read-only instance of NameResolverFacet, bound to a specific deployed contract.
func NewNameResolverFacetCaller(address common.Address, caller bind.ContractCaller) (*NameResolverFacetCaller, error) {
	contract, err := bindNameResolverFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NameResolverFacetCaller{contract: contract}, nil
}

// NewNameResolverFacetTransactor creates a new write-only instance of NameResolverFacet, bound to a specific deployed contract.
func NewNameResolverFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*NameResolverFacetTransactor, error) {
	contract, err := bindNameResolverFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NameResolverFacetTransactor{contract: contract}, nil
}

// NewNameResolverFacetFilterer creates a new log filterer instance of NameResolverFacet, bound to a specific deployed contract.
func NewNameResolverFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*NameResolverFacetFilterer, error) {
	contract, err := bindNameResolverFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NameResolverFacetFilterer{contract: contract}, nil
}

// bindNameResolverFacet binds a generic wrapper to an already deployed contract.
func bindNameResolverFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NameResolverFacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameResolverFacet *NameResolverFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NameResolverFacet.Contract.NameResolverFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameResolverFacet *NameResolverFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameResolverFacet.Contract.NameResolverFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameResolverFacet *NameResolverFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameResolverFacet.Contract.NameResolverFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameResolverFacet *NameResolverFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NameResolverFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameResolverFacet *NameResolverFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameResolverFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameResolverFacet *NameResolverFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameResolverFacet.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_NameResolverFacet *NameResolverFacetCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var out []interface{}
	err := _NameResolverFacet.contract.Call(opts, &out, "name", node)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_NameResolverFacet *NameResolverFacetSession) Name(node [32]byte) (string, error) {
	return _NameResolverFacet.Contract.Name(&_NameResolverFacet.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_NameResolverFacet *NameResolverFacetCallerSession) Name(node [32]byte) (string, error) {
	return _NameResolverFacet.Contract.Name(&_NameResolverFacet.CallOpts, node)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string _name) returns()
func (_NameResolverFacet *NameResolverFacetTransactor) SetName(opts *bind.TransactOpts, node [32]byte, _name string) (*types.Transaction, error) {
	return _NameResolverFacet.contract.Transact(opts, "setName", node, _name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string _name) returns()
func (_NameResolverFacet *NameResolverFacetSession) SetName(node [32]byte, _name string) (*types.Transaction, error) {
	return _NameResolverFacet.Contract.SetName(&_NameResolverFacet.TransactOpts, node, _name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string _name) returns()
func (_NameResolverFacet *NameResolverFacetTransactorSession) SetName(node [32]byte, _name string) (*types.Transaction, error) {
	return _NameResolverFacet.Contract.SetName(&_NameResolverFacet.TransactOpts, node, _name)
}

// NameResolverFacetNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the NameResolverFacet contract.
type NameResolverFacetNameChangedIterator struct {
	Event *NameResolverFacetNameChanged // Event containing the contract specifics and raw log

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
func (it *NameResolverFacetNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameResolverFacetNameChanged)
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
		it.Event = new(NameResolverFacetNameChanged)
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
func (it *NameResolverFacetNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameResolverFacetNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameResolverFacetNameChanged represents a NameChanged event raised by the NameResolverFacet contract.
type NameResolverFacetNameChanged struct {
	Node [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_NameResolverFacet *NameResolverFacetFilterer) FilterNameChanged(opts *bind.FilterOpts, node [][32]byte) (*NameResolverFacetNameChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NameResolverFacet.contract.FilterLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &NameResolverFacetNameChangedIterator{contract: _NameResolverFacet.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_NameResolverFacet *NameResolverFacetFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *NameResolverFacetNameChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NameResolverFacet.contract.WatchLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameResolverFacetNameChanged)
				if err := _NameResolverFacet.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// ParseNameChanged is a log parse operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_NameResolverFacet *NameResolverFacetFilterer) ParseNameChanged(log types.Log) (*NameResolverFacetNameChanged, error) {
	event := new(NameResolverFacetNameChanged)
	if err := _NameResolverFacet.contract.UnpackLog(event, "NameChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

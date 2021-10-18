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

// AddressResolverFacetMetaData contains all meta data concerning the AddressResolverFacet contract.
var AddressResolverFacetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AddressResolverFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressResolverFacetMetaData.ABI instead.
var AddressResolverFacetABI = AddressResolverFacetMetaData.ABI

// AddressResolverFacet is an auto generated Go binding around an Ethereum contract.
type AddressResolverFacet struct {
	AddressResolverFacetCaller     // Read-only binding to the contract
	AddressResolverFacetTransactor // Write-only binding to the contract
	AddressResolverFacetFilterer   // Log filterer for contract events
}

// AddressResolverFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressResolverFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressResolverFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressResolverFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressResolverFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressResolverFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressResolverFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressResolverFacetSession struct {
	Contract     *AddressResolverFacet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AddressResolverFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressResolverFacetCallerSession struct {
	Contract *AddressResolverFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AddressResolverFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressResolverFacetTransactorSession struct {
	Contract     *AddressResolverFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AddressResolverFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressResolverFacetRaw struct {
	Contract *AddressResolverFacet // Generic contract binding to access the raw methods on
}

// AddressResolverFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressResolverFacetCallerRaw struct {
	Contract *AddressResolverFacetCaller // Generic read-only contract binding to access the raw methods on
}

// AddressResolverFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressResolverFacetTransactorRaw struct {
	Contract *AddressResolverFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressResolverFacet creates a new instance of AddressResolverFacet, bound to a specific deployed contract.
func NewAddressResolverFacet(address common.Address, backend bind.ContractBackend) (*AddressResolverFacet, error) {
	contract, err := bindAddressResolverFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressResolverFacet{AddressResolverFacetCaller: AddressResolverFacetCaller{contract: contract}, AddressResolverFacetTransactor: AddressResolverFacetTransactor{contract: contract}, AddressResolverFacetFilterer: AddressResolverFacetFilterer{contract: contract}}, nil
}

// NewAddressResolverFacetCaller creates a new read-only instance of AddressResolverFacet, bound to a specific deployed contract.
func NewAddressResolverFacetCaller(address common.Address, caller bind.ContractCaller) (*AddressResolverFacetCaller, error) {
	contract, err := bindAddressResolverFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressResolverFacetCaller{contract: contract}, nil
}

// NewAddressResolverFacetTransactor creates a new write-only instance of AddressResolverFacet, bound to a specific deployed contract.
func NewAddressResolverFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressResolverFacetTransactor, error) {
	contract, err := bindAddressResolverFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressResolverFacetTransactor{contract: contract}, nil
}

// NewAddressResolverFacetFilterer creates a new log filterer instance of AddressResolverFacet, bound to a specific deployed contract.
func NewAddressResolverFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressResolverFacetFilterer, error) {
	contract, err := bindAddressResolverFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressResolverFacetFilterer{contract: contract}, nil
}

// bindAddressResolverFacet binds a generic wrapper to an already deployed contract.
func bindAddressResolverFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressResolverFacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressResolverFacet *AddressResolverFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressResolverFacet.Contract.AddressResolverFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressResolverFacet *AddressResolverFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressResolverFacet.Contract.AddressResolverFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressResolverFacet *AddressResolverFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressResolverFacet.Contract.AddressResolverFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressResolverFacet *AddressResolverFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressResolverFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressResolverFacet *AddressResolverFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressResolverFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressResolverFacet *AddressResolverFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressResolverFacet.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_AddressResolverFacet *AddressResolverFacetCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AddressResolverFacet.contract.Call(opts, &out, "addr", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_AddressResolverFacet *AddressResolverFacetSession) Addr(node [32]byte) (common.Address, error) {
	return _AddressResolverFacet.Contract.Addr(&_AddressResolverFacet.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_AddressResolverFacet *AddressResolverFacetCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _AddressResolverFacet.Contract.Addr(&_AddressResolverFacet.CallOpts, node)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address _addr) returns()
func (_AddressResolverFacet *AddressResolverFacetTransactor) SetAddr(opts *bind.TransactOpts, node [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _AddressResolverFacet.contract.Transact(opts, "setAddr", node, _addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address _addr) returns()
func (_AddressResolverFacet *AddressResolverFacetSession) SetAddr(node [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _AddressResolverFacet.Contract.SetAddr(&_AddressResolverFacet.TransactOpts, node, _addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address _addr) returns()
func (_AddressResolverFacet *AddressResolverFacetTransactorSession) SetAddr(node [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _AddressResolverFacet.Contract.SetAddr(&_AddressResolverFacet.TransactOpts, node, _addr)
}

// AddressResolverFacetAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the AddressResolverFacet contract.
type AddressResolverFacetAddrChangedIterator struct {
	Event *AddressResolverFacetAddrChanged // Event containing the contract specifics and raw log

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
func (it *AddressResolverFacetAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressResolverFacetAddrChanged)
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
		it.Event = new(AddressResolverFacetAddrChanged)
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
func (it *AddressResolverFacetAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressResolverFacetAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressResolverFacetAddrChanged represents a AddrChanged event raised by the AddressResolverFacet contract.
type AddressResolverFacetAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_AddressResolverFacet *AddressResolverFacetFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*AddressResolverFacetAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AddressResolverFacet.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &AddressResolverFacetAddrChangedIterator{contract: _AddressResolverFacet.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_AddressResolverFacet *AddressResolverFacetFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *AddressResolverFacetAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AddressResolverFacet.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressResolverFacetAddrChanged)
				if err := _AddressResolverFacet.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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

// ParseAddrChanged is a log parse operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_AddressResolverFacet *AddressResolverFacetFilterer) ParseAddrChanged(log types.Log) (*AddressResolverFacetAddrChanged, error) {
	event := new(AddressResolverFacetAddrChanged)
	if err := _AddressResolverFacet.contract.UnpackLog(event, "AddrChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diamondcutabi

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

// IDiamondCutFacetCut is an auto generated low-level Go binding around an user-defined struct.
type IDiamondCutFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// IDiamondCutMetaData contains all meta data concerning the IDiamondCut contract.
var IDiamondCutMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IDiamondCutABI is the input ABI used to generate the binding from.
// Deprecated: Use IDiamondCutMetaData.ABI instead.
var IDiamondCutABI = IDiamondCutMetaData.ABI

// IDiamondCut is an auto generated Go binding around an Ethereum contract.
type IDiamondCut struct {
	IDiamondCutCaller     // Read-only binding to the contract
	IDiamondCutTransactor // Write-only binding to the contract
	IDiamondCutFilterer   // Log filterer for contract events
}

// IDiamondCutCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDiamondCutCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondCutTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDiamondCutTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondCutFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDiamondCutFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondCutSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDiamondCutSession struct {
	Contract     *IDiamondCut      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDiamondCutCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDiamondCutCallerSession struct {
	Contract *IDiamondCutCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IDiamondCutTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDiamondCutTransactorSession struct {
	Contract     *IDiamondCutTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IDiamondCutRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDiamondCutRaw struct {
	Contract *IDiamondCut // Generic contract binding to access the raw methods on
}

// IDiamondCutCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDiamondCutCallerRaw struct {
	Contract *IDiamondCutCaller // Generic read-only contract binding to access the raw methods on
}

// IDiamondCutTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDiamondCutTransactorRaw struct {
	Contract *IDiamondCutTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDiamondCut creates a new instance of IDiamondCut, bound to a specific deployed contract.
func NewIDiamondCut(address common.Address, backend bind.ContractBackend) (*IDiamondCut, error) {
	contract, err := bindIDiamondCut(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDiamondCut{IDiamondCutCaller: IDiamondCutCaller{contract: contract}, IDiamondCutTransactor: IDiamondCutTransactor{contract: contract}, IDiamondCutFilterer: IDiamondCutFilterer{contract: contract}}, nil
}

// NewIDiamondCutCaller creates a new read-only instance of IDiamondCut, bound to a specific deployed contract.
func NewIDiamondCutCaller(address common.Address, caller bind.ContractCaller) (*IDiamondCutCaller, error) {
	contract, err := bindIDiamondCut(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDiamondCutCaller{contract: contract}, nil
}

// NewIDiamondCutTransactor creates a new write-only instance of IDiamondCut, bound to a specific deployed contract.
func NewIDiamondCutTransactor(address common.Address, transactor bind.ContractTransactor) (*IDiamondCutTransactor, error) {
	contract, err := bindIDiamondCut(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDiamondCutTransactor{contract: contract}, nil
}

// NewIDiamondCutFilterer creates a new log filterer instance of IDiamondCut, bound to a specific deployed contract.
func NewIDiamondCutFilterer(address common.Address, filterer bind.ContractFilterer) (*IDiamondCutFilterer, error) {
	contract, err := bindIDiamondCut(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDiamondCutFilterer{contract: contract}, nil
}

// bindIDiamondCut binds a generic wrapper to an already deployed contract.
func bindIDiamondCut(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDiamondCutABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDiamondCut *IDiamondCutRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDiamondCut.Contract.IDiamondCutCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDiamondCut *IDiamondCutRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDiamondCut.Contract.IDiamondCutTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDiamondCut *IDiamondCutRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDiamondCut.Contract.IDiamondCutTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDiamondCut *IDiamondCutCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDiamondCut.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDiamondCut *IDiamondCutTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDiamondCut.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDiamondCut *IDiamondCutTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDiamondCut.Contract.contract.Transact(opts, method, params...)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_IDiamondCut *IDiamondCutTransactor) DiamondCut(opts *bind.TransactOpts, _diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _IDiamondCut.contract.Transact(opts, "diamondCut", _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_IDiamondCut *IDiamondCutSession) DiamondCut(_diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _IDiamondCut.Contract.DiamondCut(&_IDiamondCut.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_IDiamondCut *IDiamondCutTransactorSession) DiamondCut(_diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _IDiamondCut.Contract.DiamondCut(&_IDiamondCut.TransactOpts, _diamondCut, _init, _calldata)
}

// IDiamondCutDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the IDiamondCut contract.
type IDiamondCutDiamondCutIterator struct {
	Event *IDiamondCutDiamondCut // Event containing the contract specifics and raw log

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
func (it *IDiamondCutDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDiamondCutDiamondCut)
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
		it.Event = new(IDiamondCutDiamondCut)
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
func (it *IDiamondCutDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDiamondCutDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDiamondCutDiamondCut represents a DiamondCut event raised by the IDiamondCut contract.
type IDiamondCutDiamondCut struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_IDiamondCut *IDiamondCutFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*IDiamondCutDiamondCutIterator, error) {

	logs, sub, err := _IDiamondCut.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &IDiamondCutDiamondCutIterator{contract: _IDiamondCut.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_IDiamondCut *IDiamondCutFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *IDiamondCutDiamondCut) (event.Subscription, error) {

	logs, sub, err := _IDiamondCut.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDiamondCutDiamondCut)
				if err := _IDiamondCut.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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

// ParseDiamondCut is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_IDiamondCut *IDiamondCutFilterer) ParseDiamondCut(log types.Log) (*IDiamondCutDiamondCut, error) {
	event := new(IDiamondCutDiamondCut)
	if err := _IDiamondCut.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

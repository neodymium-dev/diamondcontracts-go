// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ownershipabi

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

// DiamondOwnershipMetaData contains all meta data concerning the DiamondOwnership contract.
var DiamondOwnershipMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DiamondOwnershipABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondOwnershipMetaData.ABI instead.
var DiamondOwnershipABI = DiamondOwnershipMetaData.ABI

// DiamondOwnership is an auto generated Go binding around an Ethereum contract.
type DiamondOwnership struct {
	DiamondOwnershipCaller     // Read-only binding to the contract
	DiamondOwnershipTransactor // Write-only binding to the contract
	DiamondOwnershipFilterer   // Log filterer for contract events
}

// DiamondOwnershipCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondOwnershipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondOwnershipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondOwnershipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondOwnershipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondOwnershipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondOwnershipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondOwnershipSession struct {
	Contract     *DiamondOwnership // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiamondOwnershipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondOwnershipCallerSession struct {
	Contract *DiamondOwnershipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DiamondOwnershipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondOwnershipTransactorSession struct {
	Contract     *DiamondOwnershipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DiamondOwnershipRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondOwnershipRaw struct {
	Contract *DiamondOwnership // Generic contract binding to access the raw methods on
}

// DiamondOwnershipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondOwnershipCallerRaw struct {
	Contract *DiamondOwnershipCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondOwnershipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondOwnershipTransactorRaw struct {
	Contract *DiamondOwnershipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamondOwnership creates a new instance of DiamondOwnership, bound to a specific deployed contract.
func NewDiamondOwnership(address common.Address, backend bind.ContractBackend) (*DiamondOwnership, error) {
	contract, err := bindDiamondOwnership(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiamondOwnership{DiamondOwnershipCaller: DiamondOwnershipCaller{contract: contract}, DiamondOwnershipTransactor: DiamondOwnershipTransactor{contract: contract}, DiamondOwnershipFilterer: DiamondOwnershipFilterer{contract: contract}}, nil
}

// NewDiamondOwnershipCaller creates a new read-only instance of DiamondOwnership, bound to a specific deployed contract.
func NewDiamondOwnershipCaller(address common.Address, caller bind.ContractCaller) (*DiamondOwnershipCaller, error) {
	contract, err := bindDiamondOwnership(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondOwnershipCaller{contract: contract}, nil
}

// NewDiamondOwnershipTransactor creates a new write-only instance of DiamondOwnership, bound to a specific deployed contract.
func NewDiamondOwnershipTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondOwnershipTransactor, error) {
	contract, err := bindDiamondOwnership(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondOwnershipTransactor{contract: contract}, nil
}

// NewDiamondOwnershipFilterer creates a new log filterer instance of DiamondOwnership, bound to a specific deployed contract.
func NewDiamondOwnershipFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondOwnershipFilterer, error) {
	contract, err := bindDiamondOwnership(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondOwnershipFilterer{contract: contract}, nil
}

// bindDiamondOwnership binds a generic wrapper to an already deployed contract.
func bindDiamondOwnership(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiamondOwnershipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondOwnership *DiamondOwnershipRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondOwnership.Contract.DiamondOwnershipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondOwnership *DiamondOwnershipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondOwnership.Contract.DiamondOwnershipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondOwnership *DiamondOwnershipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondOwnership.Contract.DiamondOwnershipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondOwnership *DiamondOwnershipCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondOwnership.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondOwnership *DiamondOwnershipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondOwnership.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondOwnership *DiamondOwnershipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondOwnership.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DiamondOwnership *DiamondOwnershipCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DiamondOwnership.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DiamondOwnership *DiamondOwnershipSession) Owner() (common.Address, error) {
	return _DiamondOwnership.Contract.Owner(&_DiamondOwnership.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DiamondOwnership *DiamondOwnershipCallerSession) Owner() (common.Address, error) {
	return _DiamondOwnership.Contract.Owner(&_DiamondOwnership.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_DiamondOwnership *DiamondOwnershipTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _DiamondOwnership.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_DiamondOwnership *DiamondOwnershipSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DiamondOwnership.Contract.TransferOwnership(&_DiamondOwnership.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_DiamondOwnership *DiamondOwnershipTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DiamondOwnership.Contract.TransferOwnership(&_DiamondOwnership.TransactOpts, _newOwner)
}

// DiamondOwnershipOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DiamondOwnership contract.
type DiamondOwnershipOwnershipTransferredIterator struct {
	Event *DiamondOwnershipOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DiamondOwnershipOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondOwnershipOwnershipTransferred)
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
		it.Event = new(DiamondOwnershipOwnershipTransferred)
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
func (it *DiamondOwnershipOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondOwnershipOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondOwnershipOwnershipTransferred represents a OwnershipTransferred event raised by the DiamondOwnership contract.
type DiamondOwnershipOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DiamondOwnership *DiamondOwnershipFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DiamondOwnershipOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DiamondOwnership.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DiamondOwnershipOwnershipTransferredIterator{contract: _DiamondOwnership.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DiamondOwnership *DiamondOwnershipFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DiamondOwnershipOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DiamondOwnership.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondOwnershipOwnershipTransferred)
				if err := _DiamondOwnership.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DiamondOwnership *DiamondOwnershipFilterer) ParseOwnershipTransferred(log types.Log) (*DiamondOwnershipOwnershipTransferred, error) {
	event := new(DiamondOwnershipOwnershipTransferred)
	if err := _DiamondOwnership.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

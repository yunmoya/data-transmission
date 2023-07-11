// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package inforeq

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
	_ = abi.ConvertType
)

// InforeqMetaData contains all meta data concerning the Inforeq contract.
var InforeqMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"reqType\",\"type\":\"uint32\"}],\"name\":\"NewRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"newestRequests\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"}],\"name\":\"getCategory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"}],\"name\":\"getPubKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// InforeqABI is the input ABI used to generate the binding from.
// Deprecated: Use InforeqMetaData.ABI instead.
var InforeqABI = InforeqMetaData.ABI

// Inforeq is an auto generated Go binding around an Ethereum contract.
type Inforeq struct {
	InforeqCaller     // Read-only binding to the contract
	InforeqTransactor // Write-only binding to the contract
	InforeqFilterer   // Log filterer for contract events
}

// InforeqCaller is an auto generated read-only Go binding around an Ethereum contract.
type InforeqCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InforeqTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InforeqTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InforeqFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InforeqFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InforeqSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InforeqSession struct {
	Contract     *Inforeq          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InforeqCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InforeqCallerSession struct {
	Contract *InforeqCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// InforeqTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InforeqTransactorSession struct {
	Contract     *InforeqTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// InforeqRaw is an auto generated low-level Go binding around an Ethereum contract.
type InforeqRaw struct {
	Contract *Inforeq // Generic contract binding to access the raw methods on
}

// InforeqCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InforeqCallerRaw struct {
	Contract *InforeqCaller // Generic read-only contract binding to access the raw methods on
}

// InforeqTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InforeqTransactorRaw struct {
	Contract *InforeqTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInforeq creates a new instance of Inforeq, bound to a specific deployed contract.
func NewInforeq(address common.Address, backend bind.ContractBackend) (*Inforeq, error) {
	contract, err := bindInforeq(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Inforeq{InforeqCaller: InforeqCaller{contract: contract}, InforeqTransactor: InforeqTransactor{contract: contract}, InforeqFilterer: InforeqFilterer{contract: contract}}, nil
}

// NewInforeqCaller creates a new read-only instance of Inforeq, bound to a specific deployed contract.
func NewInforeqCaller(address common.Address, caller bind.ContractCaller) (*InforeqCaller, error) {
	contract, err := bindInforeq(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InforeqCaller{contract: contract}, nil
}

// NewInforeqTransactor creates a new write-only instance of Inforeq, bound to a specific deployed contract.
func NewInforeqTransactor(address common.Address, transactor bind.ContractTransactor) (*InforeqTransactor, error) {
	contract, err := bindInforeq(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InforeqTransactor{contract: contract}, nil
}

// NewInforeqFilterer creates a new log filterer instance of Inforeq, bound to a specific deployed contract.
func NewInforeqFilterer(address common.Address, filterer bind.ContractFilterer) (*InforeqFilterer, error) {
	contract, err := bindInforeq(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InforeqFilterer{contract: contract}, nil
}

// bindInforeq binds a generic wrapper to an already deployed contract.
func bindInforeq(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InforeqMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inforeq *InforeqRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inforeq.Contract.InforeqCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inforeq *InforeqRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inforeq.Contract.InforeqTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inforeq *InforeqRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inforeq.Contract.InforeqTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inforeq *InforeqCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inforeq.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inforeq *InforeqTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inforeq.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inforeq *InforeqTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inforeq.Contract.contract.Transact(opts, method, params...)
}

// NewestRequests is a free data retrieval call binding the contract method 0x3dbc94c7.
//
// Solidity: function newestRequests(address ) view returns(string)
func (_Inforeq *InforeqCaller) NewestRequests(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Inforeq.contract.Call(opts, &out, "newestRequests", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NewestRequests is a free data retrieval call binding the contract method 0x3dbc94c7.
//
// Solidity: function newestRequests(address ) view returns(string)
func (_Inforeq *InforeqSession) NewestRequests(arg0 common.Address) (string, error) {
	return _Inforeq.Contract.NewestRequests(&_Inforeq.CallOpts, arg0)
}

// NewestRequests is a free data retrieval call binding the contract method 0x3dbc94c7.
//
// Solidity: function newestRequests(address ) view returns(string)
func (_Inforeq *InforeqCallerSession) NewestRequests(arg0 common.Address) (string, error) {
	return _Inforeq.Contract.NewestRequests(&_Inforeq.CallOpts, arg0)
}

// GetCategory is a paid mutator transaction binding the contract method 0x190377a5.
//
// Solidity: function getCategory(string userId) returns()
func (_Inforeq *InforeqTransactor) GetCategory(opts *bind.TransactOpts, userId string) (*types.Transaction, error) {
	return _Inforeq.contract.Transact(opts, "getCategory", userId)
}

// GetCategory is a paid mutator transaction binding the contract method 0x190377a5.
//
// Solidity: function getCategory(string userId) returns()
func (_Inforeq *InforeqSession) GetCategory(userId string) (*types.Transaction, error) {
	return _Inforeq.Contract.GetCategory(&_Inforeq.TransactOpts, userId)
}

// GetCategory is a paid mutator transaction binding the contract method 0x190377a5.
//
// Solidity: function getCategory(string userId) returns()
func (_Inforeq *InforeqTransactorSession) GetCategory(userId string) (*types.Transaction, error) {
	return _Inforeq.Contract.GetCategory(&_Inforeq.TransactOpts, userId)
}

// GetPubKey is a paid mutator transaction binding the contract method 0x0d5ecf2e.
//
// Solidity: function getPubKey(string userId) returns()
func (_Inforeq *InforeqTransactor) GetPubKey(opts *bind.TransactOpts, userId string) (*types.Transaction, error) {
	return _Inforeq.contract.Transact(opts, "getPubKey", userId)
}

// GetPubKey is a paid mutator transaction binding the contract method 0x0d5ecf2e.
//
// Solidity: function getPubKey(string userId) returns()
func (_Inforeq *InforeqSession) GetPubKey(userId string) (*types.Transaction, error) {
	return _Inforeq.Contract.GetPubKey(&_Inforeq.TransactOpts, userId)
}

// GetPubKey is a paid mutator transaction binding the contract method 0x0d5ecf2e.
//
// Solidity: function getPubKey(string userId) returns()
func (_Inforeq *InforeqTransactorSession) GetPubKey(userId string) (*types.Transaction, error) {
	return _Inforeq.Contract.GetPubKey(&_Inforeq.TransactOpts, userId)
}

// InforeqNewRequestIterator is returned from FilterNewRequest and is used to iterate over the raw logs and unpacked data for NewRequest events raised by the Inforeq contract.
type InforeqNewRequestIterator struct {
	Event *InforeqNewRequest // Event containing the contract specifics and raw log

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
func (it *InforeqNewRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InforeqNewRequest)
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
		it.Event = new(InforeqNewRequest)
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
func (it *InforeqNewRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InforeqNewRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InforeqNewRequest represents a NewRequest event raised by the Inforeq contract.
type InforeqNewRequest struct {
	User    common.Address
	UserId  string
	ReqType uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewRequest is a free log retrieval operation binding the contract event 0x9cb1c72cd20f8506d26e23e97ad33da8ce94c16da438a1c9f784100570d67be8.
//
// Solidity: event NewRequest(address user, string userId, uint32 reqType)
func (_Inforeq *InforeqFilterer) FilterNewRequest(opts *bind.FilterOpts) (*InforeqNewRequestIterator, error) {

	logs, sub, err := _Inforeq.contract.FilterLogs(opts, "NewRequest")
	if err != nil {
		return nil, err
	}
	return &InforeqNewRequestIterator{contract: _Inforeq.contract, event: "NewRequest", logs: logs, sub: sub}, nil
}

// WatchNewRequest is a free log subscription operation binding the contract event 0x9cb1c72cd20f8506d26e23e97ad33da8ce94c16da438a1c9f784100570d67be8.
//
// Solidity: event NewRequest(address user, string userId, uint32 reqType)
func (_Inforeq *InforeqFilterer) WatchNewRequest(opts *bind.WatchOpts, sink chan<- *InforeqNewRequest) (event.Subscription, error) {

	logs, sub, err := _Inforeq.contract.WatchLogs(opts, "NewRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InforeqNewRequest)
				if err := _Inforeq.contract.UnpackLog(event, "NewRequest", log); err != nil {
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

// ParseNewRequest is a log parse operation binding the contract event 0x9cb1c72cd20f8506d26e23e97ad33da8ce94c16da438a1c9f784100570d67be8.
//
// Solidity: event NewRequest(address user, string userId, uint32 reqType)
func (_Inforeq *InforeqFilterer) ParseNewRequest(log types.Log) (*InforeqNewRequest, error) {
	event := new(InforeqNewRequest)
	if err := _Inforeq.contract.UnpackLog(event, "NewRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

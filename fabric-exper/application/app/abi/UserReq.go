// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package myabi

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

// UserreqMetaData contains all meta data concerning the Userreq contract.
var UserreqMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"config\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"encryptRequired\",\"type\":\"bool\"}],\"name\":\"NewRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"newestRequests\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"config\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"encryptRequired\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"config\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"}],\"name\":\"requestVMWithEncryption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"config\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"}],\"name\":\"requestVMWithoutEncryption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UserreqABI is the input ABI used to generate the binding from.
// Deprecated: Use UserreqMetaData.ABI instead.
var UserreqABI = UserreqMetaData.ABI

// Userreq is an auto generated Go binding around an Ethereum contract.
type Userreq struct {
	UserreqCaller     // Read-only binding to the contract
	UserreqTransactor // Write-only binding to the contract
	UserreqFilterer   // Log filterer for contract events
}

// UserreqCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserreqCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserreqTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserreqTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserreqFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserreqFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserreqSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserreqSession struct {
	Contract     *Userreq          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserreqCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserreqCallerSession struct {
	Contract *UserreqCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// UserreqTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserreqTransactorSession struct {
	Contract     *UserreqTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UserreqRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserreqRaw struct {
	Contract *Userreq // Generic contract binding to access the raw methods on
}

// UserreqCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserreqCallerRaw struct {
	Contract *UserreqCaller // Generic read-only contract binding to access the raw methods on
}

// UserreqTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserreqTransactorRaw struct {
	Contract *UserreqTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserreq creates a new instance of Userreq, bound to a specific deployed contract.
func NewUserreq(address common.Address, backend bind.ContractBackend) (*Userreq, error) {
	contract, err := bindUserreq(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Userreq{UserreqCaller: UserreqCaller{contract: contract}, UserreqTransactor: UserreqTransactor{contract: contract}, UserreqFilterer: UserreqFilterer{contract: contract}}, nil
}

// NewUserreqCaller creates a new read-only instance of Userreq, bound to a specific deployed contract.
func NewUserreqCaller(address common.Address, caller bind.ContractCaller) (*UserreqCaller, error) {
	contract, err := bindUserreq(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserreqCaller{contract: contract}, nil
}

// NewUserreqTransactor creates a new write-only instance of Userreq, bound to a specific deployed contract.
func NewUserreqTransactor(address common.Address, transactor bind.ContractTransactor) (*UserreqTransactor, error) {
	contract, err := bindUserreq(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserreqTransactor{contract: contract}, nil
}

// NewUserreqFilterer creates a new log filterer instance of Userreq, bound to a specific deployed contract.
func NewUserreqFilterer(address common.Address, filterer bind.ContractFilterer) (*UserreqFilterer, error) {
	contract, err := bindUserreq(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserreqFilterer{contract: contract}, nil
}

// bindUserreq binds a generic wrapper to an already deployed contract.
func bindUserreq(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UserreqMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Userreq *UserreqRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Userreq.Contract.UserreqCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Userreq *UserreqRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Userreq.Contract.UserreqTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Userreq *UserreqRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Userreq.Contract.UserreqTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Userreq *UserreqCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Userreq.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Userreq *UserreqTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Userreq.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Userreq *UserreqTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Userreq.Contract.contract.Transact(opts, method, params...)
}

// NewestRequests is a free data retrieval call binding the contract method 0x3dbc94c7.
//
// Solidity: function newestRequests(address ) view returns(string userId, string pubKey, uint32 config, uint32 duration, bool encryptRequired)
func (_Userreq *UserreqCaller) NewestRequests(opts *bind.CallOpts, arg0 common.Address) (struct {
	UserId          string
	PubKey          string
	Config          uint32
	Duration        uint32
	EncryptRequired bool
}, error) {
	var out []interface{}
	err := _Userreq.contract.Call(opts, &out, "newestRequests", arg0)

	outstruct := new(struct {
		UserId          string
		PubKey          string
		Config          uint32
		Duration        uint32
		EncryptRequired bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.PubKey = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Config = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Duration = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.EncryptRequired = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// NewestRequests is a free data retrieval call binding the contract method 0x3dbc94c7.
//
// Solidity: function newestRequests(address ) view returns(string userId, string pubKey, uint32 config, uint32 duration, bool encryptRequired)
func (_Userreq *UserreqSession) NewestRequests(arg0 common.Address) (struct {
	UserId          string
	PubKey          string
	Config          uint32
	Duration        uint32
	EncryptRequired bool
}, error) {
	return _Userreq.Contract.NewestRequests(&_Userreq.CallOpts, arg0)
}

// NewestRequests is a free data retrieval call binding the contract method 0x3dbc94c7.
//
// Solidity: function newestRequests(address ) view returns(string userId, string pubKey, uint32 config, uint32 duration, bool encryptRequired)
func (_Userreq *UserreqCallerSession) NewestRequests(arg0 common.Address) (struct {
	UserId          string
	PubKey          string
	Config          uint32
	Duration        uint32
	EncryptRequired bool
}, error) {
	return _Userreq.Contract.NewestRequests(&_Userreq.CallOpts, arg0)
}

// RequestVMWithEncryption is a paid mutator transaction binding the contract method 0x65626f6d.
//
// Solidity: function requestVMWithEncryption(string userId, uint32 config, uint32 duration, string pubKey) returns()
func (_Userreq *UserreqTransactor) RequestVMWithEncryption(opts *bind.TransactOpts, userId string, config uint32, duration uint32, pubKey string) (*types.Transaction, error) {
	return _Userreq.contract.Transact(opts, "requestVMWithEncryption", userId, config, duration, pubKey)
}

// RequestVMWithEncryption is a paid mutator transaction binding the contract method 0x65626f6d.
//
// Solidity: function requestVMWithEncryption(string userId, uint32 config, uint32 duration, string pubKey) returns()
func (_Userreq *UserreqSession) RequestVMWithEncryption(userId string, config uint32, duration uint32, pubKey string) (*types.Transaction, error) {
	return _Userreq.Contract.RequestVMWithEncryption(&_Userreq.TransactOpts, userId, config, duration, pubKey)
}

// RequestVMWithEncryption is a paid mutator transaction binding the contract method 0x65626f6d.
//
// Solidity: function requestVMWithEncryption(string userId, uint32 config, uint32 duration, string pubKey) returns()
func (_Userreq *UserreqTransactorSession) RequestVMWithEncryption(userId string, config uint32, duration uint32, pubKey string) (*types.Transaction, error) {
	return _Userreq.Contract.RequestVMWithEncryption(&_Userreq.TransactOpts, userId, config, duration, pubKey)
}

// RequestVMWithoutEncryption is a paid mutator transaction binding the contract method 0xb679d883.
//
// Solidity: function requestVMWithoutEncryption(string userId, uint32 config, uint32 duration) returns()
func (_Userreq *UserreqTransactor) RequestVMWithoutEncryption(opts *bind.TransactOpts, userId string, config uint32, duration uint32) (*types.Transaction, error) {
	return _Userreq.contract.Transact(opts, "requestVMWithoutEncryption", userId, config, duration)
}

// RequestVMWithoutEncryption is a paid mutator transaction binding the contract method 0xb679d883.
//
// Solidity: function requestVMWithoutEncryption(string userId, uint32 config, uint32 duration) returns()
func (_Userreq *UserreqSession) RequestVMWithoutEncryption(userId string, config uint32, duration uint32) (*types.Transaction, error) {
	return _Userreq.Contract.RequestVMWithoutEncryption(&_Userreq.TransactOpts, userId, config, duration)
}

// RequestVMWithoutEncryption is a paid mutator transaction binding the contract method 0xb679d883.
//
// Solidity: function requestVMWithoutEncryption(string userId, uint32 config, uint32 duration) returns()
func (_Userreq *UserreqTransactorSession) RequestVMWithoutEncryption(userId string, config uint32, duration uint32) (*types.Transaction, error) {
	return _Userreq.Contract.RequestVMWithoutEncryption(&_Userreq.TransactOpts, userId, config, duration)
}

// UserreqNewRequestIterator is returned from FilterNewRequest and is used to iterate over the raw logs and unpacked data for NewRequest events raised by the Userreq contract.
type UserreqNewRequestIterator struct {
	Event *UserreqNewRequest // Event containing the contract specifics and raw log

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
func (it *UserreqNewRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserreqNewRequest)
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
		it.Event = new(UserreqNewRequest)
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
func (it *UserreqNewRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserreqNewRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserreqNewRequest represents a NewRequest event raised by the Userreq contract.
type UserreqNewRequest struct {
	User            common.Address
	UserId          string
	PubKey          string
	Config          uint32
	Duration        uint32
	EncryptRequired bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewRequest is a free log retrieval operation binding the contract event 0x535ea70a9b46705e4726ca22562803c32f169182640b3e63d91f413c46c772c9.
//
// Solidity: event NewRequest(address user, string userId, string pubKey, uint32 config, uint32 duration, bool encryptRequired)
func (_Userreq *UserreqFilterer) FilterNewRequest(opts *bind.FilterOpts) (*UserreqNewRequestIterator, error) {

	logs, sub, err := _Userreq.contract.FilterLogs(opts, "NewRequest")
	if err != nil {
		return nil, err
	}
	return &UserreqNewRequestIterator{contract: _Userreq.contract, event: "NewRequest", logs: logs, sub: sub}, nil
}

// WatchNewRequest is a free log subscription operation binding the contract event 0x535ea70a9b46705e4726ca22562803c32f169182640b3e63d91f413c46c772c9.
//
// Solidity: event NewRequest(address user, string userId, string pubKey, uint32 config, uint32 duration, bool encryptRequired)
func (_Userreq *UserreqFilterer) WatchNewRequest(opts *bind.WatchOpts, sink chan<- *UserreqNewRequest) (event.Subscription, error) {

	logs, sub, err := _Userreq.contract.WatchLogs(opts, "NewRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserreqNewRequest)
				if err := _Userreq.contract.UnpackLog(event, "NewRequest", log); err != nil {
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

// ParseNewRequest is a log parse operation binding the contract event 0x535ea70a9b46705e4726ca22562803c32f169182640b3e63d91f413c46c772c9.
//
// Solidity: event NewRequest(address user, string userId, string pubKey, uint32 config, uint32 duration, bool encryptRequired)
func (_Userreq *UserreqFilterer) ParseNewRequest(log types.Log) (*UserreqNewRequest, error) {
	event := new(UserreqNewRequest)
	if err := _Userreq.contract.UnpackLog(event, "NewRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

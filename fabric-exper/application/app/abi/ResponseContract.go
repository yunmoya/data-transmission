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

// ResponseMetaData contains all meta data concerning the Response contract.
var ResponseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dataHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dataSig\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"encryptRequired\",\"type\":\"bool\"}],\"name\":\"NewResponse\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"newestResponse\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataSig\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"encryptRequired\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataSig\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"encryptRequired\",\"type\":\"bool\"}],\"name\":\"responseToUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ResponseABI is the input ABI used to generate the binding from.
// Deprecated: Use ResponseMetaData.ABI instead.
var ResponseABI = ResponseMetaData.ABI

// Response is an auto generated Go binding around an Ethereum contract.
type Response struct {
	ResponseCaller     // Read-only binding to the contract
	ResponseTransactor // Write-only binding to the contract
	ResponseFilterer   // Log filterer for contract events
}

// ResponseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResponseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResponseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResponseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResponseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResponseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResponseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResponseSession struct {
	Contract     *Response         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResponseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResponseCallerSession struct {
	Contract *ResponseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ResponseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResponseTransactorSession struct {
	Contract     *ResponseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ResponseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResponseRaw struct {
	Contract *Response // Generic contract binding to access the raw methods on
}

// ResponseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResponseCallerRaw struct {
	Contract *ResponseCaller // Generic read-only contract binding to access the raw methods on
}

// ResponseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResponseTransactorRaw struct {
	Contract *ResponseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResponse creates a new instance of Response, bound to a specific deployed contract.
func NewResponse(address common.Address, backend bind.ContractBackend) (*Response, error) {
	contract, err := bindResponse(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Response{ResponseCaller: ResponseCaller{contract: contract}, ResponseTransactor: ResponseTransactor{contract: contract}, ResponseFilterer: ResponseFilterer{contract: contract}}, nil
}

// NewResponseCaller creates a new read-only instance of Response, bound to a specific deployed contract.
func NewResponseCaller(address common.Address, caller bind.ContractCaller) (*ResponseCaller, error) {
	contract, err := bindResponse(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResponseCaller{contract: contract}, nil
}

// NewResponseTransactor creates a new write-only instance of Response, bound to a specific deployed contract.
func NewResponseTransactor(address common.Address, transactor bind.ContractTransactor) (*ResponseTransactor, error) {
	contract, err := bindResponse(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResponseTransactor{contract: contract}, nil
}

// NewResponseFilterer creates a new log filterer instance of Response, bound to a specific deployed contract.
func NewResponseFilterer(address common.Address, filterer bind.ContractFilterer) (*ResponseFilterer, error) {
	contract, err := bindResponse(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResponseFilterer{contract: contract}, nil
}

// bindResponse binds a generic wrapper to an already deployed contract.
func bindResponse(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ResponseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Response *ResponseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Response.Contract.ResponseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Response *ResponseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Response.Contract.ResponseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Response *ResponseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Response.Contract.ResponseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Response *ResponseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Response.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Response *ResponseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Response.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Response *ResponseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Response.Contract.contract.Transact(opts, method, params...)
}

// NewestResponse is a free data retrieval call binding the contract method 0xa4839fb8.
//
// Solidity: function newestResponse(address ) view returns(string userId, string dataHash, string dataSig, bytes data, bool encryptRequired)
func (_Response *ResponseCaller) NewestResponse(opts *bind.CallOpts, arg0 common.Address) (struct {
	UserId          string
	DataHash        string
	DataSig         string
	Data            []byte
	EncryptRequired bool
}, error) {
	var out []interface{}
	err := _Response.contract.Call(opts, &out, "newestResponse", arg0)

	outstruct := new(struct {
		UserId          string
		DataHash        string
		DataSig         string
		Data            []byte
		EncryptRequired bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DataHash = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.DataSig = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Data = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.EncryptRequired = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// NewestResponse is a free data retrieval call binding the contract method 0xa4839fb8.
//
// Solidity: function newestResponse(address ) view returns(string userId, string dataHash, string dataSig, bytes data, bool encryptRequired)
func (_Response *ResponseSession) NewestResponse(arg0 common.Address) (struct {
	UserId          string
	DataHash        string
	DataSig         string
	Data            []byte
	EncryptRequired bool
}, error) {
	return _Response.Contract.NewestResponse(&_Response.CallOpts, arg0)
}

// NewestResponse is a free data retrieval call binding the contract method 0xa4839fb8.
//
// Solidity: function newestResponse(address ) view returns(string userId, string dataHash, string dataSig, bytes data, bool encryptRequired)
func (_Response *ResponseCallerSession) NewestResponse(arg0 common.Address) (struct {
	UserId          string
	DataHash        string
	DataSig         string
	Data            []byte
	EncryptRequired bool
}, error) {
	return _Response.Contract.NewestResponse(&_Response.CallOpts, arg0)
}

// ResponseToUser is a paid mutator transaction binding the contract method 0x3771e25f.
//
// Solidity: function responseToUser(address to, string userId, string dataHash, string dataSig, bytes data, bool encryptRequired) returns()
func (_Response *ResponseTransactor) ResponseToUser(opts *bind.TransactOpts, to common.Address, userId string, dataHash string, dataSig string, data []byte, encryptRequired bool) (*types.Transaction, error) {
	return _Response.contract.Transact(opts, "responseToUser", to, userId, dataHash, dataSig, data, encryptRequired)
}

// ResponseToUser is a paid mutator transaction binding the contract method 0x3771e25f.
//
// Solidity: function responseToUser(address to, string userId, string dataHash, string dataSig, bytes data, bool encryptRequired) returns()
func (_Response *ResponseSession) ResponseToUser(to common.Address, userId string, dataHash string, dataSig string, data []byte, encryptRequired bool) (*types.Transaction, error) {
	return _Response.Contract.ResponseToUser(&_Response.TransactOpts, to, userId, dataHash, dataSig, data, encryptRequired)
}

// ResponseToUser is a paid mutator transaction binding the contract method 0x3771e25f.
//
// Solidity: function responseToUser(address to, string userId, string dataHash, string dataSig, bytes data, bool encryptRequired) returns()
func (_Response *ResponseTransactorSession) ResponseToUser(to common.Address, userId string, dataHash string, dataSig string, data []byte, encryptRequired bool) (*types.Transaction, error) {
	return _Response.Contract.ResponseToUser(&_Response.TransactOpts, to, userId, dataHash, dataSig, data, encryptRequired)
}

// ResponseNewResponseIterator is returned from FilterNewResponse and is used to iterate over the raw logs and unpacked data for NewResponse events raised by the Response contract.
type ResponseNewResponseIterator struct {
	Event *ResponseNewResponse // Event containing the contract specifics and raw log

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
func (it *ResponseNewResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResponseNewResponse)
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
		it.Event = new(ResponseNewResponse)
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
func (it *ResponseNewResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResponseNewResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResponseNewResponse represents a NewResponse event raised by the Response contract.
type ResponseNewResponse struct {
	From            common.Address
	To              common.Address
	UserId          string
	DataHash        string
	DataSig         string
	Data            []byte
	EncryptRequired bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewResponse is a free log retrieval operation binding the contract event 0xd970b922793008b27c4e128de0e9f5cde8460d63078a633d9963e43713795330.
//
// Solidity: event NewResponse(address from, address to, string userId, string dataHash, string dataSig, bytes data, bool encryptRequired)
func (_Response *ResponseFilterer) FilterNewResponse(opts *bind.FilterOpts) (*ResponseNewResponseIterator, error) {

	logs, sub, err := _Response.contract.FilterLogs(opts, "NewResponse")
	if err != nil {
		return nil, err
	}
	return &ResponseNewResponseIterator{contract: _Response.contract, event: "NewResponse", logs: logs, sub: sub}, nil
}

// WatchNewResponse is a free log subscription operation binding the contract event 0xd970b922793008b27c4e128de0e9f5cde8460d63078a633d9963e43713795330.
//
// Solidity: event NewResponse(address from, address to, string userId, string dataHash, string dataSig, bytes data, bool encryptRequired)
func (_Response *ResponseFilterer) WatchNewResponse(opts *bind.WatchOpts, sink chan<- *ResponseNewResponse) (event.Subscription, error) {

	logs, sub, err := _Response.contract.WatchLogs(opts, "NewResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResponseNewResponse)
				if err := _Response.contract.UnpackLog(event, "NewResponse", log); err != nil {
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

// ParseNewResponse is a log parse operation binding the contract event 0xd970b922793008b27c4e128de0e9f5cde8460d63078a633d9963e43713795330.
//
// Solidity: event NewResponse(address from, address to, string userId, string dataHash, string dataSig, bytes data, bool encryptRequired)
func (_Response *ResponseFilterer) ParseNewResponse(log types.Log) (*ResponseNewResponse, error) {
	event := new(ResponseNewResponse)
	if err := _Response.contract.UnpackLog(event, "NewResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

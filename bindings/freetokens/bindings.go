// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package freetokens

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FreetokensABI is the input ABI used to generate the binding from.
const FreetokensABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amunt\",\"type\":\"uint256\"}],\"name\":\"getFreeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Freetokens is an auto generated Go binding around an Ethereum contract.
type Freetokens struct {
	FreetokensCaller     // Read-only binding to the contract
	FreetokensTransactor // Write-only binding to the contract
	FreetokensFilterer   // Log filterer for contract events
}

// FreetokensCaller is an auto generated read-only Go binding around an Ethereum contract.
type FreetokensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreetokensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FreetokensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreetokensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FreetokensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreetokensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FreetokensSession struct {
	Contract     *Freetokens       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FreetokensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FreetokensCallerSession struct {
	Contract *FreetokensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FreetokensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FreetokensTransactorSession struct {
	Contract     *FreetokensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FreetokensRaw is an auto generated low-level Go binding around an Ethereum contract.
type FreetokensRaw struct {
	Contract *Freetokens // Generic contract binding to access the raw methods on
}

// FreetokensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FreetokensCallerRaw struct {
	Contract *FreetokensCaller // Generic read-only contract binding to access the raw methods on
}

// FreetokensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FreetokensTransactorRaw struct {
	Contract *FreetokensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFreetokens creates a new instance of Freetokens, bound to a specific deployed contract.
func NewFreetokens(address common.Address, backend bind.ContractBackend) (*Freetokens, error) {
	contract, err := bindFreetokens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Freetokens{FreetokensCaller: FreetokensCaller{contract: contract}, FreetokensTransactor: FreetokensTransactor{contract: contract}, FreetokensFilterer: FreetokensFilterer{contract: contract}}, nil
}

// NewFreetokensCaller creates a new read-only instance of Freetokens, bound to a specific deployed contract.
func NewFreetokensCaller(address common.Address, caller bind.ContractCaller) (*FreetokensCaller, error) {
	contract, err := bindFreetokens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FreetokensCaller{contract: contract}, nil
}

// NewFreetokensTransactor creates a new write-only instance of Freetokens, bound to a specific deployed contract.
func NewFreetokensTransactor(address common.Address, transactor bind.ContractTransactor) (*FreetokensTransactor, error) {
	contract, err := bindFreetokens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FreetokensTransactor{contract: contract}, nil
}

// NewFreetokensFilterer creates a new log filterer instance of Freetokens, bound to a specific deployed contract.
func NewFreetokensFilterer(address common.Address, filterer bind.ContractFilterer) (*FreetokensFilterer, error) {
	contract, err := bindFreetokens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FreetokensFilterer{contract: contract}, nil
}

// bindFreetokens binds a generic wrapper to an already deployed contract.
func bindFreetokens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FreetokensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Freetokens *FreetokensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Freetokens.Contract.FreetokensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Freetokens *FreetokensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Freetokens.Contract.FreetokensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Freetokens *FreetokensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Freetokens.Contract.FreetokensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Freetokens *FreetokensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Freetokens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Freetokens *FreetokensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Freetokens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Freetokens *FreetokensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Freetokens.Contract.contract.Transact(opts, method, params...)
}

// GetFreeTokens is a paid mutator transaction binding the contract method 0xccddff47.
//
// Solidity: function getFreeTokens(address recipient, uint256 amunt) returns()
func (_Freetokens *FreetokensTransactor) GetFreeTokens(opts *bind.TransactOpts, recipient common.Address, amunt *big.Int) (*types.Transaction, error) {
	return _Freetokens.contract.Transact(opts, "getFreeTokens", recipient, amunt)
}

// GetFreeTokens is a paid mutator transaction binding the contract method 0xccddff47.
//
// Solidity: function getFreeTokens(address recipient, uint256 amunt) returns()
func (_Freetokens *FreetokensSession) GetFreeTokens(recipient common.Address, amunt *big.Int) (*types.Transaction, error) {
	return _Freetokens.Contract.GetFreeTokens(&_Freetokens.TransactOpts, recipient, amunt)
}

// GetFreeTokens is a paid mutator transaction binding the contract method 0xccddff47.
//
// Solidity: function getFreeTokens(address recipient, uint256 amunt) returns()
func (_Freetokens *FreetokensTransactorSession) GetFreeTokens(recipient common.Address, amunt *big.Int) (*types.Transaction, error) {
	return _Freetokens.Contract.GetFreeTokens(&_Freetokens.TransactOpts, recipient, amunt)
}

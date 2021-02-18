// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package controller

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

// ControllerABI is the input ABI used to generate the binding from.
const ControllerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"CategoryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"CategorySorted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"}],\"name\":\"NewPoolInitializer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"unboundTokenSeller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"}],\"name\":\"PoolInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"addTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"categoryIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"computeAverageMarketCap\",\"outputs\":[{\"internalType\":\"uint144\",\"name\":\"\",\"type\":\"uint144\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"computeAverageMarketCaps\",\"outputs\":[{\"internalType\":\"uint144[]\",\"name\":\"\",\"type\":\"uint144[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"createCategory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"getCategoryMarketCaps\",\"outputs\":[{\"internalType\":\"uint144[]\",\"name\":\"\",\"type\":\"uint144[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"getCategoryTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"getLastCategoryUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"getTopCategoryTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"hasCategory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenInCategory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIIndexedUniswapV2Oracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"orderCategoryTokensByMarketCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"updateCategoryPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// CategoryIndex is a free data retrieval call binding the contract method 0xea99fc04.
//
// Solidity: function categoryIndex() view returns(uint256)
func (_Controller *ControllerCaller) CategoryIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "categoryIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CategoryIndex is a free data retrieval call binding the contract method 0xea99fc04.
//
// Solidity: function categoryIndex() view returns(uint256)
func (_Controller *ControllerSession) CategoryIndex() (*big.Int, error) {
	return _Controller.Contract.CategoryIndex(&_Controller.CallOpts)
}

// CategoryIndex is a free data retrieval call binding the contract method 0xea99fc04.
//
// Solidity: function categoryIndex() view returns(uint256)
func (_Controller *ControllerCallerSession) CategoryIndex() (*big.Int, error) {
	return _Controller.Contract.CategoryIndex(&_Controller.CallOpts)
}

// ComputeAverageMarketCap is a free data retrieval call binding the contract method 0xfd371138.
//
// Solidity: function computeAverageMarketCap(address token) view returns(uint144)
func (_Controller *ControllerCaller) ComputeAverageMarketCap(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "computeAverageMarketCap", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeAverageMarketCap is a free data retrieval call binding the contract method 0xfd371138.
//
// Solidity: function computeAverageMarketCap(address token) view returns(uint144)
func (_Controller *ControllerSession) ComputeAverageMarketCap(token common.Address) (*big.Int, error) {
	return _Controller.Contract.ComputeAverageMarketCap(&_Controller.CallOpts, token)
}

// ComputeAverageMarketCap is a free data retrieval call binding the contract method 0xfd371138.
//
// Solidity: function computeAverageMarketCap(address token) view returns(uint144)
func (_Controller *ControllerCallerSession) ComputeAverageMarketCap(token common.Address) (*big.Int, error) {
	return _Controller.Contract.ComputeAverageMarketCap(&_Controller.CallOpts, token)
}

// ComputeAverageMarketCaps is a free data retrieval call binding the contract method 0x6bfc3d40.
//
// Solidity: function computeAverageMarketCaps(address[] tokens) view returns(uint144[])
func (_Controller *ControllerCaller) ComputeAverageMarketCaps(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "computeAverageMarketCaps", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ComputeAverageMarketCaps is a free data retrieval call binding the contract method 0x6bfc3d40.
//
// Solidity: function computeAverageMarketCaps(address[] tokens) view returns(uint144[])
func (_Controller *ControllerSession) ComputeAverageMarketCaps(tokens []common.Address) ([]*big.Int, error) {
	return _Controller.Contract.ComputeAverageMarketCaps(&_Controller.CallOpts, tokens)
}

// ComputeAverageMarketCaps is a free data retrieval call binding the contract method 0x6bfc3d40.
//
// Solidity: function computeAverageMarketCaps(address[] tokens) view returns(uint144[])
func (_Controller *ControllerCallerSession) ComputeAverageMarketCaps(tokens []common.Address) ([]*big.Int, error) {
	return _Controller.Contract.ComputeAverageMarketCaps(&_Controller.CallOpts, tokens)
}

// GetCategoryMarketCaps is a free data retrieval call binding the contract method 0xf3cd354f.
//
// Solidity: function getCategoryMarketCaps(uint256 categoryID) view returns(uint144[])
func (_Controller *ControllerCaller) GetCategoryMarketCaps(opts *bind.CallOpts, categoryID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getCategoryMarketCaps", categoryID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCategoryMarketCaps is a free data retrieval call binding the contract method 0xf3cd354f.
//
// Solidity: function getCategoryMarketCaps(uint256 categoryID) view returns(uint144[])
func (_Controller *ControllerSession) GetCategoryMarketCaps(categoryID *big.Int) ([]*big.Int, error) {
	return _Controller.Contract.GetCategoryMarketCaps(&_Controller.CallOpts, categoryID)
}

// GetCategoryMarketCaps is a free data retrieval call binding the contract method 0xf3cd354f.
//
// Solidity: function getCategoryMarketCaps(uint256 categoryID) view returns(uint144[])
func (_Controller *ControllerCallerSession) GetCategoryMarketCaps(categoryID *big.Int) ([]*big.Int, error) {
	return _Controller.Contract.GetCategoryMarketCaps(&_Controller.CallOpts, categoryID)
}

// GetCategoryTokens is a free data retrieval call binding the contract method 0xde105dbd.
//
// Solidity: function getCategoryTokens(uint256 categoryID) view returns(address[])
func (_Controller *ControllerCaller) GetCategoryTokens(opts *bind.CallOpts, categoryID *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getCategoryTokens", categoryID)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCategoryTokens is a free data retrieval call binding the contract method 0xde105dbd.
//
// Solidity: function getCategoryTokens(uint256 categoryID) view returns(address[])
func (_Controller *ControllerSession) GetCategoryTokens(categoryID *big.Int) ([]common.Address, error) {
	return _Controller.Contract.GetCategoryTokens(&_Controller.CallOpts, categoryID)
}

// GetCategoryTokens is a free data retrieval call binding the contract method 0xde105dbd.
//
// Solidity: function getCategoryTokens(uint256 categoryID) view returns(address[])
func (_Controller *ControllerCallerSession) GetCategoryTokens(categoryID *big.Int) ([]common.Address, error) {
	return _Controller.Contract.GetCategoryTokens(&_Controller.CallOpts, categoryID)
}

// GetLastCategoryUpdate is a free data retrieval call binding the contract method 0xf256a9d0.
//
// Solidity: function getLastCategoryUpdate(uint256 categoryID) view returns(uint256)
func (_Controller *ControllerCaller) GetLastCategoryUpdate(opts *bind.CallOpts, categoryID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getLastCategoryUpdate", categoryID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCategoryUpdate is a free data retrieval call binding the contract method 0xf256a9d0.
//
// Solidity: function getLastCategoryUpdate(uint256 categoryID) view returns(uint256)
func (_Controller *ControllerSession) GetLastCategoryUpdate(categoryID *big.Int) (*big.Int, error) {
	return _Controller.Contract.GetLastCategoryUpdate(&_Controller.CallOpts, categoryID)
}

// GetLastCategoryUpdate is a free data retrieval call binding the contract method 0xf256a9d0.
//
// Solidity: function getLastCategoryUpdate(uint256 categoryID) view returns(uint256)
func (_Controller *ControllerCallerSession) GetLastCategoryUpdate(categoryID *big.Int) (*big.Int, error) {
	return _Controller.Contract.GetLastCategoryUpdate(&_Controller.CallOpts, categoryID)
}

// GetTopCategoryTokens is a free data retrieval call binding the contract method 0x030fb7f4.
//
// Solidity: function getTopCategoryTokens(uint256 categoryID, uint256 num) view returns(address[])
func (_Controller *ControllerCaller) GetTopCategoryTokens(opts *bind.CallOpts, categoryID *big.Int, num *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getTopCategoryTokens", categoryID, num)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTopCategoryTokens is a free data retrieval call binding the contract method 0x030fb7f4.
//
// Solidity: function getTopCategoryTokens(uint256 categoryID, uint256 num) view returns(address[])
func (_Controller *ControllerSession) GetTopCategoryTokens(categoryID *big.Int, num *big.Int) ([]common.Address, error) {
	return _Controller.Contract.GetTopCategoryTokens(&_Controller.CallOpts, categoryID, num)
}

// GetTopCategoryTokens is a free data retrieval call binding the contract method 0x030fb7f4.
//
// Solidity: function getTopCategoryTokens(uint256 categoryID, uint256 num) view returns(address[])
func (_Controller *ControllerCallerSession) GetTopCategoryTokens(categoryID *big.Int, num *big.Int) ([]common.Address, error) {
	return _Controller.Contract.GetTopCategoryTokens(&_Controller.CallOpts, categoryID, num)
}

// HasCategory is a free data retrieval call binding the contract method 0x1ab66317.
//
// Solidity: function hasCategory(uint256 categoryID) view returns(bool)
func (_Controller *ControllerCaller) HasCategory(opts *bind.CallOpts, categoryID *big.Int) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "hasCategory", categoryID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasCategory is a free data retrieval call binding the contract method 0x1ab66317.
//
// Solidity: function hasCategory(uint256 categoryID) view returns(bool)
func (_Controller *ControllerSession) HasCategory(categoryID *big.Int) (bool, error) {
	return _Controller.Contract.HasCategory(&_Controller.CallOpts, categoryID)
}

// HasCategory is a free data retrieval call binding the contract method 0x1ab66317.
//
// Solidity: function hasCategory(uint256 categoryID) view returns(bool)
func (_Controller *ControllerCallerSession) HasCategory(categoryID *big.Int) (bool, error) {
	return _Controller.Contract.HasCategory(&_Controller.CallOpts, categoryID)
}

// IsTokenInCategory is a free data retrieval call binding the contract method 0x1fbaf275.
//
// Solidity: function isTokenInCategory(uint256 categoryID, address token) view returns(bool)
func (_Controller *ControllerCaller) IsTokenInCategory(opts *bind.CallOpts, categoryID *big.Int, token common.Address) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "isTokenInCategory", categoryID, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenInCategory is a free data retrieval call binding the contract method 0x1fbaf275.
//
// Solidity: function isTokenInCategory(uint256 categoryID, address token) view returns(bool)
func (_Controller *ControllerSession) IsTokenInCategory(categoryID *big.Int, token common.Address) (bool, error) {
	return _Controller.Contract.IsTokenInCategory(&_Controller.CallOpts, categoryID, token)
}

// IsTokenInCategory is a free data retrieval call binding the contract method 0x1fbaf275.
//
// Solidity: function isTokenInCategory(uint256 categoryID, address token) view returns(bool)
func (_Controller *ControllerCallerSession) IsTokenInCategory(categoryID *big.Int, token common.Address) (bool, error) {
	return _Controller.Contract.IsTokenInCategory(&_Controller.CallOpts, categoryID, token)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Controller *ControllerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Controller *ControllerSession) Oracle() (common.Address, error) {
	return _Controller.Contract.Oracle(&_Controller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Controller *ControllerCallerSession) Oracle() (common.Address, error) {
	return _Controller.Contract.Oracle(&_Controller.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0x57016b0a.
//
// Solidity: function addToken(uint256 categoryID, address token) returns()
func (_Controller *ControllerTransactor) AddToken(opts *bind.TransactOpts, categoryID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addToken", categoryID, token)
}

// AddToken is a paid mutator transaction binding the contract method 0x57016b0a.
//
// Solidity: function addToken(uint256 categoryID, address token) returns()
func (_Controller *ControllerSession) AddToken(categoryID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddToken(&_Controller.TransactOpts, categoryID, token)
}

// AddToken is a paid mutator transaction binding the contract method 0x57016b0a.
//
// Solidity: function addToken(uint256 categoryID, address token) returns()
func (_Controller *ControllerTransactorSession) AddToken(categoryID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddToken(&_Controller.TransactOpts, categoryID, token)
}

// AddTokens is a paid mutator transaction binding the contract method 0x6b5b8808.
//
// Solidity: function addTokens(uint256 categoryID, address[] tokens) returns()
func (_Controller *ControllerTransactor) AddTokens(opts *bind.TransactOpts, categoryID *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addTokens", categoryID, tokens)
}

// AddTokens is a paid mutator transaction binding the contract method 0x6b5b8808.
//
// Solidity: function addTokens(uint256 categoryID, address[] tokens) returns()
func (_Controller *ControllerSession) AddTokens(categoryID *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddTokens(&_Controller.TransactOpts, categoryID, tokens)
}

// AddTokens is a paid mutator transaction binding the contract method 0x6b5b8808.
//
// Solidity: function addTokens(uint256 categoryID, address[] tokens) returns()
func (_Controller *ControllerTransactorSession) AddTokens(categoryID *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddTokens(&_Controller.TransactOpts, categoryID, tokens)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x6696bf11.
//
// Solidity: function createCategory(bytes32 metadataHash) returns()
func (_Controller *ControllerTransactor) CreateCategory(opts *bind.TransactOpts, metadataHash [32]byte) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "createCategory", metadataHash)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x6696bf11.
//
// Solidity: function createCategory(bytes32 metadataHash) returns()
func (_Controller *ControllerSession) CreateCategory(metadataHash [32]byte) (*types.Transaction, error) {
	return _Controller.Contract.CreateCategory(&_Controller.TransactOpts, metadataHash)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x6696bf11.
//
// Solidity: function createCategory(bytes32 metadataHash) returns()
func (_Controller *ControllerTransactorSession) CreateCategory(metadataHash [32]byte) (*types.Transaction, error) {
	return _Controller.Contract.CreateCategory(&_Controller.TransactOpts, metadataHash)
}

// OrderCategoryTokensByMarketCap is a paid mutator transaction binding the contract method 0xb25f86f9.
//
// Solidity: function orderCategoryTokensByMarketCap(uint256 categoryID) returns()
func (_Controller *ControllerTransactor) OrderCategoryTokensByMarketCap(opts *bind.TransactOpts, categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "orderCategoryTokensByMarketCap", categoryID)
}

// OrderCategoryTokensByMarketCap is a paid mutator transaction binding the contract method 0xb25f86f9.
//
// Solidity: function orderCategoryTokensByMarketCap(uint256 categoryID) returns()
func (_Controller *ControllerSession) OrderCategoryTokensByMarketCap(categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.OrderCategoryTokensByMarketCap(&_Controller.TransactOpts, categoryID)
}

// OrderCategoryTokensByMarketCap is a paid mutator transaction binding the contract method 0xb25f86f9.
//
// Solidity: function orderCategoryTokensByMarketCap(uint256 categoryID) returns()
func (_Controller *ControllerTransactorSession) OrderCategoryTokensByMarketCap(categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.OrderCategoryTokensByMarketCap(&_Controller.TransactOpts, categoryID)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x33fe5676.
//
// Solidity: function removeToken(uint256 categoryID, address token) returns()
func (_Controller *ControllerTransactor) RemoveToken(opts *bind.TransactOpts, categoryID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "removeToken", categoryID, token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x33fe5676.
//
// Solidity: function removeToken(uint256 categoryID, address token) returns()
func (_Controller *ControllerSession) RemoveToken(categoryID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RemoveToken(&_Controller.TransactOpts, categoryID, token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x33fe5676.
//
// Solidity: function removeToken(uint256 categoryID, address token) returns()
func (_Controller *ControllerTransactorSession) RemoveToken(categoryID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RemoveToken(&_Controller.TransactOpts, categoryID, token)
}

// UpdateCategoryPrices is a paid mutator transaction binding the contract method 0x6866ada8.
//
// Solidity: function updateCategoryPrices(uint256 categoryID) returns()
func (_Controller *ControllerTransactor) UpdateCategoryPrices(opts *bind.TransactOpts, categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "updateCategoryPrices", categoryID)
}

// UpdateCategoryPrices is a paid mutator transaction binding the contract method 0x6866ada8.
//
// Solidity: function updateCategoryPrices(uint256 categoryID) returns()
func (_Controller *ControllerSession) UpdateCategoryPrices(categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.UpdateCategoryPrices(&_Controller.TransactOpts, categoryID)
}

// UpdateCategoryPrices is a paid mutator transaction binding the contract method 0x6866ada8.
//
// Solidity: function updateCategoryPrices(uint256 categoryID) returns()
func (_Controller *ControllerTransactorSession) UpdateCategoryPrices(categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.UpdateCategoryPrices(&_Controller.TransactOpts, categoryID)
}

// ControllerCategoryAddedIterator is returned from FilterCategoryAdded and is used to iterate over the raw logs and unpacked data for CategoryAdded events raised by the Controller contract.
type ControllerCategoryAddedIterator struct {
	Event *ControllerCategoryAdded // Event containing the contract specifics and raw log

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
func (it *ControllerCategoryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerCategoryAdded)
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
		it.Event = new(ControllerCategoryAdded)
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
func (it *ControllerCategoryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerCategoryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerCategoryAdded represents a CategoryAdded event raised by the Controller contract.
type ControllerCategoryAdded struct {
	CategoryID   *big.Int
	MetadataHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCategoryAdded is a free log retrieval operation binding the contract event 0x087cd6515ab9d8ad2d847c4b093743cd73acf162a4260bd0ea2429ec6709a632.
//
// Solidity: event CategoryAdded(uint256 categoryID, bytes32 metadataHash)
func (_Controller *ControllerFilterer) FilterCategoryAdded(opts *bind.FilterOpts) (*ControllerCategoryAddedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "CategoryAdded")
	if err != nil {
		return nil, err
	}
	return &ControllerCategoryAddedIterator{contract: _Controller.contract, event: "CategoryAdded", logs: logs, sub: sub}, nil
}

// WatchCategoryAdded is a free log subscription operation binding the contract event 0x087cd6515ab9d8ad2d847c4b093743cd73acf162a4260bd0ea2429ec6709a632.
//
// Solidity: event CategoryAdded(uint256 categoryID, bytes32 metadataHash)
func (_Controller *ControllerFilterer) WatchCategoryAdded(opts *bind.WatchOpts, sink chan<- *ControllerCategoryAdded) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "CategoryAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerCategoryAdded)
				if err := _Controller.contract.UnpackLog(event, "CategoryAdded", log); err != nil {
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

// ParseCategoryAdded is a log parse operation binding the contract event 0x087cd6515ab9d8ad2d847c4b093743cd73acf162a4260bd0ea2429ec6709a632.
//
// Solidity: event CategoryAdded(uint256 categoryID, bytes32 metadataHash)
func (_Controller *ControllerFilterer) ParseCategoryAdded(log types.Log) (*ControllerCategoryAdded, error) {
	event := new(ControllerCategoryAdded)
	if err := _Controller.contract.UnpackLog(event, "CategoryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerCategorySortedIterator is returned from FilterCategorySorted and is used to iterate over the raw logs and unpacked data for CategorySorted events raised by the Controller contract.
type ControllerCategorySortedIterator struct {
	Event *ControllerCategorySorted // Event containing the contract specifics and raw log

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
func (it *ControllerCategorySortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerCategorySorted)
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
		it.Event = new(ControllerCategorySorted)
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
func (it *ControllerCategorySortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerCategorySortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerCategorySorted represents a CategorySorted event raised by the Controller contract.
type ControllerCategorySorted struct {
	CategoryID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCategorySorted is a free log retrieval operation binding the contract event 0x60d2f5fd1812c906f738651cc42bfcb9c52908b88265a1ec508ce32367af72c3.
//
// Solidity: event CategorySorted(uint256 categoryID)
func (_Controller *ControllerFilterer) FilterCategorySorted(opts *bind.FilterOpts) (*ControllerCategorySortedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "CategorySorted")
	if err != nil {
		return nil, err
	}
	return &ControllerCategorySortedIterator{contract: _Controller.contract, event: "CategorySorted", logs: logs, sub: sub}, nil
}

// WatchCategorySorted is a free log subscription operation binding the contract event 0x60d2f5fd1812c906f738651cc42bfcb9c52908b88265a1ec508ce32367af72c3.
//
// Solidity: event CategorySorted(uint256 categoryID)
func (_Controller *ControllerFilterer) WatchCategorySorted(opts *bind.WatchOpts, sink chan<- *ControllerCategorySorted) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "CategorySorted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerCategorySorted)
				if err := _Controller.contract.UnpackLog(event, "CategorySorted", log); err != nil {
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

// ParseCategorySorted is a log parse operation binding the contract event 0x60d2f5fd1812c906f738651cc42bfcb9c52908b88265a1ec508ce32367af72c3.
//
// Solidity: event CategorySorted(uint256 categoryID)
func (_Controller *ControllerFilterer) ParseCategorySorted(log types.Log) (*ControllerCategorySorted, error) {
	event := new(ControllerCategorySorted)
	if err := _Controller.contract.UnpackLog(event, "CategorySorted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerNewPoolInitializerIterator is returned from FilterNewPoolInitializer and is used to iterate over the raw logs and unpacked data for NewPoolInitializer events raised by the Controller contract.
type ControllerNewPoolInitializerIterator struct {
	Event *ControllerNewPoolInitializer // Event containing the contract specifics and raw log

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
func (it *ControllerNewPoolInitializerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerNewPoolInitializer)
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
		it.Event = new(ControllerNewPoolInitializer)
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
func (it *ControllerNewPoolInitializerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerNewPoolInitializerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerNewPoolInitializer represents a NewPoolInitializer event raised by the Controller contract.
type ControllerNewPoolInitializer struct {
	Pool        common.Address
	Initializer common.Address
	CategoryID  *big.Int
	IndexSize   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewPoolInitializer is a free log retrieval operation binding the contract event 0x7ad23833dba658b2bdc6f260fb60a3240b3868086ad60b4e42276f5eeba73e6a.
//
// Solidity: event NewPoolInitializer(address pool, address initializer, uint256 categoryID, uint256 indexSize)
func (_Controller *ControllerFilterer) FilterNewPoolInitializer(opts *bind.FilterOpts) (*ControllerNewPoolInitializerIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NewPoolInitializer")
	if err != nil {
		return nil, err
	}
	return &ControllerNewPoolInitializerIterator{contract: _Controller.contract, event: "NewPoolInitializer", logs: logs, sub: sub}, nil
}

// WatchNewPoolInitializer is a free log subscription operation binding the contract event 0x7ad23833dba658b2bdc6f260fb60a3240b3868086ad60b4e42276f5eeba73e6a.
//
// Solidity: event NewPoolInitializer(address pool, address initializer, uint256 categoryID, uint256 indexSize)
func (_Controller *ControllerFilterer) WatchNewPoolInitializer(opts *bind.WatchOpts, sink chan<- *ControllerNewPoolInitializer) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "NewPoolInitializer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerNewPoolInitializer)
				if err := _Controller.contract.UnpackLog(event, "NewPoolInitializer", log); err != nil {
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

// ParseNewPoolInitializer is a log parse operation binding the contract event 0x7ad23833dba658b2bdc6f260fb60a3240b3868086ad60b4e42276f5eeba73e6a.
//
// Solidity: event NewPoolInitializer(address pool, address initializer, uint256 categoryID, uint256 indexSize)
func (_Controller *ControllerFilterer) ParseNewPoolInitializer(log types.Log) (*ControllerNewPoolInitializer, error) {
	event := new(ControllerNewPoolInitializer)
	if err := _Controller.contract.UnpackLog(event, "NewPoolInitializer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerPoolInitializedIterator is returned from FilterPoolInitialized and is used to iterate over the raw logs and unpacked data for PoolInitialized events raised by the Controller contract.
type ControllerPoolInitializedIterator struct {
	Event *ControllerPoolInitialized // Event containing the contract specifics and raw log

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
func (it *ControllerPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerPoolInitialized)
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
		it.Event = new(ControllerPoolInitialized)
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
func (it *ControllerPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerPoolInitialized represents a PoolInitialized event raised by the Controller contract.
type ControllerPoolInitialized struct {
	Pool               common.Address
	UnboundTokenSeller common.Address
	CategoryID         *big.Int
	IndexSize          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPoolInitialized is a free log retrieval operation binding the contract event 0x9ff2050ac7faae9dc192e1cc9abe73b18a9b849f9b43f509914d80fa104d5903.
//
// Solidity: event PoolInitialized(address pool, address unboundTokenSeller, uint256 categoryID, uint256 indexSize)
func (_Controller *ControllerFilterer) FilterPoolInitialized(opts *bind.FilterOpts) (*ControllerPoolInitializedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "PoolInitialized")
	if err != nil {
		return nil, err
	}
	return &ControllerPoolInitializedIterator{contract: _Controller.contract, event: "PoolInitialized", logs: logs, sub: sub}, nil
}

// WatchPoolInitialized is a free log subscription operation binding the contract event 0x9ff2050ac7faae9dc192e1cc9abe73b18a9b849f9b43f509914d80fa104d5903.
//
// Solidity: event PoolInitialized(address pool, address unboundTokenSeller, uint256 categoryID, uint256 indexSize)
func (_Controller *ControllerFilterer) WatchPoolInitialized(opts *bind.WatchOpts, sink chan<- *ControllerPoolInitialized) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "PoolInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerPoolInitialized)
				if err := _Controller.contract.UnpackLog(event, "PoolInitialized", log); err != nil {
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

// ParsePoolInitialized is a log parse operation binding the contract event 0x9ff2050ac7faae9dc192e1cc9abe73b18a9b849f9b43f509914d80fa104d5903.
//
// Solidity: event PoolInitialized(address pool, address unboundTokenSeller, uint256 categoryID, uint256 indexSize)
func (_Controller *ControllerFilterer) ParsePoolInitialized(log types.Log) (*ControllerPoolInitialized, error) {
	event := new(ControllerPoolInitialized)
	if err := _Controller.contract.UnpackLog(event, "PoolInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the Controller contract.
type ControllerTokenAddedIterator struct {
	Event *ControllerTokenAdded // Event containing the contract specifics and raw log

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
func (it *ControllerTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerTokenAdded)
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
		it.Event = new(ControllerTokenAdded)
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
func (it *ControllerTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerTokenAdded represents a TokenAdded event raised by the Controller contract.
type ControllerTokenAdded struct {
	Token      common.Address
	CategoryID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0xf4c563a3ea86ff1f4275e8c207df0375a51963f2b831b7bf4da8be938d92876c.
//
// Solidity: event TokenAdded(address token, uint256 categoryID)
func (_Controller *ControllerFilterer) FilterTokenAdded(opts *bind.FilterOpts) (*ControllerTokenAddedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return &ControllerTokenAddedIterator{contract: _Controller.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0xf4c563a3ea86ff1f4275e8c207df0375a51963f2b831b7bf4da8be938d92876c.
//
// Solidity: event TokenAdded(address token, uint256 categoryID)
func (_Controller *ControllerFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *ControllerTokenAdded) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerTokenAdded)
				if err := _Controller.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0xf4c563a3ea86ff1f4275e8c207df0375a51963f2b831b7bf4da8be938d92876c.
//
// Solidity: event TokenAdded(address token, uint256 categoryID)
func (_Controller *ControllerFilterer) ParseTokenAdded(log types.Log) (*ControllerTokenAdded, error) {
	event := new(ControllerTokenAdded)
	if err := _Controller.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

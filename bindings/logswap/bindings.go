// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package logswap

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

// LogswapABI is the input ABI used to generate the binding from.
const LogswapABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_SWAP\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_swapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimal\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"outAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inAmount\",\"type\":\"uint256\"}],\"name\":\"emitLogSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getUsedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incRando\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setPublicSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LogswapBin is the compiled bytecode used for deploying new contracts.
var LogswapBin = "0x608060405260016000556040518060400160405280600481526020017f43433130000000000000000000000000000000000000000000000000000000008152506001908051906020019061005492919061008e565b506012600260006101000a81548160ff021916908360ff1602179055506658d15e1762800060035534801561008857600080fd5b50610139565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826100c4576000855561010b565b82601f106100dd57805160ff191683800117855561010b565b8280016001018555821561010b579182015b8281111561010a5782518255916020019190600101906100ef565b5b509050610118919061011c565b5090565b5b8082111561013557600081600090555060010161011d565b5090565b6109e3806101486000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063832781551161008c578063ad4a5c4b11610066578063ad4a5c4b14610316578063cc77828d1461034e578063d4cadf68146103ad578063f8b2cb4f146103cb576100cf565b8063832781551461021d578063948d8ce61461023b57806395d89b4114610293576100cf565b806315e84af9146100d457806318160ddd1461014c57806349b595521461016a5780634aa4e0b51461019a5780635d93a794146101f257806376809ce3146101fc575b600080fd5b610136600480360360408110156100ea57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610423565b6040518082815260200191505060405180910390f35b610154610567565b6040518082815260200191505060405180910390f35b6101986004803603602081101561018057600080fd5b810190808035151590602001909291905050506105d3565b005b6101dc600480360360208110156101b057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105f0565b6040518082815260200191505060405180910390f35b6101fa610684565b005b610204610696565b604051808260ff16815260200191505060405180910390f35b6102256106a9565b6040518082815260200191505060405180910390f35b61027d6004803603602081101561025157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106af565b6040518082815260200191505060405180910390f35b61029b61074d565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102db5780820151818401526020810190506102c0565b50505050905090810190601f1680156103085780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61034c6004803603604081101561032c57600080fd5b8101908080359060200190929190803590602001909291905050506107eb565b005b610356610873565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561039957808201518184015260208101905061037e565b505050509050019250505060405180910390f35b6103b561090f565b6040518082815260200191505060405180910390f35b61040d600480360360208110156103e157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610919565b6040518082815260200191505060405180910390f35b60006003431061043d576701aa535d3d0c00009050610561565b600060fa423360005487604051602001808581526020018473ffffffffffffffffffffffffffffffffffffffff1660601b81526014018381526020018273ffffffffffffffffffffffffffffffffffffffff1660601b81526014019450505050506040516020818303038152906040528051906020012060001c816104be57fe5b069050600060fa423360005487604051602001808581526020018473ffffffffffffffffffffffffffffffffffffffff1660601b81526014018381526020018273ffffffffffffffffffffffffffffffffffffffff1660601b81526014019450505050506040516020818303038152906040528051906020012060001c8161054257fe5b069050670de0b6b3a76400008102670de0b6b3a7640000830201925050505b92915050565b6000804233600054604051602001808481526020018373ffffffffffffffffffffffffffffffffffffffff1660601b815260140182815260200193505050506040516020818303038152906040528051906020012060001c905068056bc75e2d63100000810291505090565b80600460006101000a81548160ff02191690831515021790555050565b600080423360005485604051602001808581526020018473ffffffffffffffffffffffffffffffffffffffff1660601b81526014018381526020018273ffffffffffffffffffffffffffffffffffffffff1660601b81526014019450505050506040516020818303038152906040528051906020012060001c905060006703782dace9d90000820290508092505050919050565b60016000808282540192505081905550565b600260009054906101000a900460ff1681565b60035481565b6000806064423360005486604051602001808581526020018473ffffffffffffffffffffffffffffffffffffffff1660601b81526014018381526020018273ffffffffffffffffffffffffffffffffffffffff1660601b81526014019450505050506040516020818303038152906040528051906020012060001c8161073157fe5b0690506000670de0b6b3a7640000820290508092505050919050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107e35780601f106107b8576101008083540402835291602001916107e3565b820191906000526020600020905b8154815290600101906020018083116107c657829003601f168201915b505050505081565b3373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d433788486604051808381526020018281526020019250505060405180910390a45050565b606080600167ffffffffffffffff8111801561088e57600080fd5b506040519080825280602002602001820160405280156108bd5781602001602082028036833780820191505090505b50905030816000815181106108ce57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508091505090565b6000600354905090565b600080423360005485604051602001808581526020018473ffffffffffffffffffffffffffffffffffffffff1660601b81526014018381526020018273ffffffffffffffffffffffffffffffffffffffff1660601b81526014019450505050506040516020818303038152906040528051906020012060001c90506000678ac7230489e8000082029050809250505091905056fea26469706673582212205841405646a3b04c8a47daedd0dcd74a2fff427160a1539c92a7b2d10b57d51164736f6c63430007040033"

// DeployLogswap deploys a new Ethereum contract, binding an instance of Logswap to it.
func DeployLogswap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Logswap, error) {
	parsed, err := abi.JSON(strings.NewReader(LogswapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LogswapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Logswap{LogswapCaller: LogswapCaller{contract: contract}, LogswapTransactor: LogswapTransactor{contract: contract}, LogswapFilterer: LogswapFilterer{contract: contract}}, nil
}

// Logswap is an auto generated Go binding around an Ethereum contract.
type Logswap struct {
	LogswapCaller     // Read-only binding to the contract
	LogswapTransactor // Write-only binding to the contract
	LogswapFilterer   // Log filterer for contract events
}

// LogswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type LogswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LogswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LogswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LogswapSession struct {
	Contract     *Logswap          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LogswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LogswapCallerSession struct {
	Contract *LogswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LogswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LogswapTransactorSession struct {
	Contract     *LogswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LogswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type LogswapRaw struct {
	Contract *Logswap // Generic contract binding to access the raw methods on
}

// LogswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LogswapCallerRaw struct {
	Contract *LogswapCaller // Generic read-only contract binding to access the raw methods on
}

// LogswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LogswapTransactorRaw struct {
	Contract *LogswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLogswap creates a new instance of Logswap, bound to a specific deployed contract.
func NewLogswap(address common.Address, backend bind.ContractBackend) (*Logswap, error) {
	contract, err := bindLogswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Logswap{LogswapCaller: LogswapCaller{contract: contract}, LogswapTransactor: LogswapTransactor{contract: contract}, LogswapFilterer: LogswapFilterer{contract: contract}}, nil
}

// NewLogswapCaller creates a new read-only instance of Logswap, bound to a specific deployed contract.
func NewLogswapCaller(address common.Address, caller bind.ContractCaller) (*LogswapCaller, error) {
	contract, err := bindLogswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LogswapCaller{contract: contract}, nil
}

// NewLogswapTransactor creates a new write-only instance of Logswap, bound to a specific deployed contract.
func NewLogswapTransactor(address common.Address, transactor bind.ContractTransactor) (*LogswapTransactor, error) {
	contract, err := bindLogswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LogswapTransactor{contract: contract}, nil
}

// NewLogswapFilterer creates a new log filterer instance of Logswap, bound to a specific deployed contract.
func NewLogswapFilterer(address common.Address, filterer bind.ContractFilterer) (*LogswapFilterer, error) {
	contract, err := bindLogswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LogswapFilterer{contract: contract}, nil
}

// bindLogswap binds a generic wrapper to an already deployed contract.
func bindLogswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LogswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logswap *LogswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Logswap.Contract.LogswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logswap *LogswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logswap.Contract.LogswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logswap *LogswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logswap.Contract.LogswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logswap *LogswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Logswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logswap *LogswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logswap *LogswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logswap.Contract.contract.Transact(opts, method, params...)
}

// SwapFee is a free data retrieval call binding the contract method 0x83278155.
//
// Solidity: function _swapFee() view returns(uint256)
func (_Logswap *LogswapCaller) SwapFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Logswap.contract.Call(opts, &out, "_swapFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapFee is a free data retrieval call binding the contract method 0x83278155.
//
// Solidity: function _swapFee() view returns(uint256)
func (_Logswap *LogswapSession) SwapFee() (*big.Int, error) {
	return _Logswap.Contract.SwapFee(&_Logswap.CallOpts)
}

// SwapFee is a free data retrieval call binding the contract method 0x83278155.
//
// Solidity: function _swapFee() view returns(uint256)
func (_Logswap *LogswapCallerSession) SwapFee() (*big.Int, error) {
	return _Logswap.Contract.SwapFee(&_Logswap.CallOpts)
}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() view returns(uint8)
func (_Logswap *LogswapCaller) Decimal(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Logswap.contract.Call(opts, &out, "decimal")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() view returns(uint8)
func (_Logswap *LogswapSession) Decimal() (uint8, error) {
	return _Logswap.Contract.Decimal(&_Logswap.CallOpts)
}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() view returns(uint8)
func (_Logswap *LogswapCallerSession) Decimal() (uint8, error) {
	return _Logswap.Contract.Decimal(&_Logswap.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Logswap *LogswapCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Logswap.contract.Call(opts, &out, "getBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Logswap *LogswapSession) GetBalance(token common.Address) (*big.Int, error) {
	return _Logswap.Contract.GetBalance(&_Logswap.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Logswap *LogswapCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _Logswap.Contract.GetBalance(&_Logswap.CallOpts, token)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[])
func (_Logswap *LogswapCaller) GetCurrentTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Logswap.contract.Call(opts, &out, "getCurrentTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[])
func (_Logswap *LogswapSession) GetCurrentTokens() ([]common.Address, error) {
	return _Logswap.Contract.GetCurrentTokens(&_Logswap.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[])
func (_Logswap *LogswapCallerSession) GetCurrentTokens() ([]common.Address, error) {
	return _Logswap.Contract.GetCurrentTokens(&_Logswap.CallOpts)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Logswap *LogswapCaller) GetDenormalizedWeight(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Logswap.contract.Call(opts, &out, "getDenormalizedWeight", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Logswap *LogswapSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _Logswap.Contract.GetDenormalizedWeight(&_Logswap.CallOpts, token)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Logswap *LogswapCallerSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _Logswap.Contract.GetDenormalizedWeight(&_Logswap.CallOpts, token)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256)
func (_Logswap *LogswapCaller) GetSpotPrice(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Logswap.contract.Call(opts, &out, "getSpotPrice", tokenIn, tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256)
func (_Logswap *LogswapSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Logswap.Contract.GetSpotPrice(&_Logswap.CallOpts, tokenIn, tokenOut)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256)
func (_Logswap *LogswapCallerSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Logswap.Contract.GetSpotPrice(&_Logswap.CallOpts, tokenIn, tokenOut)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Logswap *LogswapCaller) GetSwapFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Logswap.contract.Call(opts, &out, "getSwapFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Logswap *LogswapSession) GetSwapFee() (*big.Int, error) {
	return _Logswap.Contract.GetSwapFee(&_Logswap.CallOpts)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Logswap *LogswapCallerSession) GetSwapFee() (*big.Int, error) {
	return _Logswap.Contract.GetSwapFee(&_Logswap.CallOpts)
}

// GetUsedBalance is a free data retrieval call binding the contract method 0x4aa4e0b5.
//
// Solidity: function getUsedBalance(address token) view returns(uint256)
func (_Logswap *LogswapCaller) GetUsedBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Logswap.contract.Call(opts, &out, "getUsedBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsedBalance is a free data retrieval call binding the contract method 0x4aa4e0b5.
//
// Solidity: function getUsedBalance(address token) view returns(uint256)
func (_Logswap *LogswapSession) GetUsedBalance(token common.Address) (*big.Int, error) {
	return _Logswap.Contract.GetUsedBalance(&_Logswap.CallOpts, token)
}

// GetUsedBalance is a free data retrieval call binding the contract method 0x4aa4e0b5.
//
// Solidity: function getUsedBalance(address token) view returns(uint256)
func (_Logswap *LogswapCallerSession) GetUsedBalance(token common.Address) (*big.Int, error) {
	return _Logswap.Contract.GetUsedBalance(&_Logswap.CallOpts, token)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Logswap *LogswapCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Logswap.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Logswap *LogswapSession) Symbol() (string, error) {
	return _Logswap.Contract.Symbol(&_Logswap.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Logswap *LogswapCallerSession) Symbol() (string, error) {
	return _Logswap.Contract.Symbol(&_Logswap.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Logswap *LogswapCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Logswap.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Logswap *LogswapSession) TotalSupply() (*big.Int, error) {
	return _Logswap.Contract.TotalSupply(&_Logswap.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Logswap *LogswapCallerSession) TotalSupply() (*big.Int, error) {
	return _Logswap.Contract.TotalSupply(&_Logswap.CallOpts)
}

// EmitLogSwap is a paid mutator transaction binding the contract method 0xad4a5c4b.
//
// Solidity: function emitLogSwap(uint256 outAmount, uint256 inAmount) returns()
func (_Logswap *LogswapTransactor) EmitLogSwap(opts *bind.TransactOpts, outAmount *big.Int, inAmount *big.Int) (*types.Transaction, error) {
	return _Logswap.contract.Transact(opts, "emitLogSwap", outAmount, inAmount)
}

// EmitLogSwap is a paid mutator transaction binding the contract method 0xad4a5c4b.
//
// Solidity: function emitLogSwap(uint256 outAmount, uint256 inAmount) returns()
func (_Logswap *LogswapSession) EmitLogSwap(outAmount *big.Int, inAmount *big.Int) (*types.Transaction, error) {
	return _Logswap.Contract.EmitLogSwap(&_Logswap.TransactOpts, outAmount, inAmount)
}

// EmitLogSwap is a paid mutator transaction binding the contract method 0xad4a5c4b.
//
// Solidity: function emitLogSwap(uint256 outAmount, uint256 inAmount) returns()
func (_Logswap *LogswapTransactorSession) EmitLogSwap(outAmount *big.Int, inAmount *big.Int) (*types.Transaction, error) {
	return _Logswap.Contract.EmitLogSwap(&_Logswap.TransactOpts, outAmount, inAmount)
}

// IncRando is a paid mutator transaction binding the contract method 0x5d93a794.
//
// Solidity: function incRando() returns()
func (_Logswap *LogswapTransactor) IncRando(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logswap.contract.Transact(opts, "incRando")
}

// IncRando is a paid mutator transaction binding the contract method 0x5d93a794.
//
// Solidity: function incRando() returns()
func (_Logswap *LogswapSession) IncRando() (*types.Transaction, error) {
	return _Logswap.Contract.IncRando(&_Logswap.TransactOpts)
}

// IncRando is a paid mutator transaction binding the contract method 0x5d93a794.
//
// Solidity: function incRando() returns()
func (_Logswap *LogswapTransactorSession) IncRando() (*types.Transaction, error) {
	return _Logswap.Contract.IncRando(&_Logswap.TransactOpts)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool enabled) returns()
func (_Logswap *LogswapTransactor) SetPublicSwap(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Logswap.contract.Transact(opts, "setPublicSwap", enabled)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool enabled) returns()
func (_Logswap *LogswapSession) SetPublicSwap(enabled bool) (*types.Transaction, error) {
	return _Logswap.Contract.SetPublicSwap(&_Logswap.TransactOpts, enabled)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool enabled) returns()
func (_Logswap *LogswapTransactorSession) SetPublicSwap(enabled bool) (*types.Transaction, error) {
	return _Logswap.Contract.SetPublicSwap(&_Logswap.TransactOpts, enabled)
}

// LogswapLOGSWAPIterator is returned from FilterLOGSWAP and is used to iterate over the raw logs and unpacked data for LOGSWAP events raised by the Logswap contract.
type LogswapLOGSWAPIterator struct {
	Event *LogswapLOGSWAP // Event containing the contract specifics and raw log

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
func (it *LogswapLOGSWAPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogswapLOGSWAP)
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
		it.Event = new(LogswapLOGSWAP)
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
func (it *LogswapLOGSWAPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogswapLOGSWAPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogswapLOGSWAP represents a LOGSWAP event raised by the Logswap contract.
type LogswapLOGSWAP struct {
	Caller         common.Address
	TokenIn        common.Address
	TokenOut       common.Address
	TokenAmountIn  *big.Int
	TokenAmountOut *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGSWAP is a free log retrieval operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_Logswap *LogswapFilterer) FilterLOGSWAP(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (*LogswapLOGSWAPIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Logswap.contract.FilterLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &LogswapLOGSWAPIterator{contract: _Logswap.contract, event: "LOG_SWAP", logs: logs, sub: sub}, nil
}

// WatchLOGSWAP is a free log subscription operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_Logswap *LogswapFilterer) WatchLOGSWAP(opts *bind.WatchOpts, sink chan<- *LogswapLOGSWAP, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Logswap.contract.WatchLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogswapLOGSWAP)
				if err := _Logswap.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
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

// ParseLOGSWAP is a log parse operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_Logswap *LogswapFilterer) ParseLOGSWAP(log types.Log) (*LogswapLOGSWAP, error) {
	event := new(LogswapLOGSWAP)
	if err := _Logswap.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

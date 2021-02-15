// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sigmacore

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

// IIndexPoolRecord is an auto generated low-level Go binding around an user-defined struct.
type IIndexPoolRecord struct {
	Bound            bool
	Ready            bool
	LastDenormUpdate *big.Int
	Denorm           *big.Int
	DesiredDenorm    *big.Int
	Index            uint8
	Balance          *big.Int
}

// SigmacoreABI is the input ABI used to generate the binding from.
const SigmacoreABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDenorm\",\"type\":\"uint256\"}],\"name\":\"LOG_DENORM_UPDATED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"desiredDenorm\",\"type\":\"uint256\"}],\"name\":\"LOG_DESIRED_DENORM_SET\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_EXIT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"name\":\"LOG_JOIN\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxPoolTokens\",\"type\":\"uint256\"}],\"name\":\"LOG_MAX_TOKENS_UPDATED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumBalance\",\"type\":\"uint256\"}],\"name\":\"LOG_MINIMUM_BALANCE_UPDATED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"LOG_PUBLIC_SWAP_TOGGLED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_SWAP\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"LOG_SWAP_FEE_UPDATED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"desiredDenorm\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumBalance\",\"type\":\"uint256\"}],\"name\":\"LOG_TOKEN_ADDED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"LOG_TOKEN_READY\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"LOG_TOKEN_REMOVED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"configure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegateCompLikeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountsOut\",\"type\":\"uint256[]\"}],\"name\":\"exitPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPoolAmountIn\",\"type\":\"uint256\"}],\"name\":\"exitswapExternAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"exitswapPoolAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extrapolatePoolValueFromToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentDesiredTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExitFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getMinimumBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"bound\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"ready\",\"type\":\"bool\"},{\"internalType\":\"uint40\",\"name\":\"lastDenormUpdate\",\"type\":\"uint40\"},{\"internalType\":\"uint96\",\"name\":\"denorm\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"desiredDenorm\",\"type\":\"uint96\"},{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structIIndexPool.Record\",\"name\":\"record\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getUsedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"gulp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint96[]\",\"name\":\"denorms\",\"type\":\"uint96[]\"},{\"internalType\":\"address\",\"name\":\"tokenProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"unbindHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exitFeeRecipient\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"isBound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPublicSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\"}],\"name\":\"joinPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPoolAmountOut\",\"type\":\"uint256\"}],\"name\":\"joinswapExternAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"joinswapPoolAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint96[]\",\"name\":\"desiredDenorms\",\"type\":\"uint96[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minimumBalances\",\"type\":\"uint256[]\"}],\"name\":\"reindexTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint96[]\",\"name\":\"desiredDenorms\",\"type\":\"uint96[]\"}],\"name\":\"reweighTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"exitFeeRecipient\",\"type\":\"address\"}],\"name\":\"setExitFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumBalance\",\"type\":\"uint256\"}],\"name\":\"setMinimumBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setPublicSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Sigmacore is an auto generated Go binding around an Ethereum contract.
type Sigmacore struct {
	SigmacoreCaller     // Read-only binding to the contract
	SigmacoreTransactor // Write-only binding to the contract
	SigmacoreFilterer   // Log filterer for contract events
}

// SigmacoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigmacoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigmacoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigmacoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigmacoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigmacoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigmacoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigmacoreSession struct {
	Contract     *Sigmacore        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigmacoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigmacoreCallerSession struct {
	Contract *SigmacoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SigmacoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigmacoreTransactorSession struct {
	Contract     *SigmacoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SigmacoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigmacoreRaw struct {
	Contract *Sigmacore // Generic contract binding to access the raw methods on
}

// SigmacoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigmacoreCallerRaw struct {
	Contract *SigmacoreCaller // Generic read-only contract binding to access the raw methods on
}

// SigmacoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigmacoreTransactorRaw struct {
	Contract *SigmacoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigmacore creates a new instance of Sigmacore, bound to a specific deployed contract.
func NewSigmacore(address common.Address, backend bind.ContractBackend) (*Sigmacore, error) {
	contract, err := bindSigmacore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sigmacore{SigmacoreCaller: SigmacoreCaller{contract: contract}, SigmacoreTransactor: SigmacoreTransactor{contract: contract}, SigmacoreFilterer: SigmacoreFilterer{contract: contract}}, nil
}

// NewSigmacoreCaller creates a new read-only instance of Sigmacore, bound to a specific deployed contract.
func NewSigmacoreCaller(address common.Address, caller bind.ContractCaller) (*SigmacoreCaller, error) {
	contract, err := bindSigmacore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigmacoreCaller{contract: contract}, nil
}

// NewSigmacoreTransactor creates a new write-only instance of Sigmacore, bound to a specific deployed contract.
func NewSigmacoreTransactor(address common.Address, transactor bind.ContractTransactor) (*SigmacoreTransactor, error) {
	contract, err := bindSigmacore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigmacoreTransactor{contract: contract}, nil
}

// NewSigmacoreFilterer creates a new log filterer instance of Sigmacore, bound to a specific deployed contract.
func NewSigmacoreFilterer(address common.Address, filterer bind.ContractFilterer) (*SigmacoreFilterer, error) {
	contract, err := bindSigmacore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigmacoreFilterer{contract: contract}, nil
}

// bindSigmacore binds a generic wrapper to an already deployed contract.
func bindSigmacore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigmacoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sigmacore *SigmacoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sigmacore.Contract.SigmacoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sigmacore *SigmacoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sigmacore.Contract.SigmacoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sigmacore *SigmacoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sigmacore.Contract.SigmacoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sigmacore *SigmacoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sigmacore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sigmacore *SigmacoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sigmacore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sigmacore *SigmacoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sigmacore.Contract.contract.Transact(opts, method, params...)
}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_Sigmacore *SigmacoreCaller) VERSIONNUMBER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "VERSION_NUMBER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_Sigmacore *SigmacoreSession) VERSIONNUMBER() (*big.Int, error) {
	return _Sigmacore.Contract.VERSIONNUMBER(&_Sigmacore.CallOpts)
}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_Sigmacore *SigmacoreCallerSession) VERSIONNUMBER() (*big.Int, error) {
	return _Sigmacore.Contract.VERSIONNUMBER(&_Sigmacore.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_Sigmacore *SigmacoreCaller) Allowance(opts *bind.CallOpts, src common.Address, dst common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "allowance", src, dst)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_Sigmacore *SigmacoreSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _Sigmacore.Contract.Allowance(&_Sigmacore.CallOpts, src, dst)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_Sigmacore *SigmacoreCallerSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _Sigmacore.Contract.Allowance(&_Sigmacore.CallOpts, src, dst)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_Sigmacore *SigmacoreCaller) BalanceOf(opts *bind.CallOpts, whom common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "balanceOf", whom)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_Sigmacore *SigmacoreSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _Sigmacore.Contract.BalanceOf(&_Sigmacore.CallOpts, whom)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_Sigmacore *SigmacoreCallerSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _Sigmacore.Contract.BalanceOf(&_Sigmacore.CallOpts, whom)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Sigmacore *SigmacoreCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Sigmacore *SigmacoreSession) Decimals() (uint8, error) {
	return _Sigmacore.Contract.Decimals(&_Sigmacore.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Sigmacore *SigmacoreCallerSession) Decimals() (uint8, error) {
	return _Sigmacore.Contract.Decimals(&_Sigmacore.CallOpts)
}

// ExtrapolatePoolValueFromToken is a free data retrieval call binding the contract method 0x98836f08.
//
// Solidity: function extrapolatePoolValueFromToken() view returns(address, uint256)
func (_Sigmacore *SigmacoreCaller) ExtrapolatePoolValueFromToken(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "extrapolatePoolValueFromToken")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ExtrapolatePoolValueFromToken is a free data retrieval call binding the contract method 0x98836f08.
//
// Solidity: function extrapolatePoolValueFromToken() view returns(address, uint256)
func (_Sigmacore *SigmacoreSession) ExtrapolatePoolValueFromToken() (common.Address, *big.Int, error) {
	return _Sigmacore.Contract.ExtrapolatePoolValueFromToken(&_Sigmacore.CallOpts)
}

// ExtrapolatePoolValueFromToken is a free data retrieval call binding the contract method 0x98836f08.
//
// Solidity: function extrapolatePoolValueFromToken() view returns(address, uint256)
func (_Sigmacore *SigmacoreCallerSession) ExtrapolatePoolValueFromToken() (common.Address, *big.Int, error) {
	return _Sigmacore.Contract.ExtrapolatePoolValueFromToken(&_Sigmacore.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Sigmacore *SigmacoreCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "getBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Sigmacore *SigmacoreSession) GetBalance(token common.Address) (*big.Int, error) {
	return _Sigmacore.Contract.GetBalance(&_Sigmacore.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Sigmacore *SigmacoreCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _Sigmacore.Contract.GetBalance(&_Sigmacore.CallOpts, token)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Sigmacore *SigmacoreCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "getController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Sigmacore *SigmacoreSession) GetController() (common.Address, error) {
	return _Sigmacore.Contract.GetController(&_Sigmacore.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Sigmacore *SigmacoreCallerSession) GetController() (common.Address, error) {
	return _Sigmacore.Contract.GetController(&_Sigmacore.CallOpts)
}

// GetCurrentDesiredTokens is a free data retrieval call binding the contract method 0x039209af.
//
// Solidity: function getCurrentDesiredTokens() view returns(address[] tokens)
func (_Sigmacore *SigmacoreCaller) GetCurrentDesiredTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "getCurrentDesiredTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentDesiredTokens is a free data retrieval call binding the contract method 0x039209af.
//
// Solidity: function getCurrentDesiredTokens() view returns(address[] tokens)
func (_Sigmacore *SigmacoreSession) GetCurrentDesiredTokens() ([]common.Address, error) {
	return _Sigmacore.Contract.GetCurrentDesiredTokens(&_Sigmacore.CallOpts)
}

// GetCurrentDesiredTokens is a free data retrieval call binding the contract method 0x039209af.
//
// Solidity: function getCurrentDesiredTokens() view returns(address[] tokens)
func (_Sigmacore *SigmacoreCallerSession) GetCurrentDesiredTokens() ([]common.Address, error) {
	return _Sigmacore.Contract.GetCurrentDesiredTokens(&_Sigmacore.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_Sigmacore *SigmacoreCaller) GetCurrentTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "getCurrentTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_Sigmacore *SigmacoreSession) GetCurrentTokens() ([]common.Address, error) {
	return _Sigmacore.Contract.GetCurrentTokens(&_Sigmacore.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_Sigmacore *SigmacoreCallerSession) GetCurrentTokens() ([]common.Address, error) {
	return _Sigmacore.Contract.GetCurrentTokens(&_Sigmacore.CallOpts)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Sigmacore *SigmacoreCaller) GetDenormalizedWeight(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "getDenormalizedWeight", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Sigmacore *SigmacoreSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _Sigmacore.Contract.GetDenormalizedWeight(&_Sigmacore.CallOpts, token)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Sigmacore *SigmacoreCallerSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _Sigmacore.Contract.GetDenormalizedWeight(&_Sigmacore.CallOpts, token)
}

// GetExitFeeRecipient is a free data retrieval call binding the contract method 0x865bcccb.
//
// Solidity: function getExitFeeRecipient() view returns(address)
func (_Sigmacore *SigmacoreCaller) GetExitFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "getExitFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExitFeeRecipient is a free data retrieval call binding the contract method 0x865bcccb.
//
// Solidity: function getExitFeeRecipient() view returns(address)
func (_Sigmacore *SigmacoreSession) GetExitFeeRecipient() (common.Address, error) {
	return _Sigmacore.Contract.GetExitFeeRecipient(&_Sigmacore.CallOpts)
}

// GetExitFeeRecipient is a free data retrieval call binding the contract method 0x865bcccb.
//
// Solidity: function getExitFeeRecipient() view returns(address)
func (_Sigmacore *SigmacoreCallerSession) GetExitFeeRecipient() (common.Address, error) {
	return _Sigmacore.Contract.GetExitFeeRecipient(&_Sigmacore.CallOpts)
}

// GetMinimumBalance is a free data retrieval call binding the contract method 0x91bfa2bf.
//
// Solidity: function getMinimumBalance(address token) view returns(uint256)
func (_Sigmacore *SigmacoreCaller) GetMinimumBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "getMinimumBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumBalance is a free data retrieval call binding the contract method 0x91bfa2bf.
//
// Solidity: function getMinimumBalance(address token) view returns(uint256)
func (_Sigmacore *SigmacoreSession) GetMinimumBalance(token common.Address) (*big.Int, error) {
	return _Sigmacore.Contract.GetMinimumBalance(&_Sigmacore.CallOpts, token)
}

// GetMinimumBalance is a free data retrieval call binding the contract method 0x91bfa2bf.
//
// Solidity: function getMinimumBalance(address token) view returns(uint256)
func (_Sigmacore *SigmacoreCallerSession) GetMinimumBalance(token common.Address) (*big.Int, error) {
	return _Sigmacore.Contract.GetMinimumBalance(&_Sigmacore.CallOpts, token)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_Sigmacore *SigmacoreCaller) GetNumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "getNumTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_Sigmacore *SigmacoreSession) GetNumTokens() (*big.Int, error) {
	return _Sigmacore.Contract.GetNumTokens(&_Sigmacore.CallOpts)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_Sigmacore *SigmacoreCallerSession) GetNumTokens() (*big.Int, error) {
	return _Sigmacore.Contract.GetNumTokens(&_Sigmacore.CallOpts)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256)
func (_Sigmacore *SigmacoreCaller) GetSpotPrice(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "getSpotPrice", tokenIn, tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256)
func (_Sigmacore *SigmacoreSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Sigmacore.Contract.GetSpotPrice(&_Sigmacore.CallOpts, tokenIn, tokenOut)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256)
func (_Sigmacore *SigmacoreCallerSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Sigmacore.Contract.GetSpotPrice(&_Sigmacore.CallOpts, tokenIn, tokenOut)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Sigmacore *SigmacoreCaller) GetSwapFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "getSwapFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Sigmacore *SigmacoreSession) GetSwapFee() (*big.Int, error) {
	return _Sigmacore.Contract.GetSwapFee(&_Sigmacore.CallOpts)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Sigmacore *SigmacoreCallerSession) GetSwapFee() (*big.Int, error) {
	return _Sigmacore.Contract.GetSwapFee(&_Sigmacore.CallOpts)
}

// GetTokenRecord is a free data retrieval call binding the contract method 0x64c7d661.
//
// Solidity: function getTokenRecord(address token) view returns((bool,bool,uint40,uint96,uint96,uint8,uint256) record)
func (_Sigmacore *SigmacoreCaller) GetTokenRecord(opts *bind.CallOpts, token common.Address) (IIndexPoolRecord, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "getTokenRecord", token)

	if err != nil {
		return *new(IIndexPoolRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(IIndexPoolRecord)).(*IIndexPoolRecord)

	return out0, err

}

// GetTokenRecord is a free data retrieval call binding the contract method 0x64c7d661.
//
// Solidity: function getTokenRecord(address token) view returns((bool,bool,uint40,uint96,uint96,uint8,uint256) record)
func (_Sigmacore *SigmacoreSession) GetTokenRecord(token common.Address) (IIndexPoolRecord, error) {
	return _Sigmacore.Contract.GetTokenRecord(&_Sigmacore.CallOpts, token)
}

// GetTokenRecord is a free data retrieval call binding the contract method 0x64c7d661.
//
// Solidity: function getTokenRecord(address token) view returns((bool,bool,uint40,uint96,uint96,uint8,uint256) record)
func (_Sigmacore *SigmacoreCallerSession) GetTokenRecord(token common.Address) (IIndexPoolRecord, error) {
	return _Sigmacore.Contract.GetTokenRecord(&_Sigmacore.CallOpts, token)
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_Sigmacore *SigmacoreCaller) GetTotalDenormalizedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "getTotalDenormalizedWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_Sigmacore *SigmacoreSession) GetTotalDenormalizedWeight() (*big.Int, error) {
	return _Sigmacore.Contract.GetTotalDenormalizedWeight(&_Sigmacore.CallOpts)
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_Sigmacore *SigmacoreCallerSession) GetTotalDenormalizedWeight() (*big.Int, error) {
	return _Sigmacore.Contract.GetTotalDenormalizedWeight(&_Sigmacore.CallOpts)
}

// GetUsedBalance is a free data retrieval call binding the contract method 0x4aa4e0b5.
//
// Solidity: function getUsedBalance(address token) view returns(uint256)
func (_Sigmacore *SigmacoreCaller) GetUsedBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "getUsedBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsedBalance is a free data retrieval call binding the contract method 0x4aa4e0b5.
//
// Solidity: function getUsedBalance(address token) view returns(uint256)
func (_Sigmacore *SigmacoreSession) GetUsedBalance(token common.Address) (*big.Int, error) {
	return _Sigmacore.Contract.GetUsedBalance(&_Sigmacore.CallOpts, token)
}

// GetUsedBalance is a free data retrieval call binding the contract method 0x4aa4e0b5.
//
// Solidity: function getUsedBalance(address token) view returns(uint256)
func (_Sigmacore *SigmacoreCallerSession) GetUsedBalance(token common.Address) (*big.Int, error) {
	return _Sigmacore.Contract.GetUsedBalance(&_Sigmacore.CallOpts, token)
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_Sigmacore *SigmacoreCaller) IsBound(opts *bind.CallOpts, t common.Address) (bool, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "isBound", t)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_Sigmacore *SigmacoreSession) IsBound(t common.Address) (bool, error) {
	return _Sigmacore.Contract.IsBound(&_Sigmacore.CallOpts, t)
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_Sigmacore *SigmacoreCallerSession) IsBound(t common.Address) (bool, error) {
	return _Sigmacore.Contract.IsBound(&_Sigmacore.CallOpts, t)
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_Sigmacore *SigmacoreCaller) IsPublicSwap(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "isPublicSwap")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_Sigmacore *SigmacoreSession) IsPublicSwap() (bool, error) {
	return _Sigmacore.Contract.IsPublicSwap(&_Sigmacore.CallOpts)
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_Sigmacore *SigmacoreCallerSession) IsPublicSwap() (bool, error) {
	return _Sigmacore.Contract.IsPublicSwap(&_Sigmacore.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Sigmacore *SigmacoreCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Sigmacore *SigmacoreSession) Name() (string, error) {
	return _Sigmacore.Contract.Name(&_Sigmacore.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Sigmacore *SigmacoreCallerSession) Name() (string, error) {
	return _Sigmacore.Contract.Name(&_Sigmacore.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sigmacore *SigmacoreCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sigmacore *SigmacoreSession) Symbol() (string, error) {
	return _Sigmacore.Contract.Symbol(&_Sigmacore.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sigmacore *SigmacoreCallerSession) Symbol() (string, error) {
	return _Sigmacore.Contract.Symbol(&_Sigmacore.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Sigmacore *SigmacoreCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sigmacore.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Sigmacore *SigmacoreSession) TotalSupply() (*big.Int, error) {
	return _Sigmacore.Contract.TotalSupply(&_Sigmacore.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Sigmacore *SigmacoreCallerSession) TotalSupply() (*big.Int, error) {
	return _Sigmacore.Contract.TotalSupply(&_Sigmacore.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreTransactor) Approve(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "approve", dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.Approve(&_Sigmacore.TransactOpts, dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreTransactorSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.Approve(&_Sigmacore.TransactOpts, dst, amt)
}

// Configure is a paid mutator transaction binding the contract method 0x19f0f849.
//
// Solidity: function configure(address controller, string name, string symbol) returns()
func (_Sigmacore *SigmacoreTransactor) Configure(opts *bind.TransactOpts, controller common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "configure", controller, name, symbol)
}

// Configure is a paid mutator transaction binding the contract method 0x19f0f849.
//
// Solidity: function configure(address controller, string name, string symbol) returns()
func (_Sigmacore *SigmacoreSession) Configure(controller common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Sigmacore.Contract.Configure(&_Sigmacore.TransactOpts, controller, name, symbol)
}

// Configure is a paid mutator transaction binding the contract method 0x19f0f849.
//
// Solidity: function configure(address controller, string name, string symbol) returns()
func (_Sigmacore *SigmacoreTransactorSession) Configure(controller common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Sigmacore.Contract.Configure(&_Sigmacore.TransactOpts, controller, name, symbol)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreTransactor) DecreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "decreaseApproval", dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.DecreaseApproval(&_Sigmacore.TransactOpts, dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreTransactorSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.DecreaseApproval(&_Sigmacore.TransactOpts, dst, amt)
}

// DelegateCompLikeToken is a paid mutator transaction binding the contract method 0xf8b5db09.
//
// Solidity: function delegateCompLikeToken(address token, address delegatee) returns()
func (_Sigmacore *SigmacoreTransactor) DelegateCompLikeToken(opts *bind.TransactOpts, token common.Address, delegatee common.Address) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "delegateCompLikeToken", token, delegatee)
}

// DelegateCompLikeToken is a paid mutator transaction binding the contract method 0xf8b5db09.
//
// Solidity: function delegateCompLikeToken(address token, address delegatee) returns()
func (_Sigmacore *SigmacoreSession) DelegateCompLikeToken(token common.Address, delegatee common.Address) (*types.Transaction, error) {
	return _Sigmacore.Contract.DelegateCompLikeToken(&_Sigmacore.TransactOpts, token, delegatee)
}

// DelegateCompLikeToken is a paid mutator transaction binding the contract method 0xf8b5db09.
//
// Solidity: function delegateCompLikeToken(address token, address delegatee) returns()
func (_Sigmacore *SigmacoreTransactorSession) DelegateCompLikeToken(token common.Address, delegatee common.Address) (*types.Transaction, error) {
	return _Sigmacore.Contract.DelegateCompLikeToken(&_Sigmacore.TransactOpts, token, delegatee)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_Sigmacore *SigmacoreTransactor) ExitPool(opts *bind.TransactOpts, poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "exitPool", poolAmountIn, minAmountsOut)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_Sigmacore *SigmacoreSession) ExitPool(poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.ExitPool(&_Sigmacore.TransactOpts, poolAmountIn, minAmountsOut)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_Sigmacore *SigmacoreTransactorSession) ExitPool(poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.ExitPool(&_Sigmacore.TransactOpts, poolAmountIn, minAmountsOut)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256)
func (_Sigmacore *SigmacoreTransactor) ExitswapExternAmountOut(opts *bind.TransactOpts, tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "exitswapExternAmountOut", tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256)
func (_Sigmacore *SigmacoreSession) ExitswapExternAmountOut(tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.ExitswapExternAmountOut(&_Sigmacore.TransactOpts, tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256)
func (_Sigmacore *SigmacoreTransactorSession) ExitswapExternAmountOut(tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.ExitswapExternAmountOut(&_Sigmacore.TransactOpts, tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256)
func (_Sigmacore *SigmacoreTransactor) ExitswapPoolAmountIn(opts *bind.TransactOpts, tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "exitswapPoolAmountIn", tokenOut, poolAmountIn, minAmountOut)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256)
func (_Sigmacore *SigmacoreSession) ExitswapPoolAmountIn(tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.ExitswapPoolAmountIn(&_Sigmacore.TransactOpts, tokenOut, poolAmountIn, minAmountOut)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256)
func (_Sigmacore *SigmacoreTransactorSession) ExitswapPoolAmountIn(tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.ExitswapPoolAmountIn(&_Sigmacore.TransactOpts, tokenOut, poolAmountIn, minAmountOut)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_Sigmacore *SigmacoreTransactor) Gulp(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "gulp", token)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_Sigmacore *SigmacoreSession) Gulp(token common.Address) (*types.Transaction, error) {
	return _Sigmacore.Contract.Gulp(&_Sigmacore.TransactOpts, token)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_Sigmacore *SigmacoreTransactorSession) Gulp(token common.Address) (*types.Transaction, error) {
	return _Sigmacore.Contract.Gulp(&_Sigmacore.TransactOpts, token)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreTransactor) IncreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "increaseApproval", dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.IncreaseApproval(&_Sigmacore.TransactOpts, dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreTransactorSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.IncreaseApproval(&_Sigmacore.TransactOpts, dst, amt)
}

// Initialize is a paid mutator transaction binding the contract method 0x2b110c74.
//
// Solidity: function initialize(address[] tokens, uint256[] balances, uint96[] denorms, address tokenProvider, address unbindHandler, address exitFeeRecipient) returns()
func (_Sigmacore *SigmacoreTransactor) Initialize(opts *bind.TransactOpts, tokens []common.Address, balances []*big.Int, denorms []*big.Int, tokenProvider common.Address, unbindHandler common.Address, exitFeeRecipient common.Address) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "initialize", tokens, balances, denorms, tokenProvider, unbindHandler, exitFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0x2b110c74.
//
// Solidity: function initialize(address[] tokens, uint256[] balances, uint96[] denorms, address tokenProvider, address unbindHandler, address exitFeeRecipient) returns()
func (_Sigmacore *SigmacoreSession) Initialize(tokens []common.Address, balances []*big.Int, denorms []*big.Int, tokenProvider common.Address, unbindHandler common.Address, exitFeeRecipient common.Address) (*types.Transaction, error) {
	return _Sigmacore.Contract.Initialize(&_Sigmacore.TransactOpts, tokens, balances, denorms, tokenProvider, unbindHandler, exitFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0x2b110c74.
//
// Solidity: function initialize(address[] tokens, uint256[] balances, uint96[] denorms, address tokenProvider, address unbindHandler, address exitFeeRecipient) returns()
func (_Sigmacore *SigmacoreTransactorSession) Initialize(tokens []common.Address, balances []*big.Int, denorms []*big.Int, tokenProvider common.Address, unbindHandler common.Address, exitFeeRecipient common.Address) (*types.Transaction, error) {
	return _Sigmacore.Contract.Initialize(&_Sigmacore.TransactOpts, tokens, balances, denorms, tokenProvider, unbindHandler, exitFeeRecipient)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_Sigmacore *SigmacoreTransactor) JoinPool(opts *bind.TransactOpts, poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "joinPool", poolAmountOut, maxAmountsIn)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_Sigmacore *SigmacoreSession) JoinPool(poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.JoinPool(&_Sigmacore.TransactOpts, poolAmountOut, maxAmountsIn)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_Sigmacore *SigmacoreTransactorSession) JoinPool(poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.JoinPool(&_Sigmacore.TransactOpts, poolAmountOut, maxAmountsIn)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256)
func (_Sigmacore *SigmacoreTransactor) JoinswapExternAmountIn(opts *bind.TransactOpts, tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "joinswapExternAmountIn", tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256)
func (_Sigmacore *SigmacoreSession) JoinswapExternAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.JoinswapExternAmountIn(&_Sigmacore.TransactOpts, tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256)
func (_Sigmacore *SigmacoreTransactorSession) JoinswapExternAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.JoinswapExternAmountIn(&_Sigmacore.TransactOpts, tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256)
func (_Sigmacore *SigmacoreTransactor) JoinswapPoolAmountOut(opts *bind.TransactOpts, tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "joinswapPoolAmountOut", tokenIn, poolAmountOut, maxAmountIn)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256)
func (_Sigmacore *SigmacoreSession) JoinswapPoolAmountOut(tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.JoinswapPoolAmountOut(&_Sigmacore.TransactOpts, tokenIn, poolAmountOut, maxAmountIn)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256)
func (_Sigmacore *SigmacoreTransactorSession) JoinswapPoolAmountOut(tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.JoinswapPoolAmountOut(&_Sigmacore.TransactOpts, tokenIn, poolAmountOut, maxAmountIn)
}

// ReindexTokens is a paid mutator transaction binding the contract method 0xc3f46810.
//
// Solidity: function reindexTokens(address[] tokens, uint96[] desiredDenorms, uint256[] minimumBalances) returns()
func (_Sigmacore *SigmacoreTransactor) ReindexTokens(opts *bind.TransactOpts, tokens []common.Address, desiredDenorms []*big.Int, minimumBalances []*big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "reindexTokens", tokens, desiredDenorms, minimumBalances)
}

// ReindexTokens is a paid mutator transaction binding the contract method 0xc3f46810.
//
// Solidity: function reindexTokens(address[] tokens, uint96[] desiredDenorms, uint256[] minimumBalances) returns()
func (_Sigmacore *SigmacoreSession) ReindexTokens(tokens []common.Address, desiredDenorms []*big.Int, minimumBalances []*big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.ReindexTokens(&_Sigmacore.TransactOpts, tokens, desiredDenorms, minimumBalances)
}

// ReindexTokens is a paid mutator transaction binding the contract method 0xc3f46810.
//
// Solidity: function reindexTokens(address[] tokens, uint96[] desiredDenorms, uint256[] minimumBalances) returns()
func (_Sigmacore *SigmacoreTransactorSession) ReindexTokens(tokens []common.Address, desiredDenorms []*big.Int, minimumBalances []*big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.ReindexTokens(&_Sigmacore.TransactOpts, tokens, desiredDenorms, minimumBalances)
}

// ReweighTokens is a paid mutator transaction binding the contract method 0x5d5e8ce7.
//
// Solidity: function reweighTokens(address[] tokens, uint96[] desiredDenorms) returns()
func (_Sigmacore *SigmacoreTransactor) ReweighTokens(opts *bind.TransactOpts, tokens []common.Address, desiredDenorms []*big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "reweighTokens", tokens, desiredDenorms)
}

// ReweighTokens is a paid mutator transaction binding the contract method 0x5d5e8ce7.
//
// Solidity: function reweighTokens(address[] tokens, uint96[] desiredDenorms) returns()
func (_Sigmacore *SigmacoreSession) ReweighTokens(tokens []common.Address, desiredDenorms []*big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.ReweighTokens(&_Sigmacore.TransactOpts, tokens, desiredDenorms)
}

// ReweighTokens is a paid mutator transaction binding the contract method 0x5d5e8ce7.
//
// Solidity: function reweighTokens(address[] tokens, uint96[] desiredDenorms) returns()
func (_Sigmacore *SigmacoreTransactorSession) ReweighTokens(tokens []common.Address, desiredDenorms []*big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.ReweighTokens(&_Sigmacore.TransactOpts, tokens, desiredDenorms)
}

// SetExitFeeRecipient is a paid mutator transaction binding the contract method 0xf9f97c98.
//
// Solidity: function setExitFeeRecipient(address exitFeeRecipient) returns()
func (_Sigmacore *SigmacoreTransactor) SetExitFeeRecipient(opts *bind.TransactOpts, exitFeeRecipient common.Address) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "setExitFeeRecipient", exitFeeRecipient)
}

// SetExitFeeRecipient is a paid mutator transaction binding the contract method 0xf9f97c98.
//
// Solidity: function setExitFeeRecipient(address exitFeeRecipient) returns()
func (_Sigmacore *SigmacoreSession) SetExitFeeRecipient(exitFeeRecipient common.Address) (*types.Transaction, error) {
	return _Sigmacore.Contract.SetExitFeeRecipient(&_Sigmacore.TransactOpts, exitFeeRecipient)
}

// SetExitFeeRecipient is a paid mutator transaction binding the contract method 0xf9f97c98.
//
// Solidity: function setExitFeeRecipient(address exitFeeRecipient) returns()
func (_Sigmacore *SigmacoreTransactorSession) SetExitFeeRecipient(exitFeeRecipient common.Address) (*types.Transaction, error) {
	return _Sigmacore.Contract.SetExitFeeRecipient(&_Sigmacore.TransactOpts, exitFeeRecipient)
}

// SetMinimumBalance is a paid mutator transaction binding the contract method 0xa49c44d7.
//
// Solidity: function setMinimumBalance(address token, uint256 minimumBalance) returns()
func (_Sigmacore *SigmacoreTransactor) SetMinimumBalance(opts *bind.TransactOpts, token common.Address, minimumBalance *big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "setMinimumBalance", token, minimumBalance)
}

// SetMinimumBalance is a paid mutator transaction binding the contract method 0xa49c44d7.
//
// Solidity: function setMinimumBalance(address token, uint256 minimumBalance) returns()
func (_Sigmacore *SigmacoreSession) SetMinimumBalance(token common.Address, minimumBalance *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.SetMinimumBalance(&_Sigmacore.TransactOpts, token, minimumBalance)
}

// SetMinimumBalance is a paid mutator transaction binding the contract method 0xa49c44d7.
//
// Solidity: function setMinimumBalance(address token, uint256 minimumBalance) returns()
func (_Sigmacore *SigmacoreTransactorSession) SetMinimumBalance(token common.Address, minimumBalance *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.SetMinimumBalance(&_Sigmacore.TransactOpts, token, minimumBalance)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool enabled) returns()
func (_Sigmacore *SigmacoreTransactor) SetPublicSwap(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "setPublicSwap", enabled)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool enabled) returns()
func (_Sigmacore *SigmacoreSession) SetPublicSwap(enabled bool) (*types.Transaction, error) {
	return _Sigmacore.Contract.SetPublicSwap(&_Sigmacore.TransactOpts, enabled)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool enabled) returns()
func (_Sigmacore *SigmacoreTransactorSession) SetPublicSwap(enabled bool) (*types.Transaction, error) {
	return _Sigmacore.Contract.SetPublicSwap(&_Sigmacore.TransactOpts, enabled)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_Sigmacore *SigmacoreTransactor) SetSwapFee(opts *bind.TransactOpts, swapFee *big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "setSwapFee", swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_Sigmacore *SigmacoreSession) SetSwapFee(swapFee *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.SetSwapFee(&_Sigmacore.TransactOpts, swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_Sigmacore *SigmacoreTransactorSession) SetSwapFee(swapFee *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.SetSwapFee(&_Sigmacore.TransactOpts, swapFee)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256, uint256)
func (_Sigmacore *SigmacoreTransactor) SwapExactAmountIn(opts *bind.TransactOpts, tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "swapExactAmountIn", tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256, uint256)
func (_Sigmacore *SigmacoreSession) SwapExactAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.SwapExactAmountIn(&_Sigmacore.TransactOpts, tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256, uint256)
func (_Sigmacore *SigmacoreTransactorSession) SwapExactAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.SwapExactAmountIn(&_Sigmacore.TransactOpts, tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256, uint256)
func (_Sigmacore *SigmacoreTransactor) SwapExactAmountOut(opts *bind.TransactOpts, tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "swapExactAmountOut", tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256, uint256)
func (_Sigmacore *SigmacoreSession) SwapExactAmountOut(tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.SwapExactAmountOut(&_Sigmacore.TransactOpts, tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256, uint256)
func (_Sigmacore *SigmacoreTransactorSession) SwapExactAmountOut(tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.SwapExactAmountOut(&_Sigmacore.TransactOpts, tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "transfer", dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.Transfer(&_Sigmacore.TransactOpts, dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreTransactorSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.Transfer(&_Sigmacore.TransactOpts, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.contract.Transact(opts, "transferFrom", src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.TransferFrom(&_Sigmacore.TransactOpts, src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_Sigmacore *SigmacoreTransactorSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Sigmacore.Contract.TransferFrom(&_Sigmacore.TransactOpts, src, dst, amt)
}

// SigmacoreApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Sigmacore contract.
type SigmacoreApprovalIterator struct {
	Event *SigmacoreApproval // Event containing the contract specifics and raw log

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
func (it *SigmacoreApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SigmacoreApproval)
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
		it.Event = new(SigmacoreApproval)
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
func (it *SigmacoreApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SigmacoreApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SigmacoreApproval represents a Approval event raised by the Sigmacore contract.
type SigmacoreApproval struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_Sigmacore *SigmacoreFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*SigmacoreApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Sigmacore.contract.FilterLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &SigmacoreApprovalIterator{contract: _Sigmacore.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_Sigmacore *SigmacoreFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SigmacoreApproval, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Sigmacore.contract.WatchLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SigmacoreApproval)
				if err := _Sigmacore.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_Sigmacore *SigmacoreFilterer) ParseApproval(log types.Log) (*SigmacoreApproval, error) {
	event := new(SigmacoreApproval)
	if err := _Sigmacore.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SigmacoreLOGDENORMUPDATEDIterator is returned from FilterLOGDENORMUPDATED and is used to iterate over the raw logs and unpacked data for LOGDENORMUPDATED events raised by the Sigmacore contract.
type SigmacoreLOGDENORMUPDATEDIterator struct {
	Event *SigmacoreLOGDENORMUPDATED // Event containing the contract specifics and raw log

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
func (it *SigmacoreLOGDENORMUPDATEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SigmacoreLOGDENORMUPDATED)
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
		it.Event = new(SigmacoreLOGDENORMUPDATED)
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
func (it *SigmacoreLOGDENORMUPDATEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SigmacoreLOGDENORMUPDATEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SigmacoreLOGDENORMUPDATED represents a LOGDENORMUPDATED event raised by the Sigmacore contract.
type SigmacoreLOGDENORMUPDATED struct {
	Token     common.Address
	NewDenorm *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLOGDENORMUPDATED is a free log retrieval operation binding the contract event 0x21b12aed5d425f5675450ffeeae01039085e5323974c3099e1828155d9b51e77.
//
// Solidity: event LOG_DENORM_UPDATED(address indexed token, uint256 newDenorm)
func (_Sigmacore *SigmacoreFilterer) FilterLOGDENORMUPDATED(opts *bind.FilterOpts, token []common.Address) (*SigmacoreLOGDENORMUPDATEDIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Sigmacore.contract.FilterLogs(opts, "LOG_DENORM_UPDATED", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SigmacoreLOGDENORMUPDATEDIterator{contract: _Sigmacore.contract, event: "LOG_DENORM_UPDATED", logs: logs, sub: sub}, nil
}

// WatchLOGDENORMUPDATED is a free log subscription operation binding the contract event 0x21b12aed5d425f5675450ffeeae01039085e5323974c3099e1828155d9b51e77.
//
// Solidity: event LOG_DENORM_UPDATED(address indexed token, uint256 newDenorm)
func (_Sigmacore *SigmacoreFilterer) WatchLOGDENORMUPDATED(opts *bind.WatchOpts, sink chan<- *SigmacoreLOGDENORMUPDATED, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Sigmacore.contract.WatchLogs(opts, "LOG_DENORM_UPDATED", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SigmacoreLOGDENORMUPDATED)
				if err := _Sigmacore.contract.UnpackLog(event, "LOG_DENORM_UPDATED", log); err != nil {
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

// ParseLOGDENORMUPDATED is a log parse operation binding the contract event 0x21b12aed5d425f5675450ffeeae01039085e5323974c3099e1828155d9b51e77.
//
// Solidity: event LOG_DENORM_UPDATED(address indexed token, uint256 newDenorm)
func (_Sigmacore *SigmacoreFilterer) ParseLOGDENORMUPDATED(log types.Log) (*SigmacoreLOGDENORMUPDATED, error) {
	event := new(SigmacoreLOGDENORMUPDATED)
	if err := _Sigmacore.contract.UnpackLog(event, "LOG_DENORM_UPDATED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SigmacoreLOGDESIREDDENORMSETIterator is returned from FilterLOGDESIREDDENORMSET and is used to iterate over the raw logs and unpacked data for LOGDESIREDDENORMSET events raised by the Sigmacore contract.
type SigmacoreLOGDESIREDDENORMSETIterator struct {
	Event *SigmacoreLOGDESIREDDENORMSET // Event containing the contract specifics and raw log

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
func (it *SigmacoreLOGDESIREDDENORMSETIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SigmacoreLOGDESIREDDENORMSET)
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
		it.Event = new(SigmacoreLOGDESIREDDENORMSET)
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
func (it *SigmacoreLOGDESIREDDENORMSETIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SigmacoreLOGDESIREDDENORMSETIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SigmacoreLOGDESIREDDENORMSET represents a LOGDESIREDDENORMSET event raised by the Sigmacore contract.
type SigmacoreLOGDESIREDDENORMSET struct {
	Token         common.Address
	DesiredDenorm *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLOGDESIREDDENORMSET is a free log retrieval operation binding the contract event 0xc7ea88f3376e27ce6ebc2025310023327f743a8377d438258c36b166dd8b2983.
//
// Solidity: event LOG_DESIRED_DENORM_SET(address indexed token, uint256 desiredDenorm)
func (_Sigmacore *SigmacoreFilterer) FilterLOGDESIREDDENORMSET(opts *bind.FilterOpts, token []common.Address) (*SigmacoreLOGDESIREDDENORMSETIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Sigmacore.contract.FilterLogs(opts, "LOG_DESIRED_DENORM_SET", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SigmacoreLOGDESIREDDENORMSETIterator{contract: _Sigmacore.contract, event: "LOG_DESIRED_DENORM_SET", logs: logs, sub: sub}, nil
}

// WatchLOGDESIREDDENORMSET is a free log subscription operation binding the contract event 0xc7ea88f3376e27ce6ebc2025310023327f743a8377d438258c36b166dd8b2983.
//
// Solidity: event LOG_DESIRED_DENORM_SET(address indexed token, uint256 desiredDenorm)
func (_Sigmacore *SigmacoreFilterer) WatchLOGDESIREDDENORMSET(opts *bind.WatchOpts, sink chan<- *SigmacoreLOGDESIREDDENORMSET, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Sigmacore.contract.WatchLogs(opts, "LOG_DESIRED_DENORM_SET", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SigmacoreLOGDESIREDDENORMSET)
				if err := _Sigmacore.contract.UnpackLog(event, "LOG_DESIRED_DENORM_SET", log); err != nil {
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

// ParseLOGDESIREDDENORMSET is a log parse operation binding the contract event 0xc7ea88f3376e27ce6ebc2025310023327f743a8377d438258c36b166dd8b2983.
//
// Solidity: event LOG_DESIRED_DENORM_SET(address indexed token, uint256 desiredDenorm)
func (_Sigmacore *SigmacoreFilterer) ParseLOGDESIREDDENORMSET(log types.Log) (*SigmacoreLOGDESIREDDENORMSET, error) {
	event := new(SigmacoreLOGDESIREDDENORMSET)
	if err := _Sigmacore.contract.UnpackLog(event, "LOG_DESIRED_DENORM_SET", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SigmacoreLOGEXITIterator is returned from FilterLOGEXIT and is used to iterate over the raw logs and unpacked data for LOGEXIT events raised by the Sigmacore contract.
type SigmacoreLOGEXITIterator struct {
	Event *SigmacoreLOGEXIT // Event containing the contract specifics and raw log

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
func (it *SigmacoreLOGEXITIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SigmacoreLOGEXIT)
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
		it.Event = new(SigmacoreLOGEXIT)
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
func (it *SigmacoreLOGEXITIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SigmacoreLOGEXITIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SigmacoreLOGEXIT represents a LOGEXIT event raised by the Sigmacore contract.
type SigmacoreLOGEXIT struct {
	Caller         common.Address
	TokenOut       common.Address
	TokenAmountOut *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGEXIT is a free log retrieval operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_Sigmacore *SigmacoreFilterer) FilterLOGEXIT(opts *bind.FilterOpts, caller []common.Address, tokenOut []common.Address) (*SigmacoreLOGEXITIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Sigmacore.contract.FilterLogs(opts, "LOG_EXIT", callerRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &SigmacoreLOGEXITIterator{contract: _Sigmacore.contract, event: "LOG_EXIT", logs: logs, sub: sub}, nil
}

// WatchLOGEXIT is a free log subscription operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_Sigmacore *SigmacoreFilterer) WatchLOGEXIT(opts *bind.WatchOpts, sink chan<- *SigmacoreLOGEXIT, caller []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Sigmacore.contract.WatchLogs(opts, "LOG_EXIT", callerRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SigmacoreLOGEXIT)
				if err := _Sigmacore.contract.UnpackLog(event, "LOG_EXIT", log); err != nil {
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

// ParseLOGEXIT is a log parse operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_Sigmacore *SigmacoreFilterer) ParseLOGEXIT(log types.Log) (*SigmacoreLOGEXIT, error) {
	event := new(SigmacoreLOGEXIT)
	if err := _Sigmacore.contract.UnpackLog(event, "LOG_EXIT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SigmacoreLOGJOINIterator is returned from FilterLOGJOIN and is used to iterate over the raw logs and unpacked data for LOGJOIN events raised by the Sigmacore contract.
type SigmacoreLOGJOINIterator struct {
	Event *SigmacoreLOGJOIN // Event containing the contract specifics and raw log

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
func (it *SigmacoreLOGJOINIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SigmacoreLOGJOIN)
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
		it.Event = new(SigmacoreLOGJOIN)
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
func (it *SigmacoreLOGJOINIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SigmacoreLOGJOINIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SigmacoreLOGJOIN represents a LOGJOIN event raised by the Sigmacore contract.
type SigmacoreLOGJOIN struct {
	Caller        common.Address
	TokenIn       common.Address
	TokenAmountIn *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLOGJOIN is a free log retrieval operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_Sigmacore *SigmacoreFilterer) FilterLOGJOIN(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address) (*SigmacoreLOGJOINIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Sigmacore.contract.FilterLogs(opts, "LOG_JOIN", callerRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return &SigmacoreLOGJOINIterator{contract: _Sigmacore.contract, event: "LOG_JOIN", logs: logs, sub: sub}, nil
}

// WatchLOGJOIN is a free log subscription operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_Sigmacore *SigmacoreFilterer) WatchLOGJOIN(opts *bind.WatchOpts, sink chan<- *SigmacoreLOGJOIN, caller []common.Address, tokenIn []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Sigmacore.contract.WatchLogs(opts, "LOG_JOIN", callerRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SigmacoreLOGJOIN)
				if err := _Sigmacore.contract.UnpackLog(event, "LOG_JOIN", log); err != nil {
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

// ParseLOGJOIN is a log parse operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_Sigmacore *SigmacoreFilterer) ParseLOGJOIN(log types.Log) (*SigmacoreLOGJOIN, error) {
	event := new(SigmacoreLOGJOIN)
	if err := _Sigmacore.contract.UnpackLog(event, "LOG_JOIN", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SigmacoreLOGMAXTOKENSUPDATEDIterator is returned from FilterLOGMAXTOKENSUPDATED and is used to iterate over the raw logs and unpacked data for LOGMAXTOKENSUPDATED events raised by the Sigmacore contract.
type SigmacoreLOGMAXTOKENSUPDATEDIterator struct {
	Event *SigmacoreLOGMAXTOKENSUPDATED // Event containing the contract specifics and raw log

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
func (it *SigmacoreLOGMAXTOKENSUPDATEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SigmacoreLOGMAXTOKENSUPDATED)
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
		it.Event = new(SigmacoreLOGMAXTOKENSUPDATED)
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
func (it *SigmacoreLOGMAXTOKENSUPDATEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SigmacoreLOGMAXTOKENSUPDATEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SigmacoreLOGMAXTOKENSUPDATED represents a LOGMAXTOKENSUPDATED event raised by the Sigmacore contract.
type SigmacoreLOGMAXTOKENSUPDATED struct {
	MaxPoolTokens *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLOGMAXTOKENSUPDATED is a free log retrieval operation binding the contract event 0x65492266ae9a1f46497efddd24b6946862569680a511c543590d4587ca800d05.
//
// Solidity: event LOG_MAX_TOKENS_UPDATED(uint256 maxPoolTokens)
func (_Sigmacore *SigmacoreFilterer) FilterLOGMAXTOKENSUPDATED(opts *bind.FilterOpts) (*SigmacoreLOGMAXTOKENSUPDATEDIterator, error) {

	logs, sub, err := _Sigmacore.contract.FilterLogs(opts, "LOG_MAX_TOKENS_UPDATED")
	if err != nil {
		return nil, err
	}
	return &SigmacoreLOGMAXTOKENSUPDATEDIterator{contract: _Sigmacore.contract, event: "LOG_MAX_TOKENS_UPDATED", logs: logs, sub: sub}, nil
}

// WatchLOGMAXTOKENSUPDATED is a free log subscription operation binding the contract event 0x65492266ae9a1f46497efddd24b6946862569680a511c543590d4587ca800d05.
//
// Solidity: event LOG_MAX_TOKENS_UPDATED(uint256 maxPoolTokens)
func (_Sigmacore *SigmacoreFilterer) WatchLOGMAXTOKENSUPDATED(opts *bind.WatchOpts, sink chan<- *SigmacoreLOGMAXTOKENSUPDATED) (event.Subscription, error) {

	logs, sub, err := _Sigmacore.contract.WatchLogs(opts, "LOG_MAX_TOKENS_UPDATED")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SigmacoreLOGMAXTOKENSUPDATED)
				if err := _Sigmacore.contract.UnpackLog(event, "LOG_MAX_TOKENS_UPDATED", log); err != nil {
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

// ParseLOGMAXTOKENSUPDATED is a log parse operation binding the contract event 0x65492266ae9a1f46497efddd24b6946862569680a511c543590d4587ca800d05.
//
// Solidity: event LOG_MAX_TOKENS_UPDATED(uint256 maxPoolTokens)
func (_Sigmacore *SigmacoreFilterer) ParseLOGMAXTOKENSUPDATED(log types.Log) (*SigmacoreLOGMAXTOKENSUPDATED, error) {
	event := new(SigmacoreLOGMAXTOKENSUPDATED)
	if err := _Sigmacore.contract.UnpackLog(event, "LOG_MAX_TOKENS_UPDATED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SigmacoreLOGMINIMUMBALANCEUPDATEDIterator is returned from FilterLOGMINIMUMBALANCEUPDATED and is used to iterate over the raw logs and unpacked data for LOGMINIMUMBALANCEUPDATED events raised by the Sigmacore contract.
type SigmacoreLOGMINIMUMBALANCEUPDATEDIterator struct {
	Event *SigmacoreLOGMINIMUMBALANCEUPDATED // Event containing the contract specifics and raw log

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
func (it *SigmacoreLOGMINIMUMBALANCEUPDATEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SigmacoreLOGMINIMUMBALANCEUPDATED)
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
		it.Event = new(SigmacoreLOGMINIMUMBALANCEUPDATED)
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
func (it *SigmacoreLOGMINIMUMBALANCEUPDATEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SigmacoreLOGMINIMUMBALANCEUPDATEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SigmacoreLOGMINIMUMBALANCEUPDATED represents a LOGMINIMUMBALANCEUPDATED event raised by the Sigmacore contract.
type SigmacoreLOGMINIMUMBALANCEUPDATED struct {
	Token          common.Address
	MinimumBalance *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGMINIMUMBALANCEUPDATED is a free log retrieval operation binding the contract event 0x000c7a55677231b335e6dea005fa240ac2aeafbd62f188372a7d66892b722c52.
//
// Solidity: event LOG_MINIMUM_BALANCE_UPDATED(address token, uint256 minimumBalance)
func (_Sigmacore *SigmacoreFilterer) FilterLOGMINIMUMBALANCEUPDATED(opts *bind.FilterOpts) (*SigmacoreLOGMINIMUMBALANCEUPDATEDIterator, error) {

	logs, sub, err := _Sigmacore.contract.FilterLogs(opts, "LOG_MINIMUM_BALANCE_UPDATED")
	if err != nil {
		return nil, err
	}
	return &SigmacoreLOGMINIMUMBALANCEUPDATEDIterator{contract: _Sigmacore.contract, event: "LOG_MINIMUM_BALANCE_UPDATED", logs: logs, sub: sub}, nil
}

// WatchLOGMINIMUMBALANCEUPDATED is a free log subscription operation binding the contract event 0x000c7a55677231b335e6dea005fa240ac2aeafbd62f188372a7d66892b722c52.
//
// Solidity: event LOG_MINIMUM_BALANCE_UPDATED(address token, uint256 minimumBalance)
func (_Sigmacore *SigmacoreFilterer) WatchLOGMINIMUMBALANCEUPDATED(opts *bind.WatchOpts, sink chan<- *SigmacoreLOGMINIMUMBALANCEUPDATED) (event.Subscription, error) {

	logs, sub, err := _Sigmacore.contract.WatchLogs(opts, "LOG_MINIMUM_BALANCE_UPDATED")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SigmacoreLOGMINIMUMBALANCEUPDATED)
				if err := _Sigmacore.contract.UnpackLog(event, "LOG_MINIMUM_BALANCE_UPDATED", log); err != nil {
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

// ParseLOGMINIMUMBALANCEUPDATED is a log parse operation binding the contract event 0x000c7a55677231b335e6dea005fa240ac2aeafbd62f188372a7d66892b722c52.
//
// Solidity: event LOG_MINIMUM_BALANCE_UPDATED(address token, uint256 minimumBalance)
func (_Sigmacore *SigmacoreFilterer) ParseLOGMINIMUMBALANCEUPDATED(log types.Log) (*SigmacoreLOGMINIMUMBALANCEUPDATED, error) {
	event := new(SigmacoreLOGMINIMUMBALANCEUPDATED)
	if err := _Sigmacore.contract.UnpackLog(event, "LOG_MINIMUM_BALANCE_UPDATED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SigmacoreLOGPUBLICSWAPTOGGLEDIterator is returned from FilterLOGPUBLICSWAPTOGGLED and is used to iterate over the raw logs and unpacked data for LOGPUBLICSWAPTOGGLED events raised by the Sigmacore contract.
type SigmacoreLOGPUBLICSWAPTOGGLEDIterator struct {
	Event *SigmacoreLOGPUBLICSWAPTOGGLED // Event containing the contract specifics and raw log

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
func (it *SigmacoreLOGPUBLICSWAPTOGGLEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SigmacoreLOGPUBLICSWAPTOGGLED)
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
		it.Event = new(SigmacoreLOGPUBLICSWAPTOGGLED)
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
func (it *SigmacoreLOGPUBLICSWAPTOGGLEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SigmacoreLOGPUBLICSWAPTOGGLEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SigmacoreLOGPUBLICSWAPTOGGLED represents a LOGPUBLICSWAPTOGGLED event raised by the Sigmacore contract.
type SigmacoreLOGPUBLICSWAPTOGGLED struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLOGPUBLICSWAPTOGGLED is a free log retrieval operation binding the contract event 0x40fc85fbff9305015298ba6fcee88b7e442a64cc803ddb889327680bbd62270a.
//
// Solidity: event LOG_PUBLIC_SWAP_TOGGLED(bool enabled)
func (_Sigmacore *SigmacoreFilterer) FilterLOGPUBLICSWAPTOGGLED(opts *bind.FilterOpts) (*SigmacoreLOGPUBLICSWAPTOGGLEDIterator, error) {

	logs, sub, err := _Sigmacore.contract.FilterLogs(opts, "LOG_PUBLIC_SWAP_TOGGLED")
	if err != nil {
		return nil, err
	}
	return &SigmacoreLOGPUBLICSWAPTOGGLEDIterator{contract: _Sigmacore.contract, event: "LOG_PUBLIC_SWAP_TOGGLED", logs: logs, sub: sub}, nil
}

// WatchLOGPUBLICSWAPTOGGLED is a free log subscription operation binding the contract event 0x40fc85fbff9305015298ba6fcee88b7e442a64cc803ddb889327680bbd62270a.
//
// Solidity: event LOG_PUBLIC_SWAP_TOGGLED(bool enabled)
func (_Sigmacore *SigmacoreFilterer) WatchLOGPUBLICSWAPTOGGLED(opts *bind.WatchOpts, sink chan<- *SigmacoreLOGPUBLICSWAPTOGGLED) (event.Subscription, error) {

	logs, sub, err := _Sigmacore.contract.WatchLogs(opts, "LOG_PUBLIC_SWAP_TOGGLED")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SigmacoreLOGPUBLICSWAPTOGGLED)
				if err := _Sigmacore.contract.UnpackLog(event, "LOG_PUBLIC_SWAP_TOGGLED", log); err != nil {
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

// ParseLOGPUBLICSWAPTOGGLED is a log parse operation binding the contract event 0x40fc85fbff9305015298ba6fcee88b7e442a64cc803ddb889327680bbd62270a.
//
// Solidity: event LOG_PUBLIC_SWAP_TOGGLED(bool enabled)
func (_Sigmacore *SigmacoreFilterer) ParseLOGPUBLICSWAPTOGGLED(log types.Log) (*SigmacoreLOGPUBLICSWAPTOGGLED, error) {
	event := new(SigmacoreLOGPUBLICSWAPTOGGLED)
	if err := _Sigmacore.contract.UnpackLog(event, "LOG_PUBLIC_SWAP_TOGGLED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SigmacoreLOGSWAPIterator is returned from FilterLOGSWAP and is used to iterate over the raw logs and unpacked data for LOGSWAP events raised by the Sigmacore contract.
type SigmacoreLOGSWAPIterator struct {
	Event *SigmacoreLOGSWAP // Event containing the contract specifics and raw log

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
func (it *SigmacoreLOGSWAPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SigmacoreLOGSWAP)
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
		it.Event = new(SigmacoreLOGSWAP)
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
func (it *SigmacoreLOGSWAPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SigmacoreLOGSWAPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SigmacoreLOGSWAP represents a LOGSWAP event raised by the Sigmacore contract.
type SigmacoreLOGSWAP struct {
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
func (_Sigmacore *SigmacoreFilterer) FilterLOGSWAP(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (*SigmacoreLOGSWAPIterator, error) {

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

	logs, sub, err := _Sigmacore.contract.FilterLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &SigmacoreLOGSWAPIterator{contract: _Sigmacore.contract, event: "LOG_SWAP", logs: logs, sub: sub}, nil
}

// WatchLOGSWAP is a free log subscription operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_Sigmacore *SigmacoreFilterer) WatchLOGSWAP(opts *bind.WatchOpts, sink chan<- *SigmacoreLOGSWAP, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Sigmacore.contract.WatchLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SigmacoreLOGSWAP)
				if err := _Sigmacore.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
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
func (_Sigmacore *SigmacoreFilterer) ParseLOGSWAP(log types.Log) (*SigmacoreLOGSWAP, error) {
	event := new(SigmacoreLOGSWAP)
	if err := _Sigmacore.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SigmacoreLOGSWAPFEEUPDATEDIterator is returned from FilterLOGSWAPFEEUPDATED and is used to iterate over the raw logs and unpacked data for LOGSWAPFEEUPDATED events raised by the Sigmacore contract.
type SigmacoreLOGSWAPFEEUPDATEDIterator struct {
	Event *SigmacoreLOGSWAPFEEUPDATED // Event containing the contract specifics and raw log

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
func (it *SigmacoreLOGSWAPFEEUPDATEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SigmacoreLOGSWAPFEEUPDATED)
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
		it.Event = new(SigmacoreLOGSWAPFEEUPDATED)
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
func (it *SigmacoreLOGSWAPFEEUPDATEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SigmacoreLOGSWAPFEEUPDATEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SigmacoreLOGSWAPFEEUPDATED represents a LOGSWAPFEEUPDATED event raised by the Sigmacore contract.
type SigmacoreLOGSWAPFEEUPDATED struct {
	SwapFee *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLOGSWAPFEEUPDATED is a free log retrieval operation binding the contract event 0xccfe595973efc7c1f6c29e31974d380470b9431d7770290185b7129419c7e63e.
//
// Solidity: event LOG_SWAP_FEE_UPDATED(uint256 swapFee)
func (_Sigmacore *SigmacoreFilterer) FilterLOGSWAPFEEUPDATED(opts *bind.FilterOpts) (*SigmacoreLOGSWAPFEEUPDATEDIterator, error) {

	logs, sub, err := _Sigmacore.contract.FilterLogs(opts, "LOG_SWAP_FEE_UPDATED")
	if err != nil {
		return nil, err
	}
	return &SigmacoreLOGSWAPFEEUPDATEDIterator{contract: _Sigmacore.contract, event: "LOG_SWAP_FEE_UPDATED", logs: logs, sub: sub}, nil
}

// WatchLOGSWAPFEEUPDATED is a free log subscription operation binding the contract event 0xccfe595973efc7c1f6c29e31974d380470b9431d7770290185b7129419c7e63e.
//
// Solidity: event LOG_SWAP_FEE_UPDATED(uint256 swapFee)
func (_Sigmacore *SigmacoreFilterer) WatchLOGSWAPFEEUPDATED(opts *bind.WatchOpts, sink chan<- *SigmacoreLOGSWAPFEEUPDATED) (event.Subscription, error) {

	logs, sub, err := _Sigmacore.contract.WatchLogs(opts, "LOG_SWAP_FEE_UPDATED")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SigmacoreLOGSWAPFEEUPDATED)
				if err := _Sigmacore.contract.UnpackLog(event, "LOG_SWAP_FEE_UPDATED", log); err != nil {
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

// ParseLOGSWAPFEEUPDATED is a log parse operation binding the contract event 0xccfe595973efc7c1f6c29e31974d380470b9431d7770290185b7129419c7e63e.
//
// Solidity: event LOG_SWAP_FEE_UPDATED(uint256 swapFee)
func (_Sigmacore *SigmacoreFilterer) ParseLOGSWAPFEEUPDATED(log types.Log) (*SigmacoreLOGSWAPFEEUPDATED, error) {
	event := new(SigmacoreLOGSWAPFEEUPDATED)
	if err := _Sigmacore.contract.UnpackLog(event, "LOG_SWAP_FEE_UPDATED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SigmacoreLOGTOKENADDEDIterator is returned from FilterLOGTOKENADDED and is used to iterate over the raw logs and unpacked data for LOGTOKENADDED events raised by the Sigmacore contract.
type SigmacoreLOGTOKENADDEDIterator struct {
	Event *SigmacoreLOGTOKENADDED // Event containing the contract specifics and raw log

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
func (it *SigmacoreLOGTOKENADDEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SigmacoreLOGTOKENADDED)
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
		it.Event = new(SigmacoreLOGTOKENADDED)
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
func (it *SigmacoreLOGTOKENADDEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SigmacoreLOGTOKENADDEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SigmacoreLOGTOKENADDED represents a LOGTOKENADDED event raised by the Sigmacore contract.
type SigmacoreLOGTOKENADDED struct {
	Token          common.Address
	DesiredDenorm  *big.Int
	MinimumBalance *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGTOKENADDED is a free log retrieval operation binding the contract event 0xb2daf560899f6307b318aecfb57eb2812c488da4a4c1cad2019b482fa63294ed.
//
// Solidity: event LOG_TOKEN_ADDED(address indexed token, uint256 desiredDenorm, uint256 minimumBalance)
func (_Sigmacore *SigmacoreFilterer) FilterLOGTOKENADDED(opts *bind.FilterOpts, token []common.Address) (*SigmacoreLOGTOKENADDEDIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Sigmacore.contract.FilterLogs(opts, "LOG_TOKEN_ADDED", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SigmacoreLOGTOKENADDEDIterator{contract: _Sigmacore.contract, event: "LOG_TOKEN_ADDED", logs: logs, sub: sub}, nil
}

// WatchLOGTOKENADDED is a free log subscription operation binding the contract event 0xb2daf560899f6307b318aecfb57eb2812c488da4a4c1cad2019b482fa63294ed.
//
// Solidity: event LOG_TOKEN_ADDED(address indexed token, uint256 desiredDenorm, uint256 minimumBalance)
func (_Sigmacore *SigmacoreFilterer) WatchLOGTOKENADDED(opts *bind.WatchOpts, sink chan<- *SigmacoreLOGTOKENADDED, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Sigmacore.contract.WatchLogs(opts, "LOG_TOKEN_ADDED", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SigmacoreLOGTOKENADDED)
				if err := _Sigmacore.contract.UnpackLog(event, "LOG_TOKEN_ADDED", log); err != nil {
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

// ParseLOGTOKENADDED is a log parse operation binding the contract event 0xb2daf560899f6307b318aecfb57eb2812c488da4a4c1cad2019b482fa63294ed.
//
// Solidity: event LOG_TOKEN_ADDED(address indexed token, uint256 desiredDenorm, uint256 minimumBalance)
func (_Sigmacore *SigmacoreFilterer) ParseLOGTOKENADDED(log types.Log) (*SigmacoreLOGTOKENADDED, error) {
	event := new(SigmacoreLOGTOKENADDED)
	if err := _Sigmacore.contract.UnpackLog(event, "LOG_TOKEN_ADDED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SigmacoreLOGTOKENREADYIterator is returned from FilterLOGTOKENREADY and is used to iterate over the raw logs and unpacked data for LOGTOKENREADY events raised by the Sigmacore contract.
type SigmacoreLOGTOKENREADYIterator struct {
	Event *SigmacoreLOGTOKENREADY // Event containing the contract specifics and raw log

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
func (it *SigmacoreLOGTOKENREADYIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SigmacoreLOGTOKENREADY)
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
		it.Event = new(SigmacoreLOGTOKENREADY)
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
func (it *SigmacoreLOGTOKENREADYIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SigmacoreLOGTOKENREADYIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SigmacoreLOGTOKENREADY represents a LOGTOKENREADY event raised by the Sigmacore contract.
type SigmacoreLOGTOKENREADY struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLOGTOKENREADY is a free log retrieval operation binding the contract event 0xf7bb8e57ffdfd9a31e7580ee84f68757f44fb4a8a913f44520d22f2da1c955e5.
//
// Solidity: event LOG_TOKEN_READY(address indexed token)
func (_Sigmacore *SigmacoreFilterer) FilterLOGTOKENREADY(opts *bind.FilterOpts, token []common.Address) (*SigmacoreLOGTOKENREADYIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Sigmacore.contract.FilterLogs(opts, "LOG_TOKEN_READY", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SigmacoreLOGTOKENREADYIterator{contract: _Sigmacore.contract, event: "LOG_TOKEN_READY", logs: logs, sub: sub}, nil
}

// WatchLOGTOKENREADY is a free log subscription operation binding the contract event 0xf7bb8e57ffdfd9a31e7580ee84f68757f44fb4a8a913f44520d22f2da1c955e5.
//
// Solidity: event LOG_TOKEN_READY(address indexed token)
func (_Sigmacore *SigmacoreFilterer) WatchLOGTOKENREADY(opts *bind.WatchOpts, sink chan<- *SigmacoreLOGTOKENREADY, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Sigmacore.contract.WatchLogs(opts, "LOG_TOKEN_READY", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SigmacoreLOGTOKENREADY)
				if err := _Sigmacore.contract.UnpackLog(event, "LOG_TOKEN_READY", log); err != nil {
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

// ParseLOGTOKENREADY is a log parse operation binding the contract event 0xf7bb8e57ffdfd9a31e7580ee84f68757f44fb4a8a913f44520d22f2da1c955e5.
//
// Solidity: event LOG_TOKEN_READY(address indexed token)
func (_Sigmacore *SigmacoreFilterer) ParseLOGTOKENREADY(log types.Log) (*SigmacoreLOGTOKENREADY, error) {
	event := new(SigmacoreLOGTOKENREADY)
	if err := _Sigmacore.contract.UnpackLog(event, "LOG_TOKEN_READY", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SigmacoreLOGTOKENREMOVEDIterator is returned from FilterLOGTOKENREMOVED and is used to iterate over the raw logs and unpacked data for LOGTOKENREMOVED events raised by the Sigmacore contract.
type SigmacoreLOGTOKENREMOVEDIterator struct {
	Event *SigmacoreLOGTOKENREMOVED // Event containing the contract specifics and raw log

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
func (it *SigmacoreLOGTOKENREMOVEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SigmacoreLOGTOKENREMOVED)
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
		it.Event = new(SigmacoreLOGTOKENREMOVED)
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
func (it *SigmacoreLOGTOKENREMOVEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SigmacoreLOGTOKENREMOVEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SigmacoreLOGTOKENREMOVED represents a LOGTOKENREMOVED event raised by the Sigmacore contract.
type SigmacoreLOGTOKENREMOVED struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLOGTOKENREMOVED is a free log retrieval operation binding the contract event 0x12a8262eb28ee8a8c11e6cf411b3af6ce5bea42abb36e051bf0a65ae602d52ec.
//
// Solidity: event LOG_TOKEN_REMOVED(address token)
func (_Sigmacore *SigmacoreFilterer) FilterLOGTOKENREMOVED(opts *bind.FilterOpts) (*SigmacoreLOGTOKENREMOVEDIterator, error) {

	logs, sub, err := _Sigmacore.contract.FilterLogs(opts, "LOG_TOKEN_REMOVED")
	if err != nil {
		return nil, err
	}
	return &SigmacoreLOGTOKENREMOVEDIterator{contract: _Sigmacore.contract, event: "LOG_TOKEN_REMOVED", logs: logs, sub: sub}, nil
}

// WatchLOGTOKENREMOVED is a free log subscription operation binding the contract event 0x12a8262eb28ee8a8c11e6cf411b3af6ce5bea42abb36e051bf0a65ae602d52ec.
//
// Solidity: event LOG_TOKEN_REMOVED(address token)
func (_Sigmacore *SigmacoreFilterer) WatchLOGTOKENREMOVED(opts *bind.WatchOpts, sink chan<- *SigmacoreLOGTOKENREMOVED) (event.Subscription, error) {

	logs, sub, err := _Sigmacore.contract.WatchLogs(opts, "LOG_TOKEN_REMOVED")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SigmacoreLOGTOKENREMOVED)
				if err := _Sigmacore.contract.UnpackLog(event, "LOG_TOKEN_REMOVED", log); err != nil {
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

// ParseLOGTOKENREMOVED is a log parse operation binding the contract event 0x12a8262eb28ee8a8c11e6cf411b3af6ce5bea42abb36e051bf0a65ae602d52ec.
//
// Solidity: event LOG_TOKEN_REMOVED(address token)
func (_Sigmacore *SigmacoreFilterer) ParseLOGTOKENREMOVED(log types.Log) (*SigmacoreLOGTOKENREMOVED, error) {
	event := new(SigmacoreLOGTOKENREMOVED)
	if err := _Sigmacore.contract.UnpackLog(event, "LOG_TOKEN_REMOVED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SigmacoreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Sigmacore contract.
type SigmacoreTransferIterator struct {
	Event *SigmacoreTransfer // Event containing the contract specifics and raw log

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
func (it *SigmacoreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SigmacoreTransfer)
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
		it.Event = new(SigmacoreTransfer)
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
func (it *SigmacoreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SigmacoreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SigmacoreTransfer represents a Transfer event raised by the Sigmacore contract.
type SigmacoreTransfer struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_Sigmacore *SigmacoreFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*SigmacoreTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Sigmacore.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &SigmacoreTransferIterator{contract: _Sigmacore.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_Sigmacore *SigmacoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SigmacoreTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Sigmacore.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SigmacoreTransfer)
				if err := _Sigmacore.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_Sigmacore *SigmacoreFilterer) ParseTransfer(log types.Log) (*SigmacoreTransfer, error) {
	event := new(SigmacoreTransfer)
	if err := _Sigmacore.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

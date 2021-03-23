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
const ControllerABI = "[{\"inputs\":[{\"internalType\":\"contractIIndexedUniswapV2Oracle\",\"name\":\"uniswapOracle_\",\"type\":\"address\"},{\"internalType\":\"contractIPoolFactory\",\"name\":\"poolFactory_\",\"type\":\"address\"},{\"internalType\":\"contractIDelegateCallProxyManager\",\"name\":\"proxyManager_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defaultExitFeeRecipient_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"}],\"name\":\"NewPoolInitializer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"unboundTokenSeller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"}],\"name\":\"PoolInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"scoringStrategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"minimumScore\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"maximumScore\",\"type\":\"uint128\"}],\"name\":\"TokenListAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"}],\"name\":\"TokenListSorted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INITIALIZER_IMPLEMENTATION_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_INDEX_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LIST_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_INDEX_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_IMPLEMENTATION_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_REWEIGH_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWEIGHS_BEFORE_REINDEX\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SELLER_IMPLEMENTATION_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHORT_TWAP_MAX_TIME_ELAPSED\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHORT_TWAP_MIN_TIME_ELAPSED\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"addTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circuitBreaker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circulatingMarketCapOracle\",\"outputs\":[{\"internalType\":\"contractICirculatingMarketCapOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"computeInitializerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"initializerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"}],\"name\":\"computePoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"computeSellerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sellerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"scoringStrategy\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"minimumScore\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maximumScore\",\"type\":\"uint128\"}],\"name\":\"createTokenList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultExitFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultSellerPremium\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegateCompLikeTokenFromPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"name\":\"finishPreparedIndexPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"forceReindexPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wethValue\",\"type\":\"uint256\"}],\"name\":\"getInitialTokensAndBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"}],\"name\":\"getSortedAndFilteredTokensAndScores\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"scores\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"}],\"name\":\"getTokenList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"}],\"name\":\"getTokenListConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"scoringStrategy\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"minimumScore\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maximumScore\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getTokenScores\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"scores\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getTopTokensAndScores\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"scores\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"indexPoolMetadata\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"listID\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"indexSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"reweighIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"lastReweigh\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"circuitBreaker_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenInlist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolFactory\",\"outputs\":[{\"internalType\":\"contractIPoolFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialWethValue\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"prepareIndexPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initializerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyManager\",\"outputs\":[{\"internalType\":\"contractIDelegateCallProxyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"reindexPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"reweighPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"circuitBreaker_\",\"type\":\"address\"}],\"name\":\"setCircuitBreaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_defaultSellerPremium\",\"type\":\"uint8\"}],\"name\":\"setDefaultSellerPremium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"indexPool_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"publicSwap\",\"type\":\"bool\"}],\"name\":\"setPublicSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"}],\"name\":\"sortAndFilterTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenListCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapOracle\",\"outputs\":[{\"internalType\":\"contractIIndexedUniswapV2Oracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"updateMinimumBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenSeller\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"premiumPercent\",\"type\":\"uint8\"}],\"name\":\"updateSellerPremium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listID\",\"type\":\"uint256\"}],\"name\":\"updateTokenPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"pricesUpdated\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// INITIALIZERIMPLEMENTATIONID is a free data retrieval call binding the contract method 0xf52a4f26.
//
// Solidity: function INITIALIZER_IMPLEMENTATION_ID() view returns(bytes32)
func (_Controller *ControllerCaller) INITIALIZERIMPLEMENTATIONID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "INITIALIZER_IMPLEMENTATION_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITIALIZERIMPLEMENTATIONID is a free data retrieval call binding the contract method 0xf52a4f26.
//
// Solidity: function INITIALIZER_IMPLEMENTATION_ID() view returns(bytes32)
func (_Controller *ControllerSession) INITIALIZERIMPLEMENTATIONID() ([32]byte, error) {
	return _Controller.Contract.INITIALIZERIMPLEMENTATIONID(&_Controller.CallOpts)
}

// INITIALIZERIMPLEMENTATIONID is a free data retrieval call binding the contract method 0xf52a4f26.
//
// Solidity: function INITIALIZER_IMPLEMENTATION_ID() view returns(bytes32)
func (_Controller *ControllerCallerSession) INITIALIZERIMPLEMENTATIONID() ([32]byte, error) {
	return _Controller.Contract.INITIALIZERIMPLEMENTATIONID(&_Controller.CallOpts)
}

// MAXINDEXSIZE is a free data retrieval call binding the contract method 0xacbfc96d.
//
// Solidity: function MAX_INDEX_SIZE() view returns(uint256)
func (_Controller *ControllerCaller) MAXINDEXSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "MAX_INDEX_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXINDEXSIZE is a free data retrieval call binding the contract method 0xacbfc96d.
//
// Solidity: function MAX_INDEX_SIZE() view returns(uint256)
func (_Controller *ControllerSession) MAXINDEXSIZE() (*big.Int, error) {
	return _Controller.Contract.MAXINDEXSIZE(&_Controller.CallOpts)
}

// MAXINDEXSIZE is a free data retrieval call binding the contract method 0xacbfc96d.
//
// Solidity: function MAX_INDEX_SIZE() view returns(uint256)
func (_Controller *ControllerCallerSession) MAXINDEXSIZE() (*big.Int, error) {
	return _Controller.Contract.MAXINDEXSIZE(&_Controller.CallOpts)
}

// MAXLISTTOKENS is a free data retrieval call binding the contract method 0xec0915c7.
//
// Solidity: function MAX_LIST_TOKENS() view returns(uint256)
func (_Controller *ControllerCaller) MAXLISTTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "MAX_LIST_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLISTTOKENS is a free data retrieval call binding the contract method 0xec0915c7.
//
// Solidity: function MAX_LIST_TOKENS() view returns(uint256)
func (_Controller *ControllerSession) MAXLISTTOKENS() (*big.Int, error) {
	return _Controller.Contract.MAXLISTTOKENS(&_Controller.CallOpts)
}

// MAXLISTTOKENS is a free data retrieval call binding the contract method 0xec0915c7.
//
// Solidity: function MAX_LIST_TOKENS() view returns(uint256)
func (_Controller *ControllerCallerSession) MAXLISTTOKENS() (*big.Int, error) {
	return _Controller.Contract.MAXLISTTOKENS(&_Controller.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_Controller *ControllerCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "MIN_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_Controller *ControllerSession) MINBALANCE() (*big.Int, error) {
	return _Controller.Contract.MINBALANCE(&_Controller.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_Controller *ControllerCallerSession) MINBALANCE() (*big.Int, error) {
	return _Controller.Contract.MINBALANCE(&_Controller.CallOpts)
}

// MININDEXSIZE is a free data retrieval call binding the contract method 0xb3ec862e.
//
// Solidity: function MIN_INDEX_SIZE() view returns(uint256)
func (_Controller *ControllerCaller) MININDEXSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "MIN_INDEX_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MININDEXSIZE is a free data retrieval call binding the contract method 0xb3ec862e.
//
// Solidity: function MIN_INDEX_SIZE() view returns(uint256)
func (_Controller *ControllerSession) MININDEXSIZE() (*big.Int, error) {
	return _Controller.Contract.MININDEXSIZE(&_Controller.CallOpts)
}

// MININDEXSIZE is a free data retrieval call binding the contract method 0xb3ec862e.
//
// Solidity: function MIN_INDEX_SIZE() view returns(uint256)
func (_Controller *ControllerCallerSession) MININDEXSIZE() (*big.Int, error) {
	return _Controller.Contract.MININDEXSIZE(&_Controller.CallOpts)
}

// POOLIMPLEMENTATIONID is a free data retrieval call binding the contract method 0x250436fd.
//
// Solidity: function POOL_IMPLEMENTATION_ID() view returns(bytes32)
func (_Controller *ControllerCaller) POOLIMPLEMENTATIONID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "POOL_IMPLEMENTATION_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POOLIMPLEMENTATIONID is a free data retrieval call binding the contract method 0x250436fd.
//
// Solidity: function POOL_IMPLEMENTATION_ID() view returns(bytes32)
func (_Controller *ControllerSession) POOLIMPLEMENTATIONID() ([32]byte, error) {
	return _Controller.Contract.POOLIMPLEMENTATIONID(&_Controller.CallOpts)
}

// POOLIMPLEMENTATIONID is a free data retrieval call binding the contract method 0x250436fd.
//
// Solidity: function POOL_IMPLEMENTATION_ID() view returns(bytes32)
func (_Controller *ControllerCallerSession) POOLIMPLEMENTATIONID() ([32]byte, error) {
	return _Controller.Contract.POOLIMPLEMENTATIONID(&_Controller.CallOpts)
}

// POOLREWEIGHDELAY is a free data retrieval call binding the contract method 0x830b6b0b.
//
// Solidity: function POOL_REWEIGH_DELAY() view returns(uint256)
func (_Controller *ControllerCaller) POOLREWEIGHDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "POOL_REWEIGH_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// POOLREWEIGHDELAY is a free data retrieval call binding the contract method 0x830b6b0b.
//
// Solidity: function POOL_REWEIGH_DELAY() view returns(uint256)
func (_Controller *ControllerSession) POOLREWEIGHDELAY() (*big.Int, error) {
	return _Controller.Contract.POOLREWEIGHDELAY(&_Controller.CallOpts)
}

// POOLREWEIGHDELAY is a free data retrieval call binding the contract method 0x830b6b0b.
//
// Solidity: function POOL_REWEIGH_DELAY() view returns(uint256)
func (_Controller *ControllerCallerSession) POOLREWEIGHDELAY() (*big.Int, error) {
	return _Controller.Contract.POOLREWEIGHDELAY(&_Controller.CallOpts)
}

// REWEIGHSBEFOREREINDEX is a free data retrieval call binding the contract method 0x3d89da25.
//
// Solidity: function REWEIGHS_BEFORE_REINDEX() view returns(uint8)
func (_Controller *ControllerCaller) REWEIGHSBEFOREREINDEX(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "REWEIGHS_BEFORE_REINDEX")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// REWEIGHSBEFOREREINDEX is a free data retrieval call binding the contract method 0x3d89da25.
//
// Solidity: function REWEIGHS_BEFORE_REINDEX() view returns(uint8)
func (_Controller *ControllerSession) REWEIGHSBEFOREREINDEX() (uint8, error) {
	return _Controller.Contract.REWEIGHSBEFOREREINDEX(&_Controller.CallOpts)
}

// REWEIGHSBEFOREREINDEX is a free data retrieval call binding the contract method 0x3d89da25.
//
// Solidity: function REWEIGHS_BEFORE_REINDEX() view returns(uint8)
func (_Controller *ControllerCallerSession) REWEIGHSBEFOREREINDEX() (uint8, error) {
	return _Controller.Contract.REWEIGHSBEFOREREINDEX(&_Controller.CallOpts)
}

// SELLERIMPLEMENTATIONID is a free data retrieval call binding the contract method 0xc5ea91e1.
//
// Solidity: function SELLER_IMPLEMENTATION_ID() view returns(bytes32)
func (_Controller *ControllerCaller) SELLERIMPLEMENTATIONID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "SELLER_IMPLEMENTATION_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SELLERIMPLEMENTATIONID is a free data retrieval call binding the contract method 0xc5ea91e1.
//
// Solidity: function SELLER_IMPLEMENTATION_ID() view returns(bytes32)
func (_Controller *ControllerSession) SELLERIMPLEMENTATIONID() ([32]byte, error) {
	return _Controller.Contract.SELLERIMPLEMENTATIONID(&_Controller.CallOpts)
}

// SELLERIMPLEMENTATIONID is a free data retrieval call binding the contract method 0xc5ea91e1.
//
// Solidity: function SELLER_IMPLEMENTATION_ID() view returns(bytes32)
func (_Controller *ControllerCallerSession) SELLERIMPLEMENTATIONID() ([32]byte, error) {
	return _Controller.Contract.SELLERIMPLEMENTATIONID(&_Controller.CallOpts)
}

// SHORTTWAPMAXTIMEELAPSED is a free data retrieval call binding the contract method 0xa34fefdd.
//
// Solidity: function SHORT_TWAP_MAX_TIME_ELAPSED() view returns(uint32)
func (_Controller *ControllerCaller) SHORTTWAPMAXTIMEELAPSED(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "SHORT_TWAP_MAX_TIME_ELAPSED")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SHORTTWAPMAXTIMEELAPSED is a free data retrieval call binding the contract method 0xa34fefdd.
//
// Solidity: function SHORT_TWAP_MAX_TIME_ELAPSED() view returns(uint32)
func (_Controller *ControllerSession) SHORTTWAPMAXTIMEELAPSED() (uint32, error) {
	return _Controller.Contract.SHORTTWAPMAXTIMEELAPSED(&_Controller.CallOpts)
}

// SHORTTWAPMAXTIMEELAPSED is a free data retrieval call binding the contract method 0xa34fefdd.
//
// Solidity: function SHORT_TWAP_MAX_TIME_ELAPSED() view returns(uint32)
func (_Controller *ControllerCallerSession) SHORTTWAPMAXTIMEELAPSED() (uint32, error) {
	return _Controller.Contract.SHORTTWAPMAXTIMEELAPSED(&_Controller.CallOpts)
}

// SHORTTWAPMINTIMEELAPSED is a free data retrieval call binding the contract method 0x53304eb9.
//
// Solidity: function SHORT_TWAP_MIN_TIME_ELAPSED() view returns(uint32)
func (_Controller *ControllerCaller) SHORTTWAPMINTIMEELAPSED(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "SHORT_TWAP_MIN_TIME_ELAPSED")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SHORTTWAPMINTIMEELAPSED is a free data retrieval call binding the contract method 0x53304eb9.
//
// Solidity: function SHORT_TWAP_MIN_TIME_ELAPSED() view returns(uint32)
func (_Controller *ControllerSession) SHORTTWAPMINTIMEELAPSED() (uint32, error) {
	return _Controller.Contract.SHORTTWAPMINTIMEELAPSED(&_Controller.CallOpts)
}

// SHORTTWAPMINTIMEELAPSED is a free data retrieval call binding the contract method 0x53304eb9.
//
// Solidity: function SHORT_TWAP_MIN_TIME_ELAPSED() view returns(uint32)
func (_Controller *ControllerCallerSession) SHORTTWAPMINTIMEELAPSED() (uint32, error) {
	return _Controller.Contract.SHORTTWAPMINTIMEELAPSED(&_Controller.CallOpts)
}

// CircuitBreaker is a free data retrieval call binding the contract method 0x16efd941.
//
// Solidity: function circuitBreaker() view returns(address)
func (_Controller *ControllerCaller) CircuitBreaker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "circuitBreaker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CircuitBreaker is a free data retrieval call binding the contract method 0x16efd941.
//
// Solidity: function circuitBreaker() view returns(address)
func (_Controller *ControllerSession) CircuitBreaker() (common.Address, error) {
	return _Controller.Contract.CircuitBreaker(&_Controller.CallOpts)
}

// CircuitBreaker is a free data retrieval call binding the contract method 0x16efd941.
//
// Solidity: function circuitBreaker() view returns(address)
func (_Controller *ControllerCallerSession) CircuitBreaker() (common.Address, error) {
	return _Controller.Contract.CircuitBreaker(&_Controller.CallOpts)
}

// CirculatingMarketCapOracle is a free data retrieval call binding the contract method 0xd600751d.
//
// Solidity: function circulatingMarketCapOracle() view returns(address)
func (_Controller *ControllerCaller) CirculatingMarketCapOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "circulatingMarketCapOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CirculatingMarketCapOracle is a free data retrieval call binding the contract method 0xd600751d.
//
// Solidity: function circulatingMarketCapOracle() view returns(address)
func (_Controller *ControllerSession) CirculatingMarketCapOracle() (common.Address, error) {
	return _Controller.Contract.CirculatingMarketCapOracle(&_Controller.CallOpts)
}

// CirculatingMarketCapOracle is a free data retrieval call binding the contract method 0xd600751d.
//
// Solidity: function circulatingMarketCapOracle() view returns(address)
func (_Controller *ControllerCallerSession) CirculatingMarketCapOracle() (common.Address, error) {
	return _Controller.Contract.CirculatingMarketCapOracle(&_Controller.CallOpts)
}

// ComputeInitializerAddress is a free data retrieval call binding the contract method 0xfa7c0b5e.
//
// Solidity: function computeInitializerAddress(address poolAddress) view returns(address initializerAddress)
func (_Controller *ControllerCaller) ComputeInitializerAddress(opts *bind.CallOpts, poolAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "computeInitializerAddress", poolAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeInitializerAddress is a free data retrieval call binding the contract method 0xfa7c0b5e.
//
// Solidity: function computeInitializerAddress(address poolAddress) view returns(address initializerAddress)
func (_Controller *ControllerSession) ComputeInitializerAddress(poolAddress common.Address) (common.Address, error) {
	return _Controller.Contract.ComputeInitializerAddress(&_Controller.CallOpts, poolAddress)
}

// ComputeInitializerAddress is a free data retrieval call binding the contract method 0xfa7c0b5e.
//
// Solidity: function computeInitializerAddress(address poolAddress) view returns(address initializerAddress)
func (_Controller *ControllerCallerSession) ComputeInitializerAddress(poolAddress common.Address) (common.Address, error) {
	return _Controller.Contract.ComputeInitializerAddress(&_Controller.CallOpts, poolAddress)
}

// ComputePoolAddress is a free data retrieval call binding the contract method 0x9dd19a52.
//
// Solidity: function computePoolAddress(uint256 listID, uint256 indexSize) view returns(address poolAddress)
func (_Controller *ControllerCaller) ComputePoolAddress(opts *bind.CallOpts, listID *big.Int, indexSize *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "computePoolAddress", listID, indexSize)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputePoolAddress is a free data retrieval call binding the contract method 0x9dd19a52.
//
// Solidity: function computePoolAddress(uint256 listID, uint256 indexSize) view returns(address poolAddress)
func (_Controller *ControllerSession) ComputePoolAddress(listID *big.Int, indexSize *big.Int) (common.Address, error) {
	return _Controller.Contract.ComputePoolAddress(&_Controller.CallOpts, listID, indexSize)
}

// ComputePoolAddress is a free data retrieval call binding the contract method 0x9dd19a52.
//
// Solidity: function computePoolAddress(uint256 listID, uint256 indexSize) view returns(address poolAddress)
func (_Controller *ControllerCallerSession) ComputePoolAddress(listID *big.Int, indexSize *big.Int) (common.Address, error) {
	return _Controller.Contract.ComputePoolAddress(&_Controller.CallOpts, listID, indexSize)
}

// ComputeSellerAddress is a free data retrieval call binding the contract method 0x582edc9b.
//
// Solidity: function computeSellerAddress(address poolAddress) view returns(address sellerAddress)
func (_Controller *ControllerCaller) ComputeSellerAddress(opts *bind.CallOpts, poolAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "computeSellerAddress", poolAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeSellerAddress is a free data retrieval call binding the contract method 0x582edc9b.
//
// Solidity: function computeSellerAddress(address poolAddress) view returns(address sellerAddress)
func (_Controller *ControllerSession) ComputeSellerAddress(poolAddress common.Address) (common.Address, error) {
	return _Controller.Contract.ComputeSellerAddress(&_Controller.CallOpts, poolAddress)
}

// ComputeSellerAddress is a free data retrieval call binding the contract method 0x582edc9b.
//
// Solidity: function computeSellerAddress(address poolAddress) view returns(address sellerAddress)
func (_Controller *ControllerCallerSession) ComputeSellerAddress(poolAddress common.Address) (common.Address, error) {
	return _Controller.Contract.ComputeSellerAddress(&_Controller.CallOpts, poolAddress)
}

// DefaultExitFeeRecipient is a free data retrieval call binding the contract method 0xa8eb80ab.
//
// Solidity: function defaultExitFeeRecipient() view returns(address)
func (_Controller *ControllerCaller) DefaultExitFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "defaultExitFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultExitFeeRecipient is a free data retrieval call binding the contract method 0xa8eb80ab.
//
// Solidity: function defaultExitFeeRecipient() view returns(address)
func (_Controller *ControllerSession) DefaultExitFeeRecipient() (common.Address, error) {
	return _Controller.Contract.DefaultExitFeeRecipient(&_Controller.CallOpts)
}

// DefaultExitFeeRecipient is a free data retrieval call binding the contract method 0xa8eb80ab.
//
// Solidity: function defaultExitFeeRecipient() view returns(address)
func (_Controller *ControllerCallerSession) DefaultExitFeeRecipient() (common.Address, error) {
	return _Controller.Contract.DefaultExitFeeRecipient(&_Controller.CallOpts)
}

// DefaultSellerPremium is a free data retrieval call binding the contract method 0x372accd2.
//
// Solidity: function defaultSellerPremium() view returns(uint8)
func (_Controller *ControllerCaller) DefaultSellerPremium(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "defaultSellerPremium")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultSellerPremium is a free data retrieval call binding the contract method 0x372accd2.
//
// Solidity: function defaultSellerPremium() view returns(uint8)
func (_Controller *ControllerSession) DefaultSellerPremium() (uint8, error) {
	return _Controller.Contract.DefaultSellerPremium(&_Controller.CallOpts)
}

// DefaultSellerPremium is a free data retrieval call binding the contract method 0x372accd2.
//
// Solidity: function defaultSellerPremium() view returns(uint8)
func (_Controller *ControllerCallerSession) DefaultSellerPremium() (uint8, error) {
	return _Controller.Contract.DefaultSellerPremium(&_Controller.CallOpts)
}

// GetInitialTokensAndBalances is a free data retrieval call binding the contract method 0x774cf0cb.
//
// Solidity: function getInitialTokensAndBalances(uint256 listID, uint256 indexSize, uint256 wethValue) view returns(address[] tokens, uint256[] balances)
func (_Controller *ControllerCaller) GetInitialTokensAndBalances(opts *bind.CallOpts, listID *big.Int, indexSize *big.Int, wethValue *big.Int) (struct {
	Tokens   []common.Address
	Balances []*big.Int
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getInitialTokensAndBalances", listID, indexSize, wethValue)

	outstruct := new(struct {
		Tokens   []common.Address
		Balances []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = out[0].([]common.Address)
	outstruct.Balances = out[1].([]*big.Int)

	return *outstruct, err

}

// GetInitialTokensAndBalances is a free data retrieval call binding the contract method 0x774cf0cb.
//
// Solidity: function getInitialTokensAndBalances(uint256 listID, uint256 indexSize, uint256 wethValue) view returns(address[] tokens, uint256[] balances)
func (_Controller *ControllerSession) GetInitialTokensAndBalances(listID *big.Int, indexSize *big.Int, wethValue *big.Int) (struct {
	Tokens   []common.Address
	Balances []*big.Int
}, error) {
	return _Controller.Contract.GetInitialTokensAndBalances(&_Controller.CallOpts, listID, indexSize, wethValue)
}

// GetInitialTokensAndBalances is a free data retrieval call binding the contract method 0x774cf0cb.
//
// Solidity: function getInitialTokensAndBalances(uint256 listID, uint256 indexSize, uint256 wethValue) view returns(address[] tokens, uint256[] balances)
func (_Controller *ControllerCallerSession) GetInitialTokensAndBalances(listID *big.Int, indexSize *big.Int, wethValue *big.Int) (struct {
	Tokens   []common.Address
	Balances []*big.Int
}, error) {
	return _Controller.Contract.GetInitialTokensAndBalances(&_Controller.CallOpts, listID, indexSize, wethValue)
}

// GetSortedAndFilteredTokensAndScores is a free data retrieval call binding the contract method 0xf2d75471.
//
// Solidity: function getSortedAndFilteredTokensAndScores(uint256 listID) view returns(address[] tokens, uint256[] scores)
func (_Controller *ControllerCaller) GetSortedAndFilteredTokensAndScores(opts *bind.CallOpts, listID *big.Int) (struct {
	Tokens []common.Address
	Scores []*big.Int
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getSortedAndFilteredTokensAndScores", listID)

	outstruct := new(struct {
		Tokens []common.Address
		Scores []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = out[0].([]common.Address)
	outstruct.Scores = out[1].([]*big.Int)

	return *outstruct, err

}

// GetSortedAndFilteredTokensAndScores is a free data retrieval call binding the contract method 0xf2d75471.
//
// Solidity: function getSortedAndFilteredTokensAndScores(uint256 listID) view returns(address[] tokens, uint256[] scores)
func (_Controller *ControllerSession) GetSortedAndFilteredTokensAndScores(listID *big.Int) (struct {
	Tokens []common.Address
	Scores []*big.Int
}, error) {
	return _Controller.Contract.GetSortedAndFilteredTokensAndScores(&_Controller.CallOpts, listID)
}

// GetSortedAndFilteredTokensAndScores is a free data retrieval call binding the contract method 0xf2d75471.
//
// Solidity: function getSortedAndFilteredTokensAndScores(uint256 listID) view returns(address[] tokens, uint256[] scores)
func (_Controller *ControllerCallerSession) GetSortedAndFilteredTokensAndScores(listID *big.Int) (struct {
	Tokens []common.Address
	Scores []*big.Int
}, error) {
	return _Controller.Contract.GetSortedAndFilteredTokensAndScores(&_Controller.CallOpts, listID)
}

// GetTokenList is a free data retrieval call binding the contract method 0x41b63bd8.
//
// Solidity: function getTokenList(uint256 listID) view returns(address[] tokens)
func (_Controller *ControllerCaller) GetTokenList(opts *bind.CallOpts, listID *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getTokenList", listID)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenList is a free data retrieval call binding the contract method 0x41b63bd8.
//
// Solidity: function getTokenList(uint256 listID) view returns(address[] tokens)
func (_Controller *ControllerSession) GetTokenList(listID *big.Int) ([]common.Address, error) {
	return _Controller.Contract.GetTokenList(&_Controller.CallOpts, listID)
}

// GetTokenList is a free data retrieval call binding the contract method 0x41b63bd8.
//
// Solidity: function getTokenList(uint256 listID) view returns(address[] tokens)
func (_Controller *ControllerCallerSession) GetTokenList(listID *big.Int) ([]common.Address, error) {
	return _Controller.Contract.GetTokenList(&_Controller.CallOpts, listID)
}

// GetTokenListConfig is a free data retrieval call binding the contract method 0x13b75360.
//
// Solidity: function getTokenListConfig(uint256 listID) view returns(address scoringStrategy, uint128 minimumScore, uint128 maximumScore)
func (_Controller *ControllerCaller) GetTokenListConfig(opts *bind.CallOpts, listID *big.Int) (struct {
	ScoringStrategy common.Address
	MinimumScore    *big.Int
	MaximumScore    *big.Int
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getTokenListConfig", listID)

	outstruct := new(struct {
		ScoringStrategy common.Address
		MinimumScore    *big.Int
		MaximumScore    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScoringStrategy = out[0].(common.Address)
	outstruct.MinimumScore = out[1].(*big.Int)
	outstruct.MaximumScore = out[2].(*big.Int)

	return *outstruct, err

}

// GetTokenListConfig is a free data retrieval call binding the contract method 0x13b75360.
//
// Solidity: function getTokenListConfig(uint256 listID) view returns(address scoringStrategy, uint128 minimumScore, uint128 maximumScore)
func (_Controller *ControllerSession) GetTokenListConfig(listID *big.Int) (struct {
	ScoringStrategy common.Address
	MinimumScore    *big.Int
	MaximumScore    *big.Int
}, error) {
	return _Controller.Contract.GetTokenListConfig(&_Controller.CallOpts, listID)
}

// GetTokenListConfig is a free data retrieval call binding the contract method 0x13b75360.
//
// Solidity: function getTokenListConfig(uint256 listID) view returns(address scoringStrategy, uint128 minimumScore, uint128 maximumScore)
func (_Controller *ControllerCallerSession) GetTokenListConfig(listID *big.Int) (struct {
	ScoringStrategy common.Address
	MinimumScore    *big.Int
	MaximumScore    *big.Int
}, error) {
	return _Controller.Contract.GetTokenListConfig(&_Controller.CallOpts, listID)
}

// GetTokenScores is a free data retrieval call binding the contract method 0x297c3430.
//
// Solidity: function getTokenScores(uint256 listID, address[] tokens) view returns(uint256[] scores)
func (_Controller *ControllerCaller) GetTokenScores(opts *bind.CallOpts, listID *big.Int, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getTokenScores", listID, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokenScores is a free data retrieval call binding the contract method 0x297c3430.
//
// Solidity: function getTokenScores(uint256 listID, address[] tokens) view returns(uint256[] scores)
func (_Controller *ControllerSession) GetTokenScores(listID *big.Int, tokens []common.Address) ([]*big.Int, error) {
	return _Controller.Contract.GetTokenScores(&_Controller.CallOpts, listID, tokens)
}

// GetTokenScores is a free data retrieval call binding the contract method 0x297c3430.
//
// Solidity: function getTokenScores(uint256 listID, address[] tokens) view returns(uint256[] scores)
func (_Controller *ControllerCallerSession) GetTokenScores(listID *big.Int, tokens []common.Address) ([]*big.Int, error) {
	return _Controller.Contract.GetTokenScores(&_Controller.CallOpts, listID, tokens)
}

// GetTopTokensAndScores is a free data retrieval call binding the contract method 0x861a9ec9.
//
// Solidity: function getTopTokensAndScores(uint256 listID, uint256 count) view returns(address[] tokens, uint256[] scores)
func (_Controller *ControllerCaller) GetTopTokensAndScores(opts *bind.CallOpts, listID *big.Int, count *big.Int) (struct {
	Tokens []common.Address
	Scores []*big.Int
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getTopTokensAndScores", listID, count)

	outstruct := new(struct {
		Tokens []common.Address
		Scores []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = out[0].([]common.Address)
	outstruct.Scores = out[1].([]*big.Int)

	return *outstruct, err

}

// GetTopTokensAndScores is a free data retrieval call binding the contract method 0x861a9ec9.
//
// Solidity: function getTopTokensAndScores(uint256 listID, uint256 count) view returns(address[] tokens, uint256[] scores)
func (_Controller *ControllerSession) GetTopTokensAndScores(listID *big.Int, count *big.Int) (struct {
	Tokens []common.Address
	Scores []*big.Int
}, error) {
	return _Controller.Contract.GetTopTokensAndScores(&_Controller.CallOpts, listID, count)
}

// GetTopTokensAndScores is a free data retrieval call binding the contract method 0x861a9ec9.
//
// Solidity: function getTopTokensAndScores(uint256 listID, uint256 count) view returns(address[] tokens, uint256[] scores)
func (_Controller *ControllerCallerSession) GetTopTokensAndScores(listID *big.Int, count *big.Int) (struct {
	Tokens []common.Address
	Scores []*big.Int
}, error) {
	return _Controller.Contract.GetTopTokensAndScores(&_Controller.CallOpts, listID, count)
}

// IndexPoolMetadata is a free data retrieval call binding the contract method 0x9dfab092.
//
// Solidity: function indexPoolMetadata(address ) view returns(bool initialized, uint16 listID, uint8 indexSize, uint8 reweighIndex, uint64 lastReweigh)
func (_Controller *ControllerCaller) IndexPoolMetadata(opts *bind.CallOpts, arg0 common.Address) (struct {
	Initialized  bool
	ListID       uint16
	IndexSize    uint8
	ReweighIndex uint8
	LastReweigh  uint64
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "indexPoolMetadata", arg0)

	outstruct := new(struct {
		Initialized  bool
		ListID       uint16
		IndexSize    uint8
		ReweighIndex uint8
		LastReweigh  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Initialized = out[0].(bool)
	outstruct.ListID = out[1].(uint16)
	outstruct.IndexSize = out[2].(uint8)
	outstruct.ReweighIndex = out[3].(uint8)
	outstruct.LastReweigh = out[4].(uint64)

	return *outstruct, err

}

// IndexPoolMetadata is a free data retrieval call binding the contract method 0x9dfab092.
//
// Solidity: function indexPoolMetadata(address ) view returns(bool initialized, uint16 listID, uint8 indexSize, uint8 reweighIndex, uint64 lastReweigh)
func (_Controller *ControllerSession) IndexPoolMetadata(arg0 common.Address) (struct {
	Initialized  bool
	ListID       uint16
	IndexSize    uint8
	ReweighIndex uint8
	LastReweigh  uint64
}, error) {
	return _Controller.Contract.IndexPoolMetadata(&_Controller.CallOpts, arg0)
}

// IndexPoolMetadata is a free data retrieval call binding the contract method 0x9dfab092.
//
// Solidity: function indexPoolMetadata(address ) view returns(bool initialized, uint16 listID, uint8 indexSize, uint8 reweighIndex, uint64 lastReweigh)
func (_Controller *ControllerCallerSession) IndexPoolMetadata(arg0 common.Address) (struct {
	Initialized  bool
	ListID       uint16
	IndexSize    uint8
	ReweighIndex uint8
	LastReweigh  uint64
}, error) {
	return _Controller.Contract.IndexPoolMetadata(&_Controller.CallOpts, arg0)
}

// IsTokenInlist is a free data retrieval call binding the contract method 0x8a15f964.
//
// Solidity: function isTokenInlist(uint256 listID, address token) view returns(bool)
func (_Controller *ControllerCaller) IsTokenInlist(opts *bind.CallOpts, listID *big.Int, token common.Address) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "isTokenInlist", listID, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenInlist is a free data retrieval call binding the contract method 0x8a15f964.
//
// Solidity: function isTokenInlist(uint256 listID, address token) view returns(bool)
func (_Controller *ControllerSession) IsTokenInlist(listID *big.Int, token common.Address) (bool, error) {
	return _Controller.Contract.IsTokenInlist(&_Controller.CallOpts, listID, token)
}

// IsTokenInlist is a free data retrieval call binding the contract method 0x8a15f964.
//
// Solidity: function isTokenInlist(uint256 listID, address token) view returns(bool)
func (_Controller *ControllerCallerSession) IsTokenInlist(listID *big.Int, token common.Address) (bool, error) {
	return _Controller.Contract.IsTokenInlist(&_Controller.CallOpts, listID, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerCallerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_Controller *ControllerCaller) PoolFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "poolFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_Controller *ControllerSession) PoolFactory() (common.Address, error) {
	return _Controller.Contract.PoolFactory(&_Controller.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_Controller *ControllerCallerSession) PoolFactory() (common.Address, error) {
	return _Controller.Contract.PoolFactory(&_Controller.CallOpts)
}

// ProxyManager is a free data retrieval call binding the contract method 0x96b4ca70.
//
// Solidity: function proxyManager() view returns(address)
func (_Controller *ControllerCaller) ProxyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "proxyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyManager is a free data retrieval call binding the contract method 0x96b4ca70.
//
// Solidity: function proxyManager() view returns(address)
func (_Controller *ControllerSession) ProxyManager() (common.Address, error) {
	return _Controller.Contract.ProxyManager(&_Controller.CallOpts)
}

// ProxyManager is a free data retrieval call binding the contract method 0x96b4ca70.
//
// Solidity: function proxyManager() view returns(address)
func (_Controller *ControllerCallerSession) ProxyManager() (common.Address, error) {
	return _Controller.Contract.ProxyManager(&_Controller.CallOpts)
}

// TokenListCount is a free data retrieval call binding the contract method 0xa43ad7e7.
//
// Solidity: function tokenListCount() view returns(uint256)
func (_Controller *ControllerCaller) TokenListCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "tokenListCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenListCount is a free data retrieval call binding the contract method 0xa43ad7e7.
//
// Solidity: function tokenListCount() view returns(uint256)
func (_Controller *ControllerSession) TokenListCount() (*big.Int, error) {
	return _Controller.Contract.TokenListCount(&_Controller.CallOpts)
}

// TokenListCount is a free data retrieval call binding the contract method 0xa43ad7e7.
//
// Solidity: function tokenListCount() view returns(uint256)
func (_Controller *ControllerCallerSession) TokenListCount() (*big.Int, error) {
	return _Controller.Contract.TokenListCount(&_Controller.CallOpts)
}

// UniswapOracle is a free data retrieval call binding the contract method 0x120c6c5b.
//
// Solidity: function uniswapOracle() view returns(address)
func (_Controller *ControllerCaller) UniswapOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "uniswapOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapOracle is a free data retrieval call binding the contract method 0x120c6c5b.
//
// Solidity: function uniswapOracle() view returns(address)
func (_Controller *ControllerSession) UniswapOracle() (common.Address, error) {
	return _Controller.Contract.UniswapOracle(&_Controller.CallOpts)
}

// UniswapOracle is a free data retrieval call binding the contract method 0x120c6c5b.
//
// Solidity: function uniswapOracle() view returns(address)
func (_Controller *ControllerCallerSession) UniswapOracle() (common.Address, error) {
	return _Controller.Contract.UniswapOracle(&_Controller.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0x57016b0a.
//
// Solidity: function addToken(uint256 listID, address token) returns()
func (_Controller *ControllerTransactor) AddToken(opts *bind.TransactOpts, listID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addToken", listID, token)
}

// AddToken is a paid mutator transaction binding the contract method 0x57016b0a.
//
// Solidity: function addToken(uint256 listID, address token) returns()
func (_Controller *ControllerSession) AddToken(listID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddToken(&_Controller.TransactOpts, listID, token)
}

// AddToken is a paid mutator transaction binding the contract method 0x57016b0a.
//
// Solidity: function addToken(uint256 listID, address token) returns()
func (_Controller *ControllerTransactorSession) AddToken(listID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddToken(&_Controller.TransactOpts, listID, token)
}

// AddTokens is a paid mutator transaction binding the contract method 0x6b5b8808.
//
// Solidity: function addTokens(uint256 listID, address[] tokens) returns()
func (_Controller *ControllerTransactor) AddTokens(opts *bind.TransactOpts, listID *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addTokens", listID, tokens)
}

// AddTokens is a paid mutator transaction binding the contract method 0x6b5b8808.
//
// Solidity: function addTokens(uint256 listID, address[] tokens) returns()
func (_Controller *ControllerSession) AddTokens(listID *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddTokens(&_Controller.TransactOpts, listID, tokens)
}

// AddTokens is a paid mutator transaction binding the contract method 0x6b5b8808.
//
// Solidity: function addTokens(uint256 listID, address[] tokens) returns()
func (_Controller *ControllerTransactorSession) AddTokens(listID *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddTokens(&_Controller.TransactOpts, listID, tokens)
}

// CreateTokenList is a paid mutator transaction binding the contract method 0x88d6c096.
//
// Solidity: function createTokenList(bytes32 metadataHash, address scoringStrategy, uint128 minimumScore, uint128 maximumScore) returns()
func (_Controller *ControllerTransactor) CreateTokenList(opts *bind.TransactOpts, metadataHash [32]byte, scoringStrategy common.Address, minimumScore *big.Int, maximumScore *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "createTokenList", metadataHash, scoringStrategy, minimumScore, maximumScore)
}

// CreateTokenList is a paid mutator transaction binding the contract method 0x88d6c096.
//
// Solidity: function createTokenList(bytes32 metadataHash, address scoringStrategy, uint128 minimumScore, uint128 maximumScore) returns()
func (_Controller *ControllerSession) CreateTokenList(metadataHash [32]byte, scoringStrategy common.Address, minimumScore *big.Int, maximumScore *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.CreateTokenList(&_Controller.TransactOpts, metadataHash, scoringStrategy, minimumScore, maximumScore)
}

// CreateTokenList is a paid mutator transaction binding the contract method 0x88d6c096.
//
// Solidity: function createTokenList(bytes32 metadataHash, address scoringStrategy, uint128 minimumScore, uint128 maximumScore) returns()
func (_Controller *ControllerTransactorSession) CreateTokenList(metadataHash [32]byte, scoringStrategy common.Address, minimumScore *big.Int, maximumScore *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.CreateTokenList(&_Controller.TransactOpts, metadataHash, scoringStrategy, minimumScore, maximumScore)
}

// DelegateCompLikeTokenFromPool is a paid mutator transaction binding the contract method 0xc6acb34f.
//
// Solidity: function delegateCompLikeTokenFromPool(address pool, address token, address delegatee) returns()
func (_Controller *ControllerTransactor) DelegateCompLikeTokenFromPool(opts *bind.TransactOpts, pool common.Address, token common.Address, delegatee common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "delegateCompLikeTokenFromPool", pool, token, delegatee)
}

// DelegateCompLikeTokenFromPool is a paid mutator transaction binding the contract method 0xc6acb34f.
//
// Solidity: function delegateCompLikeTokenFromPool(address pool, address token, address delegatee) returns()
func (_Controller *ControllerSession) DelegateCompLikeTokenFromPool(pool common.Address, token common.Address, delegatee common.Address) (*types.Transaction, error) {
	return _Controller.Contract.DelegateCompLikeTokenFromPool(&_Controller.TransactOpts, pool, token, delegatee)
}

// DelegateCompLikeTokenFromPool is a paid mutator transaction binding the contract method 0xc6acb34f.
//
// Solidity: function delegateCompLikeTokenFromPool(address pool, address token, address delegatee) returns()
func (_Controller *ControllerTransactorSession) DelegateCompLikeTokenFromPool(pool common.Address, token common.Address, delegatee common.Address) (*types.Transaction, error) {
	return _Controller.Contract.DelegateCompLikeTokenFromPool(&_Controller.TransactOpts, pool, token, delegatee)
}

// FinishPreparedIndexPool is a paid mutator transaction binding the contract method 0x6a7a93ed.
//
// Solidity: function finishPreparedIndexPool(address poolAddress, address[] tokens, uint256[] balances) returns()
func (_Controller *ControllerTransactor) FinishPreparedIndexPool(opts *bind.TransactOpts, poolAddress common.Address, tokens []common.Address, balances []*big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "finishPreparedIndexPool", poolAddress, tokens, balances)
}

// FinishPreparedIndexPool is a paid mutator transaction binding the contract method 0x6a7a93ed.
//
// Solidity: function finishPreparedIndexPool(address poolAddress, address[] tokens, uint256[] balances) returns()
func (_Controller *ControllerSession) FinishPreparedIndexPool(poolAddress common.Address, tokens []common.Address, balances []*big.Int) (*types.Transaction, error) {
	return _Controller.Contract.FinishPreparedIndexPool(&_Controller.TransactOpts, poolAddress, tokens, balances)
}

// FinishPreparedIndexPool is a paid mutator transaction binding the contract method 0x6a7a93ed.
//
// Solidity: function finishPreparedIndexPool(address poolAddress, address[] tokens, uint256[] balances) returns()
func (_Controller *ControllerTransactorSession) FinishPreparedIndexPool(poolAddress common.Address, tokens []common.Address, balances []*big.Int) (*types.Transaction, error) {
	return _Controller.Contract.FinishPreparedIndexPool(&_Controller.TransactOpts, poolAddress, tokens, balances)
}

// ForceReindexPool is a paid mutator transaction binding the contract method 0x0fae3828.
//
// Solidity: function forceReindexPool(address poolAddress) returns()
func (_Controller *ControllerTransactor) ForceReindexPool(opts *bind.TransactOpts, poolAddress common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "forceReindexPool", poolAddress)
}

// ForceReindexPool is a paid mutator transaction binding the contract method 0x0fae3828.
//
// Solidity: function forceReindexPool(address poolAddress) returns()
func (_Controller *ControllerSession) ForceReindexPool(poolAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.ForceReindexPool(&_Controller.TransactOpts, poolAddress)
}

// ForceReindexPool is a paid mutator transaction binding the contract method 0x0fae3828.
//
// Solidity: function forceReindexPool(address poolAddress) returns()
func (_Controller *ControllerTransactorSession) ForceReindexPool(poolAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.ForceReindexPool(&_Controller.TransactOpts, poolAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Controller *ControllerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Controller *ControllerSession) Initialize() (*types.Transaction, error) {
	return _Controller.Contract.Initialize(&_Controller.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Controller *ControllerTransactorSession) Initialize() (*types.Transaction, error) {
	return _Controller.Contract.Initialize(&_Controller.TransactOpts)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address circuitBreaker_) returns()
func (_Controller *ControllerTransactor) Initialize0(opts *bind.TransactOpts, circuitBreaker_ common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "initialize0", circuitBreaker_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address circuitBreaker_) returns()
func (_Controller *ControllerSession) Initialize0(circuitBreaker_ common.Address) (*types.Transaction, error) {
	return _Controller.Contract.Initialize0(&_Controller.TransactOpts, circuitBreaker_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address circuitBreaker_) returns()
func (_Controller *ControllerTransactorSession) Initialize0(circuitBreaker_ common.Address) (*types.Transaction, error) {
	return _Controller.Contract.Initialize0(&_Controller.TransactOpts, circuitBreaker_)
}

// PrepareIndexPool is a paid mutator transaction binding the contract method 0x74a58783.
//
// Solidity: function prepareIndexPool(uint256 listID, uint256 indexSize, uint256 initialWethValue, string name, string symbol) returns(address poolAddress, address initializerAddress)
func (_Controller *ControllerTransactor) PrepareIndexPool(opts *bind.TransactOpts, listID *big.Int, indexSize *big.Int, initialWethValue *big.Int, name string, symbol string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "prepareIndexPool", listID, indexSize, initialWethValue, name, symbol)
}

// PrepareIndexPool is a paid mutator transaction binding the contract method 0x74a58783.
//
// Solidity: function prepareIndexPool(uint256 listID, uint256 indexSize, uint256 initialWethValue, string name, string symbol) returns(address poolAddress, address initializerAddress)
func (_Controller *ControllerSession) PrepareIndexPool(listID *big.Int, indexSize *big.Int, initialWethValue *big.Int, name string, symbol string) (*types.Transaction, error) {
	return _Controller.Contract.PrepareIndexPool(&_Controller.TransactOpts, listID, indexSize, initialWethValue, name, symbol)
}

// PrepareIndexPool is a paid mutator transaction binding the contract method 0x74a58783.
//
// Solidity: function prepareIndexPool(uint256 listID, uint256 indexSize, uint256 initialWethValue, string name, string symbol) returns(address poolAddress, address initializerAddress)
func (_Controller *ControllerTransactorSession) PrepareIndexPool(listID *big.Int, indexSize *big.Int, initialWethValue *big.Int, name string, symbol string) (*types.Transaction, error) {
	return _Controller.Contract.PrepareIndexPool(&_Controller.TransactOpts, listID, indexSize, initialWethValue, name, symbol)
}

// ReindexPool is a paid mutator transaction binding the contract method 0x50b1e342.
//
// Solidity: function reindexPool(address poolAddress) returns()
func (_Controller *ControllerTransactor) ReindexPool(opts *bind.TransactOpts, poolAddress common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "reindexPool", poolAddress)
}

// ReindexPool is a paid mutator transaction binding the contract method 0x50b1e342.
//
// Solidity: function reindexPool(address poolAddress) returns()
func (_Controller *ControllerSession) ReindexPool(poolAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.ReindexPool(&_Controller.TransactOpts, poolAddress)
}

// ReindexPool is a paid mutator transaction binding the contract method 0x50b1e342.
//
// Solidity: function reindexPool(address poolAddress) returns()
func (_Controller *ControllerTransactorSession) ReindexPool(poolAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.ReindexPool(&_Controller.TransactOpts, poolAddress)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x33fe5676.
//
// Solidity: function removeToken(uint256 listID, address token) returns()
func (_Controller *ControllerTransactor) RemoveToken(opts *bind.TransactOpts, listID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "removeToken", listID, token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x33fe5676.
//
// Solidity: function removeToken(uint256 listID, address token) returns()
func (_Controller *ControllerSession) RemoveToken(listID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RemoveToken(&_Controller.TransactOpts, listID, token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x33fe5676.
//
// Solidity: function removeToken(uint256 listID, address token) returns()
func (_Controller *ControllerTransactorSession) RemoveToken(listID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RemoveToken(&_Controller.TransactOpts, listID, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Controller.Contract.RenounceOwnership(&_Controller.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Controller.Contract.RenounceOwnership(&_Controller.TransactOpts)
}

// ReweighPool is a paid mutator transaction binding the contract method 0xe9819daf.
//
// Solidity: function reweighPool(address poolAddress) returns()
func (_Controller *ControllerTransactor) ReweighPool(opts *bind.TransactOpts, poolAddress common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "reweighPool", poolAddress)
}

// ReweighPool is a paid mutator transaction binding the contract method 0xe9819daf.
//
// Solidity: function reweighPool(address poolAddress) returns()
func (_Controller *ControllerSession) ReweighPool(poolAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.ReweighPool(&_Controller.TransactOpts, poolAddress)
}

// ReweighPool is a paid mutator transaction binding the contract method 0xe9819daf.
//
// Solidity: function reweighPool(address poolAddress) returns()
func (_Controller *ControllerTransactorSession) ReweighPool(poolAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.ReweighPool(&_Controller.TransactOpts, poolAddress)
}

// SetCircuitBreaker is a paid mutator transaction binding the contract method 0x82beee89.
//
// Solidity: function setCircuitBreaker(address circuitBreaker_) returns()
func (_Controller *ControllerTransactor) SetCircuitBreaker(opts *bind.TransactOpts, circuitBreaker_ common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "setCircuitBreaker", circuitBreaker_)
}

// SetCircuitBreaker is a paid mutator transaction binding the contract method 0x82beee89.
//
// Solidity: function setCircuitBreaker(address circuitBreaker_) returns()
func (_Controller *ControllerSession) SetCircuitBreaker(circuitBreaker_ common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetCircuitBreaker(&_Controller.TransactOpts, circuitBreaker_)
}

// SetCircuitBreaker is a paid mutator transaction binding the contract method 0x82beee89.
//
// Solidity: function setCircuitBreaker(address circuitBreaker_) returns()
func (_Controller *ControllerTransactorSession) SetCircuitBreaker(circuitBreaker_ common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetCircuitBreaker(&_Controller.TransactOpts, circuitBreaker_)
}

// SetDefaultSellerPremium is a paid mutator transaction binding the contract method 0xdf8e238a.
//
// Solidity: function setDefaultSellerPremium(uint8 _defaultSellerPremium) returns()
func (_Controller *ControllerTransactor) SetDefaultSellerPremium(opts *bind.TransactOpts, _defaultSellerPremium uint8) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "setDefaultSellerPremium", _defaultSellerPremium)
}

// SetDefaultSellerPremium is a paid mutator transaction binding the contract method 0xdf8e238a.
//
// Solidity: function setDefaultSellerPremium(uint8 _defaultSellerPremium) returns()
func (_Controller *ControllerSession) SetDefaultSellerPremium(_defaultSellerPremium uint8) (*types.Transaction, error) {
	return _Controller.Contract.SetDefaultSellerPremium(&_Controller.TransactOpts, _defaultSellerPremium)
}

// SetDefaultSellerPremium is a paid mutator transaction binding the contract method 0xdf8e238a.
//
// Solidity: function setDefaultSellerPremium(uint8 _defaultSellerPremium) returns()
func (_Controller *ControllerTransactorSession) SetDefaultSellerPremium(_defaultSellerPremium uint8) (*types.Transaction, error) {
	return _Controller.Contract.SetDefaultSellerPremium(&_Controller.TransactOpts, _defaultSellerPremium)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x1eccc185.
//
// Solidity: function setPublicSwap(address indexPool_, bool publicSwap) returns()
func (_Controller *ControllerTransactor) SetPublicSwap(opts *bind.TransactOpts, indexPool_ common.Address, publicSwap bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "setPublicSwap", indexPool_, publicSwap)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x1eccc185.
//
// Solidity: function setPublicSwap(address indexPool_, bool publicSwap) returns()
func (_Controller *ControllerSession) SetPublicSwap(indexPool_ common.Address, publicSwap bool) (*types.Transaction, error) {
	return _Controller.Contract.SetPublicSwap(&_Controller.TransactOpts, indexPool_, publicSwap)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x1eccc185.
//
// Solidity: function setPublicSwap(address indexPool_, bool publicSwap) returns()
func (_Controller *ControllerTransactorSession) SetPublicSwap(indexPool_ common.Address, publicSwap bool) (*types.Transaction, error) {
	return _Controller.Contract.SetPublicSwap(&_Controller.TransactOpts, indexPool_, publicSwap)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x991991c7.
//
// Solidity: function setSwapFee(address poolAddress, uint256 swapFee) returns()
func (_Controller *ControllerTransactor) SetSwapFee(opts *bind.TransactOpts, poolAddress common.Address, swapFee *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "setSwapFee", poolAddress, swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x991991c7.
//
// Solidity: function setSwapFee(address poolAddress, uint256 swapFee) returns()
func (_Controller *ControllerSession) SetSwapFee(poolAddress common.Address, swapFee *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetSwapFee(&_Controller.TransactOpts, poolAddress, swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x991991c7.
//
// Solidity: function setSwapFee(address poolAddress, uint256 swapFee) returns()
func (_Controller *ControllerTransactorSession) SetSwapFee(poolAddress common.Address, swapFee *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetSwapFee(&_Controller.TransactOpts, poolAddress, swapFee)
}

// SortAndFilterTokens is a paid mutator transaction binding the contract method 0x88b9813a.
//
// Solidity: function sortAndFilterTokens(uint256 listID) returns()
func (_Controller *ControllerTransactor) SortAndFilterTokens(opts *bind.TransactOpts, listID *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "sortAndFilterTokens", listID)
}

// SortAndFilterTokens is a paid mutator transaction binding the contract method 0x88b9813a.
//
// Solidity: function sortAndFilterTokens(uint256 listID) returns()
func (_Controller *ControllerSession) SortAndFilterTokens(listID *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SortAndFilterTokens(&_Controller.TransactOpts, listID)
}

// SortAndFilterTokens is a paid mutator transaction binding the contract method 0x88b9813a.
//
// Solidity: function sortAndFilterTokens(uint256 listID) returns()
func (_Controller *ControllerTransactorSession) SortAndFilterTokens(listID *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SortAndFilterTokens(&_Controller.TransactOpts, listID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Controller *ControllerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Controller *ControllerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Controller.Contract.TransferOwnership(&_Controller.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Controller *ControllerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Controller.Contract.TransferOwnership(&_Controller.TransactOpts, newOwner)
}

// UpdateMinimumBalance is a paid mutator transaction binding the contract method 0x034b904e.
//
// Solidity: function updateMinimumBalance(address pool, address tokenAddress) returns()
func (_Controller *ControllerTransactor) UpdateMinimumBalance(opts *bind.TransactOpts, pool common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "updateMinimumBalance", pool, tokenAddress)
}

// UpdateMinimumBalance is a paid mutator transaction binding the contract method 0x034b904e.
//
// Solidity: function updateMinimumBalance(address pool, address tokenAddress) returns()
func (_Controller *ControllerSession) UpdateMinimumBalance(pool common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.UpdateMinimumBalance(&_Controller.TransactOpts, pool, tokenAddress)
}

// UpdateMinimumBalance is a paid mutator transaction binding the contract method 0x034b904e.
//
// Solidity: function updateMinimumBalance(address pool, address tokenAddress) returns()
func (_Controller *ControllerTransactorSession) UpdateMinimumBalance(pool common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.UpdateMinimumBalance(&_Controller.TransactOpts, pool, tokenAddress)
}

// UpdateSellerPremium is a paid mutator transaction binding the contract method 0xdcdf7e16.
//
// Solidity: function updateSellerPremium(address tokenSeller, uint8 premiumPercent) returns()
func (_Controller *ControllerTransactor) UpdateSellerPremium(opts *bind.TransactOpts, tokenSeller common.Address, premiumPercent uint8) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "updateSellerPremium", tokenSeller, premiumPercent)
}

// UpdateSellerPremium is a paid mutator transaction binding the contract method 0xdcdf7e16.
//
// Solidity: function updateSellerPremium(address tokenSeller, uint8 premiumPercent) returns()
func (_Controller *ControllerSession) UpdateSellerPremium(tokenSeller common.Address, premiumPercent uint8) (*types.Transaction, error) {
	return _Controller.Contract.UpdateSellerPremium(&_Controller.TransactOpts, tokenSeller, premiumPercent)
}

// UpdateSellerPremium is a paid mutator transaction binding the contract method 0xdcdf7e16.
//
// Solidity: function updateSellerPremium(address tokenSeller, uint8 premiumPercent) returns()
func (_Controller *ControllerTransactorSession) UpdateSellerPremium(tokenSeller common.Address, premiumPercent uint8) (*types.Transaction, error) {
	return _Controller.Contract.UpdateSellerPremium(&_Controller.TransactOpts, tokenSeller, premiumPercent)
}

// UpdateTokenPrices is a paid mutator transaction binding the contract method 0xde16eb99.
//
// Solidity: function updateTokenPrices(uint256 listID) returns(bool[] pricesUpdated)
func (_Controller *ControllerTransactor) UpdateTokenPrices(opts *bind.TransactOpts, listID *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "updateTokenPrices", listID)
}

// UpdateTokenPrices is a paid mutator transaction binding the contract method 0xde16eb99.
//
// Solidity: function updateTokenPrices(uint256 listID) returns(bool[] pricesUpdated)
func (_Controller *ControllerSession) UpdateTokenPrices(listID *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.UpdateTokenPrices(&_Controller.TransactOpts, listID)
}

// UpdateTokenPrices is a paid mutator transaction binding the contract method 0xde16eb99.
//
// Solidity: function updateTokenPrices(uint256 listID) returns(bool[] pricesUpdated)
func (_Controller *ControllerTransactorSession) UpdateTokenPrices(listID *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.UpdateTokenPrices(&_Controller.TransactOpts, listID)
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
	ListID      *big.Int
	IndexSize   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewPoolInitializer is a free log retrieval operation binding the contract event 0x7ad23833dba658b2bdc6f260fb60a3240b3868086ad60b4e42276f5eeba73e6a.
//
// Solidity: event NewPoolInitializer(address pool, address initializer, uint256 listID, uint256 indexSize)
func (_Controller *ControllerFilterer) FilterNewPoolInitializer(opts *bind.FilterOpts) (*ControllerNewPoolInitializerIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NewPoolInitializer")
	if err != nil {
		return nil, err
	}
	return &ControllerNewPoolInitializerIterator{contract: _Controller.contract, event: "NewPoolInitializer", logs: logs, sub: sub}, nil
}

// WatchNewPoolInitializer is a free log subscription operation binding the contract event 0x7ad23833dba658b2bdc6f260fb60a3240b3868086ad60b4e42276f5eeba73e6a.
//
// Solidity: event NewPoolInitializer(address pool, address initializer, uint256 listID, uint256 indexSize)
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
// Solidity: event NewPoolInitializer(address pool, address initializer, uint256 listID, uint256 indexSize)
func (_Controller *ControllerFilterer) ParseNewPoolInitializer(log types.Log) (*ControllerNewPoolInitializer, error) {
	event := new(ControllerNewPoolInitializer)
	if err := _Controller.contract.UnpackLog(event, "NewPoolInitializer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Controller contract.
type ControllerOwnershipTransferredIterator struct {
	Event *ControllerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerOwnershipTransferred)
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
		it.Event = new(ControllerOwnershipTransferred)
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
func (it *ControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerOwnershipTransferred represents a OwnershipTransferred event raised by the Controller contract.
type ControllerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Controller *ControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ControllerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ControllerOwnershipTransferredIterator{contract: _Controller.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Controller *ControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ControllerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerOwnershipTransferred)
				if err := _Controller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Controller *ControllerFilterer) ParseOwnershipTransferred(log types.Log) (*ControllerOwnershipTransferred, error) {
	event := new(ControllerOwnershipTransferred)
	if err := _Controller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
	ListID             *big.Int
	IndexSize          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPoolInitialized is a free log retrieval operation binding the contract event 0x9ff2050ac7faae9dc192e1cc9abe73b18a9b849f9b43f509914d80fa104d5903.
//
// Solidity: event PoolInitialized(address pool, address unboundTokenSeller, uint256 listID, uint256 indexSize)
func (_Controller *ControllerFilterer) FilterPoolInitialized(opts *bind.FilterOpts) (*ControllerPoolInitializedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "PoolInitialized")
	if err != nil {
		return nil, err
	}
	return &ControllerPoolInitializedIterator{contract: _Controller.contract, event: "PoolInitialized", logs: logs, sub: sub}, nil
}

// WatchPoolInitialized is a free log subscription operation binding the contract event 0x9ff2050ac7faae9dc192e1cc9abe73b18a9b849f9b43f509914d80fa104d5903.
//
// Solidity: event PoolInitialized(address pool, address unboundTokenSeller, uint256 listID, uint256 indexSize)
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
// Solidity: event PoolInitialized(address pool, address unboundTokenSeller, uint256 listID, uint256 indexSize)
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
	Token  common.Address
	ListID *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0xf4c563a3ea86ff1f4275e8c207df0375a51963f2b831b7bf4da8be938d92876c.
//
// Solidity: event TokenAdded(address token, uint256 listID)
func (_Controller *ControllerFilterer) FilterTokenAdded(opts *bind.FilterOpts) (*ControllerTokenAddedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return &ControllerTokenAddedIterator{contract: _Controller.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0xf4c563a3ea86ff1f4275e8c207df0375a51963f2b831b7bf4da8be938d92876c.
//
// Solidity: event TokenAdded(address token, uint256 listID)
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
// Solidity: event TokenAdded(address token, uint256 listID)
func (_Controller *ControllerFilterer) ParseTokenAdded(log types.Log) (*ControllerTokenAdded, error) {
	event := new(ControllerTokenAdded)
	if err := _Controller.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerTokenListAddedIterator is returned from FilterTokenListAdded and is used to iterate over the raw logs and unpacked data for TokenListAdded events raised by the Controller contract.
type ControllerTokenListAddedIterator struct {
	Event *ControllerTokenListAdded // Event containing the contract specifics and raw log

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
func (it *ControllerTokenListAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerTokenListAdded)
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
		it.Event = new(ControllerTokenListAdded)
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
func (it *ControllerTokenListAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerTokenListAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerTokenListAdded represents a TokenListAdded event raised by the Controller contract.
type ControllerTokenListAdded struct {
	ListID          *big.Int
	MetadataHash    [32]byte
	ScoringStrategy common.Address
	MinimumScore    *big.Int
	MaximumScore    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenListAdded is a free log retrieval operation binding the contract event 0xd5e8e673f6e24dc7176b58a5620dcee488e08e985c6fd9db96b1b6d214004117.
//
// Solidity: event TokenListAdded(uint256 listID, bytes32 metadataHash, address scoringStrategy, uint128 minimumScore, uint128 maximumScore)
func (_Controller *ControllerFilterer) FilterTokenListAdded(opts *bind.FilterOpts) (*ControllerTokenListAddedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "TokenListAdded")
	if err != nil {
		return nil, err
	}
	return &ControllerTokenListAddedIterator{contract: _Controller.contract, event: "TokenListAdded", logs: logs, sub: sub}, nil
}

// WatchTokenListAdded is a free log subscription operation binding the contract event 0xd5e8e673f6e24dc7176b58a5620dcee488e08e985c6fd9db96b1b6d214004117.
//
// Solidity: event TokenListAdded(uint256 listID, bytes32 metadataHash, address scoringStrategy, uint128 minimumScore, uint128 maximumScore)
func (_Controller *ControllerFilterer) WatchTokenListAdded(opts *bind.WatchOpts, sink chan<- *ControllerTokenListAdded) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "TokenListAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerTokenListAdded)
				if err := _Controller.contract.UnpackLog(event, "TokenListAdded", log); err != nil {
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

// ParseTokenListAdded is a log parse operation binding the contract event 0xd5e8e673f6e24dc7176b58a5620dcee488e08e985c6fd9db96b1b6d214004117.
//
// Solidity: event TokenListAdded(uint256 listID, bytes32 metadataHash, address scoringStrategy, uint128 minimumScore, uint128 maximumScore)
func (_Controller *ControllerFilterer) ParseTokenListAdded(log types.Log) (*ControllerTokenListAdded, error) {
	event := new(ControllerTokenListAdded)
	if err := _Controller.contract.UnpackLog(event, "TokenListAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerTokenListSortedIterator is returned from FilterTokenListSorted and is used to iterate over the raw logs and unpacked data for TokenListSorted events raised by the Controller contract.
type ControllerTokenListSortedIterator struct {
	Event *ControllerTokenListSorted // Event containing the contract specifics and raw log

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
func (it *ControllerTokenListSortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerTokenListSorted)
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
		it.Event = new(ControllerTokenListSorted)
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
func (it *ControllerTokenListSortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerTokenListSortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerTokenListSorted represents a TokenListSorted event raised by the Controller contract.
type ControllerTokenListSorted struct {
	ListID *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenListSorted is a free log retrieval operation binding the contract event 0xf53ac44f7bda4111b4677882e4e57cfab74f80d4de1e8b3245754369b48db601.
//
// Solidity: event TokenListSorted(uint256 listID)
func (_Controller *ControllerFilterer) FilterTokenListSorted(opts *bind.FilterOpts) (*ControllerTokenListSortedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "TokenListSorted")
	if err != nil {
		return nil, err
	}
	return &ControllerTokenListSortedIterator{contract: _Controller.contract, event: "TokenListSorted", logs: logs, sub: sub}, nil
}

// WatchTokenListSorted is a free log subscription operation binding the contract event 0xf53ac44f7bda4111b4677882e4e57cfab74f80d4de1e8b3245754369b48db601.
//
// Solidity: event TokenListSorted(uint256 listID)
func (_Controller *ControllerFilterer) WatchTokenListSorted(opts *bind.WatchOpts, sink chan<- *ControllerTokenListSorted) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "TokenListSorted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerTokenListSorted)
				if err := _Controller.contract.UnpackLog(event, "TokenListSorted", log); err != nil {
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

// ParseTokenListSorted is a log parse operation binding the contract event 0xf53ac44f7bda4111b4677882e4e57cfab74f80d4de1e8b3245754369b48db601.
//
// Solidity: event TokenListSorted(uint256 listID)
func (_Controller *ControllerFilterer) ParseTokenListSorted(log types.Log) (*ControllerTokenListSorted, error) {
	event := new(ControllerTokenListSorted)
	if err := _Controller.contract.UnpackLog(event, "TokenListSorted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the Controller contract.
type ControllerTokenRemovedIterator struct {
	Event *ControllerTokenRemoved // Event containing the contract specifics and raw log

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
func (it *ControllerTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerTokenRemoved)
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
		it.Event = new(ControllerTokenRemoved)
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
func (it *ControllerTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerTokenRemoved represents a TokenRemoved event raised by the Controller contract.
type ControllerTokenRemoved struct {
	Token  common.Address
	ListID *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0xbe9bb4bdca0a094babd75e3a98b1d2e2390633430d0a2f6e2b9970e2ee03fb2e.
//
// Solidity: event TokenRemoved(address token, uint256 listID)
func (_Controller *ControllerFilterer) FilterTokenRemoved(opts *bind.FilterOpts) (*ControllerTokenRemovedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return &ControllerTokenRemovedIterator{contract: _Controller.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0xbe9bb4bdca0a094babd75e3a98b1d2e2390633430d0a2f6e2b9970e2ee03fb2e.
//
// Solidity: event TokenRemoved(address token, uint256 listID)
func (_Controller *ControllerFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *ControllerTokenRemoved) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerTokenRemoved)
				if err := _Controller.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0xbe9bb4bdca0a094babd75e3a98b1d2e2390633430d0a2f6e2b9970e2ee03fb2e.
//
// Solidity: event TokenRemoved(address token, uint256 listID)
func (_Controller *ControllerFilterer) ParseTokenRemoved(log types.Log) (*ControllerTokenRemoved, error) {
	event := new(ControllerTokenRemoved)
	if err := _Controller.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

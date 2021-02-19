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

// FixedPointuq112x112 is an auto generated low-level Go binding around an user-defined struct.
type FixedPointuq112x112 struct {
	X *big.Int
}

// ControllerABI is the input ABI used to generate the binding from.
const ControllerABI = "[{\"inputs\":[{\"internalType\":\"contractIIndexedUniswapV2Oracle\",\"name\":\"uniswapOracle_\",\"type\":\"address\"},{\"internalType\":\"contractIPoolFactory\",\"name\":\"poolFactory_\",\"type\":\"address\"},{\"internalType\":\"contractIDelegateCallProxyManager\",\"name\":\"proxyManager_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defaultExitFeeRecipient_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useFullyDilutedMarketCaps\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"minMarketCap\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"maxMarketCap\",\"type\":\"uint128\"}],\"name\":\"CategoryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"CategorySorted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumMarketCapSqrtController.WeightingFormula\",\"name\":\"formula\",\"type\":\"uint8\"}],\"name\":\"NewPoolInitializer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"unboundTokenSeller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"}],\"name\":\"PoolInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INITIALIZER_IMPLEMENTATION_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LONG_TWAP_MAX_TIME_ELAPSED\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LONG_TWAP_MIN_TIME_ELAPSED\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CATEGORY_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_INDEX_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SORT_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_INDEX_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_IMPLEMENTATION_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_REWEIGH_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWEIGHS_BEFORE_REINDEX\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SELLER_IMPLEMENTATION_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHORT_TWAP_MAX_TIME_ELAPSED\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHORT_TWAP_MIN_TIME_ELAPSED\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WEIGHT_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"addTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"categoryIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circuitBreaker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circulatingMarketCapOracle\",\"outputs\":[{\"internalType\":\"contractICirculatingMarketCapOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"computeInitializerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"initializerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"}],\"name\":\"computePoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"computeSellerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sellerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumMarketCapSqrtController.WeightingFormula\",\"name\":\"formula\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"marketCaps\",\"type\":\"uint256[]\"}],\"name\":\"computeWeights\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"_x\",\"type\":\"uint224\"}],\"internalType\":\"structFixedPoint.uq112x112[]\",\"name\":\"fractionalWeights\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"useFullyDilutedMarketCaps\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"minMarketCap\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"maxMarketCap\",\"type\":\"uint112\"}],\"name\":\"createCategory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultExitFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultSellerPremium\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegateCompLikeTokenFromPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"name\":\"finishPreparedIndexPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"forceReindexPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"getCategoryConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"useFullyDilutedMarketCaps\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"minMarketCap\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"maxMarketCap\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"getCategoryTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getCirculatingMarketCaps\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"marketCaps\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getFullyDilutedMarketCaps\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"marketCaps\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"enumMarketCapSqrtController.WeightingFormula\",\"name\":\"formula\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"},{\"internalType\":\"uint144\",\"name\":\"wethValue\",\"type\":\"uint144\"}],\"name\":\"getInitialTokensAndBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"useFullyDilutedMarketCaps\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getMarketCaps\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"marketCaps\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"getSortedAndFilteredTokensAndMarketCaps\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"categoryTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"marketCaps\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getTopCategoryTokensAndMarketCaps\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"categoryTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"marketCaps\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"hasCategory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"indexPoolMetadata\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"categoryID\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"indexSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"reweighIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"lastReweigh\",\"type\":\"uint64\"},{\"internalType\":\"enumMarketCapSqrtController.WeightingFormula\",\"name\":\"formula\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"circulatingMarketCapOracle_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"circuitBreaker_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"circulatingMarketCapOracle_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenInCategory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolFactory\",\"outputs\":[{\"internalType\":\"contractIPoolFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialWethValue\",\"type\":\"uint256\"},{\"internalType\":\"enumMarketCapSqrtController.WeightingFormula\",\"name\":\"formula\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"prepareIndexPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initializerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyManager\",\"outputs\":[{\"internalType\":\"contractIDelegateCallProxyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"reindexPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"reweighPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"circuitBreaker_\",\"type\":\"address\"}],\"name\":\"setCircuitBreaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"circulatingMarketCapOracle_\",\"type\":\"address\"}],\"name\":\"setCirculatingMarketCapOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_defaultSellerPremium\",\"type\":\"uint8\"}],\"name\":\"setDefaultSellerPremium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"indexPool_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"publicSwap\",\"type\":\"bool\"}],\"name\":\"setPublicSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"sortAndFilterTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapOracle\",\"outputs\":[{\"internalType\":\"contractIIndexedUniswapV2Oracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"updateCategoryMarketCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"updateCategoryPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"pricesUpdated\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIIndexPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"updateMinimumBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenSeller\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"premiumPercent\",\"type\":\"uint8\"}],\"name\":\"updateSellerPremium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// LONGTWAPMAXTIMEELAPSED is a free data retrieval call binding the contract method 0x43ec8923.
//
// Solidity: function LONG_TWAP_MAX_TIME_ELAPSED() view returns(uint32)
func (_Controller *ControllerCaller) LONGTWAPMAXTIMEELAPSED(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "LONG_TWAP_MAX_TIME_ELAPSED")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LONGTWAPMAXTIMEELAPSED is a free data retrieval call binding the contract method 0x43ec8923.
//
// Solidity: function LONG_TWAP_MAX_TIME_ELAPSED() view returns(uint32)
func (_Controller *ControllerSession) LONGTWAPMAXTIMEELAPSED() (uint32, error) {
	return _Controller.Contract.LONGTWAPMAXTIMEELAPSED(&_Controller.CallOpts)
}

// LONGTWAPMAXTIMEELAPSED is a free data retrieval call binding the contract method 0x43ec8923.
//
// Solidity: function LONG_TWAP_MAX_TIME_ELAPSED() view returns(uint32)
func (_Controller *ControllerCallerSession) LONGTWAPMAXTIMEELAPSED() (uint32, error) {
	return _Controller.Contract.LONGTWAPMAXTIMEELAPSED(&_Controller.CallOpts)
}

// LONGTWAPMINTIMEELAPSED is a free data retrieval call binding the contract method 0x70ffd3e0.
//
// Solidity: function LONG_TWAP_MIN_TIME_ELAPSED() view returns(uint32)
func (_Controller *ControllerCaller) LONGTWAPMINTIMEELAPSED(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "LONG_TWAP_MIN_TIME_ELAPSED")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LONGTWAPMINTIMEELAPSED is a free data retrieval call binding the contract method 0x70ffd3e0.
//
// Solidity: function LONG_TWAP_MIN_TIME_ELAPSED() view returns(uint32)
func (_Controller *ControllerSession) LONGTWAPMINTIMEELAPSED() (uint32, error) {
	return _Controller.Contract.LONGTWAPMINTIMEELAPSED(&_Controller.CallOpts)
}

// LONGTWAPMINTIMEELAPSED is a free data retrieval call binding the contract method 0x70ffd3e0.
//
// Solidity: function LONG_TWAP_MIN_TIME_ELAPSED() view returns(uint32)
func (_Controller *ControllerCallerSession) LONGTWAPMINTIMEELAPSED() (uint32, error) {
	return _Controller.Contract.LONGTWAPMINTIMEELAPSED(&_Controller.CallOpts)
}

// MAXCATEGORYTOKENS is a free data retrieval call binding the contract method 0x84fca9cd.
//
// Solidity: function MAX_CATEGORY_TOKENS() view returns(uint256)
func (_Controller *ControllerCaller) MAXCATEGORYTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "MAX_CATEGORY_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCATEGORYTOKENS is a free data retrieval call binding the contract method 0x84fca9cd.
//
// Solidity: function MAX_CATEGORY_TOKENS() view returns(uint256)
func (_Controller *ControllerSession) MAXCATEGORYTOKENS() (*big.Int, error) {
	return _Controller.Contract.MAXCATEGORYTOKENS(&_Controller.CallOpts)
}

// MAXCATEGORYTOKENS is a free data retrieval call binding the contract method 0x84fca9cd.
//
// Solidity: function MAX_CATEGORY_TOKENS() view returns(uint256)
func (_Controller *ControllerCallerSession) MAXCATEGORYTOKENS() (*big.Int, error) {
	return _Controller.Contract.MAXCATEGORYTOKENS(&_Controller.CallOpts)
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

// MAXSORTDELAY is a free data retrieval call binding the contract method 0xfdcfa2ec.
//
// Solidity: function MAX_SORT_DELAY() view returns(uint256)
func (_Controller *ControllerCaller) MAXSORTDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "MAX_SORT_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSORTDELAY is a free data retrieval call binding the contract method 0xfdcfa2ec.
//
// Solidity: function MAX_SORT_DELAY() view returns(uint256)
func (_Controller *ControllerSession) MAXSORTDELAY() (*big.Int, error) {
	return _Controller.Contract.MAXSORTDELAY(&_Controller.CallOpts)
}

// MAXSORTDELAY is a free data retrieval call binding the contract method 0xfdcfa2ec.
//
// Solidity: function MAX_SORT_DELAY() view returns(uint256)
func (_Controller *ControllerCallerSession) MAXSORTDELAY() (*big.Int, error) {
	return _Controller.Contract.MAXSORTDELAY(&_Controller.CallOpts)
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

// WEIGHTMULTIPLIER is a free data retrieval call binding the contract method 0xc290d8bd.
//
// Solidity: function WEIGHT_MULTIPLIER() view returns(uint256)
func (_Controller *ControllerCaller) WEIGHTMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "WEIGHT_MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WEIGHTMULTIPLIER is a free data retrieval call binding the contract method 0xc290d8bd.
//
// Solidity: function WEIGHT_MULTIPLIER() view returns(uint256)
func (_Controller *ControllerSession) WEIGHTMULTIPLIER() (*big.Int, error) {
	return _Controller.Contract.WEIGHTMULTIPLIER(&_Controller.CallOpts)
}

// WEIGHTMULTIPLIER is a free data retrieval call binding the contract method 0xc290d8bd.
//
// Solidity: function WEIGHT_MULTIPLIER() view returns(uint256)
func (_Controller *ControllerCallerSession) WEIGHTMULTIPLIER() (*big.Int, error) {
	return _Controller.Contract.WEIGHTMULTIPLIER(&_Controller.CallOpts)
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
// Solidity: function computePoolAddress(uint256 categoryID, uint256 indexSize) view returns(address poolAddress)
func (_Controller *ControllerCaller) ComputePoolAddress(opts *bind.CallOpts, categoryID *big.Int, indexSize *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "computePoolAddress", categoryID, indexSize)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputePoolAddress is a free data retrieval call binding the contract method 0x9dd19a52.
//
// Solidity: function computePoolAddress(uint256 categoryID, uint256 indexSize) view returns(address poolAddress)
func (_Controller *ControllerSession) ComputePoolAddress(categoryID *big.Int, indexSize *big.Int) (common.Address, error) {
	return _Controller.Contract.ComputePoolAddress(&_Controller.CallOpts, categoryID, indexSize)
}

// ComputePoolAddress is a free data retrieval call binding the contract method 0x9dd19a52.
//
// Solidity: function computePoolAddress(uint256 categoryID, uint256 indexSize) view returns(address poolAddress)
func (_Controller *ControllerCallerSession) ComputePoolAddress(categoryID *big.Int, indexSize *big.Int) (common.Address, error) {
	return _Controller.Contract.ComputePoolAddress(&_Controller.CallOpts, categoryID, indexSize)
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

// ComputeWeights is a free data retrieval call binding the contract method 0xa4d09816.
//
// Solidity: function computeWeights(uint8 formula, uint256[] marketCaps) view returns((uint224)[] fractionalWeights)
func (_Controller *ControllerCaller) ComputeWeights(opts *bind.CallOpts, formula uint8, marketCaps []*big.Int) ([]FixedPointuq112x112, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "computeWeights", formula, marketCaps)

	if err != nil {
		return *new([]FixedPointuq112x112), err
	}

	out0 := *abi.ConvertType(out[0], new([]FixedPointuq112x112)).(*[]FixedPointuq112x112)

	return out0, err

}

// ComputeWeights is a free data retrieval call binding the contract method 0xa4d09816.
//
// Solidity: function computeWeights(uint8 formula, uint256[] marketCaps) view returns((uint224)[] fractionalWeights)
func (_Controller *ControllerSession) ComputeWeights(formula uint8, marketCaps []*big.Int) ([]FixedPointuq112x112, error) {
	return _Controller.Contract.ComputeWeights(&_Controller.CallOpts, formula, marketCaps)
}

// ComputeWeights is a free data retrieval call binding the contract method 0xa4d09816.
//
// Solidity: function computeWeights(uint8 formula, uint256[] marketCaps) view returns((uint224)[] fractionalWeights)
func (_Controller *ControllerCallerSession) ComputeWeights(formula uint8, marketCaps []*big.Int) ([]FixedPointuq112x112, error) {
	return _Controller.Contract.ComputeWeights(&_Controller.CallOpts, formula, marketCaps)
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

// GetCategoryConfig is a free data retrieval call binding the contract method 0x28b996c1.
//
// Solidity: function getCategoryConfig(uint256 categoryID) view returns(bool useFullyDilutedMarketCaps, uint112 minMarketCap, uint112 maxMarketCap)
func (_Controller *ControllerCaller) GetCategoryConfig(opts *bind.CallOpts, categoryID *big.Int) (struct {
	UseFullyDilutedMarketCaps bool
	MinMarketCap              *big.Int
	MaxMarketCap              *big.Int
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getCategoryConfig", categoryID)

	outstruct := new(struct {
		UseFullyDilutedMarketCaps bool
		MinMarketCap              *big.Int
		MaxMarketCap              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UseFullyDilutedMarketCaps = out[0].(bool)
	outstruct.MinMarketCap = out[1].(*big.Int)
	outstruct.MaxMarketCap = out[2].(*big.Int)

	return *outstruct, err

}

// GetCategoryConfig is a free data retrieval call binding the contract method 0x28b996c1.
//
// Solidity: function getCategoryConfig(uint256 categoryID) view returns(bool useFullyDilutedMarketCaps, uint112 minMarketCap, uint112 maxMarketCap)
func (_Controller *ControllerSession) GetCategoryConfig(categoryID *big.Int) (struct {
	UseFullyDilutedMarketCaps bool
	MinMarketCap              *big.Int
	MaxMarketCap              *big.Int
}, error) {
	return _Controller.Contract.GetCategoryConfig(&_Controller.CallOpts, categoryID)
}

// GetCategoryConfig is a free data retrieval call binding the contract method 0x28b996c1.
//
// Solidity: function getCategoryConfig(uint256 categoryID) view returns(bool useFullyDilutedMarketCaps, uint112 minMarketCap, uint112 maxMarketCap)
func (_Controller *ControllerCallerSession) GetCategoryConfig(categoryID *big.Int) (struct {
	UseFullyDilutedMarketCaps bool
	MinMarketCap              *big.Int
	MaxMarketCap              *big.Int
}, error) {
	return _Controller.Contract.GetCategoryConfig(&_Controller.CallOpts, categoryID)
}

// GetCategoryTokens is a free data retrieval call binding the contract method 0xde105dbd.
//
// Solidity: function getCategoryTokens(uint256 categoryID) view returns(address[] tokens)
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
// Solidity: function getCategoryTokens(uint256 categoryID) view returns(address[] tokens)
func (_Controller *ControllerSession) GetCategoryTokens(categoryID *big.Int) ([]common.Address, error) {
	return _Controller.Contract.GetCategoryTokens(&_Controller.CallOpts, categoryID)
}

// GetCategoryTokens is a free data retrieval call binding the contract method 0xde105dbd.
//
// Solidity: function getCategoryTokens(uint256 categoryID) view returns(address[] tokens)
func (_Controller *ControllerCallerSession) GetCategoryTokens(categoryID *big.Int) ([]common.Address, error) {
	return _Controller.Contract.GetCategoryTokens(&_Controller.CallOpts, categoryID)
}

// GetCirculatingMarketCaps is a free data retrieval call binding the contract method 0xbf99568f.
//
// Solidity: function getCirculatingMarketCaps(address[] tokens) view returns(uint256[] marketCaps)
func (_Controller *ControllerCaller) GetCirculatingMarketCaps(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getCirculatingMarketCaps", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCirculatingMarketCaps is a free data retrieval call binding the contract method 0xbf99568f.
//
// Solidity: function getCirculatingMarketCaps(address[] tokens) view returns(uint256[] marketCaps)
func (_Controller *ControllerSession) GetCirculatingMarketCaps(tokens []common.Address) ([]*big.Int, error) {
	return _Controller.Contract.GetCirculatingMarketCaps(&_Controller.CallOpts, tokens)
}

// GetCirculatingMarketCaps is a free data retrieval call binding the contract method 0xbf99568f.
//
// Solidity: function getCirculatingMarketCaps(address[] tokens) view returns(uint256[] marketCaps)
func (_Controller *ControllerCallerSession) GetCirculatingMarketCaps(tokens []common.Address) ([]*big.Int, error) {
	return _Controller.Contract.GetCirculatingMarketCaps(&_Controller.CallOpts, tokens)
}

// GetFullyDilutedMarketCaps is a free data retrieval call binding the contract method 0xc0f40a6b.
//
// Solidity: function getFullyDilutedMarketCaps(address[] tokens) view returns(uint256[] marketCaps)
func (_Controller *ControllerCaller) GetFullyDilutedMarketCaps(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getFullyDilutedMarketCaps", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetFullyDilutedMarketCaps is a free data retrieval call binding the contract method 0xc0f40a6b.
//
// Solidity: function getFullyDilutedMarketCaps(address[] tokens) view returns(uint256[] marketCaps)
func (_Controller *ControllerSession) GetFullyDilutedMarketCaps(tokens []common.Address) ([]*big.Int, error) {
	return _Controller.Contract.GetFullyDilutedMarketCaps(&_Controller.CallOpts, tokens)
}

// GetFullyDilutedMarketCaps is a free data retrieval call binding the contract method 0xc0f40a6b.
//
// Solidity: function getFullyDilutedMarketCaps(address[] tokens) view returns(uint256[] marketCaps)
func (_Controller *ControllerCallerSession) GetFullyDilutedMarketCaps(tokens []common.Address) ([]*big.Int, error) {
	return _Controller.Contract.GetFullyDilutedMarketCaps(&_Controller.CallOpts, tokens)
}

// GetInitialTokensAndBalances is a free data retrieval call binding the contract method 0xeef34dfe.
//
// Solidity: function getInitialTokensAndBalances(uint256 categoryID, uint8 formula, uint256 indexSize, uint144 wethValue) view returns(address[] tokens, uint256[] balances)
func (_Controller *ControllerCaller) GetInitialTokensAndBalances(opts *bind.CallOpts, categoryID *big.Int, formula uint8, indexSize *big.Int, wethValue *big.Int) (struct {
	Tokens   []common.Address
	Balances []*big.Int
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getInitialTokensAndBalances", categoryID, formula, indexSize, wethValue)

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

// GetInitialTokensAndBalances is a free data retrieval call binding the contract method 0xeef34dfe.
//
// Solidity: function getInitialTokensAndBalances(uint256 categoryID, uint8 formula, uint256 indexSize, uint144 wethValue) view returns(address[] tokens, uint256[] balances)
func (_Controller *ControllerSession) GetInitialTokensAndBalances(categoryID *big.Int, formula uint8, indexSize *big.Int, wethValue *big.Int) (struct {
	Tokens   []common.Address
	Balances []*big.Int
}, error) {
	return _Controller.Contract.GetInitialTokensAndBalances(&_Controller.CallOpts, categoryID, formula, indexSize, wethValue)
}

// GetInitialTokensAndBalances is a free data retrieval call binding the contract method 0xeef34dfe.
//
// Solidity: function getInitialTokensAndBalances(uint256 categoryID, uint8 formula, uint256 indexSize, uint144 wethValue) view returns(address[] tokens, uint256[] balances)
func (_Controller *ControllerCallerSession) GetInitialTokensAndBalances(categoryID *big.Int, formula uint8, indexSize *big.Int, wethValue *big.Int) (struct {
	Tokens   []common.Address
	Balances []*big.Int
}, error) {
	return _Controller.Contract.GetInitialTokensAndBalances(&_Controller.CallOpts, categoryID, formula, indexSize, wethValue)
}

// GetMarketCaps is a free data retrieval call binding the contract method 0x7645e3f2.
//
// Solidity: function getMarketCaps(bool useFullyDilutedMarketCaps, address[] tokens) view returns(uint256[] marketCaps)
func (_Controller *ControllerCaller) GetMarketCaps(opts *bind.CallOpts, useFullyDilutedMarketCaps bool, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getMarketCaps", useFullyDilutedMarketCaps, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetMarketCaps is a free data retrieval call binding the contract method 0x7645e3f2.
//
// Solidity: function getMarketCaps(bool useFullyDilutedMarketCaps, address[] tokens) view returns(uint256[] marketCaps)
func (_Controller *ControllerSession) GetMarketCaps(useFullyDilutedMarketCaps bool, tokens []common.Address) ([]*big.Int, error) {
	return _Controller.Contract.GetMarketCaps(&_Controller.CallOpts, useFullyDilutedMarketCaps, tokens)
}

// GetMarketCaps is a free data retrieval call binding the contract method 0x7645e3f2.
//
// Solidity: function getMarketCaps(bool useFullyDilutedMarketCaps, address[] tokens) view returns(uint256[] marketCaps)
func (_Controller *ControllerCallerSession) GetMarketCaps(useFullyDilutedMarketCaps bool, tokens []common.Address) ([]*big.Int, error) {
	return _Controller.Contract.GetMarketCaps(&_Controller.CallOpts, useFullyDilutedMarketCaps, tokens)
}

// GetSortedAndFilteredTokensAndMarketCaps is a free data retrieval call binding the contract method 0x7a92d7ce.
//
// Solidity: function getSortedAndFilteredTokensAndMarketCaps(uint256 categoryID) view returns(address[] categoryTokens, uint256[] marketCaps)
func (_Controller *ControllerCaller) GetSortedAndFilteredTokensAndMarketCaps(opts *bind.CallOpts, categoryID *big.Int) (struct {
	CategoryTokens []common.Address
	MarketCaps     []*big.Int
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getSortedAndFilteredTokensAndMarketCaps", categoryID)

	outstruct := new(struct {
		CategoryTokens []common.Address
		MarketCaps     []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CategoryTokens = out[0].([]common.Address)
	outstruct.MarketCaps = out[1].([]*big.Int)

	return *outstruct, err

}

// GetSortedAndFilteredTokensAndMarketCaps is a free data retrieval call binding the contract method 0x7a92d7ce.
//
// Solidity: function getSortedAndFilteredTokensAndMarketCaps(uint256 categoryID) view returns(address[] categoryTokens, uint256[] marketCaps)
func (_Controller *ControllerSession) GetSortedAndFilteredTokensAndMarketCaps(categoryID *big.Int) (struct {
	CategoryTokens []common.Address
	MarketCaps     []*big.Int
}, error) {
	return _Controller.Contract.GetSortedAndFilteredTokensAndMarketCaps(&_Controller.CallOpts, categoryID)
}

// GetSortedAndFilteredTokensAndMarketCaps is a free data retrieval call binding the contract method 0x7a92d7ce.
//
// Solidity: function getSortedAndFilteredTokensAndMarketCaps(uint256 categoryID) view returns(address[] categoryTokens, uint256[] marketCaps)
func (_Controller *ControllerCallerSession) GetSortedAndFilteredTokensAndMarketCaps(categoryID *big.Int) (struct {
	CategoryTokens []common.Address
	MarketCaps     []*big.Int
}, error) {
	return _Controller.Contract.GetSortedAndFilteredTokensAndMarketCaps(&_Controller.CallOpts, categoryID)
}

// GetTopCategoryTokensAndMarketCaps is a free data retrieval call binding the contract method 0xb61beade.
//
// Solidity: function getTopCategoryTokensAndMarketCaps(uint256 categoryID, uint256 count) view returns(address[] categoryTokens, uint256[] marketCaps)
func (_Controller *ControllerCaller) GetTopCategoryTokensAndMarketCaps(opts *bind.CallOpts, categoryID *big.Int, count *big.Int) (struct {
	CategoryTokens []common.Address
	MarketCaps     []*big.Int
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getTopCategoryTokensAndMarketCaps", categoryID, count)

	outstruct := new(struct {
		CategoryTokens []common.Address
		MarketCaps     []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CategoryTokens = out[0].([]common.Address)
	outstruct.MarketCaps = out[1].([]*big.Int)

	return *outstruct, err

}

// GetTopCategoryTokensAndMarketCaps is a free data retrieval call binding the contract method 0xb61beade.
//
// Solidity: function getTopCategoryTokensAndMarketCaps(uint256 categoryID, uint256 count) view returns(address[] categoryTokens, uint256[] marketCaps)
func (_Controller *ControllerSession) GetTopCategoryTokensAndMarketCaps(categoryID *big.Int, count *big.Int) (struct {
	CategoryTokens []common.Address
	MarketCaps     []*big.Int
}, error) {
	return _Controller.Contract.GetTopCategoryTokensAndMarketCaps(&_Controller.CallOpts, categoryID, count)
}

// GetTopCategoryTokensAndMarketCaps is a free data retrieval call binding the contract method 0xb61beade.
//
// Solidity: function getTopCategoryTokensAndMarketCaps(uint256 categoryID, uint256 count) view returns(address[] categoryTokens, uint256[] marketCaps)
func (_Controller *ControllerCallerSession) GetTopCategoryTokensAndMarketCaps(categoryID *big.Int, count *big.Int) (struct {
	CategoryTokens []common.Address
	MarketCaps     []*big.Int
}, error) {
	return _Controller.Contract.GetTopCategoryTokensAndMarketCaps(&_Controller.CallOpts, categoryID, count)
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

// IndexPoolMetadata is a free data retrieval call binding the contract method 0x9dfab092.
//
// Solidity: function indexPoolMetadata(address ) view returns(bool initialized, uint16 categoryID, uint8 indexSize, uint8 reweighIndex, uint64 lastReweigh, uint8 formula)
func (_Controller *ControllerCaller) IndexPoolMetadata(opts *bind.CallOpts, arg0 common.Address) (struct {
	Initialized  bool
	CategoryID   uint16
	IndexSize    uint8
	ReweighIndex uint8
	LastReweigh  uint64
	Formula      uint8
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "indexPoolMetadata", arg0)

	outstruct := new(struct {
		Initialized  bool
		CategoryID   uint16
		IndexSize    uint8
		ReweighIndex uint8
		LastReweigh  uint64
		Formula      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Initialized = out[0].(bool)
	outstruct.CategoryID = out[1].(uint16)
	outstruct.IndexSize = out[2].(uint8)
	outstruct.ReweighIndex = out[3].(uint8)
	outstruct.LastReweigh = out[4].(uint64)
	outstruct.Formula = out[5].(uint8)

	return *outstruct, err

}

// IndexPoolMetadata is a free data retrieval call binding the contract method 0x9dfab092.
//
// Solidity: function indexPoolMetadata(address ) view returns(bool initialized, uint16 categoryID, uint8 indexSize, uint8 reweighIndex, uint64 lastReweigh, uint8 formula)
func (_Controller *ControllerSession) IndexPoolMetadata(arg0 common.Address) (struct {
	Initialized  bool
	CategoryID   uint16
	IndexSize    uint8
	ReweighIndex uint8
	LastReweigh  uint64
	Formula      uint8
}, error) {
	return _Controller.Contract.IndexPoolMetadata(&_Controller.CallOpts, arg0)
}

// IndexPoolMetadata is a free data retrieval call binding the contract method 0x9dfab092.
//
// Solidity: function indexPoolMetadata(address ) view returns(bool initialized, uint16 categoryID, uint8 indexSize, uint8 reweighIndex, uint64 lastReweigh, uint8 formula)
func (_Controller *ControllerCallerSession) IndexPoolMetadata(arg0 common.Address) (struct {
	Initialized  bool
	CategoryID   uint16
	IndexSize    uint8
	ReweighIndex uint8
	LastReweigh  uint64
	Formula      uint8
}, error) {
	return _Controller.Contract.IndexPoolMetadata(&_Controller.CallOpts, arg0)
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

// CreateCategory is a paid mutator transaction binding the contract method 0xa151e0f4.
//
// Solidity: function createCategory(bytes32 metadataHash, bool useFullyDilutedMarketCaps, uint112 minMarketCap, uint112 maxMarketCap) returns()
func (_Controller *ControllerTransactor) CreateCategory(opts *bind.TransactOpts, metadataHash [32]byte, useFullyDilutedMarketCaps bool, minMarketCap *big.Int, maxMarketCap *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "createCategory", metadataHash, useFullyDilutedMarketCaps, minMarketCap, maxMarketCap)
}

// CreateCategory is a paid mutator transaction binding the contract method 0xa151e0f4.
//
// Solidity: function createCategory(bytes32 metadataHash, bool useFullyDilutedMarketCaps, uint112 minMarketCap, uint112 maxMarketCap) returns()
func (_Controller *ControllerSession) CreateCategory(metadataHash [32]byte, useFullyDilutedMarketCaps bool, minMarketCap *big.Int, maxMarketCap *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.CreateCategory(&_Controller.TransactOpts, metadataHash, useFullyDilutedMarketCaps, minMarketCap, maxMarketCap)
}

// CreateCategory is a paid mutator transaction binding the contract method 0xa151e0f4.
//
// Solidity: function createCategory(bytes32 metadataHash, bool useFullyDilutedMarketCaps, uint112 minMarketCap, uint112 maxMarketCap) returns()
func (_Controller *ControllerTransactorSession) CreateCategory(metadataHash [32]byte, useFullyDilutedMarketCaps bool, minMarketCap *big.Int, maxMarketCap *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.CreateCategory(&_Controller.TransactOpts, metadataHash, useFullyDilutedMarketCaps, minMarketCap, maxMarketCap)
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

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address circulatingMarketCapOracle_, address circuitBreaker_) returns()
func (_Controller *ControllerTransactor) Initialize(opts *bind.TransactOpts, circulatingMarketCapOracle_ common.Address, circuitBreaker_ common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "initialize", circulatingMarketCapOracle_, circuitBreaker_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address circulatingMarketCapOracle_, address circuitBreaker_) returns()
func (_Controller *ControllerSession) Initialize(circulatingMarketCapOracle_ common.Address, circuitBreaker_ common.Address) (*types.Transaction, error) {
	return _Controller.Contract.Initialize(&_Controller.TransactOpts, circulatingMarketCapOracle_, circuitBreaker_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address circulatingMarketCapOracle_, address circuitBreaker_) returns()
func (_Controller *ControllerTransactorSession) Initialize(circulatingMarketCapOracle_ common.Address, circuitBreaker_ common.Address) (*types.Transaction, error) {
	return _Controller.Contract.Initialize(&_Controller.TransactOpts, circulatingMarketCapOracle_, circuitBreaker_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address circulatingMarketCapOracle_) returns()
func (_Controller *ControllerTransactor) Initialize0(opts *bind.TransactOpts, circulatingMarketCapOracle_ common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "initialize0", circulatingMarketCapOracle_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address circulatingMarketCapOracle_) returns()
func (_Controller *ControllerSession) Initialize0(circulatingMarketCapOracle_ common.Address) (*types.Transaction, error) {
	return _Controller.Contract.Initialize0(&_Controller.TransactOpts, circulatingMarketCapOracle_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address circulatingMarketCapOracle_) returns()
func (_Controller *ControllerTransactorSession) Initialize0(circulatingMarketCapOracle_ common.Address) (*types.Transaction, error) {
	return _Controller.Contract.Initialize0(&_Controller.TransactOpts, circulatingMarketCapOracle_)
}

// PrepareIndexPool is a paid mutator transaction binding the contract method 0xf5bde611.
//
// Solidity: function prepareIndexPool(uint256 categoryID, uint256 indexSize, uint256 initialWethValue, uint8 formula, string name, string symbol) returns(address poolAddress, address initializerAddress)
func (_Controller *ControllerTransactor) PrepareIndexPool(opts *bind.TransactOpts, categoryID *big.Int, indexSize *big.Int, initialWethValue *big.Int, formula uint8, name string, symbol string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "prepareIndexPool", categoryID, indexSize, initialWethValue, formula, name, symbol)
}

// PrepareIndexPool is a paid mutator transaction binding the contract method 0xf5bde611.
//
// Solidity: function prepareIndexPool(uint256 categoryID, uint256 indexSize, uint256 initialWethValue, uint8 formula, string name, string symbol) returns(address poolAddress, address initializerAddress)
func (_Controller *ControllerSession) PrepareIndexPool(categoryID *big.Int, indexSize *big.Int, initialWethValue *big.Int, formula uint8, name string, symbol string) (*types.Transaction, error) {
	return _Controller.Contract.PrepareIndexPool(&_Controller.TransactOpts, categoryID, indexSize, initialWethValue, formula, name, symbol)
}

// PrepareIndexPool is a paid mutator transaction binding the contract method 0xf5bde611.
//
// Solidity: function prepareIndexPool(uint256 categoryID, uint256 indexSize, uint256 initialWethValue, uint8 formula, string name, string symbol) returns(address poolAddress, address initializerAddress)
func (_Controller *ControllerTransactorSession) PrepareIndexPool(categoryID *big.Int, indexSize *big.Int, initialWethValue *big.Int, formula uint8, name string, symbol string) (*types.Transaction, error) {
	return _Controller.Contract.PrepareIndexPool(&_Controller.TransactOpts, categoryID, indexSize, initialWethValue, formula, name, symbol)
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

// SetCirculatingMarketCapOracle is a paid mutator transaction binding the contract method 0x4ee6411e.
//
// Solidity: function setCirculatingMarketCapOracle(address circulatingMarketCapOracle_) returns()
func (_Controller *ControllerTransactor) SetCirculatingMarketCapOracle(opts *bind.TransactOpts, circulatingMarketCapOracle_ common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "setCirculatingMarketCapOracle", circulatingMarketCapOracle_)
}

// SetCirculatingMarketCapOracle is a paid mutator transaction binding the contract method 0x4ee6411e.
//
// Solidity: function setCirculatingMarketCapOracle(address circulatingMarketCapOracle_) returns()
func (_Controller *ControllerSession) SetCirculatingMarketCapOracle(circulatingMarketCapOracle_ common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetCirculatingMarketCapOracle(&_Controller.TransactOpts, circulatingMarketCapOracle_)
}

// SetCirculatingMarketCapOracle is a paid mutator transaction binding the contract method 0x4ee6411e.
//
// Solidity: function setCirculatingMarketCapOracle(address circulatingMarketCapOracle_) returns()
func (_Controller *ControllerTransactorSession) SetCirculatingMarketCapOracle(circulatingMarketCapOracle_ common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetCirculatingMarketCapOracle(&_Controller.TransactOpts, circulatingMarketCapOracle_)
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
// Solidity: function sortAndFilterTokens(uint256 categoryID) returns()
func (_Controller *ControllerTransactor) SortAndFilterTokens(opts *bind.TransactOpts, categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "sortAndFilterTokens", categoryID)
}

// SortAndFilterTokens is a paid mutator transaction binding the contract method 0x88b9813a.
//
// Solidity: function sortAndFilterTokens(uint256 categoryID) returns()
func (_Controller *ControllerSession) SortAndFilterTokens(categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SortAndFilterTokens(&_Controller.TransactOpts, categoryID)
}

// SortAndFilterTokens is a paid mutator transaction binding the contract method 0x88b9813a.
//
// Solidity: function sortAndFilterTokens(uint256 categoryID) returns()
func (_Controller *ControllerTransactorSession) SortAndFilterTokens(categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SortAndFilterTokens(&_Controller.TransactOpts, categoryID)
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

// UpdateCategoryMarketCaps is a paid mutator transaction binding the contract method 0x90bec8ac.
//
// Solidity: function updateCategoryMarketCaps(uint256 categoryID) returns()
func (_Controller *ControllerTransactor) UpdateCategoryMarketCaps(opts *bind.TransactOpts, categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "updateCategoryMarketCaps", categoryID)
}

// UpdateCategoryMarketCaps is a paid mutator transaction binding the contract method 0x90bec8ac.
//
// Solidity: function updateCategoryMarketCaps(uint256 categoryID) returns()
func (_Controller *ControllerSession) UpdateCategoryMarketCaps(categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.UpdateCategoryMarketCaps(&_Controller.TransactOpts, categoryID)
}

// UpdateCategoryMarketCaps is a paid mutator transaction binding the contract method 0x90bec8ac.
//
// Solidity: function updateCategoryMarketCaps(uint256 categoryID) returns()
func (_Controller *ControllerTransactorSession) UpdateCategoryMarketCaps(categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.UpdateCategoryMarketCaps(&_Controller.TransactOpts, categoryID)
}

// UpdateCategoryPrices is a paid mutator transaction binding the contract method 0x6866ada8.
//
// Solidity: function updateCategoryPrices(uint256 categoryID) returns(bool[] pricesUpdated)
func (_Controller *ControllerTransactor) UpdateCategoryPrices(opts *bind.TransactOpts, categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "updateCategoryPrices", categoryID)
}

// UpdateCategoryPrices is a paid mutator transaction binding the contract method 0x6866ada8.
//
// Solidity: function updateCategoryPrices(uint256 categoryID) returns(bool[] pricesUpdated)
func (_Controller *ControllerSession) UpdateCategoryPrices(categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.UpdateCategoryPrices(&_Controller.TransactOpts, categoryID)
}

// UpdateCategoryPrices is a paid mutator transaction binding the contract method 0x6866ada8.
//
// Solidity: function updateCategoryPrices(uint256 categoryID) returns(bool[] pricesUpdated)
func (_Controller *ControllerTransactorSession) UpdateCategoryPrices(categoryID *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.UpdateCategoryPrices(&_Controller.TransactOpts, categoryID)
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
	CategoryID                *big.Int
	MetadataHash              [32]byte
	UseFullyDilutedMarketCaps bool
	MinMarketCap              *big.Int
	MaxMarketCap              *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterCategoryAdded is a free log retrieval operation binding the contract event 0x204e69e7c4d704d675be18559b3a36c31d58c335c1e1027f9a98c65d762441d3.
//
// Solidity: event CategoryAdded(uint256 categoryID, bytes32 metadataHash, bool useFullyDilutedMarketCaps, uint128 minMarketCap, uint128 maxMarketCap)
func (_Controller *ControllerFilterer) FilterCategoryAdded(opts *bind.FilterOpts) (*ControllerCategoryAddedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "CategoryAdded")
	if err != nil {
		return nil, err
	}
	return &ControllerCategoryAddedIterator{contract: _Controller.contract, event: "CategoryAdded", logs: logs, sub: sub}, nil
}

// WatchCategoryAdded is a free log subscription operation binding the contract event 0x204e69e7c4d704d675be18559b3a36c31d58c335c1e1027f9a98c65d762441d3.
//
// Solidity: event CategoryAdded(uint256 categoryID, bytes32 metadataHash, bool useFullyDilutedMarketCaps, uint128 minMarketCap, uint128 maxMarketCap)
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

// ParseCategoryAdded is a log parse operation binding the contract event 0x204e69e7c4d704d675be18559b3a36c31d58c335c1e1027f9a98c65d762441d3.
//
// Solidity: event CategoryAdded(uint256 categoryID, bytes32 metadataHash, bool useFullyDilutedMarketCaps, uint128 minMarketCap, uint128 maxMarketCap)
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
	Formula     uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewPoolInitializer is a free log retrieval operation binding the contract event 0x290721d63c4f9911fe368948944f0df2c9ebce49e6a204da0d0886a758733115.
//
// Solidity: event NewPoolInitializer(address pool, address initializer, uint256 categoryID, uint256 indexSize, uint8 formula)
func (_Controller *ControllerFilterer) FilterNewPoolInitializer(opts *bind.FilterOpts) (*ControllerNewPoolInitializerIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NewPoolInitializer")
	if err != nil {
		return nil, err
	}
	return &ControllerNewPoolInitializerIterator{contract: _Controller.contract, event: "NewPoolInitializer", logs: logs, sub: sub}, nil
}

// WatchNewPoolInitializer is a free log subscription operation binding the contract event 0x290721d63c4f9911fe368948944f0df2c9ebce49e6a204da0d0886a758733115.
//
// Solidity: event NewPoolInitializer(address pool, address initializer, uint256 categoryID, uint256 indexSize, uint8 formula)
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

// ParseNewPoolInitializer is a log parse operation binding the contract event 0x290721d63c4f9911fe368948944f0df2c9ebce49e6a204da0d0886a758733115.
//
// Solidity: event NewPoolInitializer(address pool, address initializer, uint256 categoryID, uint256 indexSize, uint8 formula)
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
	Token      common.Address
	CategoryID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0xbe9bb4bdca0a094babd75e3a98b1d2e2390633430d0a2f6e2b9970e2ee03fb2e.
//
// Solidity: event TokenRemoved(address token, uint256 categoryID)
func (_Controller *ControllerFilterer) FilterTokenRemoved(opts *bind.FilterOpts) (*ControllerTokenRemovedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return &ControllerTokenRemovedIterator{contract: _Controller.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0xbe9bb4bdca0a094babd75e3a98b1d2e2390633430d0a2f6e2b9970e2ee03fb2e.
//
// Solidity: event TokenRemoved(address token, uint256 categoryID)
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
// Solidity: event TokenRemoved(address token, uint256 categoryID)
func (_Controller *ControllerFilterer) ParseTokenRemoved(log types.Log) (*ControllerTokenRemoved, error) {
	event := new(ControllerTokenRemoved)
	if err := _Controller.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

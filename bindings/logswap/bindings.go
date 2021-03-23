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
const LogswapABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"LOG_PUBLIC_SWAP_TOGGLED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_SWAP\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_swapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circuitBreaker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimal\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"outAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inAmount\",\"type\":\"uint256\"}],\"name\":\"emitLogSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"emitPublicSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getUsedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incRando\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setPublicSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setPublicSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LogswapBin is the compiled bytecode used for deploying new contracts.
var LogswapBin = "0x608060405260016000556040518060400160405280600481526020017f43433130000000000000000000000000000000000000000000000000000000008152506001908051906020019062000056929190620000d3565b506012600260006101000a81548160ff021916908360ff1602179055506658d15e176280006003553480156200008b57600080fd5b5030600460016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620001e8565b828054620000e19062000183565b90600052602060002090601f01602090048101928262000105576000855562000151565b82601f106200012057805160ff191683800117855562000151565b8280016001018555821562000151579182015b828111156200015057825182559160200191906001019062000133565b5b50905062000160919062000164565b5090565b5b808211156200017f57600081600090555060010162000165565b5090565b600060028204905060018216806200019c57607f821691505b60208210811415620001b357620001b2620001b9565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b610f7480620001f86000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806376809ce3116100a2578063ad4a5c4b11610071578063ad4a5c4b14610296578063c94ddf8c146102b2578063cc77828d146102ce578063d4cadf68146102ec578063f8b2cb4f1461030a5761010b565b806376809ce31461020c578063832781551461022a578063948d8ce61461024857806395d89b41146102785761010b565b80633018205f116100de5780633018205f1461019857806349b59552146101b65780634aa4e0b5146101d25780635d93a794146102025761010b565b806315e84af91461011057806316efd9411461014057806318160ddd1461015e5780631eccc1851461017c575b600080fd5b61012a600480360381019061012591906108c7565b61033a565b6040516101379190610bd1565b60405180910390f35b610148610415565b6040516101559190610b57565b60405180910390f35b61016661043b565b6040516101739190610bd1565b60405180910390f35b61019660048036038101906101919190610903565b61048b565b005b6101a06104a9565b6040516101ad9190610b57565b60405180910390f35b6101d060048036038101906101cb919061093f565b6104b1565b005b6101ec60048036038101906101e7919061089e565b6104ce565b6040516101f99190610bd1565b60405180910390f35b61020a610527565b005b610214610542565b6040516102219190610c15565b60405180910390f35b610232610555565b60405161023f9190610bd1565b60405180910390f35b610262600480360381019061025d919061089e565b61055b565b60405161026f9190610bd1565b60405180910390f35b6102806105c0565b60405161028d9190610baf565b60405180910390f35b6102b060048036038101906102ab9190610968565b61064e565b005b6102cc60048036038101906102c7919061093f565b6106d0565b005b6102d661070a565b6040516102e39190610b72565b60405180910390f35b6102f46107fc565b6040516103019190610bd1565b60405180910390f35b610324600480360381019061031f919061089e565b610806565b6040516103319190610bd1565b60405180910390f35b600060034310610354576701aa535d3d0c0000905061040f565b600060fa4233600054876040516020016103719493929190610b09565b6040516020818303038152906040528051906020012060001c6103949190610e1d565b9050600060fa4233600054876040516020016103b39493929190610b09565b6040516020818303038152906040528051906020012060001c6103d69190610e1d565b9050670de0b6b3a7640000816103ec9190610cdb565b670de0b6b3a7640000836104009190610cdb565b61040a9190610c85565b925050505b92915050565b600460019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080423360005460405160200161045593929190610acc565b6040516020818303038152906040528051906020012060001c905068056bc75e2d63100000816104859190610cdb565b91505090565b80600460006101000a81548160ff0219169083151502179055505050565b600030905090565b80600460006101000a81548160ff02191690831515021790555050565b6000804233600054856040516020016104ea9493929190610b09565b6040516020818303038152906040528051906020012060001c905060006703782dace9d900008261051b9190610cdb565b90508092505050919050565b60016000808282546105399190610c85565b92505081905550565b600260009054906101000a900460ff1681565b60035481565b60008060644233600054866040516020016105799493929190610b09565b6040516020818303038152906040528051906020012060001c61059c9190610e1d565b90506000670de0b6b3a7640000826105b49190610cdb565b90508092505050919050565b600180546105cd90610dbd565b80601f01602080910402602001604051908101604052809291908181526020018280546105f990610dbd565b80156106465780601f1061061b57610100808354040283529160200191610646565b820191906000526020600020905b81548152906001019060200180831161062957829003601f168201915b505050505081565b3373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d4337884866040516106c4929190610bec565b60405180910390a45050565b7f40fc85fbff9305015298ba6fcee88b7e442a64cc803ddb889327680bbd62270a816040516106ff9190610b94565b60405180910390a150565b60606000600167ffffffffffffffff81111561074f577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561077d5781602001602082028036833780820191505090505b50905030816000815181106107bb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508091505090565b6000600354905090565b6000804233600054856040516020016108229493929190610b09565b6040516020818303038152906040528051906020012060001c90506000678ac7230489e80000826108539190610cdb565b90508092505050919050565b60008135905061086e81610ef9565b92915050565b60008135905061088381610f10565b92915050565b60008135905061089881610f27565b92915050565b6000602082840312156108b057600080fd5b60006108be8482850161085f565b91505092915050565b600080604083850312156108da57600080fd5b60006108e88582860161085f565b92505060206108f98582860161085f565b9150509250929050565b6000806040838503121561091657600080fd5b60006109248582860161085f565b925050602061093585828601610874565b9150509250929050565b60006020828403121561095157600080fd5b600061095f84828501610874565b91505092915050565b6000806040838503121561097b57600080fd5b600061098985828601610889565b925050602061099a85828601610889565b9150509250929050565b60006109b083836109bc565b60208301905092915050565b6109c581610d35565b82525050565b6109d481610d35565b82525050565b6109eb6109e682610d35565b610def565b82525050565b60006109fc82610c40565b610a068185610c63565b9350610a1183610c30565b8060005b83811015610a42578151610a2988826109a4565b9750610a3483610c56565b925050600181019050610a15565b5085935050505092915050565b610a5881610d47565b82525050565b6000610a6982610c4b565b610a738185610c74565b9350610a83818560208601610d8a565b610a8c81610edb565b840191505092915050565b610aa081610d73565b82525050565b610ab7610ab282610d73565b610e13565b82525050565b610ac681610d7d565b82525050565b6000610ad88286610aa6565b602082019150610ae882856109da565b601482019150610af88284610aa6565b602082019150819050949350505050565b6000610b158287610aa6565b602082019150610b2582866109da565b601482019150610b358285610aa6565b602082019150610b4582846109da565b60148201915081905095945050505050565b6000602082019050610b6c60008301846109cb565b92915050565b60006020820190508181036000830152610b8c81846109f1565b905092915050565b6000602082019050610ba96000830184610a4f565b92915050565b60006020820190508181036000830152610bc98184610a5e565b905092915050565b6000602082019050610be66000830184610a97565b92915050565b6000604082019050610c016000830185610a97565b610c0e6020830184610a97565b9392505050565b6000602082019050610c2a6000830184610abd565b92915050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610c9082610d73565b9150610c9b83610d73565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610cd057610ccf610e4e565b5b828201905092915050565b6000610ce682610d73565b9150610cf183610d73565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610d2a57610d29610e4e565b5b828202905092915050565b6000610d4082610d53565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015610da8578082015181840152602081019050610d8d565b83811115610db7576000848401525b50505050565b60006002820490506001821680610dd557607f821691505b60208210811415610de957610de8610eac565b5b50919050565b6000610dfa82610e01565b9050919050565b6000610e0c82610eec565b9050919050565b6000819050919050565b6000610e2882610d73565b9150610e3383610d73565b925082610e4357610e42610e7d565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b610f0281610d35565b8114610f0d57600080fd5b50565b610f1981610d47565b8114610f2457600080fd5b50565b610f3081610d73565b8114610f3b57600080fd5b5056fea2646970667358221220206ac0cc502ae40158c0d46a3cff85b029e7582527de90ef0b3595d94385663d64736f6c63430008010033"

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

// CircuitBreaker is a free data retrieval call binding the contract method 0x16efd941.
//
// Solidity: function circuitBreaker() view returns(address)
func (_Logswap *LogswapCaller) CircuitBreaker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Logswap.contract.Call(opts, &out, "circuitBreaker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CircuitBreaker is a free data retrieval call binding the contract method 0x16efd941.
//
// Solidity: function circuitBreaker() view returns(address)
func (_Logswap *LogswapSession) CircuitBreaker() (common.Address, error) {
	return _Logswap.Contract.CircuitBreaker(&_Logswap.CallOpts)
}

// CircuitBreaker is a free data retrieval call binding the contract method 0x16efd941.
//
// Solidity: function circuitBreaker() view returns(address)
func (_Logswap *LogswapCallerSession) CircuitBreaker() (common.Address, error) {
	return _Logswap.Contract.CircuitBreaker(&_Logswap.CallOpts)
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

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Logswap *LogswapCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Logswap.contract.Call(opts, &out, "getController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Logswap *LogswapSession) GetController() (common.Address, error) {
	return _Logswap.Contract.GetController(&_Logswap.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Logswap *LogswapCallerSession) GetController() (common.Address, error) {
	return _Logswap.Contract.GetController(&_Logswap.CallOpts)
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

// EmitPublicSwap is a paid mutator transaction binding the contract method 0xc94ddf8c.
//
// Solidity: function emitPublicSwap(bool enabled) returns()
func (_Logswap *LogswapTransactor) EmitPublicSwap(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Logswap.contract.Transact(opts, "emitPublicSwap", enabled)
}

// EmitPublicSwap is a paid mutator transaction binding the contract method 0xc94ddf8c.
//
// Solidity: function emitPublicSwap(bool enabled) returns()
func (_Logswap *LogswapSession) EmitPublicSwap(enabled bool) (*types.Transaction, error) {
	return _Logswap.Contract.EmitPublicSwap(&_Logswap.TransactOpts, enabled)
}

// EmitPublicSwap is a paid mutator transaction binding the contract method 0xc94ddf8c.
//
// Solidity: function emitPublicSwap(bool enabled) returns()
func (_Logswap *LogswapTransactorSession) EmitPublicSwap(enabled bool) (*types.Transaction, error) {
	return _Logswap.Contract.EmitPublicSwap(&_Logswap.TransactOpts, enabled)
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

// SetPublicSwap is a paid mutator transaction binding the contract method 0x1eccc185.
//
// Solidity: function setPublicSwap(address pool, bool enabled) returns()
func (_Logswap *LogswapTransactor) SetPublicSwap(opts *bind.TransactOpts, pool common.Address, enabled bool) (*types.Transaction, error) {
	return _Logswap.contract.Transact(opts, "setPublicSwap", pool, enabled)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x1eccc185.
//
// Solidity: function setPublicSwap(address pool, bool enabled) returns()
func (_Logswap *LogswapSession) SetPublicSwap(pool common.Address, enabled bool) (*types.Transaction, error) {
	return _Logswap.Contract.SetPublicSwap(&_Logswap.TransactOpts, pool, enabled)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x1eccc185.
//
// Solidity: function setPublicSwap(address pool, bool enabled) returns()
func (_Logswap *LogswapTransactorSession) SetPublicSwap(pool common.Address, enabled bool) (*types.Transaction, error) {
	return _Logswap.Contract.SetPublicSwap(&_Logswap.TransactOpts, pool, enabled)
}

// SetPublicSwap0 is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool enabled) returns()
func (_Logswap *LogswapTransactor) SetPublicSwap0(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Logswap.contract.Transact(opts, "setPublicSwap0", enabled)
}

// SetPublicSwap0 is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool enabled) returns()
func (_Logswap *LogswapSession) SetPublicSwap0(enabled bool) (*types.Transaction, error) {
	return _Logswap.Contract.SetPublicSwap0(&_Logswap.TransactOpts, enabled)
}

// SetPublicSwap0 is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool enabled) returns()
func (_Logswap *LogswapTransactorSession) SetPublicSwap0(enabled bool) (*types.Transaction, error) {
	return _Logswap.Contract.SetPublicSwap0(&_Logswap.TransactOpts, enabled)
}

// LogswapLOGPUBLICSWAPTOGGLEDIterator is returned from FilterLOGPUBLICSWAPTOGGLED and is used to iterate over the raw logs and unpacked data for LOGPUBLICSWAPTOGGLED events raised by the Logswap contract.
type LogswapLOGPUBLICSWAPTOGGLEDIterator struct {
	Event *LogswapLOGPUBLICSWAPTOGGLED // Event containing the contract specifics and raw log

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
func (it *LogswapLOGPUBLICSWAPTOGGLEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogswapLOGPUBLICSWAPTOGGLED)
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
		it.Event = new(LogswapLOGPUBLICSWAPTOGGLED)
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
func (it *LogswapLOGPUBLICSWAPTOGGLEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogswapLOGPUBLICSWAPTOGGLEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogswapLOGPUBLICSWAPTOGGLED represents a LOGPUBLICSWAPTOGGLED event raised by the Logswap contract.
type LogswapLOGPUBLICSWAPTOGGLED struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLOGPUBLICSWAPTOGGLED is a free log retrieval operation binding the contract event 0x40fc85fbff9305015298ba6fcee88b7e442a64cc803ddb889327680bbd62270a.
//
// Solidity: event LOG_PUBLIC_SWAP_TOGGLED(bool enabled)
func (_Logswap *LogswapFilterer) FilterLOGPUBLICSWAPTOGGLED(opts *bind.FilterOpts) (*LogswapLOGPUBLICSWAPTOGGLEDIterator, error) {

	logs, sub, err := _Logswap.contract.FilterLogs(opts, "LOG_PUBLIC_SWAP_TOGGLED")
	if err != nil {
		return nil, err
	}
	return &LogswapLOGPUBLICSWAPTOGGLEDIterator{contract: _Logswap.contract, event: "LOG_PUBLIC_SWAP_TOGGLED", logs: logs, sub: sub}, nil
}

// WatchLOGPUBLICSWAPTOGGLED is a free log subscription operation binding the contract event 0x40fc85fbff9305015298ba6fcee88b7e442a64cc803ddb889327680bbd62270a.
//
// Solidity: event LOG_PUBLIC_SWAP_TOGGLED(bool enabled)
func (_Logswap *LogswapFilterer) WatchLOGPUBLICSWAPTOGGLED(opts *bind.WatchOpts, sink chan<- *LogswapLOGPUBLICSWAPTOGGLED) (event.Subscription, error) {

	logs, sub, err := _Logswap.contract.WatchLogs(opts, "LOG_PUBLIC_SWAP_TOGGLED")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogswapLOGPUBLICSWAPTOGGLED)
				if err := _Logswap.contract.UnpackLog(event, "LOG_PUBLIC_SWAP_TOGGLED", log); err != nil {
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
func (_Logswap *LogswapFilterer) ParseLOGPUBLICSWAPTOGGLED(log types.Log) (*LogswapLOGPUBLICSWAPTOGGLED, error) {
	event := new(LogswapLOGPUBLICSWAPTOGGLED)
	if err := _Logswap.contract.UnpackLog(event, "LOG_PUBLIC_SWAP_TOGGLED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

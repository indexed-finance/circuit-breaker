// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multicall

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

// SimpleMultiCallBundle is an auto generated low-level Go binding around an user-defined struct.
type SimpleMultiCallBundle struct {
	Pool                common.Address
	Tokens              []common.Address
	DenormalizedWeights []*big.Int
	Balances            []*big.Int
	TotalSupplies       []*big.Int
}

// MulticallABI is the input ABI used to generate the binding from.
const MulticallABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getBundle\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"denormalizedWeights\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"totalSupplies\",\"type\":\"uint256[]\"}],\"internalType\":\"structSimpleMultiCall.Bundle\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"address[][]\",\"name\":\"tokens\",\"type\":\"address[][]\"}],\"name\":\"getBundles\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"denormalizedWeights\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"totalSupplies\",\"type\":\"uint256[]\"}],\"internalType\":\"structSimpleMultiCall.Bundle[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getDenormalizedWeights\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"inTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"outTokens\",\"type\":\"address[]\"}],\"name\":\"getSpotPrices\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getTotalSupplies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getUsedBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MulticallBin is the compiled bytecode used for deploying new contracts.
var MulticallBin = "0x608060405234801561001057600080fd5b50610d23806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638841576a1161005b5780638841576a146100df57806395d7710e146100ff578063db16d8ac14610112578063f7a6b6b6146101255761007d565b80632dd43d89146100825780636a385ae9146100ac578063823f5d86146100bf575b600080fd5b610095610090366004610966565b610147565b6040516100a3929190610bc9565b60405180910390f35b6100956100ba3660046108a9565b610248565b6100d26100cd3660046109a1565b610356565b6040516100a39190610bee565b6100f26100ed3660046108a9565b61043c565b6040516100a39190610c91565b61009561010d3660046108a9565b6104a4565b6100956101203660046108a9565b6105a8565b6101386101333660046108f5565b6106ac565b6040516100a393929190610b86565b6060806060835167ffffffffffffffff8111801561016457600080fd5b5060405190808252806020026020018201604052801561018e578160200160208202803683370190505b50905060005b8451811015610240578481815181106101a957fe5b60200260200101516001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156101e957600080fd5b505afa1580156101fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102219190610a51565b82828151811061022d57fe5b6020908102919091010152600101610194565b509293915050565b6060806060835167ffffffffffffffff8111801561026557600080fd5b5060405190808252806020026020018201604052801561028f578160200160208202803683370190505b50905060005b845181101561034c57856001600160a01b031663f8b2cb4f8683815181106102b957fe5b60200260200101516040518263ffffffff1660e01b81526004016102dd9190610b58565b60206040518083038186803b1580156102f557600080fd5b505afa158015610309573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061032d9190610a51565b82828151811061033957fe5b6020908102919091010152600101610295565b5092949293505050565b606081518351146103825760405162461bcd60e51b815260040161037990610c4e565b60405180910390fd5b6060835167ffffffffffffffff8111801561039c57600080fd5b506040519080825280602002602001820160405280156103d657816020015b6103c36107e0565b8152602001906001900390816103bb5790505b50905060005b8351811015610434576104158582815181106103f457fe5b602002602001015185838151811061040857fe5b602002602001015161043c565b82828151811061042157fe5b60209081029190910101526001016103dc565b509392505050565b6104446107e0565b606061045084846105a8565b915050606061045f8585610248565b915050606061046d85610147565b6040805160a0810182526001600160a01b039990991689526020890197909752958701939093525060608501525050608082015290565b6060806060835167ffffffffffffffff811180156104c157600080fd5b506040519080825280602002602001820160405280156104eb578160200160208202803683370190505b50905060005b845181101561034c57856001600160a01b0316634aa4e0b586838151811061051557fe5b60200260200101516040518263ffffffff1660e01b81526004016105399190610b58565b60206040518083038186803b15801561055157600080fd5b505afa158015610565573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105899190610a51565b82828151811061059557fe5b60209081029190910101526001016104f1565b6060806060835167ffffffffffffffff811180156105c557600080fd5b506040519080825280602002602001820160405280156105ef578160200160208202803683370190505b50905060005b845181101561034c57856001600160a01b031663948d8ce686838151811061061957fe5b60200260200101516040518263ffffffff1660e01b815260040161063d9190610b58565b60206040518083038186803b15801561065557600080fd5b505afa158015610669573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061068d9190610a51565b82828151811061069957fe5b60209081029190910101526001016105f5565b606080606083518551146106bf57600080fd5b6060855167ffffffffffffffff811180156106d957600080fd5b50604051908082528060200260200182016040528015610703578160200160208202803683370190505b50905060005b86518110156107d557876001600160a01b03166315e84af988838151811061072d57fe5b602002602001015188848151811061074157fe5b60200260200101516040518363ffffffff1660e01b8152600401610766929190610b6c565b60206040518083038186803b15801561077e57600080fd5b505afa158015610792573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107b69190610a51565b8282815181106107c257fe5b6020908102919091010152600101610709565b509496939550505050565b6040518060a0016040528060006001600160a01b03168152602001606081526020016060815260200160608152602001606081525090565b80356001600160a01b038116811461082f57600080fd5b919050565b600082601f830112610844578081fd5b813561085761085282610ccf565b610cab565b81815291506020808301908481018184028601820187101561087857600080fd5b60005b8481101561089e5761088c82610818565b8452928201929082019060010161087b565b505050505092915050565b600080604083850312156108bb578182fd5b6108c483610818565b9150602083013567ffffffffffffffff8111156108df578182fd5b6108eb85828601610834565b9150509250929050565b600080600060608486031215610909578081fd5b61091284610818565b9250602084013567ffffffffffffffff8082111561092e578283fd5b61093a87838801610834565b9350604086013591508082111561094f578283fd5b5061095c86828701610834565b9150509250925092565b600060208284031215610977578081fd5b813567ffffffffffffffff81111561098d578182fd5b61099984828501610834565b949350505050565b600080604083850312156109b3578182fd5b823567ffffffffffffffff808211156109ca578384fd5b6109d686838701610834565b93506020915081850135818111156109ec578384fd5b85019050601f810186136109fe578283fd5b8035610a0c61085282610ccf565b81815283810190838501865b84811015610a4157610a2f8b888435890101610834565b84529286019290860190600101610a18565b5096999098509650505050505050565b600060208284031215610a62578081fd5b5051919050565b6000815180845260208085019450808401835b83811015610aa15781516001600160a01b031687529582019590820190600101610a7c565b509495945050505050565b6000815180845260208085019450808401835b83811015610aa157815187529582019590820190600101610abf565b600060018060a01b038251168352602082015160a06020850152610b0260a0850182610a69565b905060408301518482036040860152610b1b8282610aac565b91505060608301518482036060860152610b358282610aac565b91505060808301518482036080860152610b4f8282610aac565b95945050505050565b6001600160a01b0391909116815260200190565b6001600160a01b0392831681529116602082015260400190565b600060608252610b996060830186610a69565b8281036020840152610bab8186610a69565b90508281036040840152610bbf8185610aac565b9695505050505050565b600060408252610bdc6040830185610a69565b8281036020840152610b4f8185610aac565b6000602080830181845280855180835260408601915060408482028701019250838701855b82811015610c4157603f19888603018452610c2f858351610adb565b94509285019290850190600101610c13565b5092979650505050505050565b60208082526023908201527f6d69736d61746368696e6720706f6f6c7320616e6420746f6b656e73206c656e6040820152620cee8d60eb1b606082015260800190565b600060208252610ca46020830184610adb565b9392505050565b60405181810167ffffffffffffffff81118282101715610cc757fe5b604052919050565b600067ffffffffffffffff821115610ce357fe5b506020908102019056fea2646970667358221220b3c953d0d2d2bcd671362bf4ccd06657f492e46b3c9698df9d6920fbb929c66a64736f6c63430007040033"

// DeployMulticall deploys a new Ethereum contract, binding an instance of Multicall to it.
func DeployMulticall(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Multicall, error) {
	parsed, err := abi.JSON(strings.NewReader(MulticallABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MulticallBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Multicall{MulticallCaller: MulticallCaller{contract: contract}, MulticallTransactor: MulticallTransactor{contract: contract}, MulticallFilterer: MulticallFilterer{contract: contract}}, nil
}

// Multicall is an auto generated Go binding around an Ethereum contract.
type Multicall struct {
	MulticallCaller     // Read-only binding to the contract
	MulticallTransactor // Write-only binding to the contract
	MulticallFilterer   // Log filterer for contract events
}

// MulticallCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallSession struct {
	Contract     *Multicall        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulticallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallCallerSession struct {
	Contract *MulticallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MulticallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallTransactorSession struct {
	Contract     *MulticallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MulticallRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallRaw struct {
	Contract *Multicall // Generic contract binding to access the raw methods on
}

// MulticallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallCallerRaw struct {
	Contract *MulticallCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallTransactorRaw struct {
	Contract *MulticallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticall creates a new instance of Multicall, bound to a specific deployed contract.
func NewMulticall(address common.Address, backend bind.ContractBackend) (*Multicall, error) {
	contract, err := bindMulticall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multicall{MulticallCaller: MulticallCaller{contract: contract}, MulticallTransactor: MulticallTransactor{contract: contract}, MulticallFilterer: MulticallFilterer{contract: contract}}, nil
}

// NewMulticallCaller creates a new read-only instance of Multicall, bound to a specific deployed contract.
func NewMulticallCaller(address common.Address, caller bind.ContractCaller) (*MulticallCaller, error) {
	contract, err := bindMulticall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallCaller{contract: contract}, nil
}

// NewMulticallTransactor creates a new write-only instance of Multicall, bound to a specific deployed contract.
func NewMulticallTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallTransactor, error) {
	contract, err := bindMulticall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallTransactor{contract: contract}, nil
}

// NewMulticallFilterer creates a new log filterer instance of Multicall, bound to a specific deployed contract.
func NewMulticallFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallFilterer, error) {
	contract, err := bindMulticall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallFilterer{contract: contract}, nil
}

// bindMulticall binds a generic wrapper to an already deployed contract.
func bindMulticall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MulticallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall *MulticallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall.Contract.MulticallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall *MulticallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall.Contract.MulticallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall *MulticallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall.Contract.MulticallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall *MulticallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall *MulticallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall *MulticallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall.Contract.contract.Transact(opts, method, params...)
}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCaller) GetBalances(opts *bind.CallOpts, poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "getBalances", poolAddress, tokens)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallSession) GetBalances(poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetBalances(&_Multicall.CallOpts, poolAddress, tokens)
}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCallerSession) GetBalances(poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetBalances(&_Multicall.CallOpts, poolAddress, tokens)
}

// GetBundle is a free data retrieval call binding the contract method 0x8841576a.
//
// Solidity: function getBundle(address poolAddress, address[] tokens) view returns((address,address[],uint256[],uint256[],uint256[]))
func (_Multicall *MulticallCaller) GetBundle(opts *bind.CallOpts, poolAddress common.Address, tokens []common.Address) (SimpleMultiCallBundle, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "getBundle", poolAddress, tokens)

	if err != nil {
		return *new(SimpleMultiCallBundle), err
	}

	out0 := *abi.ConvertType(out[0], new(SimpleMultiCallBundle)).(*SimpleMultiCallBundle)

	return out0, err

}

// GetBundle is a free data retrieval call binding the contract method 0x8841576a.
//
// Solidity: function getBundle(address poolAddress, address[] tokens) view returns((address,address[],uint256[],uint256[],uint256[]))
func (_Multicall *MulticallSession) GetBundle(poolAddress common.Address, tokens []common.Address) (SimpleMultiCallBundle, error) {
	return _Multicall.Contract.GetBundle(&_Multicall.CallOpts, poolAddress, tokens)
}

// GetBundle is a free data retrieval call binding the contract method 0x8841576a.
//
// Solidity: function getBundle(address poolAddress, address[] tokens) view returns((address,address[],uint256[],uint256[],uint256[]))
func (_Multicall *MulticallCallerSession) GetBundle(poolAddress common.Address, tokens []common.Address) (SimpleMultiCallBundle, error) {
	return _Multicall.Contract.GetBundle(&_Multicall.CallOpts, poolAddress, tokens)
}

// GetBundles is a free data retrieval call binding the contract method 0x823f5d86.
//
// Solidity: function getBundles(address[] pools, address[][] tokens) view returns((address,address[],uint256[],uint256[],uint256[])[])
func (_Multicall *MulticallCaller) GetBundles(opts *bind.CallOpts, pools []common.Address, tokens [][]common.Address) ([]SimpleMultiCallBundle, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "getBundles", pools, tokens)

	if err != nil {
		return *new([]SimpleMultiCallBundle), err
	}

	out0 := *abi.ConvertType(out[0], new([]SimpleMultiCallBundle)).(*[]SimpleMultiCallBundle)

	return out0, err

}

// GetBundles is a free data retrieval call binding the contract method 0x823f5d86.
//
// Solidity: function getBundles(address[] pools, address[][] tokens) view returns((address,address[],uint256[],uint256[],uint256[])[])
func (_Multicall *MulticallSession) GetBundles(pools []common.Address, tokens [][]common.Address) ([]SimpleMultiCallBundle, error) {
	return _Multicall.Contract.GetBundles(&_Multicall.CallOpts, pools, tokens)
}

// GetBundles is a free data retrieval call binding the contract method 0x823f5d86.
//
// Solidity: function getBundles(address[] pools, address[][] tokens) view returns((address,address[],uint256[],uint256[],uint256[])[])
func (_Multicall *MulticallCallerSession) GetBundles(pools []common.Address, tokens [][]common.Address) ([]SimpleMultiCallBundle, error) {
	return _Multicall.Contract.GetBundles(&_Multicall.CallOpts, pools, tokens)
}

// GetDenormalizedWeights is a free data retrieval call binding the contract method 0xdb16d8ac.
//
// Solidity: function getDenormalizedWeights(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCaller) GetDenormalizedWeights(opts *bind.CallOpts, poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "getDenormalizedWeights", poolAddress, tokens)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDenormalizedWeights is a free data retrieval call binding the contract method 0xdb16d8ac.
//
// Solidity: function getDenormalizedWeights(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallSession) GetDenormalizedWeights(poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetDenormalizedWeights(&_Multicall.CallOpts, poolAddress, tokens)
}

// GetDenormalizedWeights is a free data retrieval call binding the contract method 0xdb16d8ac.
//
// Solidity: function getDenormalizedWeights(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCallerSession) GetDenormalizedWeights(poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetDenormalizedWeights(&_Multicall.CallOpts, poolAddress, tokens)
}

// GetSpotPrices is a free data retrieval call binding the contract method 0xf7a6b6b6.
//
// Solidity: function getSpotPrices(address poolAddress, address[] inTokens, address[] outTokens) view returns(address[], address[], uint256[])
func (_Multicall *MulticallCaller) GetSpotPrices(opts *bind.CallOpts, poolAddress common.Address, inTokens []common.Address, outTokens []common.Address) ([]common.Address, []common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "getSpotPrices", poolAddress, inTokens, outTokens)

	if err != nil {
		return *new([]common.Address), *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, out2, err

}

// GetSpotPrices is a free data retrieval call binding the contract method 0xf7a6b6b6.
//
// Solidity: function getSpotPrices(address poolAddress, address[] inTokens, address[] outTokens) view returns(address[], address[], uint256[])
func (_Multicall *MulticallSession) GetSpotPrices(poolAddress common.Address, inTokens []common.Address, outTokens []common.Address) ([]common.Address, []common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetSpotPrices(&_Multicall.CallOpts, poolAddress, inTokens, outTokens)
}

// GetSpotPrices is a free data retrieval call binding the contract method 0xf7a6b6b6.
//
// Solidity: function getSpotPrices(address poolAddress, address[] inTokens, address[] outTokens) view returns(address[], address[], uint256[])
func (_Multicall *MulticallCallerSession) GetSpotPrices(poolAddress common.Address, inTokens []common.Address, outTokens []common.Address) ([]common.Address, []common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetSpotPrices(&_Multicall.CallOpts, poolAddress, inTokens, outTokens)
}

// GetTotalSupplies is a free data retrieval call binding the contract method 0x2dd43d89.
//
// Solidity: function getTotalSupplies(address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCaller) GetTotalSupplies(opts *bind.CallOpts, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "getTotalSupplies", tokens)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetTotalSupplies is a free data retrieval call binding the contract method 0x2dd43d89.
//
// Solidity: function getTotalSupplies(address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallSession) GetTotalSupplies(tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetTotalSupplies(&_Multicall.CallOpts, tokens)
}

// GetTotalSupplies is a free data retrieval call binding the contract method 0x2dd43d89.
//
// Solidity: function getTotalSupplies(address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCallerSession) GetTotalSupplies(tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetTotalSupplies(&_Multicall.CallOpts, tokens)
}

// GetUsedBalances is a free data retrieval call binding the contract method 0x95d7710e.
//
// Solidity: function getUsedBalances(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCaller) GetUsedBalances(opts *bind.CallOpts, poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "getUsedBalances", poolAddress, tokens)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetUsedBalances is a free data retrieval call binding the contract method 0x95d7710e.
//
// Solidity: function getUsedBalances(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallSession) GetUsedBalances(poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetUsedBalances(&_Multicall.CallOpts, poolAddress, tokens)
}

// GetUsedBalances is a free data retrieval call binding the contract method 0x95d7710e.
//
// Solidity: function getUsedBalances(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCallerSession) GetUsedBalances(poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetUsedBalances(&_Multicall.CallOpts, poolAddress, tokens)
}

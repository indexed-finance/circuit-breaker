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
var MulticallBin = "0x608060405234801561001057600080fd5b50610d2d806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638841576a1161005b5780638841576a146100df57806395d7710e146100ff578063db16d8ac14610112578063f7a6b6b6146101255761007d565b80632dd43d89146100825780636a385ae9146100ac578063823f5d86146100bf575b600080fd5b610095610090366004610970565b610147565b6040516100a3929190610bd3565b60405180910390f35b6100956100ba3660046108b3565b610248565b6100d26100cd3660046109ab565b610356565b6040516100a39190610bf8565b6100f26100ed3660046108b3565b61043c565b6040516100a39190610c9b565b61009561010d3660046108b3565b6104ae565b6100956101203660046108b3565b6105b2565b6101386101333660046108ff565b6106b6565b6040516100a393929190610b90565b6060806060835167ffffffffffffffff8111801561016457600080fd5b5060405190808252806020026020018201604052801561018e578160200160208202803683370190505b50905060005b8451811015610240578481815181106101a957fe5b60200260200101516001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156101e957600080fd5b505afa1580156101fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102219190610a5b565b82828151811061022d57fe5b6020908102919091010152600101610194565b509293915050565b6060806060835167ffffffffffffffff8111801561026557600080fd5b5060405190808252806020026020018201604052801561028f578160200160208202803683370190505b50905060005b845181101561034c57856001600160a01b031663f8b2cb4f8683815181106102b957fe5b60200260200101516040518263ffffffff1660e01b81526004016102dd9190610b62565b60206040518083038186803b1580156102f557600080fd5b505afa158015610309573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061032d9190610a5b565b82828151811061033957fe5b6020908102919091010152600101610295565b5092949293505050565b606081518351146103825760405162461bcd60e51b815260040161037990610c58565b60405180910390fd5b6060835167ffffffffffffffff8111801561039c57600080fd5b506040519080825280602002602001820160405280156103d657816020015b6103c36107ea565b8152602001906001900390816103bb5790505b50905060005b8351811015610434576104158582815181106103f457fe5b602002602001015185838151811061040857fe5b602002602001015161043c565b82828151811061042157fe5b60209081029190910101526001016103dc565b509392505050565b6104446107ea565b606061045084846105b2565b915050606061045f8585610248565b915050606061046d85610147565b9150506104786107ea565b506040805160a0810182526001600160a01b03979097168752602087019590955293850191909152606084015250608082015290565b6060806060835167ffffffffffffffff811180156104cb57600080fd5b506040519080825280602002602001820160405280156104f5578160200160208202803683370190505b50905060005b845181101561034c57856001600160a01b0316634aa4e0b586838151811061051f57fe5b60200260200101516040518263ffffffff1660e01b81526004016105439190610b62565b60206040518083038186803b15801561055b57600080fd5b505afa15801561056f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105939190610a5b565b82828151811061059f57fe5b60209081029190910101526001016104fb565b6060806060835167ffffffffffffffff811180156105cf57600080fd5b506040519080825280602002602001820160405280156105f9578160200160208202803683370190505b50905060005b845181101561034c57856001600160a01b031663948d8ce686838151811061062357fe5b60200260200101516040518263ffffffff1660e01b81526004016106479190610b62565b60206040518083038186803b15801561065f57600080fd5b505afa158015610673573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106979190610a5b565b8282815181106106a357fe5b60209081029190910101526001016105ff565b606080606083518551146106c957600080fd5b6060855167ffffffffffffffff811180156106e357600080fd5b5060405190808252806020026020018201604052801561070d578160200160208202803683370190505b50905060005b86518110156107df57876001600160a01b03166315e84af988838151811061073757fe5b602002602001015188848151811061074b57fe5b60200260200101516040518363ffffffff1660e01b8152600401610770929190610b76565b60206040518083038186803b15801561078857600080fd5b505afa15801561079c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c09190610a5b565b8282815181106107cc57fe5b6020908102919091010152600101610713565b509496939550505050565b6040518060a0016040528060006001600160a01b03168152602001606081526020016060815260200160608152602001606081525090565b80356001600160a01b038116811461083957600080fd5b919050565b600082601f83011261084e578081fd5b813561086161085c82610cd9565b610cb5565b81815291506020808301908481018184028601820187101561088257600080fd5b60005b848110156108a85761089682610822565b84529282019290820190600101610885565b505050505092915050565b600080604083850312156108c5578182fd5b6108ce83610822565b9150602083013567ffffffffffffffff8111156108e9578182fd5b6108f58582860161083e565b9150509250929050565b600080600060608486031215610913578081fd5b61091c84610822565b9250602084013567ffffffffffffffff80821115610938578283fd5b6109448783880161083e565b93506040860135915080821115610959578283fd5b506109668682870161083e565b9150509250925092565b600060208284031215610981578081fd5b813567ffffffffffffffff811115610997578182fd5b6109a38482850161083e565b949350505050565b600080604083850312156109bd578182fd5b823567ffffffffffffffff808211156109d4578384fd5b6109e08683870161083e565b93506020915081850135818111156109f6578384fd5b85019050601f81018613610a08578283fd5b8035610a1661085c82610cd9565b81815283810190838501865b84811015610a4b57610a398b88843589010161083e565b84529286019290860190600101610a22565b5096999098509650505050505050565b600060208284031215610a6c578081fd5b5051919050565b6000815180845260208085019450808401835b83811015610aab5781516001600160a01b031687529582019590820190600101610a86565b509495945050505050565b6000815180845260208085019450808401835b83811015610aab57815187529582019590820190600101610ac9565b600060018060a01b038251168352602082015160a06020850152610b0c60a0850182610a73565b905060408301518482036040860152610b258282610ab6565b91505060608301518482036060860152610b3f8282610ab6565b91505060808301518482036080860152610b598282610ab6565b95945050505050565b6001600160a01b0391909116815260200190565b6001600160a01b0392831681529116602082015260400190565b600060608252610ba36060830186610a73565b8281036020840152610bb58186610a73565b90508281036040840152610bc98185610ab6565b9695505050505050565b600060408252610be66040830185610a73565b8281036020840152610b598185610ab6565b6000602080830181845280855180835260408601915060408482028701019250838701855b82811015610c4b57603f19888603018452610c39858351610ae5565b94509285019290850190600101610c1d565b5092979650505050505050565b60208082526023908201527f6d69736d61746368696e6720706f6f6c7320616e6420746f6b656e73206c656e6040820152620cee8d60eb1b606082015260800190565b600060208252610cae6020830184610ae5565b9392505050565b60405181810167ffffffffffffffff81118282101715610cd157fe5b604052919050565b600067ffffffffffffffff821115610ced57fe5b506020908102019056fea2646970667358221220b94c0b5be88351167153a5fec65d4f48450866e1b1c333675b7a6e7d840ce15d64736f6c63430007040033"

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

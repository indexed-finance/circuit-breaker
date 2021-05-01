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
var MulticallBin = "0x608060405234801561001057600080fd5b50610c8a806100206000396000f3fe608060405234801561001057600080fd5b506004361061006d5760003560e01c80632dd43d89146100725780636a385ae91461009c578063823f5d86146100af5780638841576a146100cf57806395d7710e146100ef578063db16d8ac14610102578063f7a6b6b614610115575b600080fd5b610085610080366004610914565b610137565b604051610093929190610b75565b60405180910390f35b6100856100aa366004610859565b610237565b6100c26100bd36600461094e565b610344565b6040516100939190610b9a565b6100e26100dd366004610859565b6103fe565b6040516100939190610bfa565b6100856100fd366004610859565b610466565b610085610110366004610859565b610569565b6101286101233660046108a4565b61066c565b60405161009393929190610b32565b606080606083516001600160401b038111801561015357600080fd5b5060405190808252806020026020018201604052801561017d578160200160208202803683370190505b50905060005b845181101561022f5784818151811061019857fe5b60200260200101516001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156101d857600080fd5b505afa1580156101ec573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061021091906109fd565b82828151811061021c57fe5b6020908102919091010152600101610183565b509293915050565b606080606083516001600160401b038111801561025357600080fd5b5060405190808252806020026020018201604052801561027d578160200160208202803683370190505b50905060005b845181101561033a57856001600160a01b031663f8b2cb4f8683815181106102a757fe5b60200260200101516040518263ffffffff1660e01b81526004016102cb9190610b04565b60206040518083038186803b1580156102e357600080fd5b505afa1580156102f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061031b91906109fd565b82828151811061032757fe5b6020908102919091010152600101610283565b5092949293505050565b60608083516001600160401b038111801561035e57600080fd5b5060405190808252806020026020018201604052801561039857816020015b610385610790565b81526020019060019003908161037d5790505b50905060005b83518110156103f6576103d78582815181106103b657fe5b60200260200101518583815181106103ca57fe5b60200260200101516103fe565b8282815181106103e357fe5b602090810291909101015260010161039e565b509392505050565b610406610790565b60606104128484610569565b91505060606104218585610237565b915050606061042f85610137565b6040805160a0810182526001600160a01b039990991689526020890197909752958701939093525060608501525050608082015290565b606080606083516001600160401b038111801561048257600080fd5b506040519080825280602002602001820160405280156104ac578160200160208202803683370190505b50905060005b845181101561033a57856001600160a01b0316634aa4e0b58683815181106104d657fe5b60200260200101516040518263ffffffff1660e01b81526004016104fa9190610b04565b60206040518083038186803b15801561051257600080fd5b505afa158015610526573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061054a91906109fd565b82828151811061055657fe5b60209081029190910101526001016104b2565b606080606083516001600160401b038111801561058557600080fd5b506040519080825280602002602001820160405280156105af578160200160208202803683370190505b50905060005b845181101561033a57856001600160a01b031663948d8ce68683815181106105d957fe5b60200260200101516040518263ffffffff1660e01b81526004016105fd9190610b04565b60206040518083038186803b15801561061557600080fd5b505afa158015610629573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064d91906109fd565b82828151811061065957fe5b60209081029190910101526001016105b5565b60608060608085516001600160401b038111801561068957600080fd5b506040519080825280602002602001820160405280156106b3578160200160208202803683370190505b50905060005b865181101561078557876001600160a01b03166315e84af98883815181106106dd57fe5b60200260200101518884815181106106f157fe5b60200260200101516040518363ffffffff1660e01b8152600401610716929190610b18565b60206040518083038186803b15801561072e57600080fd5b505afa158015610742573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076691906109fd565b82828151811061077257fe5b60209081029190910101526001016106b9565b509496939550505050565b6040518060a0016040528060006001600160a01b03168152602001606081526020016060815260200160608152602001606081525090565b80356001600160a01b03811681146107df57600080fd5b919050565b600082601f8301126107f4578081fd5b813561080761080282610c37565b610c14565b81815291506020808301908481018184028601820187101561082857600080fd5b60005b8481101561084e5761083c826107c8565b8452928201929082019060010161082b565b505050505092915050565b6000806040838503121561086b578182fd5b610874836107c8565b915060208301356001600160401b0381111561088e578182fd5b61089a858286016107e4565b9150509250929050565b6000806000606084860312156108b8578081fd5b6108c1846107c8565b925060208401356001600160401b03808211156108dc578283fd5b6108e8878388016107e4565b935060408601359150808211156108fd578283fd5b5061090a868287016107e4565b9150509250925092565b600060208284031215610925578081fd5b81356001600160401b0381111561093a578182fd5b610946848285016107e4565b949350505050565b60008060408385031215610960578182fd5b82356001600160401b0380821115610976578384fd5b610982868387016107e4565b9350602091508185013581811115610998578384fd5b85019050601f810186136109aa578283fd5b80356109b861080282610c37565b81815283810190838501865b848110156109ed576109db8b8884358901016107e4565b845292860192908601906001016109c4565b5096999098509650505050505050565b600060208284031215610a0e578081fd5b5051919050565b6000815180845260208085019450808401835b83811015610a4d5781516001600160a01b031687529582019590820190600101610a28565b509495945050505050565b6000815180845260208085019450808401835b83811015610a4d57815187529582019590820190600101610a6b565b600060018060a01b038251168352602082015160a06020850152610aae60a0850182610a15565b905060408301518482036040860152610ac78282610a58565b91505060608301518482036060860152610ae18282610a58565b91505060808301518482036080860152610afb8282610a58565b95945050505050565b6001600160a01b0391909116815260200190565b6001600160a01b0392831681529116602082015260400190565b600060608252610b456060830186610a15565b8281036020840152610b578186610a15565b90508281036040840152610b6b8185610a58565b9695505050505050565b600060408252610b886040830185610a15565b8281036020840152610afb8185610a58565b6000602080830181845280855180835260408601915060408482028701019250838701855b82811015610bed57603f19888603018452610bdb858351610a87565b94509285019290850190600101610bbf565b5092979650505050505050565b600060208252610c0d6020830184610a87565b9392505050565b6040518181016001600160401b0381118282101715610c2f57fe5b604052919050565b60006001600160401b03821115610c4a57fe5b506020908102019056fea264697066735822122079c876ad27cf77dbfaa2a0f979c22a5a8ac8b7582be1bd0d0b9587095c0ecfb564736f6c63430007040033"

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

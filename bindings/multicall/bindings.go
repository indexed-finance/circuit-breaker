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
var MulticallBin = "0x608060405234801561001057600080fd5b50610ca8806100206000396000f3fe608060405234801561001057600080fd5b506004361061006d5760003560e01c80632dd43d89146100725780636a385ae91461009c578063823f5d86146100af5780638841576a146100cf57806395d7710e146100ef578063db16d8ac14610102578063f7a6b6b614610115575b600080fd5b610085610080366004610932565b610137565b604051610093929190610b93565b60405180910390f35b6100856100aa366004610877565b610237565b6100c26100bd36600461096c565b610344565b6040516100939190610bb8565b6100e26100dd366004610877565b61040d565b6040516100939190610c18565b6100856100fd366004610877565b610475565b610085610110366004610877565b610578565b6101286101233660046108c2565b61067b565b60405161009393929190610b50565b606080606083516001600160401b038111801561015357600080fd5b5060405190808252806020026020018201604052801561017d578160200160208202803683370190505b50905060005b845181101561022f5784818151811061019857fe5b60200260200101516001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156101d857600080fd5b505afa1580156101ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102109190610a1b565b82828151811061021c57fe5b6020908102919091010152600101610183565b509293915050565b606080606083516001600160401b038111801561025357600080fd5b5060405190808252806020026020018201604052801561027d578160200160208202803683370190505b50905060005b845181101561033a57856001600160a01b031663f8b2cb4f8683815181106102a757fe5b60200260200101516040518263ffffffff1660e01b81526004016102cb9190610b22565b60206040518083038186803b1580156102e357600080fd5b505afa1580156102f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061031b9190610a1b565b82828151811061032757fe5b6020908102919091010152600101610283565b5092949293505050565b6060815183511461035457600080fd5b606083516001600160401b038111801561036d57600080fd5b506040519080825280602002602001820160405280156103a757816020015b6103946107ae565b81526020019060019003908161038c5790505b50905060005b8351811015610405576103e68582815181106103c557fe5b60200260200101518583815181106103d957fe5b602002602001015161040d565b8282815181106103f257fe5b60209081029190910101526001016103ad565b509392505050565b6104156107ae565b60606104218484610578565b91505060606104308585610237565b915050606061043e85610137565b6040805160a0810182526001600160a01b039990991689526020890197909752958701939093525060608501525050608082015290565b606080606083516001600160401b038111801561049157600080fd5b506040519080825280602002602001820160405280156104bb578160200160208202803683370190505b50905060005b845181101561033a57856001600160a01b0316634aa4e0b58683815181106104e557fe5b60200260200101516040518263ffffffff1660e01b81526004016105099190610b22565b60206040518083038186803b15801561052157600080fd5b505afa158015610535573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105599190610a1b565b82828151811061056557fe5b60209081029190910101526001016104c1565b606080606083516001600160401b038111801561059457600080fd5b506040519080825280602002602001820160405280156105be578160200160208202803683370190505b50905060005b845181101561033a57856001600160a01b031663948d8ce68683815181106105e857fe5b60200260200101516040518263ffffffff1660e01b815260040161060c9190610b22565b60206040518083038186803b15801561062457600080fd5b505afa158015610638573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061065c9190610a1b565b82828151811061066857fe5b60209081029190910101526001016105c4565b6060806060835185511461068e57600080fd5b606085516001600160401b03811180156106a757600080fd5b506040519080825280602002602001820160405280156106d1578160200160208202803683370190505b50905060005b86518110156107a357876001600160a01b03166315e84af98883815181106106fb57fe5b602002602001015188848151811061070f57fe5b60200260200101516040518363ffffffff1660e01b8152600401610734929190610b36565b60206040518083038186803b15801561074c57600080fd5b505afa158015610760573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107849190610a1b565b82828151811061079057fe5b60209081029190910101526001016106d7565b509496939550505050565b6040518060a0016040528060006001600160a01b03168152602001606081526020016060815260200160608152602001606081525090565b80356001600160a01b03811681146107fd57600080fd5b919050565b600082601f830112610812578081fd5b813561082561082082610c55565b610c32565b81815291506020808301908481018184028601820187101561084657600080fd5b60005b8481101561086c5761085a826107e6565b84529282019290820190600101610849565b505050505092915050565b60008060408385031215610889578182fd5b610892836107e6565b915060208301356001600160401b038111156108ac578182fd5b6108b885828601610802565b9150509250929050565b6000806000606084860312156108d6578081fd5b6108df846107e6565b925060208401356001600160401b03808211156108fa578283fd5b61090687838801610802565b9350604086013591508082111561091b578283fd5b5061092886828701610802565b9150509250925092565b600060208284031215610943578081fd5b81356001600160401b03811115610958578182fd5b61096484828501610802565b949350505050565b6000806040838503121561097e578182fd5b82356001600160401b0380821115610994578384fd5b6109a086838701610802565b93506020915081850135818111156109b6578384fd5b85019050601f810186136109c8578283fd5b80356109d661082082610c55565b81815283810190838501865b84811015610a0b576109f98b888435890101610802565b845292860192908601906001016109e2565b5096999098509650505050505050565b600060208284031215610a2c578081fd5b5051919050565b6000815180845260208085019450808401835b83811015610a6b5781516001600160a01b031687529582019590820190600101610a46565b509495945050505050565b6000815180845260208085019450808401835b83811015610a6b57815187529582019590820190600101610a89565b600060018060a01b038251168352602082015160a06020850152610acc60a0850182610a33565b905060408301518482036040860152610ae58282610a76565b91505060608301518482036060860152610aff8282610a76565b91505060808301518482036080860152610b198282610a76565b95945050505050565b6001600160a01b0391909116815260200190565b6001600160a01b0392831681529116602082015260400190565b600060608252610b636060830186610a33565b8281036020840152610b758186610a33565b90508281036040840152610b898185610a76565b9695505050505050565b600060408252610ba66040830185610a33565b8281036020840152610b198185610a76565b6000602080830181845280855180835260408601915060408482028701019250838701855b82811015610c0b57603f19888603018452610bf9858351610aa5565b94509285019290850190600101610bdd565b5092979650505050505050565b600060208252610c2b6020830184610aa5565b9392505050565b6040518181016001600160401b0381118282101715610c4d57fe5b604052919050565b60006001600160401b03821115610c6857fe5b506020908102019056fea2646970667358221220d1e5fc7b4f6b99479b3d0a66d697fe884a1afb84039b8f9c14f4ab731ce480b164736f6c63430007040033"

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

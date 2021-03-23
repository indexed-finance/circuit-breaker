// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bmat

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

// BmatABI is the input ABI used to generate the binding from.
const BmatABI = "[{\"inputs\":[],\"name\":\"BONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BPOW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_POOL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_IN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_OUT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WEIGHT_CHANGE_PCT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WEIGHT_UPDATE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"badd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"bdiv\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"bfloor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"bmul\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"}],\"name\":\"bpow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"precision\",\"type\":\"uint256\"}],\"name\":\"bpowApprox\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"bpowi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"bsub\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"bsubSign\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"btoi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BmatBin is the compiled bytecode used for deploying new contracts.
var BmatBin = "0x608060405234801561001057600080fd5b506110f0806100206000396000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c80638c051bf31161010f578063bc694ea2116100a2578063d011f32a11610071578063d011f32a146106f3578063e4a28a5214610749578063ec09302114610767578063ed61bdc014610785576101e5565b8063bc694ea21461064d578063c36596a61461066b578063c6580d1214610689578063cd660299146106a7576101e5565b8063b0e0d136116100de578063b0e0d136146105d5578063b7b800a4146105f3578063ba019dab14610611578063bc063e1a1461062f576101e5565b80638c051bf3146104e35780639381cd2b1461052f578063992e2a921461054d578063a221ee491461056b576101e5565b80636852b5cd116101875780637e6d71a2116101565780637e6d71a2146104475780638025e30314610465578063859ae48e14610483578063867378c5146104c5576101e5565b80636852b5cd1461036a578063743480cc146103885780637673eb11146103dd57806376c7a3c714610429576101e5565b80631d0a82df116101c35780631d0a82df14610272578063218b5382146102be578063317223bc146102dc5780633cf3c7d41461031e576101e5565b806309a3bbe4146101ea5780630b71b95c14610208578063189d00ca14610254575b600080fd5b6101f26107a3565b6040518082815260200191505060405180910390f35b61023e6004803603604081101561021e57600080fd5b8101908080359060200190929190803590602001909291905050506107b2565b6040518082815260200191505060405180910390f35b61025c61083a565b6040518082815260200191505060405180910390f35b6102a86004803603604081101561028857600080fd5b810190808035906020019092919080359060200190929190505050610854565b6040518082815260200191505060405180910390f35b6102c66109d2565b6040518082815260200191505060405180910390f35b610308600480360360208110156102f257600080fd5b81019080803590602001909291905050506109e8565b6040518082815260200191505060405180910390f35b6103546004803603604081101561033457600080fd5b810190808035906020019092919080359060200190929190505050610a04565b6040518082815260200191505060405180910390f35b610372610a96565b6040518082815260200191505060405180910390f35b6103be6004803603604081101561039e57600080fd5b810190808035906020019092919080359060200190929190505050610aac565b6040518083815260200182151581526020019250505060405180910390f35b610413600480360360408110156103f357600080fd5b810190808035906020019092919080359060200190929190505050610ad5565b6040518082815260200191505060405180910390f35b610431610c8e565b6040518082815260200191505060405180910390f35b61044f610ca6565b6040518082815260200191505060405180910390f35b61046d610cac565b6040518082815260200191505060405180910390f35b6104af6004803603602081101561049957600080fd5b8101908080359060200190929190505050610cb1565b6040518082815260200191505060405180910390f35b6104cd610ccc565b6040518082815260200191505060405180910390f35b610519600480360360408110156104f957600080fd5b810190808035906020019092919080359060200190929190505050610ce6565b6040518082815260200191505060405180910390f35b610537610e28565b6040518082815260200191505060405180910390f35b610555610e37565b6040518082815260200191505060405180910390f35b6105bf600480360360a081101561058157600080fd5b810190808035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610e50565b6040518082815260200191505060405180910390f35b6105dd610ebb565b6040518082815260200191505060405180910390f35b6105fb610ec0565b6040518082815260200191505060405180910390f35b610619610ec5565b6040518082815260200191505060405180910390f35b610637610eca565b6040518082815260200191505060405180910390f35b610655610ee0565b6040518082815260200191505060405180910390f35b610673610ef2565b6040518082815260200191505060405180910390f35b610691610efe565b6040518082815260200191505060405180910390f35b6106dd600480360360408110156106bd57600080fd5b810190808035906020019092919080359060200190929190505050610f03565b6040518082815260200191505060405180910390f35b6107336004803603606081101561070957600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050610f88565b6040518082815260200191505060405180910390f35b610751611086565b6040518082815260200191505060405180910390f35b61076f611095565b6040518082815260200191505060405180910390f35b61078d6110ab565b6040518082815260200191505060405180910390f35b601a670de0b6b3a76400000281565b600080828401905083811015610830576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f4552525f4144445f4f564552464c4f570000000000000000000000000000000081525060200191505060405180910390fd5b8091505092915050565b6402540be400670de0b6b3a76400008161085057fe5b0481565b600060018310156108cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4552525f42504f575f424153455f544f4f5f4c4f57000000000000000000000081525060200191505060405180910390fd5b6001670de0b6b3a764000060020203831115610951576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f4552525f42504f575f424153455f544f4f5f484947480000000000000000000081525060200191505060405180910390fd5b600061095c836109e8565b9050600061096a8483610a04565b905060006109808661097b85610cb1565b610f03565b90506000821415610996578093505050506109cc565b60006109b987846402540be400670de0b6b3a7640000816109b357fe5b04610f88565b90506109c58282610ce6565b9450505050505b92915050565b6004670de0b6b3a7640000816109e457fe5b0481565b6000670de0b6b3a76400006109fc83610cb1565b029050919050565b6000806000610a138585610aac565b915091508015610a8b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f4552525f5355425f554e444552464c4f5700000000000000000000000000000081525060200191505060405180910390fd5b819250505092915050565b6064670de0b6b3a764000081610aa857fe5b0481565b600080828410610ac457828403600091509150610ace565b8383036001915091505b9250929050565b600080821415610b4d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f4552525f4449565f5a45524f000000000000000000000000000000000000000081525060200191505060405180910390fd5b6000670de0b6b3a7640000840290506000841480610b7b5750670de0b6b3a7640000848281610b7857fe5b04145b610bed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f4552525f4449565f494e5445524e414c0000000000000000000000000000000081525060200191505060405180910390fd5b600060028481610bf957fe5b048201905081811015610c74576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f4552525f4449565f494e5445524e414c0000000000000000000000000000000081525060200191505060405180910390fd5b6000848281610c7f57fe5b04905080935050505092915050565b620f4240670de0b6b3a764000081610ca257fe5b0481565b610e1081565b600181565b6000670de0b6b3a76400008281610cc457fe5b049050919050565b64e8d4a51000670de0b6b3a764000081610ce257fe5b0481565b60008082840290506000841480610d05575082848281610d0257fe5b04145b610d77576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f4552525f4d554c5f4f564552464c4f570000000000000000000000000000000081525060200191505060405180910390fd5b60006002670de0b6b3a764000081610d8b57fe5b048201905081811015610e06576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f4552525f4d554c5f4f564552464c4f570000000000000000000000000000000081525060200191505060405180910390fd5b6000670de0b6b3a76400008281610e1957fe5b04905080935050505092915050565b6064670de0b6b3a76400000281565b60016003670de0b6b3a764000081610e4b57fe5b040181565b600080610e5d8787610ad5565b90506000610e6b8686610ad5565b90506000610e798383610ad5565b90506000610ea0670de0b6b3a7640000610e9b670de0b6b3a764000089610a04565b610ad5565b9050610eac8282610ce6565b94505050505095945050505050565b600a81565b600281565b600181565b600a670de0b6b3a764000081610edc57fe5b0481565b6001670de0b6b3a76400006002020381565b670de0b6b3a764000081565b600081565b600080600060028481610f1257fe5b061415610f2757670de0b6b3a7640000610f29565b835b905060028381610f3557fe5b0492505b60008314610f7e57610f4b8485610ce6565b9350600060028481610f5957fe5b0614610f6c57610f698185610ce6565b90505b60028381610f7657fe5b049250610f39565b8091505092915050565b600080839050600080610fa387670de0b6b3a7640000610aac565b915091506000670de0b6b3a764000090506000819050600080600190505b888410611075576000670de0b6b3a764000082029050600080610ff58a610ff085670de0b6b3a7640000610a04565b610aac565b9150915061100c87611007848c610ce6565b610ce6565b96506110188784610ad5565b9650600087141561102b57505050611075565b871561103657841594505b801561104157841594505b8415611058576110518688610a04565b9550611065565b61106286886107b2565b95505b5050508080600101915050610fc1565b508196505050505050509392505050565b6019670de0b6b3a76400000281565b6002670de0b6b3a7640000816110a757fe5b0481565b6019670de0b6b3a7640000028156fea2646970667358221220657cb47a3c842f88c4797931ef431874e46b3aa084b9408b875d67c7c4ed436964736f6c63430007040033"

// DeployBmat deploys a new Ethereum contract, binding an instance of Bmat to it.
func DeployBmat(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bmat, error) {
	parsed, err := abi.JSON(strings.NewReader(BmatABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BmatBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bmat{BmatCaller: BmatCaller{contract: contract}, BmatTransactor: BmatTransactor{contract: contract}, BmatFilterer: BmatFilterer{contract: contract}}, nil
}

// Bmat is an auto generated Go binding around an Ethereum contract.
type Bmat struct {
	BmatCaller     // Read-only binding to the contract
	BmatTransactor // Write-only binding to the contract
	BmatFilterer   // Log filterer for contract events
}

// BmatCaller is an auto generated read-only Go binding around an Ethereum contract.
type BmatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BmatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BmatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BmatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BmatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BmatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BmatSession struct {
	Contract     *Bmat             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BmatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BmatCallerSession struct {
	Contract *BmatCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BmatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BmatTransactorSession struct {
	Contract     *BmatTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BmatRaw is an auto generated low-level Go binding around an Ethereum contract.
type BmatRaw struct {
	Contract *Bmat // Generic contract binding to access the raw methods on
}

// BmatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BmatCallerRaw struct {
	Contract *BmatCaller // Generic read-only contract binding to access the raw methods on
}

// BmatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BmatTransactorRaw struct {
	Contract *BmatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBmat creates a new instance of Bmat, bound to a specific deployed contract.
func NewBmat(address common.Address, backend bind.ContractBackend) (*Bmat, error) {
	contract, err := bindBmat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bmat{BmatCaller: BmatCaller{contract: contract}, BmatTransactor: BmatTransactor{contract: contract}, BmatFilterer: BmatFilterer{contract: contract}}, nil
}

// NewBmatCaller creates a new read-only instance of Bmat, bound to a specific deployed contract.
func NewBmatCaller(address common.Address, caller bind.ContractCaller) (*BmatCaller, error) {
	contract, err := bindBmat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BmatCaller{contract: contract}, nil
}

// NewBmatTransactor creates a new write-only instance of Bmat, bound to a specific deployed contract.
func NewBmatTransactor(address common.Address, transactor bind.ContractTransactor) (*BmatTransactor, error) {
	contract, err := bindBmat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BmatTransactor{contract: contract}, nil
}

// NewBmatFilterer creates a new log filterer instance of Bmat, bound to a specific deployed contract.
func NewBmatFilterer(address common.Address, filterer bind.ContractFilterer) (*BmatFilterer, error) {
	contract, err := bindBmat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BmatFilterer{contract: contract}, nil
}

// bindBmat binds a generic wrapper to an already deployed contract.
func bindBmat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BmatABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bmat *BmatRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bmat.Contract.BmatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bmat *BmatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bmat.Contract.BmatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bmat *BmatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bmat.Contract.BmatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bmat *BmatCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bmat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bmat *BmatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bmat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bmat *BmatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bmat.Contract.contract.Transact(opts, method, params...)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_Bmat *BmatCaller) BONE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "BONE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_Bmat *BmatSession) BONE() (*big.Int, error) {
	return _Bmat.Contract.BONE(&_Bmat.CallOpts)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_Bmat *BmatCallerSession) BONE() (*big.Int, error) {
	return _Bmat.Contract.BONE(&_Bmat.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_Bmat *BmatCaller) BPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "BPOW_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_Bmat *BmatSession) BPOWPRECISION() (*big.Int, error) {
	return _Bmat.Contract.BPOWPRECISION(&_Bmat.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_Bmat *BmatCallerSession) BPOWPRECISION() (*big.Int, error) {
	return _Bmat.Contract.BPOWPRECISION(&_Bmat.CallOpts)
}

// DEFAULTTOTALWEIGHT is a free data retrieval call binding the contract method 0xed61bdc0.
//
// Solidity: function DEFAULT_TOTAL_WEIGHT() view returns(uint256)
func (_Bmat *BmatCaller) DEFAULTTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "DEFAULT_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTTOTALWEIGHT is a free data retrieval call binding the contract method 0xed61bdc0.
//
// Solidity: function DEFAULT_TOTAL_WEIGHT() view returns(uint256)
func (_Bmat *BmatSession) DEFAULTTOTALWEIGHT() (*big.Int, error) {
	return _Bmat.Contract.DEFAULTTOTALWEIGHT(&_Bmat.CallOpts)
}

// DEFAULTTOTALWEIGHT is a free data retrieval call binding the contract method 0xed61bdc0.
//
// Solidity: function DEFAULT_TOTAL_WEIGHT() view returns(uint256)
func (_Bmat *BmatCallerSession) DEFAULTTOTALWEIGHT() (*big.Int, error) {
	return _Bmat.Contract.DEFAULTTOTALWEIGHT(&_Bmat.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_Bmat *BmatCaller) EXITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "EXIT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_Bmat *BmatSession) EXITFEE() (*big.Int, error) {
	return _Bmat.Contract.EXITFEE(&_Bmat.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_Bmat *BmatCallerSession) EXITFEE() (*big.Int, error) {
	return _Bmat.Contract.EXITFEE(&_Bmat.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_Bmat *BmatCaller) INITPOOLSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "INIT_POOL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_Bmat *BmatSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _Bmat.Contract.INITPOOLSUPPLY(&_Bmat.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_Bmat *BmatCallerSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _Bmat.Contract.INITPOOLSUPPLY(&_Bmat.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_Bmat *BmatCaller) MAXBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "MAX_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_Bmat *BmatSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _Bmat.Contract.MAXBOUNDTOKENS(&_Bmat.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_Bmat *BmatCallerSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _Bmat.Contract.MAXBOUNDTOKENS(&_Bmat.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_Bmat *BmatCaller) MAXBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "MAX_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_Bmat *BmatSession) MAXBPOWBASE() (*big.Int, error) {
	return _Bmat.Contract.MAXBPOWBASE(&_Bmat.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_Bmat *BmatCallerSession) MAXBPOWBASE() (*big.Int, error) {
	return _Bmat.Contract.MAXBPOWBASE(&_Bmat.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Bmat *BmatCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Bmat *BmatSession) MAXFEE() (*big.Int, error) {
	return _Bmat.Contract.MAXFEE(&_Bmat.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Bmat *BmatCallerSession) MAXFEE() (*big.Int, error) {
	return _Bmat.Contract.MAXFEE(&_Bmat.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_Bmat *BmatCaller) MAXINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "MAX_IN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_Bmat *BmatSession) MAXINRATIO() (*big.Int, error) {
	return _Bmat.Contract.MAXINRATIO(&_Bmat.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_Bmat *BmatCallerSession) MAXINRATIO() (*big.Int, error) {
	return _Bmat.Contract.MAXINRATIO(&_Bmat.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_Bmat *BmatCaller) MAXOUTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "MAX_OUT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_Bmat *BmatSession) MAXOUTRATIO() (*big.Int, error) {
	return _Bmat.Contract.MAXOUTRATIO(&_Bmat.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_Bmat *BmatCallerSession) MAXOUTRATIO() (*big.Int, error) {
	return _Bmat.Contract.MAXOUTRATIO(&_Bmat.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_Bmat *BmatCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "MAX_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_Bmat *BmatSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _Bmat.Contract.MAXTOTALWEIGHT(&_Bmat.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_Bmat *BmatCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _Bmat.Contract.MAXTOTALWEIGHT(&_Bmat.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_Bmat *BmatCaller) MAXWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "MAX_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_Bmat *BmatSession) MAXWEIGHT() (*big.Int, error) {
	return _Bmat.Contract.MAXWEIGHT(&_Bmat.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_Bmat *BmatCallerSession) MAXWEIGHT() (*big.Int, error) {
	return _Bmat.Contract.MAXWEIGHT(&_Bmat.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_Bmat *BmatCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "MIN_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_Bmat *BmatSession) MINBALANCE() (*big.Int, error) {
	return _Bmat.Contract.MINBALANCE(&_Bmat.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_Bmat *BmatCallerSession) MINBALANCE() (*big.Int, error) {
	return _Bmat.Contract.MINBALANCE(&_Bmat.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_Bmat *BmatCaller) MINBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "MIN_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_Bmat *BmatSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _Bmat.Contract.MINBOUNDTOKENS(&_Bmat.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_Bmat *BmatCallerSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _Bmat.Contract.MINBOUNDTOKENS(&_Bmat.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_Bmat *BmatCaller) MINBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "MIN_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_Bmat *BmatSession) MINBPOWBASE() (*big.Int, error) {
	return _Bmat.Contract.MINBPOWBASE(&_Bmat.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_Bmat *BmatCallerSession) MINBPOWBASE() (*big.Int, error) {
	return _Bmat.Contract.MINBPOWBASE(&_Bmat.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_Bmat *BmatCaller) MINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "MIN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_Bmat *BmatSession) MINFEE() (*big.Int, error) {
	return _Bmat.Contract.MINFEE(&_Bmat.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_Bmat *BmatCallerSession) MINFEE() (*big.Int, error) {
	return _Bmat.Contract.MINFEE(&_Bmat.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_Bmat *BmatCaller) MINWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "MIN_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_Bmat *BmatSession) MINWEIGHT() (*big.Int, error) {
	return _Bmat.Contract.MINWEIGHT(&_Bmat.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_Bmat *BmatCallerSession) MINWEIGHT() (*big.Int, error) {
	return _Bmat.Contract.MINWEIGHT(&_Bmat.CallOpts)
}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_Bmat *BmatCaller) VERSIONNUMBER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "VERSION_NUMBER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_Bmat *BmatSession) VERSIONNUMBER() (*big.Int, error) {
	return _Bmat.Contract.VERSIONNUMBER(&_Bmat.CallOpts)
}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_Bmat *BmatCallerSession) VERSIONNUMBER() (*big.Int, error) {
	return _Bmat.Contract.VERSIONNUMBER(&_Bmat.CallOpts)
}

// WEIGHTCHANGEPCT is a free data retrieval call binding the contract method 0x6852b5cd.
//
// Solidity: function WEIGHT_CHANGE_PCT() view returns(uint256)
func (_Bmat *BmatCaller) WEIGHTCHANGEPCT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "WEIGHT_CHANGE_PCT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WEIGHTCHANGEPCT is a free data retrieval call binding the contract method 0x6852b5cd.
//
// Solidity: function WEIGHT_CHANGE_PCT() view returns(uint256)
func (_Bmat *BmatSession) WEIGHTCHANGEPCT() (*big.Int, error) {
	return _Bmat.Contract.WEIGHTCHANGEPCT(&_Bmat.CallOpts)
}

// WEIGHTCHANGEPCT is a free data retrieval call binding the contract method 0x6852b5cd.
//
// Solidity: function WEIGHT_CHANGE_PCT() view returns(uint256)
func (_Bmat *BmatCallerSession) WEIGHTCHANGEPCT() (*big.Int, error) {
	return _Bmat.Contract.WEIGHTCHANGEPCT(&_Bmat.CallOpts)
}

// WEIGHTUPDATEDELAY is a free data retrieval call binding the contract method 0x7e6d71a2.
//
// Solidity: function WEIGHT_UPDATE_DELAY() view returns(uint256)
func (_Bmat *BmatCaller) WEIGHTUPDATEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "WEIGHT_UPDATE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WEIGHTUPDATEDELAY is a free data retrieval call binding the contract method 0x7e6d71a2.
//
// Solidity: function WEIGHT_UPDATE_DELAY() view returns(uint256)
func (_Bmat *BmatSession) WEIGHTUPDATEDELAY() (*big.Int, error) {
	return _Bmat.Contract.WEIGHTUPDATEDELAY(&_Bmat.CallOpts)
}

// WEIGHTUPDATEDELAY is a free data retrieval call binding the contract method 0x7e6d71a2.
//
// Solidity: function WEIGHT_UPDATE_DELAY() view returns(uint256)
func (_Bmat *BmatCallerSession) WEIGHTUPDATEDELAY() (*big.Int, error) {
	return _Bmat.Contract.WEIGHTUPDATEDELAY(&_Bmat.CallOpts)
}

// Badd is a free data retrieval call binding the contract method 0x0b71b95c.
//
// Solidity: function badd(uint256 a, uint256 b) pure returns(uint256)
func (_Bmat *BmatCaller) Badd(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "badd", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Badd is a free data retrieval call binding the contract method 0x0b71b95c.
//
// Solidity: function badd(uint256 a, uint256 b) pure returns(uint256)
func (_Bmat *BmatSession) Badd(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Badd(&_Bmat.CallOpts, a, b)
}

// Badd is a free data retrieval call binding the contract method 0x0b71b95c.
//
// Solidity: function badd(uint256 a, uint256 b) pure returns(uint256)
func (_Bmat *BmatCallerSession) Badd(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Badd(&_Bmat.CallOpts, a, b)
}

// Bdiv is a free data retrieval call binding the contract method 0x7673eb11.
//
// Solidity: function bdiv(uint256 a, uint256 b) pure returns(uint256)
func (_Bmat *BmatCaller) Bdiv(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "bdiv", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bdiv is a free data retrieval call binding the contract method 0x7673eb11.
//
// Solidity: function bdiv(uint256 a, uint256 b) pure returns(uint256)
func (_Bmat *BmatSession) Bdiv(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Bdiv(&_Bmat.CallOpts, a, b)
}

// Bdiv is a free data retrieval call binding the contract method 0x7673eb11.
//
// Solidity: function bdiv(uint256 a, uint256 b) pure returns(uint256)
func (_Bmat *BmatCallerSession) Bdiv(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Bdiv(&_Bmat.CallOpts, a, b)
}

// Bfloor is a free data retrieval call binding the contract method 0x317223bc.
//
// Solidity: function bfloor(uint256 a) pure returns(uint256)
func (_Bmat *BmatCaller) Bfloor(opts *bind.CallOpts, a *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "bfloor", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bfloor is a free data retrieval call binding the contract method 0x317223bc.
//
// Solidity: function bfloor(uint256 a) pure returns(uint256)
func (_Bmat *BmatSession) Bfloor(a *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Bfloor(&_Bmat.CallOpts, a)
}

// Bfloor is a free data retrieval call binding the contract method 0x317223bc.
//
// Solidity: function bfloor(uint256 a) pure returns(uint256)
func (_Bmat *BmatCallerSession) Bfloor(a *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Bfloor(&_Bmat.CallOpts, a)
}

// Bmul is a free data retrieval call binding the contract method 0x8c051bf3.
//
// Solidity: function bmul(uint256 a, uint256 b) pure returns(uint256)
func (_Bmat *BmatCaller) Bmul(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "bmul", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bmul is a free data retrieval call binding the contract method 0x8c051bf3.
//
// Solidity: function bmul(uint256 a, uint256 b) pure returns(uint256)
func (_Bmat *BmatSession) Bmul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Bmul(&_Bmat.CallOpts, a, b)
}

// Bmul is a free data retrieval call binding the contract method 0x8c051bf3.
//
// Solidity: function bmul(uint256 a, uint256 b) pure returns(uint256)
func (_Bmat *BmatCallerSession) Bmul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Bmul(&_Bmat.CallOpts, a, b)
}

// Bpow is a free data retrieval call binding the contract method 0x1d0a82df.
//
// Solidity: function bpow(uint256 base, uint256 exp) pure returns(uint256)
func (_Bmat *BmatCaller) Bpow(opts *bind.CallOpts, base *big.Int, exp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "bpow", base, exp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bpow is a free data retrieval call binding the contract method 0x1d0a82df.
//
// Solidity: function bpow(uint256 base, uint256 exp) pure returns(uint256)
func (_Bmat *BmatSession) Bpow(base *big.Int, exp *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Bpow(&_Bmat.CallOpts, base, exp)
}

// Bpow is a free data retrieval call binding the contract method 0x1d0a82df.
//
// Solidity: function bpow(uint256 base, uint256 exp) pure returns(uint256)
func (_Bmat *BmatCallerSession) Bpow(base *big.Int, exp *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Bpow(&_Bmat.CallOpts, base, exp)
}

// BpowApprox is a free data retrieval call binding the contract method 0xd011f32a.
//
// Solidity: function bpowApprox(uint256 base, uint256 exp, uint256 precision) pure returns(uint256)
func (_Bmat *BmatCaller) BpowApprox(opts *bind.CallOpts, base *big.Int, exp *big.Int, precision *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "bpowApprox", base, exp, precision)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BpowApprox is a free data retrieval call binding the contract method 0xd011f32a.
//
// Solidity: function bpowApprox(uint256 base, uint256 exp, uint256 precision) pure returns(uint256)
func (_Bmat *BmatSession) BpowApprox(base *big.Int, exp *big.Int, precision *big.Int) (*big.Int, error) {
	return _Bmat.Contract.BpowApprox(&_Bmat.CallOpts, base, exp, precision)
}

// BpowApprox is a free data retrieval call binding the contract method 0xd011f32a.
//
// Solidity: function bpowApprox(uint256 base, uint256 exp, uint256 precision) pure returns(uint256)
func (_Bmat *BmatCallerSession) BpowApprox(base *big.Int, exp *big.Int, precision *big.Int) (*big.Int, error) {
	return _Bmat.Contract.BpowApprox(&_Bmat.CallOpts, base, exp, precision)
}

// Bpowi is a free data retrieval call binding the contract method 0xcd660299.
//
// Solidity: function bpowi(uint256 a, uint256 n) pure returns(uint256)
func (_Bmat *BmatCaller) Bpowi(opts *bind.CallOpts, a *big.Int, n *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "bpowi", a, n)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bpowi is a free data retrieval call binding the contract method 0xcd660299.
//
// Solidity: function bpowi(uint256 a, uint256 n) pure returns(uint256)
func (_Bmat *BmatSession) Bpowi(a *big.Int, n *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Bpowi(&_Bmat.CallOpts, a, n)
}

// Bpowi is a free data retrieval call binding the contract method 0xcd660299.
//
// Solidity: function bpowi(uint256 a, uint256 n) pure returns(uint256)
func (_Bmat *BmatCallerSession) Bpowi(a *big.Int, n *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Bpowi(&_Bmat.CallOpts, a, n)
}

// Bsub is a free data retrieval call binding the contract method 0x3cf3c7d4.
//
// Solidity: function bsub(uint256 a, uint256 b) pure returns(uint256)
func (_Bmat *BmatCaller) Bsub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "bsub", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bsub is a free data retrieval call binding the contract method 0x3cf3c7d4.
//
// Solidity: function bsub(uint256 a, uint256 b) pure returns(uint256)
func (_Bmat *BmatSession) Bsub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Bsub(&_Bmat.CallOpts, a, b)
}

// Bsub is a free data retrieval call binding the contract method 0x3cf3c7d4.
//
// Solidity: function bsub(uint256 a, uint256 b) pure returns(uint256)
func (_Bmat *BmatCallerSession) Bsub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Bsub(&_Bmat.CallOpts, a, b)
}

// BsubSign is a free data retrieval call binding the contract method 0x743480cc.
//
// Solidity: function bsubSign(uint256 a, uint256 b) pure returns(uint256, bool)
func (_Bmat *BmatCaller) BsubSign(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "bsubSign", a, b)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// BsubSign is a free data retrieval call binding the contract method 0x743480cc.
//
// Solidity: function bsubSign(uint256 a, uint256 b) pure returns(uint256, bool)
func (_Bmat *BmatSession) BsubSign(a *big.Int, b *big.Int) (*big.Int, bool, error) {
	return _Bmat.Contract.BsubSign(&_Bmat.CallOpts, a, b)
}

// BsubSign is a free data retrieval call binding the contract method 0x743480cc.
//
// Solidity: function bsubSign(uint256 a, uint256 b) pure returns(uint256, bool)
func (_Bmat *BmatCallerSession) BsubSign(a *big.Int, b *big.Int) (*big.Int, bool, error) {
	return _Bmat.Contract.BsubSign(&_Bmat.CallOpts, a, b)
}

// Btoi is a free data retrieval call binding the contract method 0x859ae48e.
//
// Solidity: function btoi(uint256 a) pure returns(uint256)
func (_Bmat *BmatCaller) Btoi(opts *bind.CallOpts, a *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "btoi", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Btoi is a free data retrieval call binding the contract method 0x859ae48e.
//
// Solidity: function btoi(uint256 a) pure returns(uint256)
func (_Bmat *BmatSession) Btoi(a *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Btoi(&_Bmat.CallOpts, a)
}

// Btoi is a free data retrieval call binding the contract method 0x859ae48e.
//
// Solidity: function btoi(uint256 a) pure returns(uint256)
func (_Bmat *BmatCallerSession) Btoi(a *big.Int) (*big.Int, error) {
	return _Bmat.Contract.Btoi(&_Bmat.CallOpts, a)
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_Bmat *BmatCaller) CalcSpotPrice(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bmat.contract.Call(opts, &out, "calcSpotPrice", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_Bmat *BmatSession) CalcSpotPrice(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Bmat.Contract.CalcSpotPrice(&_Bmat.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_Bmat *BmatCallerSession) CalcSpotPrice(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Bmat.Contract.CalcSpotPrice(&_Bmat.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
}

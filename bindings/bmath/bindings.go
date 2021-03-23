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
var BmatBin = "0x608060405234801561001057600080fd5b50611566806100206000396000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c80638c051bf31161010f578063bc694ea2116100a2578063d011f32a11610071578063d011f32a146105c9578063e4a28a52146105f9578063ec09302114610617578063ed61bdc014610635576101e5565b8063bc694ea21461053f578063c36596a61461055d578063c6580d121461057b578063cd66029914610599576101e5565b8063b0e0d136116100de578063b0e0d136146104c7578063b7b800a4146104e5578063ba019dab14610503578063bc063e1a14610521576101e5565b80638c051bf31461042b5780639381cd2b1461045b578063992e2a9214610479578063a221ee4914610497576101e5565b80636852b5cd116101875780637e6d71a2116101565780637e6d71a2146103a15780638025e303146103bf578063859ae48e146103dd578063867378c51461040d576101e5565b80636852b5cd14610304578063743480cc146103225780637673eb111461035357806376c7a3c714610383576101e5565b80631d0a82df116101c35780631d0a82df14610256578063218b538214610286578063317223bc146102a45780633cf3c7d4146102d4576101e5565b806309a3bbe4146101ea5780630b71b95c14610208578063189d00ca14610238575b600080fd5b6101f2610653565b6040516101ff91906111a2565b60405180910390f35b610222600480360381019061021d9190610ead565b61066b565b60405161022f91906111a2565b60405180910390f35b6102406106c9565b60405161024d91906111a2565b60405180910390f35b610270600480360381019061026b9190610ead565b6106e5565b60405161027d91906111a2565b60405180910390f35b61028e610811565b60405161029b91906111a2565b60405180910390f35b6102be60048036038101906102b99190610e84565b610829565b6040516102cb91906111a2565b60405180910390f35b6102ee60048036038101906102e99190610ead565b61084e565b6040516102fb91906111a2565b60405180910390f35b61030c6108ad565b60405161031991906111a2565b60405180910390f35b61033c60048036038101906103379190610ead565b6108c5565b60405161034a9291906111bd565b60405180910390f35b61036d60048036038101906103689190610ead565b610900565b60405161037a91906111a2565b60405180910390f35b61038b610a38565b60405161039891906111a2565b60405180910390f35b6103a9610a52565b6040516103b691906111a2565b60405180910390f35b6103c7610a58565b6040516103d491906111a2565b60405180910390f35b6103f760048036038101906103f29190610e84565b610a5d565b60405161040491906111a2565b60405180910390f35b610415610a7a565b60405161042291906111a2565b60405180910390f35b61044560048036038101906104409190610ead565b610a96565b60405161045291906111a2565b60405180910390f35b610463610b8a565b60405161047091906111a2565b60405180910390f35b610481610ba2565b60405161048e91906111a2565b60405180910390f35b6104b160048036038101906104ac9190610f38565b610bc6565b6040516104be91906111a2565b60405180910390f35b6104cf610c31565b6040516104dc91906111a2565b60405180910390f35b6104ed610c36565b6040516104fa91906111a2565b60405180910390f35b61050b610c3b565b60405161051891906111a2565b60405180910390f35b610529610c40565b60405161053691906111a2565b60405180910390f35b610547610c58565b60405161055491906111a2565b60405180910390f35b610565610c7c565b60405161057291906111a2565b60405180910390f35b610583610c88565b60405161059091906111a2565b60405180910390f35b6105b360048036038101906105ae9190610ead565b610c8d565b6040516105c091906111a2565b60405180910390f35b6105e360048036038101906105de9190610ee9565b610d1a565b6040516105f091906111a2565b60405180910390f35b610601610e27565b60405161060e91906111a2565b60405180910390f35b61061f610e3f565b60405161062c91906111a2565b60405180910390f35b61063d610e57565b60405161064a91906111a2565b60405180910390f35b601a670de0b6b3a7640000610668919061127e565b81565b600080828461067a91906111f7565b9050838110156106bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b690611162565b60405180910390fd5b8091505092915050565b6402540be400670de0b6b3a76400006106e2919061124d565b81565b6000600183101561072b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072290611122565b60405180910390fd5b6001670de0b6b3a76400006002610742919061127e565b61074c91906112d8565b83111561078e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078590611182565b60405180910390fd5b600061079983610829565b905060006107a7848361084e565b905060006107bd866107b885610a5d565b610c8d565b905060008214156107d35780935050505061080b565b60006107f887846402540be400670de0b6b3a76400006107f3919061124d565b610d1a565b90506108048282610a96565b9450505050505b92915050565b6004670de0b6b3a7640000610826919061124d565b81565b6000670de0b6b3a764000061083d83610a5d565b610847919061127e565b9050919050565b600080600061085d85856108c5565b9150915080156108a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089990611142565b60405180910390fd5b819250505092915050565b6064670de0b6b3a76400006108c2919061124d565b81565b6000808284106108e65782846108db91906112d8565b6000915091506108f9565b83836108f291906112d8565b6001915091505b9250929050565b600080821415610945576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093c90611102565b60405180910390fd5b6000670de0b6b3a76400008461095b919061127e565b9050600084148061097e5750670de0b6b3a7640000848261097c919061124d565b145b6109bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b4906110c2565b60405180910390fd5b60006002846109cc919061124d565b826109d791906111f7565b905081811015610a1c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a13906110c2565b60405180910390fd5b60008482610a2a919061124d565b905080935050505092915050565b620f4240670de0b6b3a7640000610a4f919061124d565b81565b610e1081565b600181565b6000670de0b6b3a764000082610a73919061124d565b9050919050565b64e8d4a51000670de0b6b3a7640000610a93919061124d565b81565b6000808284610aa5919061127e565b90506000841480610ac05750828482610abe919061124d565b145b610aff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af6906110e2565b60405180910390fd5b60006002670de0b6b3a7640000610b16919061124d565b82610b2191906111f7565b905081811015610b66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5d906110e2565b60405180910390fd5b6000670de0b6b3a764000082610b7c919061124d565b905080935050505092915050565b6064670de0b6b3a7640000610b9f919061127e565b81565b60016003670de0b6b3a7640000610bb9919061124d565b610bc391906111f7565b81565b600080610bd38787610900565b90506000610be18686610900565b90506000610bef8383610900565b90506000610c16670de0b6b3a7640000610c11670de0b6b3a76400008961084e565b610900565b9050610c228282610a96565b94505050505095945050505050565b600a81565b600281565b600181565b600a670de0b6b3a7640000610c55919061124d565b81565b6001670de0b6b3a76400006002610c6f919061127e565b610c7991906112d8565b81565b670de0b6b3a764000081565b600081565b6000806000600284610c9f919061136b565b1415610cb357670de0b6b3a7640000610cb5565b835b9050600283610cc4919061124d565b92505b60008314610d1057610cd98485610a96565b93506000600284610cea919061136b565b14610cfc57610cf98185610a96565b90505b600283610d09919061124d565b9250610cc7565b8091505092915050565b600080839050600080610d3587670de0b6b3a76400006108c5565b915091506000670de0b6b3a764000090506000819050600080600190505b888410610e16576000670de0b6b3a764000082610d70919061127e565b9050600080610d908a610d8b85670de0b6b3a764000061084e565b6108c5565b91509150610da787610da2848c610a96565b610a96565b9650610db38784610900565b96506000871415610dc657505050610e16565b8715610dd157841594505b8015610ddc57841594505b8415610df357610dec868861084e565b9550610e00565b610dfd868861066b565b95505b5050508080610e0e90611322565b915050610d53565b508196505050505050509392505050565b6019670de0b6b3a7640000610e3c919061127e565b81565b6002670de0b6b3a7640000610e54919061124d565b81565b6019670de0b6b3a7640000610e6c919061127e565b81565b600081359050610e7e81611519565b92915050565b600060208284031215610e9657600080fd5b6000610ea484828501610e6f565b91505092915050565b60008060408385031215610ec057600080fd5b6000610ece85828601610e6f565b9250506020610edf85828601610e6f565b9150509250929050565b600080600060608486031215610efe57600080fd5b6000610f0c86828701610e6f565b9350506020610f1d86828701610e6f565b9250506040610f2e86828701610e6f565b9150509250925092565b600080600080600060a08688031215610f5057600080fd5b6000610f5e88828901610e6f565b9550506020610f6f88828901610e6f565b9450506040610f8088828901610e6f565b9350506060610f9188828901610e6f565b9250506080610fa288828901610e6f565b9150509295509295909350565b610fb88161130c565b82525050565b6000610fcb6010836111e6565b9150610fd6826113fa565b602082019050919050565b6000610fee6010836111e6565b9150610ff982611423565b602082019050919050565b6000611011600c836111e6565b915061101c8261144c565b602082019050919050565b60006110346015836111e6565b915061103f82611475565b602082019050919050565b60006110576011836111e6565b91506110628261149e565b602082019050919050565b600061107a6010836111e6565b9150611085826114c7565b602082019050919050565b600061109d6016836111e6565b91506110a8826114f0565b602082019050919050565b6110bc81611318565b82525050565b600060208201905081810360008301526110db81610fbe565b9050919050565b600060208201905081810360008301526110fb81610fe1565b9050919050565b6000602082019050818103600083015261111b81611004565b9050919050565b6000602082019050818103600083015261113b81611027565b9050919050565b6000602082019050818103600083015261115b8161104a565b9050919050565b6000602082019050818103600083015261117b8161106d565b9050919050565b6000602082019050818103600083015261119b81611090565b9050919050565b60006020820190506111b760008301846110b3565b92915050565b60006040820190506111d260008301856110b3565b6111df6020830184610faf565b9392505050565b600082825260208201905092915050565b600061120282611318565b915061120d83611318565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156112425761124161139c565b5b828201905092915050565b600061125882611318565b915061126383611318565b925082611273576112726113cb565b5b828204905092915050565b600061128982611318565b915061129483611318565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156112cd576112cc61139c565b5b828202905092915050565b60006112e382611318565b91506112ee83611318565b9250828210156113015761130061139c565b5b828203905092915050565b60008115159050919050565b6000819050919050565b600061132d82611318565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156113605761135f61139c565b5b600182019050919050565b600061137682611318565b915061138183611318565b925082611391576113906113cb565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4552525f4449565f494e5445524e414c00000000000000000000000000000000600082015250565b7f4552525f4d554c5f4f564552464c4f5700000000000000000000000000000000600082015250565b7f4552525f4449565f5a45524f0000000000000000000000000000000000000000600082015250565b7f4552525f42504f575f424153455f544f4f5f4c4f570000000000000000000000600082015250565b7f4552525f5355425f554e444552464c4f57000000000000000000000000000000600082015250565b7f4552525f4144445f4f564552464c4f5700000000000000000000000000000000600082015250565b7f4552525f42504f575f424153455f544f4f5f4849474800000000000000000000600082015250565b61152281611318565b811461152d57600080fd5b5056fea2646970667358221220b21ce0d2bcccb5c02861668ffcefe53e1d4d1c5f82e867e1b759189f4884f48664736f6c63430008010033"

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

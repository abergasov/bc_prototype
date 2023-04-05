// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"_circleID\",\"type\":\"bytes16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CIRCLE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_SEEDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nameToSeedsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"seedsOwners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalSeedsAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516104b53803806104b583398181016040528101906100329190610109565b806fffffffffffffffffffffffffffffffff19166080816fffffffffffffffffffffffffffffffff1916815250506103e8600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050610136565b600080fd5b60007fffffffffffffffffffffffffffffffff0000000000000000000000000000000082169050919050565b6100e6816100b1565b81146100f157600080fd5b50565b600081519050610103816100dd565b92915050565b60006020828403121561011f5761011e6100ac565b5b600061012d848285016100f4565b91505092915050565b608051610365610150600039600060f001526103656000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630a8c6c5a146100515780630cb690531461006f578063668292921461009f578063f2b630f9146100bd575b600080fd5b6100596100ee565b60405161006691906101bf565b60405180910390f35b6100896004803603810190610084919061023d565b610112565b6040516100969190610283565b60405180910390f35b6100a761012a565b6040516100b49190610283565b60405180910390f35b6100d760048036038101906100d291906102ca565b610130565b6040516100e5929190610306565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b60016020528060005260406000206000915090505481565b6103e881565b6000818154811061014057600080fd5b90600052602060002090600202016000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082565b60007fffffffffffffffffffffffffffffffff0000000000000000000000000000000082169050919050565b6101b981610184565b82525050565b60006020820190506101d460008301846101b0565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061020a826101df565b9050919050565b61021a816101ff565b811461022557600080fd5b50565b60008135905061023781610211565b92915050565b600060208284031215610253576102526101da565b5b600061026184828501610228565b91505092915050565b6000819050919050565b61027d8161026a565b82525050565b60006020820190506102986000830184610274565b92915050565b6102a78161026a565b81146102b257600080fd5b50565b6000813590506102c48161029e565b92915050565b6000602082840312156102e0576102df6101da565b5b60006102ee848285016102b5565b91505092915050565b610300816101ff565b82525050565b600060408201905061031b6000830185610274565b61032860208301846102f7565b939250505056fea2646970667358221220176a39fbd848c4ec96101948c36cf88c1d93635cde732092597684b51918948a64736f6c63430008130033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _circleID [16]byte) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _circleID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// CIRCLEADDRESS is a free data retrieval call binding the contract method 0x0a8c6c5a.
//
// Solidity: function CIRCLE_ADDRESS() view returns(bytes16)
func (_Contracts *ContractsCaller) CIRCLEADDRESS(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "CIRCLE_ADDRESS")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// CIRCLEADDRESS is a free data retrieval call binding the contract method 0x0a8c6c5a.
//
// Solidity: function CIRCLE_ADDRESS() view returns(bytes16)
func (_Contracts *ContractsSession) CIRCLEADDRESS() ([16]byte, error) {
	return _Contracts.Contract.CIRCLEADDRESS(&_Contracts.CallOpts)
}

// CIRCLEADDRESS is a free data retrieval call binding the contract method 0x0a8c6c5a.
//
// Solidity: function CIRCLE_ADDRESS() view returns(bytes16)
func (_Contracts *ContractsCallerSession) CIRCLEADDRESS() ([16]byte, error) {
	return _Contracts.Contract.CIRCLEADDRESS(&_Contracts.CallOpts)
}

// TOTALSEEDS is a free data retrieval call binding the contract method 0x66829292.
//
// Solidity: function TOTAL_SEEDS() view returns(uint256)
func (_Contracts *ContractsCaller) TOTALSEEDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "TOTAL_SEEDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALSEEDS is a free data retrieval call binding the contract method 0x66829292.
//
// Solidity: function TOTAL_SEEDS() view returns(uint256)
func (_Contracts *ContractsSession) TOTALSEEDS() (*big.Int, error) {
	return _Contracts.Contract.TOTALSEEDS(&_Contracts.CallOpts)
}

// TOTALSEEDS is a free data retrieval call binding the contract method 0x66829292.
//
// Solidity: function TOTAL_SEEDS() view returns(uint256)
func (_Contracts *ContractsCallerSession) TOTALSEEDS() (*big.Int, error) {
	return _Contracts.Contract.TOTALSEEDS(&_Contracts.CallOpts)
}

// NameToSeedsAmount is a free data retrieval call binding the contract method 0x0cb69053.
//
// Solidity: function nameToSeedsAmount(address ) view returns(uint256)
func (_Contracts *ContractsCaller) NameToSeedsAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "nameToSeedsAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NameToSeedsAmount is a free data retrieval call binding the contract method 0x0cb69053.
//
// Solidity: function nameToSeedsAmount(address ) view returns(uint256)
func (_Contracts *ContractsSession) NameToSeedsAmount(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.NameToSeedsAmount(&_Contracts.CallOpts, arg0)
}

// NameToSeedsAmount is a free data retrieval call binding the contract method 0x0cb69053.
//
// Solidity: function nameToSeedsAmount(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) NameToSeedsAmount(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.NameToSeedsAmount(&_Contracts.CallOpts, arg0)
}

// SeedsOwners is a free data retrieval call binding the contract method 0xf2b630f9.
//
// Solidity: function seedsOwners(uint256 ) view returns(uint256 totalSeedsAmount, address member)
func (_Contracts *ContractsCaller) SeedsOwners(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TotalSeedsAmount *big.Int
	Member           common.Address
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "seedsOwners", arg0)

	outstruct := new(struct {
		TotalSeedsAmount *big.Int
		Member           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalSeedsAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Member = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SeedsOwners is a free data retrieval call binding the contract method 0xf2b630f9.
//
// Solidity: function seedsOwners(uint256 ) view returns(uint256 totalSeedsAmount, address member)
func (_Contracts *ContractsSession) SeedsOwners(arg0 *big.Int) (struct {
	TotalSeedsAmount *big.Int
	Member           common.Address
}, error) {
	return _Contracts.Contract.SeedsOwners(&_Contracts.CallOpts, arg0)
}

// SeedsOwners is a free data retrieval call binding the contract method 0xf2b630f9.
//
// Solidity: function seedsOwners(uint256 ) view returns(uint256 totalSeedsAmount, address member)
func (_Contracts *ContractsCallerSession) SeedsOwners(arg0 *big.Int) (struct {
	TotalSeedsAmount *big.Int
	Member           common.Address
}, error) {
	return _Contracts.Contract.SeedsOwners(&_Contracts.CallOpts, arg0)
}

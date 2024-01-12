// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"

	ethereum "github.com/ledgerwatch/erigon"
	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/accounts/abi"
	"github.com/ledgerwatch/erigon/accounts/abi/bind"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = libcommon.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PairStoreABI is the input ABI used to generate the binding from.
const PairStoreABI = "[{\"inputs\":[],\"name\":\"getFirst\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setFirst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setSecond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PairStoreBin is the compiled bytecode used for deploying new contracts.
var PairStoreBin = "0x608060405234801561001057600080fd5b5060d88061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060465760003560e01c80631b03316f14604b5780631e2231431460615780635ef3d3dd146068578063b698c12914607a575b600080fd5b6001545b60405190815260200160405180910390f35b600054604f565b60786073366004608a565b600055565b005b60786085366004608a565b600155565b600060208284031215609b57600080fd5b503591905056fea2646970667358221220a0124977f16d8e156cc1889025d5fce79e5a2d54aa1d0116577f2fe4053373ee64736f6c63430008130033"

// DeployPairStore deploys a new Ethereum contract, binding an instance of PairStore to it.
func DeployPairStore(auth *bind.TransactOpts, backend bind.ContractBackend) (libcommon.Address, types.Transaction, *PairStore, error) {
	parsed, err := abi.JSON(strings.NewReader(PairStoreABI))
	if err != nil {
		return libcommon.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, libcommon.FromHex(PairStoreBin), backend)
	if err != nil {
		return libcommon.Address{}, nil, nil, err
	}
	return address, tx, &PairStore{PairStoreCaller: PairStoreCaller{contract: contract}, PairStoreTransactor: PairStoreTransactor{contract: contract}, PairStoreFilterer: PairStoreFilterer{contract: contract}}, nil
}

// PairStore is an auto generated Go binding around an Ethereum contract.
type PairStore struct {
	PairStoreCaller     // Read-only binding to the contract
	PairStoreTransactor // Write-only binding to the contract
	PairStoreFilterer   // Log filterer for contract events
}

// PairStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairStoreSession struct {
	Contract     *PairStore        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairStoreCallerSession struct {
	Contract *PairStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PairStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairStoreTransactorSession struct {
	Contract     *PairStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PairStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairStoreRaw struct {
	Contract *PairStore // Generic contract binding to access the raw methods on
}

// PairStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairStoreCallerRaw struct {
	Contract *PairStoreCaller // Generic read-only contract binding to access the raw methods on
}

// PairStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairStoreTransactorRaw struct {
	Contract *PairStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairStore creates a new instance of PairStore, bound to a specific deployed contract.
func NewPairStore(address libcommon.Address, backend bind.ContractBackend) (*PairStore, error) {
	contract, err := bindPairStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PairStore{PairStoreCaller: PairStoreCaller{contract: contract}, PairStoreTransactor: PairStoreTransactor{contract: contract}, PairStoreFilterer: PairStoreFilterer{contract: contract}}, nil
}

// NewPairStoreCaller creates a new read-only instance of PairStore, bound to a specific deployed contract.
func NewPairStoreCaller(address libcommon.Address, caller bind.ContractCaller) (*PairStoreCaller, error) {
	contract, err := bindPairStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairStoreCaller{contract: contract}, nil
}

// NewPairStoreTransactor creates a new write-only instance of PairStore, bound to a specific deployed contract.
func NewPairStoreTransactor(address libcommon.Address, transactor bind.ContractTransactor) (*PairStoreTransactor, error) {
	contract, err := bindPairStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairStoreTransactor{contract: contract}, nil
}

// NewPairStoreFilterer creates a new log filterer instance of PairStore, bound to a specific deployed contract.
func NewPairStoreFilterer(address libcommon.Address, filterer bind.ContractFilterer) (*PairStoreFilterer, error) {
	contract, err := bindPairStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairStoreFilterer{contract: contract}, nil
}

// bindPairStore binds a generic wrapper to an already deployed contract.
func bindPairStore(address libcommon.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairStoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairStore *PairStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PairStore.Contract.PairStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairStore *PairStoreRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _PairStore.Contract.PairStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairStore *PairStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _PairStore.Contract.PairStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairStore *PairStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PairStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairStore *PairStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _PairStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairStore *PairStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _PairStore.Contract.contract.Transact(opts, method, params...)
}

// GetFirst is a free data retrieval call binding the contract method 0x1e223143.
//
// Solidity: function getFirst() view returns(uint256)
func (_PairStore *PairStoreCaller) GetFirst(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairStore.contract.Call(opts, &out, "getFirst")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFirst is a free data retrieval call binding the contract method 0x1e223143.
//
// Solidity: function getFirst() view returns(uint256)
func (_PairStore *PairStoreSession) GetFirst() (*big.Int, error) {
	return _PairStore.Contract.GetFirst(&_PairStore.CallOpts)
}

// GetFirst is a free data retrieval call binding the contract method 0x1e223143.
//
// Solidity: function getFirst() view returns(uint256)
func (_PairStore *PairStoreCallerSession) GetFirst() (*big.Int, error) {
	return _PairStore.Contract.GetFirst(&_PairStore.CallOpts)
}

// GetSecond is a free data retrieval call binding the contract method 0x1b03316f.
//
// Solidity: function getSecond() view returns(uint256)
func (_PairStore *PairStoreCaller) GetSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairStore.contract.Call(opts, &out, "getSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSecond is a free data retrieval call binding the contract method 0x1b03316f.
//
// Solidity: function getSecond() view returns(uint256)
func (_PairStore *PairStoreSession) GetSecond() (*big.Int, error) {
	return _PairStore.Contract.GetSecond(&_PairStore.CallOpts)
}

// GetSecond is a free data retrieval call binding the contract method 0x1b03316f.
//
// Solidity: function getSecond() view returns(uint256)
func (_PairStore *PairStoreCallerSession) GetSecond() (*big.Int, error) {
	return _PairStore.Contract.GetSecond(&_PairStore.CallOpts)
}

// SetFirst is a paid mutator transaction binding the contract method 0x5ef3d3dd.
//
// Solidity: function setFirst(uint256 value) returns()
func (_PairStore *PairStoreTransactor) SetFirst(opts *bind.TransactOpts, value *big.Int) (types.Transaction, error) {
	return _PairStore.contract.Transact(opts, "setFirst", value)
}

// SetFirst is a paid mutator transaction binding the contract method 0x5ef3d3dd.
//
// Solidity: function setFirst(uint256 value) returns()
func (_PairStore *PairStoreSession) SetFirst(value *big.Int) (types.Transaction, error) {
	return _PairStore.Contract.SetFirst(&_PairStore.TransactOpts, value)
}

// SetFirst is a paid mutator transaction binding the contract method 0x5ef3d3dd.
//
// Solidity: function setFirst(uint256 value) returns()
func (_PairStore *PairStoreTransactorSession) SetFirst(value *big.Int) (types.Transaction, error) {
	return _PairStore.Contract.SetFirst(&_PairStore.TransactOpts, value)
}

// SetSecond is a paid mutator transaction binding the contract method 0xb698c129.
//
// Solidity: function setSecond(uint256 value) returns()
func (_PairStore *PairStoreTransactor) SetSecond(opts *bind.TransactOpts, value *big.Int) (types.Transaction, error) {
	return _PairStore.contract.Transact(opts, "setSecond", value)
}

// SetSecond is a paid mutator transaction binding the contract method 0xb698c129.
//
// Solidity: function setSecond(uint256 value) returns()
func (_PairStore *PairStoreSession) SetSecond(value *big.Int) (types.Transaction, error) {
	return _PairStore.Contract.SetSecond(&_PairStore.TransactOpts, value)
}

// SetSecond is a paid mutator transaction binding the contract method 0xb698c129.
//
// Solidity: function setSecond(uint256 value) returns()
func (_PairStore *PairStoreTransactorSession) SetSecond(value *big.Int) (types.Transaction, error) {
	return _PairStore.Contract.SetSecond(&_PairStore.TransactOpts, value)
}

// SetFirstParams is an auto generated read-only Go binding of transcaction calldata params
type SetFirstParams struct {
	Param_value *big.Int
}

// Parse SetFirst method from calldata of a transaction
//
// Solidity: function setFirst(uint256 value) returns()
func ParseSetFirst(calldata []byte) (*SetFirstParams, error) {
	if len(calldata) <= 4 {
		return nil, fmt.Errorf("invalid calldata input")
	}

	_abi, err := abi.JSON(strings.NewReader(PairStoreABI))
	if err != nil {
		return nil, fmt.Errorf("failed to get abi of registry metadata: %w", err)
	}

	out, err := _abi.Methods["setFirst"].Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack setFirst params data: %w", err)
	}

	var paramsResult = new(SetFirstParams)
	value := reflect.ValueOf(paramsResult).Elem()

	if value.NumField() != len(out) {
		return nil, fmt.Errorf("failed to match calldata with param field number")
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &SetFirstParams{
		Param_value: out0,
	}, nil
}

// SetSecondParams is an auto generated read-only Go binding of transcaction calldata params
type SetSecondParams struct {
	Param_value *big.Int
}

// Parse SetSecond method from calldata of a transaction
//
// Solidity: function setSecond(uint256 value) returns()
func ParseSetSecond(calldata []byte) (*SetSecondParams, error) {
	if len(calldata) <= 4 {
		return nil, fmt.Errorf("invalid calldata input")
	}

	_abi, err := abi.JSON(strings.NewReader(PairStoreABI))
	if err != nil {
		return nil, fmt.Errorf("failed to get abi of registry metadata: %w", err)
	}

	out, err := _abi.Methods["setSecond"].Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack setSecond params data: %w", err)
	}

	var paramsResult = new(SetSecondParams)
	value := reflect.ValueOf(paramsResult).Elem()

	if value.NumField() != len(out) {
		return nil, fmt.Errorf("failed to match calldata with param field number")
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &SetSecondParams{
		Param_value: out0,
	}, nil
}

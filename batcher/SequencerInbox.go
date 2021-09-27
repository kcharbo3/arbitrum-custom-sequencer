// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package batcher

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
)

// BatcherMetaData contains all meta data concerning the Batcher contract.
var BatcherMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[2]\",\"name\":\"afterAccAndDelayed\",\"type\":\"bytes32[2]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"}],\"name\":\"DelayedInboxForced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSequencer\",\"type\":\"bool\"}],\"name\":\"IsSequencerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxDelayBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxDelaySeconds\",\"type\":\"uint256\"}],\"name\":\"MaxDelayUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"sectionsMetadata\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"SequencerBatchDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"}],\"name\":\"SequencerBatchDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sectionsMetadata\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"}],\"name\":\"addSequencerL2Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sectionsMetadata\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"}],\"name\":\"addSequencerL2BatchFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sectionsMetadata\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"}],\"name\":\"addSequencerL2BatchFromOriginWithGasRefunder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedInbox\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint256[2]\",\"name\":\"l1BlockAndTimestamp\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceL1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"}],\"name\":\"forceInclusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInboxAccsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_delayedInbox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDelayBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDelaySeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_messageCount\",\"type\":\"uint256\"}],\"name\":\"proveBatchContainsSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_messageCount\",\"type\":\"uint256\"}],\"name\":\"proveInboxContainsMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"newIsSequencer\",\"type\":\"bool\"}],\"name\":\"setIsSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxDelayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMaxDelaySeconds\",\"type\":\"uint256\"}],\"name\":\"setMaxDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDelayedMessagesRead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BatcherABI is the input ABI used to generate the binding from.
// Deprecated: Use BatcherMetaData.ABI instead.
var BatcherABI = BatcherMetaData.ABI

// Batcher is an auto generated Go binding around an Ethereum contract.
type Batcher struct {
	BatcherCaller     // Read-only binding to the contract
	BatcherTransactor // Write-only binding to the contract
	BatcherFilterer   // Log filterer for contract events
}

// BatcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatcherSession struct {
	Contract     *Batcher          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatcherCallerSession struct {
	Contract *BatcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BatcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatcherTransactorSession struct {
	Contract     *BatcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BatcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatcherRaw struct {
	Contract *Batcher // Generic contract binding to access the raw methods on
}

// BatcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatcherCallerRaw struct {
	Contract *BatcherCaller // Generic read-only contract binding to access the raw methods on
}

// BatcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatcherTransactorRaw struct {
	Contract *BatcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatcher creates a new instance of Batcher, bound to a specific deployed contract.
func NewBatcher(address common.Address, backend bind.ContractBackend) (*Batcher, error) {
	contract, err := bindBatcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Batcher{BatcherCaller: BatcherCaller{contract: contract}, BatcherTransactor: BatcherTransactor{contract: contract}, BatcherFilterer: BatcherFilterer{contract: contract}}, nil
}

// NewBatcherCaller creates a new read-only instance of Batcher, bound to a specific deployed contract.
func NewBatcherCaller(address common.Address, caller bind.ContractCaller) (*BatcherCaller, error) {
	contract, err := bindBatcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatcherCaller{contract: contract}, nil
}

// NewBatcherTransactor creates a new write-only instance of Batcher, bound to a specific deployed contract.
func NewBatcherTransactor(address common.Address, transactor bind.ContractTransactor) (*BatcherTransactor, error) {
	contract, err := bindBatcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatcherTransactor{contract: contract}, nil
}

// NewBatcherFilterer creates a new log filterer instance of Batcher, bound to a specific deployed contract.
func NewBatcherFilterer(address common.Address, filterer bind.ContractFilterer) (*BatcherFilterer, error) {
	contract, err := bindBatcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatcherFilterer{contract: contract}, nil
}

// bindBatcher binds a generic wrapper to an already deployed contract.
func bindBatcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BatcherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Batcher *BatcherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Batcher.Contract.BatcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Batcher *BatcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Batcher.Contract.BatcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Batcher *BatcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Batcher.Contract.BatcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Batcher *BatcherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Batcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Batcher *BatcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Batcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Batcher *BatcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Batcher.Contract.contract.Transact(opts, method, params...)
}

// DelayedInbox is a free data retrieval call binding the contract method 0xb71939b1.
//
// Solidity: function delayedInbox() view returns(address)
func (_Batcher *BatcherCaller) DelayedInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Batcher.contract.Call(opts, &out, "delayedInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayedInbox is a free data retrieval call binding the contract method 0xb71939b1.
//
// Solidity: function delayedInbox() view returns(address)
func (_Batcher *BatcherSession) DelayedInbox() (common.Address, error) {
	return _Batcher.Contract.DelayedInbox(&_Batcher.CallOpts)
}

// DelayedInbox is a free data retrieval call binding the contract method 0xb71939b1.
//
// Solidity: function delayedInbox() view returns(address)
func (_Batcher *BatcherCallerSession) DelayedInbox() (common.Address, error) {
	return _Batcher.Contract.DelayedInbox(&_Batcher.CallOpts)
}

// GetInboxAccsLength is a free data retrieval call binding the contract method 0xd9b141ff.
//
// Solidity: function getInboxAccsLength() view returns(uint256)
func (_Batcher *BatcherCaller) GetInboxAccsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Batcher.contract.Call(opts, &out, "getInboxAccsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInboxAccsLength is a free data retrieval call binding the contract method 0xd9b141ff.
//
// Solidity: function getInboxAccsLength() view returns(uint256)
func (_Batcher *BatcherSession) GetInboxAccsLength() (*big.Int, error) {
	return _Batcher.Contract.GetInboxAccsLength(&_Batcher.CallOpts)
}

// GetInboxAccsLength is a free data retrieval call binding the contract method 0xd9b141ff.
//
// Solidity: function getInboxAccsLength() view returns(uint256)
func (_Batcher *BatcherCallerSession) GetInboxAccsLength() (*big.Int, error) {
	return _Batcher.Contract.GetInboxAccsLength(&_Batcher.CallOpts)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_Batcher *BatcherCaller) InboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Batcher.contract.Call(opts, &out, "inboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_Batcher *BatcherSession) InboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _Batcher.Contract.InboxAccs(&_Batcher.CallOpts, arg0)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_Batcher *BatcherCallerSession) InboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _Batcher.Contract.InboxAccs(&_Batcher.CallOpts, arg0)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Batcher *BatcherCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Batcher.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Batcher *BatcherSession) IsMaster() (bool, error) {
	return _Batcher.Contract.IsMaster(&_Batcher.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Batcher *BatcherCallerSession) IsMaster() (bool, error) {
	return _Batcher.Contract.IsMaster(&_Batcher.CallOpts)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_Batcher *BatcherCaller) IsSequencer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Batcher.contract.Call(opts, &out, "isSequencer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_Batcher *BatcherSession) IsSequencer(arg0 common.Address) (bool, error) {
	return _Batcher.Contract.IsSequencer(&_Batcher.CallOpts, arg0)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_Batcher *BatcherCallerSession) IsSequencer(arg0 common.Address) (bool, error) {
	return _Batcher.Contract.IsSequencer(&_Batcher.CallOpts, arg0)
}

// MaxDelayBlocks is a free data retrieval call binding the contract method 0xe367a2c1.
//
// Solidity: function maxDelayBlocks() view returns(uint256)
func (_Batcher *BatcherCaller) MaxDelayBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Batcher.contract.Call(opts, &out, "maxDelayBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDelayBlocks is a free data retrieval call binding the contract method 0xe367a2c1.
//
// Solidity: function maxDelayBlocks() view returns(uint256)
func (_Batcher *BatcherSession) MaxDelayBlocks() (*big.Int, error) {
	return _Batcher.Contract.MaxDelayBlocks(&_Batcher.CallOpts)
}

// MaxDelayBlocks is a free data retrieval call binding the contract method 0xe367a2c1.
//
// Solidity: function maxDelayBlocks() view returns(uint256)
func (_Batcher *BatcherCallerSession) MaxDelayBlocks() (*big.Int, error) {
	return _Batcher.Contract.MaxDelayBlocks(&_Batcher.CallOpts)
}

// MaxDelaySeconds is a free data retrieval call binding the contract method 0x342025fa.
//
// Solidity: function maxDelaySeconds() view returns(uint256)
func (_Batcher *BatcherCaller) MaxDelaySeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Batcher.contract.Call(opts, &out, "maxDelaySeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDelaySeconds is a free data retrieval call binding the contract method 0x342025fa.
//
// Solidity: function maxDelaySeconds() view returns(uint256)
func (_Batcher *BatcherSession) MaxDelaySeconds() (*big.Int, error) {
	return _Batcher.Contract.MaxDelaySeconds(&_Batcher.CallOpts)
}

// MaxDelaySeconds is a free data retrieval call binding the contract method 0x342025fa.
//
// Solidity: function maxDelaySeconds() view returns(uint256)
func (_Batcher *BatcherCallerSession) MaxDelaySeconds() (*big.Int, error) {
	return _Batcher.Contract.MaxDelaySeconds(&_Batcher.CallOpts)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Batcher *BatcherCaller) MessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Batcher.contract.Call(opts, &out, "messageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Batcher *BatcherSession) MessageCount() (*big.Int, error) {
	return _Batcher.Contract.MessageCount(&_Batcher.CallOpts)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Batcher *BatcherCallerSession) MessageCount() (*big.Int, error) {
	return _Batcher.Contract.MessageCount(&_Batcher.CallOpts)
}

// PostUpgradeInit is a free data retrieval call binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() view returns()
func (_Batcher *BatcherCaller) PostUpgradeInit(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Batcher.contract.Call(opts, &out, "postUpgradeInit")

	if err != nil {
		return err
	}

	return err

}

// PostUpgradeInit is a free data retrieval call binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() view returns()
func (_Batcher *BatcherSession) PostUpgradeInit() error {
	return _Batcher.Contract.PostUpgradeInit(&_Batcher.CallOpts)
}

// PostUpgradeInit is a free data retrieval call binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() view returns()
func (_Batcher *BatcherCallerSession) PostUpgradeInit() error {
	return _Batcher.Contract.PostUpgradeInit(&_Batcher.CallOpts)
}

// ProveBatchContainsSequenceNumber is a free data retrieval call binding the contract method 0x06cc91b2.
//
// Solidity: function proveBatchContainsSequenceNumber(bytes proof, uint256 _messageCount) view returns(uint256, bytes32)
func (_Batcher *BatcherCaller) ProveBatchContainsSequenceNumber(opts *bind.CallOpts, proof []byte, _messageCount *big.Int) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _Batcher.contract.Call(opts, &out, "proveBatchContainsSequenceNumber", proof, _messageCount)

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// ProveBatchContainsSequenceNumber is a free data retrieval call binding the contract method 0x06cc91b2.
//
// Solidity: function proveBatchContainsSequenceNumber(bytes proof, uint256 _messageCount) view returns(uint256, bytes32)
func (_Batcher *BatcherSession) ProveBatchContainsSequenceNumber(proof []byte, _messageCount *big.Int) (*big.Int, [32]byte, error) {
	return _Batcher.Contract.ProveBatchContainsSequenceNumber(&_Batcher.CallOpts, proof, _messageCount)
}

// ProveBatchContainsSequenceNumber is a free data retrieval call binding the contract method 0x06cc91b2.
//
// Solidity: function proveBatchContainsSequenceNumber(bytes proof, uint256 _messageCount) view returns(uint256, bytes32)
func (_Batcher *BatcherCallerSession) ProveBatchContainsSequenceNumber(proof []byte, _messageCount *big.Int) (*big.Int, [32]byte, error) {
	return _Batcher.Contract.ProveBatchContainsSequenceNumber(&_Batcher.CallOpts, proof, _messageCount)
}

// ProveInboxContainsMessage is a free data retrieval call binding the contract method 0xdc1b7b1f.
//
// Solidity: function proveInboxContainsMessage(bytes proof, uint256 _messageCount) view returns(uint256, bytes32)
func (_Batcher *BatcherCaller) ProveInboxContainsMessage(opts *bind.CallOpts, proof []byte, _messageCount *big.Int) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _Batcher.contract.Call(opts, &out, "proveInboxContainsMessage", proof, _messageCount)

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// ProveInboxContainsMessage is a free data retrieval call binding the contract method 0xdc1b7b1f.
//
// Solidity: function proveInboxContainsMessage(bytes proof, uint256 _messageCount) view returns(uint256, bytes32)
func (_Batcher *BatcherSession) ProveInboxContainsMessage(proof []byte, _messageCount *big.Int) (*big.Int, [32]byte, error) {
	return _Batcher.Contract.ProveInboxContainsMessage(&_Batcher.CallOpts, proof, _messageCount)
}

// ProveInboxContainsMessage is a free data retrieval call binding the contract method 0xdc1b7b1f.
//
// Solidity: function proveInboxContainsMessage(bytes proof, uint256 _messageCount) view returns(uint256, bytes32)
func (_Batcher *BatcherCallerSession) ProveInboxContainsMessage(proof []byte, _messageCount *big.Int) (*big.Int, [32]byte, error) {
	return _Batcher.Contract.ProveInboxContainsMessage(&_Batcher.CallOpts, proof, _messageCount)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_Batcher *BatcherCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Batcher.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_Batcher *BatcherSession) Rollup() (common.Address, error) {
	return _Batcher.Contract.Rollup(&_Batcher.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_Batcher *BatcherCallerSession) Rollup() (common.Address, error) {
	return _Batcher.Contract.Rollup(&_Batcher.CallOpts)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_Batcher *BatcherCaller) Sequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Batcher.contract.Call(opts, &out, "sequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_Batcher *BatcherSession) Sequencer() (common.Address, error) {
	return _Batcher.Contract.Sequencer(&_Batcher.CallOpts)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_Batcher *BatcherCallerSession) Sequencer() (common.Address, error) {
	return _Batcher.Contract.Sequencer(&_Batcher.CallOpts)
}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_Batcher *BatcherCaller) TotalDelayedMessagesRead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Batcher.contract.Call(opts, &out, "totalDelayedMessagesRead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_Batcher *BatcherSession) TotalDelayedMessagesRead() (*big.Int, error) {
	return _Batcher.Contract.TotalDelayedMessagesRead(&_Batcher.CallOpts)
}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_Batcher *BatcherCallerSession) TotalDelayedMessagesRead() (*big.Int, error) {
	return _Batcher.Contract.TotalDelayedMessagesRead(&_Batcher.CallOpts)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x1a734229.
//
// Solidity: function addSequencerL2Batch(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc) returns()
func (_Batcher *BatcherTransactor) AddSequencerL2Batch(opts *bind.TransactOpts, transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _Batcher.contract.Transact(opts, "addSequencerL2Batch", transactions, lengths, sectionsMetadata, afterAcc)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x1a734229.
//
// Solidity: function addSequencerL2Batch(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc) returns()
func (_Batcher *BatcherSession) AddSequencerL2Batch(transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _Batcher.Contract.AddSequencerL2Batch(&_Batcher.TransactOpts, transactions, lengths, sectionsMetadata, afterAcc)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x1a734229.
//
// Solidity: function addSequencerL2Batch(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc) returns()
func (_Batcher *BatcherTransactorSession) AddSequencerL2Batch(transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _Batcher.Contract.AddSequencerL2Batch(&_Batcher.TransactOpts, transactions, lengths, sectionsMetadata, afterAcc)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x44c7cc30.
//
// Solidity: function addSequencerL2BatchFromOrigin(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc) returns()
func (_Batcher *BatcherTransactor) AddSequencerL2BatchFromOrigin(opts *bind.TransactOpts, transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _Batcher.contract.Transact(opts, "addSequencerL2BatchFromOrigin", transactions, lengths, sectionsMetadata, afterAcc)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x44c7cc30.
//
// Solidity: function addSequencerL2BatchFromOrigin(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc) returns()
func (_Batcher *BatcherSession) AddSequencerL2BatchFromOrigin(transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _Batcher.Contract.AddSequencerL2BatchFromOrigin(&_Batcher.TransactOpts, transactions, lengths, sectionsMetadata, afterAcc)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x44c7cc30.
//
// Solidity: function addSequencerL2BatchFromOrigin(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc) returns()
func (_Batcher *BatcherTransactorSession) AddSequencerL2BatchFromOrigin(transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _Batcher.Contract.AddSequencerL2BatchFromOrigin(&_Batcher.TransactOpts, transactions, lengths, sectionsMetadata, afterAcc)
}

// AddSequencerL2BatchFromOriginWithGasRefunder is a paid mutator transaction binding the contract method 0x8a2df18d.
//
// Solidity: function addSequencerL2BatchFromOriginWithGasRefunder(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc, address gasRefunder) returns()
func (_Batcher *BatcherTransactor) AddSequencerL2BatchFromOriginWithGasRefunder(opts *bind.TransactOpts, transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte, gasRefunder common.Address) (*types.Transaction, error) {
	return _Batcher.contract.Transact(opts, "addSequencerL2BatchFromOriginWithGasRefunder", transactions, lengths, sectionsMetadata, afterAcc, gasRefunder)
}

// AddSequencerL2BatchFromOriginWithGasRefunder is a paid mutator transaction binding the contract method 0x8a2df18d.
//
// Solidity: function addSequencerL2BatchFromOriginWithGasRefunder(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc, address gasRefunder) returns()
func (_Batcher *BatcherSession) AddSequencerL2BatchFromOriginWithGasRefunder(transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte, gasRefunder common.Address) (*types.Transaction, error) {
	return _Batcher.Contract.AddSequencerL2BatchFromOriginWithGasRefunder(&_Batcher.TransactOpts, transactions, lengths, sectionsMetadata, afterAcc, gasRefunder)
}

// AddSequencerL2BatchFromOriginWithGasRefunder is a paid mutator transaction binding the contract method 0x8a2df18d.
//
// Solidity: function addSequencerL2BatchFromOriginWithGasRefunder(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc, address gasRefunder) returns()
func (_Batcher *BatcherTransactorSession) AddSequencerL2BatchFromOriginWithGasRefunder(transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte, gasRefunder common.Address) (*types.Transaction, error) {
	return _Batcher.Contract.AddSequencerL2BatchFromOriginWithGasRefunder(&_Batcher.TransactOpts, transactions, lengths, sectionsMetadata, afterAcc, gasRefunder)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0x0c4a1e59.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint256[2] l1BlockAndTimestamp, uint256 inboxSeqNum, uint256 gasPriceL1, address sender, bytes32 messageDataHash, bytes32 delayedAcc) returns()
func (_Batcher *BatcherTransactor) ForceInclusion(opts *bind.TransactOpts, _totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTimestamp [2]*big.Int, inboxSeqNum *big.Int, gasPriceL1 *big.Int, sender common.Address, messageDataHash [32]byte, delayedAcc [32]byte) (*types.Transaction, error) {
	return _Batcher.contract.Transact(opts, "forceInclusion", _totalDelayedMessagesRead, kind, l1BlockAndTimestamp, inboxSeqNum, gasPriceL1, sender, messageDataHash, delayedAcc)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0x0c4a1e59.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint256[2] l1BlockAndTimestamp, uint256 inboxSeqNum, uint256 gasPriceL1, address sender, bytes32 messageDataHash, bytes32 delayedAcc) returns()
func (_Batcher *BatcherSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTimestamp [2]*big.Int, inboxSeqNum *big.Int, gasPriceL1 *big.Int, sender common.Address, messageDataHash [32]byte, delayedAcc [32]byte) (*types.Transaction, error) {
	return _Batcher.Contract.ForceInclusion(&_Batcher.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTimestamp, inboxSeqNum, gasPriceL1, sender, messageDataHash, delayedAcc)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0x0c4a1e59.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint256[2] l1BlockAndTimestamp, uint256 inboxSeqNum, uint256 gasPriceL1, address sender, bytes32 messageDataHash, bytes32 delayedAcc) returns()
func (_Batcher *BatcherTransactorSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTimestamp [2]*big.Int, inboxSeqNum *big.Int, gasPriceL1 *big.Int, sender common.Address, messageDataHash [32]byte, delayedAcc [32]byte) (*types.Transaction, error) {
	return _Batcher.Contract.ForceInclusion(&_Batcher.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTimestamp, inboxSeqNum, gasPriceL1, sender, messageDataHash, delayedAcc)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _delayedInbox, address _sequencer, address _rollup) returns()
func (_Batcher *BatcherTransactor) Initialize(opts *bind.TransactOpts, _delayedInbox common.Address, _sequencer common.Address, _rollup common.Address) (*types.Transaction, error) {
	return _Batcher.contract.Transact(opts, "initialize", _delayedInbox, _sequencer, _rollup)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _delayedInbox, address _sequencer, address _rollup) returns()
func (_Batcher *BatcherSession) Initialize(_delayedInbox common.Address, _sequencer common.Address, _rollup common.Address) (*types.Transaction, error) {
	return _Batcher.Contract.Initialize(&_Batcher.TransactOpts, _delayedInbox, _sequencer, _rollup)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _delayedInbox, address _sequencer, address _rollup) returns()
func (_Batcher *BatcherTransactorSession) Initialize(_delayedInbox common.Address, _sequencer common.Address, _rollup common.Address) (*types.Transaction, error) {
	return _Batcher.Contract.Initialize(&_Batcher.TransactOpts, _delayedInbox, _sequencer, _rollup)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool newIsSequencer) returns()
func (_Batcher *BatcherTransactor) SetIsSequencer(opts *bind.TransactOpts, addr common.Address, newIsSequencer bool) (*types.Transaction, error) {
	return _Batcher.contract.Transact(opts, "setIsSequencer", addr, newIsSequencer)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool newIsSequencer) returns()
func (_Batcher *BatcherSession) SetIsSequencer(addr common.Address, newIsSequencer bool) (*types.Transaction, error) {
	return _Batcher.Contract.SetIsSequencer(&_Batcher.TransactOpts, addr, newIsSequencer)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool newIsSequencer) returns()
func (_Batcher *BatcherTransactorSession) SetIsSequencer(addr common.Address, newIsSequencer bool) (*types.Transaction, error) {
	return _Batcher.Contract.SetIsSequencer(&_Batcher.TransactOpts, addr, newIsSequencer)
}

// SetMaxDelay is a paid mutator transaction binding the contract method 0x4d480faa.
//
// Solidity: function setMaxDelay(uint256 newMaxDelayBlocks, uint256 newMaxDelaySeconds) returns()
func (_Batcher *BatcherTransactor) SetMaxDelay(opts *bind.TransactOpts, newMaxDelayBlocks *big.Int, newMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _Batcher.contract.Transact(opts, "setMaxDelay", newMaxDelayBlocks, newMaxDelaySeconds)
}

// SetMaxDelay is a paid mutator transaction binding the contract method 0x4d480faa.
//
// Solidity: function setMaxDelay(uint256 newMaxDelayBlocks, uint256 newMaxDelaySeconds) returns()
func (_Batcher *BatcherSession) SetMaxDelay(newMaxDelayBlocks *big.Int, newMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _Batcher.Contract.SetMaxDelay(&_Batcher.TransactOpts, newMaxDelayBlocks, newMaxDelaySeconds)
}

// SetMaxDelay is a paid mutator transaction binding the contract method 0x4d480faa.
//
// Solidity: function setMaxDelay(uint256 newMaxDelayBlocks, uint256 newMaxDelaySeconds) returns()
func (_Batcher *BatcherTransactorSession) SetMaxDelay(newMaxDelayBlocks *big.Int, newMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _Batcher.Contract.SetMaxDelay(&_Batcher.TransactOpts, newMaxDelayBlocks, newMaxDelaySeconds)
}

// BatcherDelayedInboxForcedIterator is returned from FilterDelayedInboxForced and is used to iterate over the raw logs and unpacked data for DelayedInboxForced events raised by the Batcher contract.
type BatcherDelayedInboxForcedIterator struct {
	Event *BatcherDelayedInboxForced // Event containing the contract specifics and raw log

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
func (it *BatcherDelayedInboxForcedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatcherDelayedInboxForced)
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
		it.Event = new(BatcherDelayedInboxForced)
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
func (it *BatcherDelayedInboxForcedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatcherDelayedInboxForcedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatcherDelayedInboxForced represents a DelayedInboxForced event raised by the Batcher contract.
type BatcherDelayedInboxForced struct {
	FirstMessageNum          *big.Int
	BeforeAcc                [32]byte
	NewMessageCount          *big.Int
	TotalDelayedMessagesRead *big.Int
	AfterAccAndDelayed       [2][32]byte
	SeqBatchIndex            *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterDelayedInboxForced is a free log retrieval operation binding the contract event 0x85b6a949bf20bfd6bc6e20f98fb490c7944ab61dcfa5a30b5dae543412c9a8a0.
//
// Solidity: event DelayedInboxForced(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, uint256 totalDelayedMessagesRead, bytes32[2] afterAccAndDelayed, uint256 seqBatchIndex)
func (_Batcher *BatcherFilterer) FilterDelayedInboxForced(opts *bind.FilterOpts, firstMessageNum []*big.Int, beforeAcc [][32]byte) (*BatcherDelayedInboxForcedIterator, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _Batcher.contract.FilterLogs(opts, "DelayedInboxForced", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return &BatcherDelayedInboxForcedIterator{contract: _Batcher.contract, event: "DelayedInboxForced", logs: logs, sub: sub}, nil
}

// WatchDelayedInboxForced is a free log subscription operation binding the contract event 0x85b6a949bf20bfd6bc6e20f98fb490c7944ab61dcfa5a30b5dae543412c9a8a0.
//
// Solidity: event DelayedInboxForced(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, uint256 totalDelayedMessagesRead, bytes32[2] afterAccAndDelayed, uint256 seqBatchIndex)
func (_Batcher *BatcherFilterer) WatchDelayedInboxForced(opts *bind.WatchOpts, sink chan<- *BatcherDelayedInboxForced, firstMessageNum []*big.Int, beforeAcc [][32]byte) (event.Subscription, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _Batcher.contract.WatchLogs(opts, "DelayedInboxForced", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatcherDelayedInboxForced)
				if err := _Batcher.contract.UnpackLog(event, "DelayedInboxForced", log); err != nil {
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

// ParseDelayedInboxForced is a log parse operation binding the contract event 0x85b6a949bf20bfd6bc6e20f98fb490c7944ab61dcfa5a30b5dae543412c9a8a0.
//
// Solidity: event DelayedInboxForced(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, uint256 totalDelayedMessagesRead, bytes32[2] afterAccAndDelayed, uint256 seqBatchIndex)
func (_Batcher *BatcherFilterer) ParseDelayedInboxForced(log types.Log) (*BatcherDelayedInboxForced, error) {
	event := new(BatcherDelayedInboxForced)
	if err := _Batcher.contract.UnpackLog(event, "DelayedInboxForced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatcherIsSequencerUpdatedIterator is returned from FilterIsSequencerUpdated and is used to iterate over the raw logs and unpacked data for IsSequencerUpdated events raised by the Batcher contract.
type BatcherIsSequencerUpdatedIterator struct {
	Event *BatcherIsSequencerUpdated // Event containing the contract specifics and raw log

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
func (it *BatcherIsSequencerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatcherIsSequencerUpdated)
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
		it.Event = new(BatcherIsSequencerUpdated)
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
func (it *BatcherIsSequencerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatcherIsSequencerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatcherIsSequencerUpdated represents a IsSequencerUpdated event raised by the Batcher contract.
type BatcherIsSequencerUpdated struct {
	Addr        common.Address
	IsSequencer bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterIsSequencerUpdated is a free log retrieval operation binding the contract event 0xce86e570206e55533301cb66529b33afbd75e991c575b85adeaca10146be8cb4.
//
// Solidity: event IsSequencerUpdated(address addr, bool isSequencer)
func (_Batcher *BatcherFilterer) FilterIsSequencerUpdated(opts *bind.FilterOpts) (*BatcherIsSequencerUpdatedIterator, error) {

	logs, sub, err := _Batcher.contract.FilterLogs(opts, "IsSequencerUpdated")
	if err != nil {
		return nil, err
	}
	return &BatcherIsSequencerUpdatedIterator{contract: _Batcher.contract, event: "IsSequencerUpdated", logs: logs, sub: sub}, nil
}

// WatchIsSequencerUpdated is a free log subscription operation binding the contract event 0xce86e570206e55533301cb66529b33afbd75e991c575b85adeaca10146be8cb4.
//
// Solidity: event IsSequencerUpdated(address addr, bool isSequencer)
func (_Batcher *BatcherFilterer) WatchIsSequencerUpdated(opts *bind.WatchOpts, sink chan<- *BatcherIsSequencerUpdated) (event.Subscription, error) {

	logs, sub, err := _Batcher.contract.WatchLogs(opts, "IsSequencerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatcherIsSequencerUpdated)
				if err := _Batcher.contract.UnpackLog(event, "IsSequencerUpdated", log); err != nil {
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

// ParseIsSequencerUpdated is a log parse operation binding the contract event 0xce86e570206e55533301cb66529b33afbd75e991c575b85adeaca10146be8cb4.
//
// Solidity: event IsSequencerUpdated(address addr, bool isSequencer)
func (_Batcher *BatcherFilterer) ParseIsSequencerUpdated(log types.Log) (*BatcherIsSequencerUpdated, error) {
	event := new(BatcherIsSequencerUpdated)
	if err := _Batcher.contract.UnpackLog(event, "IsSequencerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatcherMaxDelayUpdatedIterator is returned from FilterMaxDelayUpdated and is used to iterate over the raw logs and unpacked data for MaxDelayUpdated events raised by the Batcher contract.
type BatcherMaxDelayUpdatedIterator struct {
	Event *BatcherMaxDelayUpdated // Event containing the contract specifics and raw log

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
func (it *BatcherMaxDelayUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatcherMaxDelayUpdated)
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
		it.Event = new(BatcherMaxDelayUpdated)
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
func (it *BatcherMaxDelayUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatcherMaxDelayUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatcherMaxDelayUpdated represents a MaxDelayUpdated event raised by the Batcher contract.
type BatcherMaxDelayUpdated struct {
	NewMaxDelayBlocks  *big.Int
	NewMaxDelaySeconds *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMaxDelayUpdated is a free log retrieval operation binding the contract event 0x3bcd3c6d4304309e4b36d94f90517baf304582bb1ac828906808577e067e6b6e.
//
// Solidity: event MaxDelayUpdated(uint256 newMaxDelayBlocks, uint256 newMaxDelaySeconds)
func (_Batcher *BatcherFilterer) FilterMaxDelayUpdated(opts *bind.FilterOpts) (*BatcherMaxDelayUpdatedIterator, error) {

	logs, sub, err := _Batcher.contract.FilterLogs(opts, "MaxDelayUpdated")
	if err != nil {
		return nil, err
	}
	return &BatcherMaxDelayUpdatedIterator{contract: _Batcher.contract, event: "MaxDelayUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxDelayUpdated is a free log subscription operation binding the contract event 0x3bcd3c6d4304309e4b36d94f90517baf304582bb1ac828906808577e067e6b6e.
//
// Solidity: event MaxDelayUpdated(uint256 newMaxDelayBlocks, uint256 newMaxDelaySeconds)
func (_Batcher *BatcherFilterer) WatchMaxDelayUpdated(opts *bind.WatchOpts, sink chan<- *BatcherMaxDelayUpdated) (event.Subscription, error) {

	logs, sub, err := _Batcher.contract.WatchLogs(opts, "MaxDelayUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatcherMaxDelayUpdated)
				if err := _Batcher.contract.UnpackLog(event, "MaxDelayUpdated", log); err != nil {
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

// ParseMaxDelayUpdated is a log parse operation binding the contract event 0x3bcd3c6d4304309e4b36d94f90517baf304582bb1ac828906808577e067e6b6e.
//
// Solidity: event MaxDelayUpdated(uint256 newMaxDelayBlocks, uint256 newMaxDelaySeconds)
func (_Batcher *BatcherFilterer) ParseMaxDelayUpdated(log types.Log) (*BatcherMaxDelayUpdated, error) {
	event := new(BatcherMaxDelayUpdated)
	if err := _Batcher.contract.UnpackLog(event, "MaxDelayUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatcherSequencerBatchDeliveredIterator is returned from FilterSequencerBatchDelivered and is used to iterate over the raw logs and unpacked data for SequencerBatchDelivered events raised by the Batcher contract.
type BatcherSequencerBatchDeliveredIterator struct {
	Event *BatcherSequencerBatchDelivered // Event containing the contract specifics and raw log

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
func (it *BatcherSequencerBatchDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatcherSequencerBatchDelivered)
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
		it.Event = new(BatcherSequencerBatchDelivered)
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
func (it *BatcherSequencerBatchDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatcherSequencerBatchDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatcherSequencerBatchDelivered represents a SequencerBatchDelivered event raised by the Batcher contract.
type BatcherSequencerBatchDelivered struct {
	FirstMessageNum  *big.Int
	BeforeAcc        [32]byte
	NewMessageCount  *big.Int
	AfterAcc         [32]byte
	Transactions     []byte
	Lengths          []*big.Int
	SectionsMetadata []*big.Int
	SeqBatchIndex    *big.Int
	Sequencer        common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchDelivered is a free log retrieval operation binding the contract event 0x3bf85aebd2a1dc6c510ffc4795a3785e786b5817ab30144f88501d4c6456c986.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, uint256 seqBatchIndex, address sequencer)
func (_Batcher *BatcherFilterer) FilterSequencerBatchDelivered(opts *bind.FilterOpts, firstMessageNum []*big.Int, beforeAcc [][32]byte) (*BatcherSequencerBatchDeliveredIterator, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _Batcher.contract.FilterLogs(opts, "SequencerBatchDelivered", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return &BatcherSequencerBatchDeliveredIterator{contract: _Batcher.contract, event: "SequencerBatchDelivered", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchDelivered is a free log subscription operation binding the contract event 0x3bf85aebd2a1dc6c510ffc4795a3785e786b5817ab30144f88501d4c6456c986.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, uint256 seqBatchIndex, address sequencer)
func (_Batcher *BatcherFilterer) WatchSequencerBatchDelivered(opts *bind.WatchOpts, sink chan<- *BatcherSequencerBatchDelivered, firstMessageNum []*big.Int, beforeAcc [][32]byte) (event.Subscription, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _Batcher.contract.WatchLogs(opts, "SequencerBatchDelivered", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatcherSequencerBatchDelivered)
				if err := _Batcher.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
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

// ParseSequencerBatchDelivered is a log parse operation binding the contract event 0x3bf85aebd2a1dc6c510ffc4795a3785e786b5817ab30144f88501d4c6456c986.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, uint256 seqBatchIndex, address sequencer)
func (_Batcher *BatcherFilterer) ParseSequencerBatchDelivered(log types.Log) (*BatcherSequencerBatchDelivered, error) {
	event := new(BatcherSequencerBatchDelivered)
	if err := _Batcher.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatcherSequencerBatchDeliveredFromOriginIterator is returned from FilterSequencerBatchDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for SequencerBatchDeliveredFromOrigin events raised by the Batcher contract.
type BatcherSequencerBatchDeliveredFromOriginIterator struct {
	Event *BatcherSequencerBatchDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *BatcherSequencerBatchDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatcherSequencerBatchDeliveredFromOrigin)
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
		it.Event = new(BatcherSequencerBatchDeliveredFromOrigin)
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
func (it *BatcherSequencerBatchDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatcherSequencerBatchDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatcherSequencerBatchDeliveredFromOrigin represents a SequencerBatchDeliveredFromOrigin event raised by the Batcher contract.
type BatcherSequencerBatchDeliveredFromOrigin struct {
	FirstMessageNum *big.Int
	BeforeAcc       [32]byte
	NewMessageCount *big.Int
	AfterAcc        [32]byte
	SeqBatchIndex   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchDeliveredFromOrigin is a free log retrieval operation binding the contract event 0x10e0571aafaf282151fd5b0215b5495521c549509cb0de3a3f8310bd2e344682.
//
// Solidity: event SequencerBatchDeliveredFromOrigin(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, uint256 seqBatchIndex)
func (_Batcher *BatcherFilterer) FilterSequencerBatchDeliveredFromOrigin(opts *bind.FilterOpts, firstMessageNum []*big.Int, beforeAcc [][32]byte) (*BatcherSequencerBatchDeliveredFromOriginIterator, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _Batcher.contract.FilterLogs(opts, "SequencerBatchDeliveredFromOrigin", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return &BatcherSequencerBatchDeliveredFromOriginIterator{contract: _Batcher.contract, event: "SequencerBatchDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchDeliveredFromOrigin is a free log subscription operation binding the contract event 0x10e0571aafaf282151fd5b0215b5495521c549509cb0de3a3f8310bd2e344682.
//
// Solidity: event SequencerBatchDeliveredFromOrigin(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, uint256 seqBatchIndex)
func (_Batcher *BatcherFilterer) WatchSequencerBatchDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *BatcherSequencerBatchDeliveredFromOrigin, firstMessageNum []*big.Int, beforeAcc [][32]byte) (event.Subscription, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _Batcher.contract.WatchLogs(opts, "SequencerBatchDeliveredFromOrigin", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatcherSequencerBatchDeliveredFromOrigin)
				if err := _Batcher.contract.UnpackLog(event, "SequencerBatchDeliveredFromOrigin", log); err != nil {
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

// ParseSequencerBatchDeliveredFromOrigin is a log parse operation binding the contract event 0x10e0571aafaf282151fd5b0215b5495521c549509cb0de3a3f8310bd2e344682.
//
// Solidity: event SequencerBatchDeliveredFromOrigin(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, uint256 seqBatchIndex)
func (_Batcher *BatcherFilterer) ParseSequencerBatchDeliveredFromOrigin(log types.Log) (*BatcherSequencerBatchDeliveredFromOrigin, error) {
	event := new(BatcherSequencerBatchDeliveredFromOrigin)
	if err := _Batcher.contract.UnpackLog(event, "SequencerBatchDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

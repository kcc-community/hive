// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package punish

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

// PunishABI is the input ABI used to generate the binding from.
const PunishABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"LogDecreaseMissedBlocksCounter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogPunishValidator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_VALIDATORS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROPOSAL_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIProposal\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBLISH_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIPunish\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVEPOOL_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIReservePool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIValidators\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"decreaseMissedBlocksCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decreaseRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getPunishRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPunishValidatorsLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorsContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_punishContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proposalContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reservePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"}],\"name\":\"punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punishThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"punishValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Punish is an auto generated Go binding around an Ethereum contract.
type Punish struct {
	PunishCaller     // Read-only binding to the contract
	PunishTransactor // Write-only binding to the contract
	PunishFilterer   // Log filterer for contract events
}

// PunishCaller is an auto generated read-only Go binding around an Ethereum contract.
type PunishCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PunishTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PunishTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PunishFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PunishFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PunishSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PunishSession struct {
	Contract     *Punish           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PunishCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PunishCallerSession struct {
	Contract *PunishCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PunishTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PunishTransactorSession struct {
	Contract     *PunishTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PunishRaw is an auto generated low-level Go binding around an Ethereum contract.
type PunishRaw struct {
	Contract *Punish // Generic contract binding to access the raw methods on
}

// PunishCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PunishCallerRaw struct {
	Contract *PunishCaller // Generic read-only contract binding to access the raw methods on
}

// PunishTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PunishTransactorRaw struct {
	Contract *PunishTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPunish creates a new instance of Punish, bound to a specific deployed contract.
func NewPunish(address common.Address, backend bind.ContractBackend) (*Punish, error) {
	contract, err := bindPunish(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Punish{PunishCaller: PunishCaller{contract: contract}, PunishTransactor: PunishTransactor{contract: contract}, PunishFilterer: PunishFilterer{contract: contract}}, nil
}

// NewPunishCaller creates a new read-only instance of Punish, bound to a specific deployed contract.
func NewPunishCaller(address common.Address, caller bind.ContractCaller) (*PunishCaller, error) {
	contract, err := bindPunish(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PunishCaller{contract: contract}, nil
}

// NewPunishTransactor creates a new write-only instance of Punish, bound to a specific deployed contract.
func NewPunishTransactor(address common.Address, transactor bind.ContractTransactor) (*PunishTransactor, error) {
	contract, err := bindPunish(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PunishTransactor{contract: contract}, nil
}

// NewPunishFilterer creates a new log filterer instance of Punish, bound to a specific deployed contract.
func NewPunishFilterer(address common.Address, filterer bind.ContractFilterer) (*PunishFilterer, error) {
	contract, err := bindPunish(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PunishFilterer{contract: contract}, nil
}

// bindPunish binds a generic wrapper to an already deployed contract.
func bindPunish(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PunishABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Punish *PunishRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Punish.Contract.PunishCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Punish *PunishRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Punish.Contract.PunishTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Punish *PunishRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Punish.Contract.PunishTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Punish *PunishCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Punish.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Punish *PunishTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Punish.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Punish *PunishTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Punish.Contract.contract.Transact(opts, method, params...)
}

// MAXVALIDATORS is a free data retrieval call binding the contract method 0x714897df.
//
// Solidity: function MAX_VALIDATORS() view returns(uint16)
func (_Punish *PunishCaller) MAXVALIDATORS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "MAX_VALIDATORS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXVALIDATORS is a free data retrieval call binding the contract method 0x714897df.
//
// Solidity: function MAX_VALIDATORS() view returns(uint16)
func (_Punish *PunishSession) MAXVALIDATORS() (uint16, error) {
	return _Punish.Contract.MAXVALIDATORS(&_Punish.CallOpts)
}

// MAXVALIDATORS is a free data retrieval call binding the contract method 0x714897df.
//
// Solidity: function MAX_VALIDATORS() view returns(uint16)
func (_Punish *PunishCallerSession) MAXVALIDATORS() (uint16, error) {
	return _Punish.Contract.MAXVALIDATORS(&_Punish.CallOpts)
}

// PROPOSALCONTRACT is a free data retrieval call binding the contract method 0x863a20b7.
//
// Solidity: function PROPOSAL_CONTRACT() view returns(address)
func (_Punish *PunishCaller) PROPOSALCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "PROPOSAL_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROPOSALCONTRACT is a free data retrieval call binding the contract method 0x863a20b7.
//
// Solidity: function PROPOSAL_CONTRACT() view returns(address)
func (_Punish *PunishSession) PROPOSALCONTRACT() (common.Address, error) {
	return _Punish.Contract.PROPOSALCONTRACT(&_Punish.CallOpts)
}

// PROPOSALCONTRACT is a free data retrieval call binding the contract method 0x863a20b7.
//
// Solidity: function PROPOSAL_CONTRACT() view returns(address)
func (_Punish *PunishCallerSession) PROPOSALCONTRACT() (common.Address, error) {
	return _Punish.Contract.PROPOSALCONTRACT(&_Punish.CallOpts)
}

// PUBLISHCONTRACT is a free data retrieval call binding the contract method 0xe5a99f4f.
//
// Solidity: function PUBLISH_CONTRACT() view returns(address)
func (_Punish *PunishCaller) PUBLISHCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "PUBLISH_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PUBLISHCONTRACT is a free data retrieval call binding the contract method 0xe5a99f4f.
//
// Solidity: function PUBLISH_CONTRACT() view returns(address)
func (_Punish *PunishSession) PUBLISHCONTRACT() (common.Address, error) {
	return _Punish.Contract.PUBLISHCONTRACT(&_Punish.CallOpts)
}

// PUBLISHCONTRACT is a free data retrieval call binding the contract method 0xe5a99f4f.
//
// Solidity: function PUBLISH_CONTRACT() view returns(address)
func (_Punish *PunishCallerSession) PUBLISHCONTRACT() (common.Address, error) {
	return _Punish.Contract.PUBLISHCONTRACT(&_Punish.CallOpts)
}

// RESERVEPOOLCONTRACT is a free data retrieval call binding the contract method 0x60c80cbf.
//
// Solidity: function RESERVEPOOL_CONTRACT() view returns(address)
func (_Punish *PunishCaller) RESERVEPOOLCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "RESERVEPOOL_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RESERVEPOOLCONTRACT is a free data retrieval call binding the contract method 0x60c80cbf.
//
// Solidity: function RESERVEPOOL_CONTRACT() view returns(address)
func (_Punish *PunishSession) RESERVEPOOLCONTRACT() (common.Address, error) {
	return _Punish.Contract.RESERVEPOOLCONTRACT(&_Punish.CallOpts)
}

// RESERVEPOOLCONTRACT is a free data retrieval call binding the contract method 0x60c80cbf.
//
// Solidity: function RESERVEPOOL_CONTRACT() view returns(address)
func (_Punish *PunishCallerSession) RESERVEPOOLCONTRACT() (common.Address, error) {
	return _Punish.Contract.RESERVEPOOLCONTRACT(&_Punish.CallOpts)
}

// VALIDATORCONTRACT is a free data retrieval call binding the contract method 0x46f75138.
//
// Solidity: function VALIDATOR_CONTRACT() view returns(address)
func (_Punish *PunishCaller) VALIDATORCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "VALIDATOR_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORCONTRACT is a free data retrieval call binding the contract method 0x46f75138.
//
// Solidity: function VALIDATOR_CONTRACT() view returns(address)
func (_Punish *PunishSession) VALIDATORCONTRACT() (common.Address, error) {
	return _Punish.Contract.VALIDATORCONTRACT(&_Punish.CallOpts)
}

// VALIDATORCONTRACT is a free data retrieval call binding the contract method 0x46f75138.
//
// Solidity: function VALIDATOR_CONTRACT() view returns(address)
func (_Punish *PunishCallerSession) VALIDATORCONTRACT() (common.Address, error) {
	return _Punish.Contract.VALIDATORCONTRACT(&_Punish.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Punish *PunishCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Punish *PunishSession) Admin() (common.Address, error) {
	return _Punish.Contract.Admin(&_Punish.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Punish *PunishCallerSession) Admin() (common.Address, error) {
	return _Punish.Contract.Admin(&_Punish.CallOpts)
}

// DecreaseRate is a free data retrieval call binding the contract method 0x2897183d.
//
// Solidity: function decreaseRate() view returns(uint256)
func (_Punish *PunishCaller) DecreaseRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "decreaseRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecreaseRate is a free data retrieval call binding the contract method 0x2897183d.
//
// Solidity: function decreaseRate() view returns(uint256)
func (_Punish *PunishSession) DecreaseRate() (*big.Int, error) {
	return _Punish.Contract.DecreaseRate(&_Punish.CallOpts)
}

// DecreaseRate is a free data retrieval call binding the contract method 0x2897183d.
//
// Solidity: function decreaseRate() view returns(uint256)
func (_Punish *PunishCallerSession) DecreaseRate() (*big.Int, error) {
	return _Punish.Contract.DecreaseRate(&_Punish.CallOpts)
}

// GetPunishRecord is a free data retrieval call binding the contract method 0x32f3c17f.
//
// Solidity: function getPunishRecord(address val) view returns(uint256)
func (_Punish *PunishCaller) GetPunishRecord(opts *bind.CallOpts, val common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "getPunishRecord", val)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPunishRecord is a free data retrieval call binding the contract method 0x32f3c17f.
//
// Solidity: function getPunishRecord(address val) view returns(uint256)
func (_Punish *PunishSession) GetPunishRecord(val common.Address) (*big.Int, error) {
	return _Punish.Contract.GetPunishRecord(&_Punish.CallOpts, val)
}

// GetPunishRecord is a free data retrieval call binding the contract method 0x32f3c17f.
//
// Solidity: function getPunishRecord(address val) view returns(uint256)
func (_Punish *PunishCallerSession) GetPunishRecord(val common.Address) (*big.Int, error) {
	return _Punish.Contract.GetPunishRecord(&_Punish.CallOpts, val)
}

// GetPunishValidatorsLen is a free data retrieval call binding the contract method 0xe0d8ea53.
//
// Solidity: function getPunishValidatorsLen() view returns(uint256)
func (_Punish *PunishCaller) GetPunishValidatorsLen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "getPunishValidatorsLen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPunishValidatorsLen is a free data retrieval call binding the contract method 0xe0d8ea53.
//
// Solidity: function getPunishValidatorsLen() view returns(uint256)
func (_Punish *PunishSession) GetPunishValidatorsLen() (*big.Int, error) {
	return _Punish.Contract.GetPunishValidatorsLen(&_Punish.CallOpts)
}

// GetPunishValidatorsLen is a free data retrieval call binding the contract method 0xe0d8ea53.
//
// Solidity: function getPunishValidatorsLen() view returns(uint256)
func (_Punish *PunishCallerSession) GetPunishValidatorsLen() (*big.Int, error) {
	return _Punish.Contract.GetPunishValidatorsLen(&_Punish.CallOpts)
}

// PunishThreshold is a free data retrieval call binding the contract method 0xcb1ea725.
//
// Solidity: function punishThreshold() view returns(uint256)
func (_Punish *PunishCaller) PunishThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "punishThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PunishThreshold is a free data retrieval call binding the contract method 0xcb1ea725.
//
// Solidity: function punishThreshold() view returns(uint256)
func (_Punish *PunishSession) PunishThreshold() (*big.Int, error) {
	return _Punish.Contract.PunishThreshold(&_Punish.CallOpts)
}

// PunishThreshold is a free data retrieval call binding the contract method 0xcb1ea725.
//
// Solidity: function punishThreshold() view returns(uint256)
func (_Punish *PunishCallerSession) PunishThreshold() (*big.Int, error) {
	return _Punish.Contract.PunishThreshold(&_Punish.CallOpts)
}

// PunishValidators is a free data retrieval call binding the contract method 0xf62af26c.
//
// Solidity: function punishValidators(uint256 ) view returns(address)
func (_Punish *PunishCaller) PunishValidators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "punishValidators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PunishValidators is a free data retrieval call binding the contract method 0xf62af26c.
//
// Solidity: function punishValidators(uint256 ) view returns(address)
func (_Punish *PunishSession) PunishValidators(arg0 *big.Int) (common.Address, error) {
	return _Punish.Contract.PunishValidators(&_Punish.CallOpts, arg0)
}

// PunishValidators is a free data retrieval call binding the contract method 0xf62af26c.
//
// Solidity: function punishValidators(uint256 ) view returns(address)
func (_Punish *PunishCallerSession) PunishValidators(arg0 *big.Int) (common.Address, error) {
	return _Punish.Contract.PunishValidators(&_Punish.CallOpts, arg0)
}

// RemoveThreshold is a free data retrieval call binding the contract method 0x44c1aa99.
//
// Solidity: function removeThreshold() view returns(uint256)
func (_Punish *PunishCaller) RemoveThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "removeThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemoveThreshold is a free data retrieval call binding the contract method 0x44c1aa99.
//
// Solidity: function removeThreshold() view returns(uint256)
func (_Punish *PunishSession) RemoveThreshold() (*big.Int, error) {
	return _Punish.Contract.RemoveThreshold(&_Punish.CallOpts)
}

// RemoveThreshold is a free data retrieval call binding the contract method 0x44c1aa99.
//
// Solidity: function removeThreshold() view returns(uint256)
func (_Punish *PunishCallerSession) RemoveThreshold() (*big.Int, error) {
	return _Punish.Contract.RemoveThreshold(&_Punish.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_Punish *PunishTransactor) ChangeAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Punish.contract.Transact(opts, "changeAdmin", _admin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_Punish *PunishSession) ChangeAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Punish.Contract.ChangeAdmin(&_Punish.TransactOpts, _admin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_Punish *PunishTransactorSession) ChangeAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Punish.Contract.ChangeAdmin(&_Punish.TransactOpts, _admin)
}

// DecreaseMissedBlocksCounter is a paid mutator transaction binding the contract method 0xd93d2cb9.
//
// Solidity: function decreaseMissedBlocksCounter(uint256 _epoch) returns()
func (_Punish *PunishTransactor) DecreaseMissedBlocksCounter(opts *bind.TransactOpts, _epoch *big.Int) (*types.Transaction, error) {
	return _Punish.contract.Transact(opts, "decreaseMissedBlocksCounter", _epoch)
}

// DecreaseMissedBlocksCounter is a paid mutator transaction binding the contract method 0xd93d2cb9.
//
// Solidity: function decreaseMissedBlocksCounter(uint256 _epoch) returns()
func (_Punish *PunishSession) DecreaseMissedBlocksCounter(_epoch *big.Int) (*types.Transaction, error) {
	return _Punish.Contract.DecreaseMissedBlocksCounter(&_Punish.TransactOpts, _epoch)
}

// DecreaseMissedBlocksCounter is a paid mutator transaction binding the contract method 0xd93d2cb9.
//
// Solidity: function decreaseMissedBlocksCounter(uint256 _epoch) returns()
func (_Punish *PunishTransactorSession) DecreaseMissedBlocksCounter(_epoch *big.Int) (*types.Transaction, error) {
	return _Punish.Contract.DecreaseMissedBlocksCounter(&_Punish.TransactOpts, _epoch)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _validatorsContract, address _punishContract, address _proposalContract, address _reservePool, address _admin) returns()
func (_Punish *PunishTransactor) Initialize(opts *bind.TransactOpts, _validatorsContract common.Address, _punishContract common.Address, _proposalContract common.Address, _reservePool common.Address, _admin common.Address) (*types.Transaction, error) {
	return _Punish.contract.Transact(opts, "initialize", _validatorsContract, _punishContract, _proposalContract, _reservePool, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _validatorsContract, address _punishContract, address _proposalContract, address _reservePool, address _admin) returns()
func (_Punish *PunishSession) Initialize(_validatorsContract common.Address, _punishContract common.Address, _proposalContract common.Address, _reservePool common.Address, _admin common.Address) (*types.Transaction, error) {
	return _Punish.Contract.Initialize(&_Punish.TransactOpts, _validatorsContract, _punishContract, _proposalContract, _reservePool, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _validatorsContract, address _punishContract, address _proposalContract, address _reservePool, address _admin) returns()
func (_Punish *PunishTransactorSession) Initialize(_validatorsContract common.Address, _punishContract common.Address, _proposalContract common.Address, _reservePool common.Address, _admin common.Address) (*types.Transaction, error) {
	return _Punish.Contract.Initialize(&_Punish.TransactOpts, _validatorsContract, _punishContract, _proposalContract, _reservePool, _admin)
}

// Punish is a paid mutator transaction binding the contract method 0xea7221a1.
//
// Solidity: function punish(address _val) returns()
func (_Punish *PunishTransactor) Punish(opts *bind.TransactOpts, _val common.Address) (*types.Transaction, error) {
	return _Punish.contract.Transact(opts, "punish", _val)
}

// Punish is a paid mutator transaction binding the contract method 0xea7221a1.
//
// Solidity: function punish(address _val) returns()
func (_Punish *PunishSession) Punish(_val common.Address) (*types.Transaction, error) {
	return _Punish.Contract.Punish(&_Punish.TransactOpts, _val)
}

// Punish is a paid mutator transaction binding the contract method 0xea7221a1.
//
// Solidity: function punish(address _val) returns()
func (_Punish *PunishTransactorSession) Punish(_val common.Address) (*types.Transaction, error) {
	return _Punish.Contract.Punish(&_Punish.TransactOpts, _val)
}

// PunishLogDecreaseMissedBlocksCounterIterator is returned from FilterLogDecreaseMissedBlocksCounter and is used to iterate over the raw logs and unpacked data for LogDecreaseMissedBlocksCounter events raised by the Punish contract.
type PunishLogDecreaseMissedBlocksCounterIterator struct {
	Event *PunishLogDecreaseMissedBlocksCounter // Event containing the contract specifics and raw log

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
func (it *PunishLogDecreaseMissedBlocksCounterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PunishLogDecreaseMissedBlocksCounter)
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
		it.Event = new(PunishLogDecreaseMissedBlocksCounter)
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
func (it *PunishLogDecreaseMissedBlocksCounterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PunishLogDecreaseMissedBlocksCounterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PunishLogDecreaseMissedBlocksCounter represents a LogDecreaseMissedBlocksCounter event raised by the Punish contract.
type PunishLogDecreaseMissedBlocksCounter struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogDecreaseMissedBlocksCounter is a free log retrieval operation binding the contract event 0x181d51be54e8e8eaca6eae0eab32d4162099236bd519e7238d015d0870db4641.
//
// Solidity: event LogDecreaseMissedBlocksCounter()
func (_Punish *PunishFilterer) FilterLogDecreaseMissedBlocksCounter(opts *bind.FilterOpts) (*PunishLogDecreaseMissedBlocksCounterIterator, error) {

	logs, sub, err := _Punish.contract.FilterLogs(opts, "LogDecreaseMissedBlocksCounter")
	if err != nil {
		return nil, err
	}
	return &PunishLogDecreaseMissedBlocksCounterIterator{contract: _Punish.contract, event: "LogDecreaseMissedBlocksCounter", logs: logs, sub: sub}, nil
}

// WatchLogDecreaseMissedBlocksCounter is a free log subscription operation binding the contract event 0x181d51be54e8e8eaca6eae0eab32d4162099236bd519e7238d015d0870db4641.
//
// Solidity: event LogDecreaseMissedBlocksCounter()
func (_Punish *PunishFilterer) WatchLogDecreaseMissedBlocksCounter(opts *bind.WatchOpts, sink chan<- *PunishLogDecreaseMissedBlocksCounter) (event.Subscription, error) {

	logs, sub, err := _Punish.contract.WatchLogs(opts, "LogDecreaseMissedBlocksCounter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PunishLogDecreaseMissedBlocksCounter)
				if err := _Punish.contract.UnpackLog(event, "LogDecreaseMissedBlocksCounter", log); err != nil {
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

// ParseLogDecreaseMissedBlocksCounter is a log parse operation binding the contract event 0x181d51be54e8e8eaca6eae0eab32d4162099236bd519e7238d015d0870db4641.
//
// Solidity: event LogDecreaseMissedBlocksCounter()
func (_Punish *PunishFilterer) ParseLogDecreaseMissedBlocksCounter(log types.Log) (*PunishLogDecreaseMissedBlocksCounter, error) {
	event := new(PunishLogDecreaseMissedBlocksCounter)
	if err := _Punish.contract.UnpackLog(event, "LogDecreaseMissedBlocksCounter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PunishLogPunishValidatorIterator is returned from FilterLogPunishValidator and is used to iterate over the raw logs and unpacked data for LogPunishValidator events raised by the Punish contract.
type PunishLogPunishValidatorIterator struct {
	Event *PunishLogPunishValidator // Event containing the contract specifics and raw log

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
func (it *PunishLogPunishValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PunishLogPunishValidator)
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
		it.Event = new(PunishLogPunishValidator)
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
func (it *PunishLogPunishValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PunishLogPunishValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PunishLogPunishValidator represents a LogPunishValidator event raised by the Punish contract.
type PunishLogPunishValidator struct {
	Val  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogPunishValidator is a free log retrieval operation binding the contract event 0x770e0cca42c35d00240986ce8d3ed438be04663c91dac6576b79537d7c180f1e.
//
// Solidity: event LogPunishValidator(address indexed val, uint256 time)
func (_Punish *PunishFilterer) FilterLogPunishValidator(opts *bind.FilterOpts, val []common.Address) (*PunishLogPunishValidatorIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Punish.contract.FilterLogs(opts, "LogPunishValidator", valRule)
	if err != nil {
		return nil, err
	}
	return &PunishLogPunishValidatorIterator{contract: _Punish.contract, event: "LogPunishValidator", logs: logs, sub: sub}, nil
}

// WatchLogPunishValidator is a free log subscription operation binding the contract event 0x770e0cca42c35d00240986ce8d3ed438be04663c91dac6576b79537d7c180f1e.
//
// Solidity: event LogPunishValidator(address indexed val, uint256 time)
func (_Punish *PunishFilterer) WatchLogPunishValidator(opts *bind.WatchOpts, sink chan<- *PunishLogPunishValidator, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Punish.contract.WatchLogs(opts, "LogPunishValidator", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PunishLogPunishValidator)
				if err := _Punish.contract.UnpackLog(event, "LogPunishValidator", log); err != nil {
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

// ParseLogPunishValidator is a log parse operation binding the contract event 0x770e0cca42c35d00240986ce8d3ed438be04663c91dac6576b79537d7c180f1e.
//
// Solidity: event LogPunishValidator(address indexed val, uint256 time)
func (_Punish *PunishFilterer) ParseLogPunishValidator(log types.Log) (*PunishLogPunishValidator, error) {
	event := new(PunishLogPunishValidator)
	if err := _Punish.contract.UnpackLog(event, "LogPunishValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reservepool

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

// ReservepoolABI is the input ABI used to generate the binding from.
const ReservepoolABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BLOCK_REWARD_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_VALIDATORS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROPOSAL_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIProposal\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBLISH_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIPunish\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVEPOOL_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIReservePool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIValidators\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorsContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_punishContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proposalContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reservePool\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setBlockRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumReservePool.State\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"setState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumReservePool.State\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawBlockReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Reservepool is an auto generated Go binding around an Ethereum contract.
type Reservepool struct {
	ReservepoolCaller     // Read-only binding to the contract
	ReservepoolTransactor // Write-only binding to the contract
	ReservepoolFilterer   // Log filterer for contract events
}

// ReservepoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReservepoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReservepoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReservepoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReservepoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReservepoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReservepoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReservepoolSession struct {
	Contract     *Reservepool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReservepoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReservepoolCallerSession struct {
	Contract *ReservepoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ReservepoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReservepoolTransactorSession struct {
	Contract     *ReservepoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ReservepoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReservepoolRaw struct {
	Contract *Reservepool // Generic contract binding to access the raw methods on
}

// ReservepoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReservepoolCallerRaw struct {
	Contract *ReservepoolCaller // Generic read-only contract binding to access the raw methods on
}

// ReservepoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReservepoolTransactorRaw struct {
	Contract *ReservepoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReservepool creates a new instance of Reservepool, bound to a specific deployed contract.
func NewReservepool(address common.Address, backend bind.ContractBackend) (*Reservepool, error) {
	contract, err := bindReservepool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reservepool{ReservepoolCaller: ReservepoolCaller{contract: contract}, ReservepoolTransactor: ReservepoolTransactor{contract: contract}, ReservepoolFilterer: ReservepoolFilterer{contract: contract}}, nil
}

// NewReservepoolCaller creates a new read-only instance of Reservepool, bound to a specific deployed contract.
func NewReservepoolCaller(address common.Address, caller bind.ContractCaller) (*ReservepoolCaller, error) {
	contract, err := bindReservepool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReservepoolCaller{contract: contract}, nil
}

// NewReservepoolTransactor creates a new write-only instance of Reservepool, bound to a specific deployed contract.
func NewReservepoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ReservepoolTransactor, error) {
	contract, err := bindReservepool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReservepoolTransactor{contract: contract}, nil
}

// NewReservepoolFilterer creates a new log filterer instance of Reservepool, bound to a specific deployed contract.
func NewReservepoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ReservepoolFilterer, error) {
	contract, err := bindReservepool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReservepoolFilterer{contract: contract}, nil
}

// bindReservepool binds a generic wrapper to an already deployed contract.
func bindReservepool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReservepoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reservepool *ReservepoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reservepool.Contract.ReservepoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reservepool *ReservepoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reservepool.Contract.ReservepoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reservepool *ReservepoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reservepool.Contract.ReservepoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reservepool *ReservepoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reservepool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reservepool *ReservepoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reservepool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reservepool *ReservepoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reservepool.Contract.contract.Transact(opts, method, params...)
}

// MAXBLOCKREWARDAMOUNT is a free data retrieval call binding the contract method 0x828090d0.
//
// Solidity: function MAX_BLOCK_REWARD_AMOUNT() view returns(uint256)
func (_Reservepool *ReservepoolCaller) MAXBLOCKREWARDAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reservepool.contract.Call(opts, &out, "MAX_BLOCK_REWARD_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBLOCKREWARDAMOUNT is a free data retrieval call binding the contract method 0x828090d0.
//
// Solidity: function MAX_BLOCK_REWARD_AMOUNT() view returns(uint256)
func (_Reservepool *ReservepoolSession) MAXBLOCKREWARDAMOUNT() (*big.Int, error) {
	return _Reservepool.Contract.MAXBLOCKREWARDAMOUNT(&_Reservepool.CallOpts)
}

// MAXBLOCKREWARDAMOUNT is a free data retrieval call binding the contract method 0x828090d0.
//
// Solidity: function MAX_BLOCK_REWARD_AMOUNT() view returns(uint256)
func (_Reservepool *ReservepoolCallerSession) MAXBLOCKREWARDAMOUNT() (*big.Int, error) {
	return _Reservepool.Contract.MAXBLOCKREWARDAMOUNT(&_Reservepool.CallOpts)
}

// MAXVALIDATORS is a free data retrieval call binding the contract method 0x714897df.
//
// Solidity: function MAX_VALIDATORS() view returns(uint16)
func (_Reservepool *ReservepoolCaller) MAXVALIDATORS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Reservepool.contract.Call(opts, &out, "MAX_VALIDATORS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXVALIDATORS is a free data retrieval call binding the contract method 0x714897df.
//
// Solidity: function MAX_VALIDATORS() view returns(uint16)
func (_Reservepool *ReservepoolSession) MAXVALIDATORS() (uint16, error) {
	return _Reservepool.Contract.MAXVALIDATORS(&_Reservepool.CallOpts)
}

// MAXVALIDATORS is a free data retrieval call binding the contract method 0x714897df.
//
// Solidity: function MAX_VALIDATORS() view returns(uint16)
func (_Reservepool *ReservepoolCallerSession) MAXVALIDATORS() (uint16, error) {
	return _Reservepool.Contract.MAXVALIDATORS(&_Reservepool.CallOpts)
}

// PROPOSALCONTRACT is a free data retrieval call binding the contract method 0x863a20b7.
//
// Solidity: function PROPOSAL_CONTRACT() view returns(address)
func (_Reservepool *ReservepoolCaller) PROPOSALCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reservepool.contract.Call(opts, &out, "PROPOSAL_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROPOSALCONTRACT is a free data retrieval call binding the contract method 0x863a20b7.
//
// Solidity: function PROPOSAL_CONTRACT() view returns(address)
func (_Reservepool *ReservepoolSession) PROPOSALCONTRACT() (common.Address, error) {
	return _Reservepool.Contract.PROPOSALCONTRACT(&_Reservepool.CallOpts)
}

// PROPOSALCONTRACT is a free data retrieval call binding the contract method 0x863a20b7.
//
// Solidity: function PROPOSAL_CONTRACT() view returns(address)
func (_Reservepool *ReservepoolCallerSession) PROPOSALCONTRACT() (common.Address, error) {
	return _Reservepool.Contract.PROPOSALCONTRACT(&_Reservepool.CallOpts)
}

// PUBLISHCONTRACT is a free data retrieval call binding the contract method 0xe5a99f4f.
//
// Solidity: function PUBLISH_CONTRACT() view returns(address)
func (_Reservepool *ReservepoolCaller) PUBLISHCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reservepool.contract.Call(opts, &out, "PUBLISH_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PUBLISHCONTRACT is a free data retrieval call binding the contract method 0xe5a99f4f.
//
// Solidity: function PUBLISH_CONTRACT() view returns(address)
func (_Reservepool *ReservepoolSession) PUBLISHCONTRACT() (common.Address, error) {
	return _Reservepool.Contract.PUBLISHCONTRACT(&_Reservepool.CallOpts)
}

// PUBLISHCONTRACT is a free data retrieval call binding the contract method 0xe5a99f4f.
//
// Solidity: function PUBLISH_CONTRACT() view returns(address)
func (_Reservepool *ReservepoolCallerSession) PUBLISHCONTRACT() (common.Address, error) {
	return _Reservepool.Contract.PUBLISHCONTRACT(&_Reservepool.CallOpts)
}

// RESERVEPOOLCONTRACT is a free data retrieval call binding the contract method 0x60c80cbf.
//
// Solidity: function RESERVEPOOL_CONTRACT() view returns(address)
func (_Reservepool *ReservepoolCaller) RESERVEPOOLCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reservepool.contract.Call(opts, &out, "RESERVEPOOL_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RESERVEPOOLCONTRACT is a free data retrieval call binding the contract method 0x60c80cbf.
//
// Solidity: function RESERVEPOOL_CONTRACT() view returns(address)
func (_Reservepool *ReservepoolSession) RESERVEPOOLCONTRACT() (common.Address, error) {
	return _Reservepool.Contract.RESERVEPOOLCONTRACT(&_Reservepool.CallOpts)
}

// RESERVEPOOLCONTRACT is a free data retrieval call binding the contract method 0x60c80cbf.
//
// Solidity: function RESERVEPOOL_CONTRACT() view returns(address)
func (_Reservepool *ReservepoolCallerSession) RESERVEPOOLCONTRACT() (common.Address, error) {
	return _Reservepool.Contract.RESERVEPOOLCONTRACT(&_Reservepool.CallOpts)
}

// VALIDATORCONTRACT is a free data retrieval call binding the contract method 0x46f75138.
//
// Solidity: function VALIDATOR_CONTRACT() view returns(address)
func (_Reservepool *ReservepoolCaller) VALIDATORCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reservepool.contract.Call(opts, &out, "VALIDATOR_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORCONTRACT is a free data retrieval call binding the contract method 0x46f75138.
//
// Solidity: function VALIDATOR_CONTRACT() view returns(address)
func (_Reservepool *ReservepoolSession) VALIDATORCONTRACT() (common.Address, error) {
	return _Reservepool.Contract.VALIDATORCONTRACT(&_Reservepool.CallOpts)
}

// VALIDATORCONTRACT is a free data retrieval call binding the contract method 0x46f75138.
//
// Solidity: function VALIDATOR_CONTRACT() view returns(address)
func (_Reservepool *ReservepoolCallerSession) VALIDATORCONTRACT() (common.Address, error) {
	return _Reservepool.Contract.VALIDATORCONTRACT(&_Reservepool.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Reservepool *ReservepoolCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reservepool.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Reservepool *ReservepoolSession) Admin() (common.Address, error) {
	return _Reservepool.Contract.Admin(&_Reservepool.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Reservepool *ReservepoolCallerSession) Admin() (common.Address, error) {
	return _Reservepool.Contract.Admin(&_Reservepool.CallOpts)
}

// BlockRewardAmount is a free data retrieval call binding the contract method 0x26cc2256.
//
// Solidity: function blockRewardAmount() view returns(uint256)
func (_Reservepool *ReservepoolCaller) BlockRewardAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reservepool.contract.Call(opts, &out, "blockRewardAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockRewardAmount is a free data retrieval call binding the contract method 0x26cc2256.
//
// Solidity: function blockRewardAmount() view returns(uint256)
func (_Reservepool *ReservepoolSession) BlockRewardAmount() (*big.Int, error) {
	return _Reservepool.Contract.BlockRewardAmount(&_Reservepool.CallOpts)
}

// BlockRewardAmount is a free data retrieval call binding the contract method 0x26cc2256.
//
// Solidity: function blockRewardAmount() view returns(uint256)
func (_Reservepool *ReservepoolCallerSession) BlockRewardAmount() (*big.Int, error) {
	return _Reservepool.Contract.BlockRewardAmount(&_Reservepool.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Reservepool *ReservepoolCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Reservepool.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Reservepool *ReservepoolSession) State() (uint8, error) {
	return _Reservepool.Contract.State(&_Reservepool.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Reservepool *ReservepoolCallerSession) State() (uint8, error) {
	return _Reservepool.Contract.State(&_Reservepool.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_Reservepool *ReservepoolTransactor) ChangeAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Reservepool.contract.Transact(opts, "changeAdmin", _admin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_Reservepool *ReservepoolSession) ChangeAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Reservepool.Contract.ChangeAdmin(&_Reservepool.TransactOpts, _admin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_Reservepool *ReservepoolTransactorSession) ChangeAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Reservepool.Contract.ChangeAdmin(&_Reservepool.TransactOpts, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _admin, address _validatorsContract, address _punishContract, address _proposalContract, address _reservePool) returns()
func (_Reservepool *ReservepoolTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _validatorsContract common.Address, _punishContract common.Address, _proposalContract common.Address, _reservePool common.Address) (*types.Transaction, error) {
	return _Reservepool.contract.Transact(opts, "initialize", _admin, _validatorsContract, _punishContract, _proposalContract, _reservePool)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _admin, address _validatorsContract, address _punishContract, address _proposalContract, address _reservePool) returns()
func (_Reservepool *ReservepoolSession) Initialize(_admin common.Address, _validatorsContract common.Address, _punishContract common.Address, _proposalContract common.Address, _reservePool common.Address) (*types.Transaction, error) {
	return _Reservepool.Contract.Initialize(&_Reservepool.TransactOpts, _admin, _validatorsContract, _punishContract, _proposalContract, _reservePool)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _admin, address _validatorsContract, address _punishContract, address _proposalContract, address _reservePool) returns()
func (_Reservepool *ReservepoolTransactorSession) Initialize(_admin common.Address, _validatorsContract common.Address, _punishContract common.Address, _proposalContract common.Address, _reservePool common.Address) (*types.Transaction, error) {
	return _Reservepool.Contract.Initialize(&_Reservepool.TransactOpts, _admin, _validatorsContract, _punishContract, _proposalContract, _reservePool)
}

// SetBlockRewardAmount is a paid mutator transaction binding the contract method 0x2ed19cd5.
//
// Solidity: function setBlockRewardAmount(uint256 amount) returns()
func (_Reservepool *ReservepoolTransactor) SetBlockRewardAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Reservepool.contract.Transact(opts, "setBlockRewardAmount", amount)
}

// SetBlockRewardAmount is a paid mutator transaction binding the contract method 0x2ed19cd5.
//
// Solidity: function setBlockRewardAmount(uint256 amount) returns()
func (_Reservepool *ReservepoolSession) SetBlockRewardAmount(amount *big.Int) (*types.Transaction, error) {
	return _Reservepool.Contract.SetBlockRewardAmount(&_Reservepool.TransactOpts, amount)
}

// SetBlockRewardAmount is a paid mutator transaction binding the contract method 0x2ed19cd5.
//
// Solidity: function setBlockRewardAmount(uint256 amount) returns()
func (_Reservepool *ReservepoolTransactorSession) SetBlockRewardAmount(amount *big.Int) (*types.Transaction, error) {
	return _Reservepool.Contract.SetBlockRewardAmount(&_Reservepool.TransactOpts, amount)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_Reservepool *ReservepoolTransactor) SetState(opts *bind.TransactOpts, newState uint8) (*types.Transaction, error) {
	return _Reservepool.contract.Transact(opts, "setState", newState)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_Reservepool *ReservepoolSession) SetState(newState uint8) (*types.Transaction, error) {
	return _Reservepool.Contract.SetState(&_Reservepool.TransactOpts, newState)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_Reservepool *ReservepoolTransactorSession) SetState(newState uint8) (*types.Transaction, error) {
	return _Reservepool.Contract.SetState(&_Reservepool.TransactOpts, newState)
}

// WithdrawBlockReward is a paid mutator transaction binding the contract method 0xcf813c1d.
//
// Solidity: function withdrawBlockReward() returns(uint256)
func (_Reservepool *ReservepoolTransactor) WithdrawBlockReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reservepool.contract.Transact(opts, "withdrawBlockReward")
}

// WithdrawBlockReward is a paid mutator transaction binding the contract method 0xcf813c1d.
//
// Solidity: function withdrawBlockReward() returns(uint256)
func (_Reservepool *ReservepoolSession) WithdrawBlockReward() (*types.Transaction, error) {
	return _Reservepool.Contract.WithdrawBlockReward(&_Reservepool.TransactOpts)
}

// WithdrawBlockReward is a paid mutator transaction binding the contract method 0xcf813c1d.
//
// Solidity: function withdrawBlockReward() returns(uint256)
func (_Reservepool *ReservepoolTransactorSession) WithdrawBlockReward() (*types.Transaction, error) {
	return _Reservepool.Contract.WithdrawBlockReward(&_Reservepool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Reservepool *ReservepoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reservepool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Reservepool *ReservepoolSession) Receive() (*types.Transaction, error) {
	return _Reservepool.Contract.Receive(&_Reservepool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Reservepool *ReservepoolTransactorSession) Receive() (*types.Transaction, error) {
	return _Reservepool.Contract.Receive(&_Reservepool.TransactOpts)
}

// ReservepoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Reservepool contract.
type ReservepoolDepositIterator struct {
	Event *ReservepoolDeposit // Event containing the contract specifics and raw log

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
func (it *ReservepoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservepoolDeposit)
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
		it.Event = new(ReservepoolDeposit)
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
func (it *ReservepoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservepoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservepoolDeposit represents a Deposit event raised by the Reservepool contract.
type ReservepoolDeposit struct {
	Actor  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed actor, uint256 amount)
func (_Reservepool *ReservepoolFilterer) FilterDeposit(opts *bind.FilterOpts, actor []common.Address) (*ReservepoolDepositIterator, error) {

	var actorRule []interface{}
	for _, actorItem := range actor {
		actorRule = append(actorRule, actorItem)
	}

	logs, sub, err := _Reservepool.contract.FilterLogs(opts, "Deposit", actorRule)
	if err != nil {
		return nil, err
	}
	return &ReservepoolDepositIterator{contract: _Reservepool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed actor, uint256 amount)
func (_Reservepool *ReservepoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ReservepoolDeposit, actor []common.Address) (event.Subscription, error) {

	var actorRule []interface{}
	for _, actorItem := range actor {
		actorRule = append(actorRule, actorItem)
	}

	logs, sub, err := _Reservepool.contract.WatchLogs(opts, "Deposit", actorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservepoolDeposit)
				if err := _Reservepool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed actor, uint256 amount)
func (_Reservepool *ReservepoolFilterer) ParseDeposit(log types.Log) (*ReservepoolDeposit, error) {
	event := new(ReservepoolDeposit)
	if err := _Reservepool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReservepoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Reservepool contract.
type ReservepoolWithdrawIterator struct {
	Event *ReservepoolWithdraw // Event containing the contract specifics and raw log

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
func (it *ReservepoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservepoolWithdraw)
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
		it.Event = new(ReservepoolWithdraw)
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
func (it *ReservepoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservepoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservepoolWithdraw represents a Withdraw event raised by the Reservepool contract.
type ReservepoolWithdraw struct {
	Actor  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed actor, uint256 amount)
func (_Reservepool *ReservepoolFilterer) FilterWithdraw(opts *bind.FilterOpts, actor []common.Address) (*ReservepoolWithdrawIterator, error) {

	var actorRule []interface{}
	for _, actorItem := range actor {
		actorRule = append(actorRule, actorItem)
	}

	logs, sub, err := _Reservepool.contract.FilterLogs(opts, "Withdraw", actorRule)
	if err != nil {
		return nil, err
	}
	return &ReservepoolWithdrawIterator{contract: _Reservepool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed actor, uint256 amount)
func (_Reservepool *ReservepoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ReservepoolWithdraw, actor []common.Address) (event.Subscription, error) {

	var actorRule []interface{}
	for _, actorItem := range actor {
		actorRule = append(actorRule, actorItem)
	}

	logs, sub, err := _Reservepool.contract.WatchLogs(opts, "Withdraw", actorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservepoolWithdraw)
				if err := _Reservepool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed actor, uint256 amount)
func (_Reservepool *ReservepoolFilterer) ParseWithdraw(log types.Log) (*ReservepoolWithdraw, error) {
	event := new(ReservepoolWithdraw)
	if err := _Reservepool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// PowerTokenABI is the input ABI used to generate the binding from.
const PowerTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PowerTokenBin is the compiled bytecode used for deploying new contracts.
const PowerTokenBin = `0x6060604052341561000f57600080fd5b6101c08061001e6000396000f3006060604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461005b578063313ce567146100e557806395d89b411461010e575b600080fd5b341561006657600080fd5b61006e610121565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100aa578082015183820152602001610092565b50505050905090810190601f1680156100d75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156100f057600080fd5b6100f8610158565b60405160ff909116815260200160405180910390f35b341561011957600080fd5b61006e61015d565b60408051908101604052600981527f506f776572436f696e0000000000000000000000000000000000000000000000602082015281565b601281565b60408051908101604052600381527f50575200000000000000000000000000000000000000000000000000000000006020820152815600a165627a7a72305820cf36229ab63f6fe3b378e1cc9576dbf8a08b98fa6d6d46d869fb0afec69d72a30029`

// DeployPowerToken deploys a new Ethereum contract, binding an instance of PowerToken to it.
func DeployPowerToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PowerToken, error) {
	parsed, err := abi.JSON(strings.NewReader(PowerTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PowerTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PowerToken{PowerTokenCaller: PowerTokenCaller{contract: contract}, PowerTokenTransactor: PowerTokenTransactor{contract: contract}}, nil
}

// PowerToken is an auto generated Go binding around an Ethereum contract.
type PowerToken struct {
	PowerTokenCaller     // Read-only binding to the contract
	PowerTokenTransactor // Write-only binding to the contract
}

// PowerTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PowerTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowerTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PowerTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowerTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PowerTokenSession struct {
	Contract     *PowerToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PowerTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PowerTokenCallerSession struct {
	Contract *PowerTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PowerTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PowerTokenTransactorSession struct {
	Contract     *PowerTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PowerTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PowerTokenRaw struct {
	Contract *PowerToken // Generic contract binding to access the raw methods on
}

// PowerTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PowerTokenCallerRaw struct {
	Contract *PowerTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PowerTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PowerTokenTransactorRaw struct {
	Contract *PowerTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPowerToken creates a new instance of PowerToken, bound to a specific deployed contract.
func NewPowerToken(address common.Address, backend bind.ContractBackend) (*PowerToken, error) {
	contract, err := bindPowerToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PowerToken{PowerTokenCaller: PowerTokenCaller{contract: contract}, PowerTokenTransactor: PowerTokenTransactor{contract: contract}}, nil
}

// NewPowerTokenCaller creates a new read-only instance of PowerToken, bound to a specific deployed contract.
func NewPowerTokenCaller(address common.Address, caller bind.ContractCaller) (*PowerTokenCaller, error) {
	contract, err := bindPowerToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PowerTokenCaller{contract: contract}, nil
}

// NewPowerTokenTransactor creates a new write-only instance of PowerToken, bound to a specific deployed contract.
func NewPowerTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PowerTokenTransactor, error) {
	contract, err := bindPowerToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PowerTokenTransactor{contract: contract}, nil
}

// bindPowerToken binds a generic wrapper to an already deployed contract.
func bindPowerToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PowerTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PowerToken *PowerTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PowerToken.Contract.PowerTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PowerToken *PowerTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PowerToken.Contract.PowerTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PowerToken *PowerTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PowerToken.Contract.PowerTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PowerToken *PowerTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PowerToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PowerToken *PowerTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PowerToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PowerToken *PowerTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PowerToken.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PowerToken *PowerTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _PowerToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PowerToken *PowerTokenSession) Decimals() (uint8, error) {
	return _PowerToken.Contract.Decimals(&_PowerToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PowerToken *PowerTokenCallerSession) Decimals() (uint8, error) {
	return _PowerToken.Contract.Decimals(&_PowerToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PowerToken *PowerTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PowerToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PowerToken *PowerTokenSession) Name() (string, error) {
	return _PowerToken.Contract.Name(&_PowerToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PowerToken *PowerTokenCallerSession) Name() (string, error) {
	return _PowerToken.Contract.Name(&_PowerToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PowerToken *PowerTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PowerToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PowerToken *PowerTokenSession) Symbol() (string, error) {
	return _PowerToken.Contract.Symbol(&_PowerToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PowerToken *PowerTokenCallerSession) Symbol() (string, error) {
	return _PowerToken.Contract.Symbol(&_PowerToken.CallOpts)
}

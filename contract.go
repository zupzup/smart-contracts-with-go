// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// WinnerTakesAllABI is the input ABI used to generate the binding from.
const WinnerTakesAllABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"deadlineCampaign\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winningAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getProjectInfo\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"url\",\"type\":\"string\"},{\"name\":\"funds\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winningFunds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"url\",\"type\":\"string\"}],\"name\":\"submitProject\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"projectAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"supportProject\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberOfProjects\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deadlineProjects\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_minimumEntryFee\",\"type\":\"uint256\"},{\"name\":\"_durationProjects\",\"type\":\"uint256\"},{\"name\":\"_durationCampaign\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"initialized\",\"type\":\"bool\"}],\"name\":\"ProjectSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProjectSupported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winningFunds\",\"type\":\"uint256\"}],\"name\":\"PayedOutTo\",\"type\":\"event\"}]"

// WinnerTakesAllBin is the compiled bytecode used for deploying new contracts.
const WinnerTakesAllBin = `0x6060604052341561000f57600080fd5b604051606080610a9183398101604052808051919060200180519190602001805191505081811161003f57600080fd5b6000928355429182016001550160025560048054600160a060020a033316600160a060020a0319909116179055600355610a138061007e6000396000f3006060604052600436106100a35763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633356331d81146100a85780633ea8b3bc146100cd57806347ee45b9146100fc5780637ecdf746146102035780637ee3e74f14610216578063cc13fd3b146102b2578063cf5e9747146102c8578063d56b2889146102dc578063db5e4c7f146102f1578063eb1f9d3914610304575b600080fd5b34156100b357600080fd5b6100bb610317565b60405190815260200160405180910390f35b34156100d857600080fd5b6100e061031d565b604051600160a060020a03909116815260200160405180910390f35b341561010757600080fd5b61011b600160a060020a036004351661032c565b604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b8381101561016257808201518382015260200161014a565b50505050905090810190601f16801561018f5780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b838110156101c55780820151838201526020016101ad565b50505050905090810190601f1680156101f25780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b341561020e57600080fd5b6100bb6104ba565b61029e60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506104c095505050505050565b604051901515815260200160405180910390f35b34156102bd57600080fd5b6100e0600435610760565b61029e600160a060020a0360043516610788565b34156102e757600080fd5b6102ef610899565b005b34156102fc57600080fd5b6100bb610905565b341561030f57600080fd5b6100bb61090b565b60025481565b600454600160a060020a031681565b610334610911565b61033c610911565b600160a060020a0383166000908152600560205260408120600481015460ff16151561036757600080fd5b80600101816002018260030154828054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104095780601f106103de57610100808354040283529160200191610409565b820191906000526020600020905b8154815290600101906020018083116103ec57829003601f168201915b50505050509250818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104a55780601f1061047a576101008083540402835291602001916104a5565b820191906000526020600020905b81548152906001019060200180831161048857829003601f168201915b50505050509150935093509350509193909250565b60035481565b600080543410156104d057600080fd5b6001544211156104df57600080fd5b600160a060020a03331660009081526005602052604090206004015460ff1615156107565760a06040519081016040908152600160a060020a03331680835260208084018790528284018690526000606085018190526001608086015291825260059052208151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0391909116178155602082015181600101908051610588929160200190610923565b506040820151816002019080516105a3929160200190610923565b50606082015181600301556080820151600491909101805460ff19169115159190911790555060068054600181016105db83826109a1565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a03811691821790925560065460075583526005909152604091829020600401547fa31614ce68f0443c0017f17939642084995e6a8b91736cac18520b829b157c33928691869160ff9091169051600160a060020a03851681528115156060820152608060208201818152906040830190830186818151815260200191508051906020019080838360005b838110156106ad578082015183820152602001610695565b50505050905090810190601f1680156106da5780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b838110156107105780820151838201526020016106f8565b50505050905090810190601f16801561073d5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a150600161075a565b5060005b92915050565b600680548290811061076e57fe5b600091825260209091200154600160a060020a0316905081565b60003481901161079757600080fd5b6002544211806107a957506001544211155b156107b357600080fd5b600160a060020a03821660009081526005602052604090206004015460ff1615156107dd57600080fd5b600160a060020a038216600090815260056020526040902060039081018054340190819055905490111561084d576004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038416908117909155600090815260056020526040902060039081015490555b7fdaf9d988cff8579fb76d980affd31da41317fddf02e5134b4c00220b2cb2f3748234604051600160a060020a03909216825260208201526040908101905180910390a1506001919050565b6002544210610903576004546003547f8fba92fdc99409a91fdfb6b6b1da7ddf633a9c3cc3fc11f237866e54321d39a991600160a060020a031690604051600160a060020a03909216825260208201526040908101905180910390a1600454600160a060020a0316ff5b565b60075481565b60015481565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061096457805160ff1916838001178555610991565b82800160010185558215610991579182015b82811115610991578251825591602001919060010190610976565b5061099d9291506109ca565b5090565b8154818355818115116109c5576000838152602090206109c59181019083016109ca565b505050565b6109e491905b8082111561099d57600081556001016109d0565b905600a165627a7a7230582084ede253ea549df90d0ff9dd3e1c87f9e93af0a77fefa9d1aa472e11fdfe33920029`

// DeployWinnerTakesAll deploys a new Ethereum contract, binding an instance of WinnerTakesAll to it.
func DeployWinnerTakesAll(auth *bind.TransactOpts, backend bind.ContractBackend, _minimumEntryFee *big.Int, _durationProjects *big.Int, _durationCampaign *big.Int) (common.Address, *types.Transaction, *WinnerTakesAll, error) {
	parsed, err := abi.JSON(strings.NewReader(WinnerTakesAllABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WinnerTakesAllBin), backend, _minimumEntryFee, _durationProjects, _durationCampaign)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WinnerTakesAll{WinnerTakesAllCaller: WinnerTakesAllCaller{contract: contract}, WinnerTakesAllTransactor: WinnerTakesAllTransactor{contract: contract}, WinnerTakesAllFilterer: WinnerTakesAllFilterer{contract: contract}}, nil
}

// WinnerTakesAll is an auto generated Go binding around an Ethereum contract.
type WinnerTakesAll struct {
	WinnerTakesAllCaller     // Read-only binding to the contract
	WinnerTakesAllTransactor // Write-only binding to the contract
	WinnerTakesAllFilterer   // Log filterer for contract events
}

// WinnerTakesAllCaller is an auto generated read-only Go binding around an Ethereum contract.
type WinnerTakesAllCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WinnerTakesAllTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WinnerTakesAllTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WinnerTakesAllFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WinnerTakesAllFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WinnerTakesAllSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WinnerTakesAllSession struct {
	Contract     *WinnerTakesAll   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WinnerTakesAllCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WinnerTakesAllCallerSession struct {
	Contract *WinnerTakesAllCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// WinnerTakesAllTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WinnerTakesAllTransactorSession struct {
	Contract     *WinnerTakesAllTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WinnerTakesAllRaw is an auto generated low-level Go binding around an Ethereum contract.
type WinnerTakesAllRaw struct {
	Contract *WinnerTakesAll // Generic contract binding to access the raw methods on
}

// WinnerTakesAllCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WinnerTakesAllCallerRaw struct {
	Contract *WinnerTakesAllCaller // Generic read-only contract binding to access the raw methods on
}

// WinnerTakesAllTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WinnerTakesAllTransactorRaw struct {
	Contract *WinnerTakesAllTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWinnerTakesAll creates a new instance of WinnerTakesAll, bound to a specific deployed contract.
func NewWinnerTakesAll(address common.Address, backend bind.ContractBackend) (*WinnerTakesAll, error) {
	contract, err := bindWinnerTakesAll(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WinnerTakesAll{WinnerTakesAllCaller: WinnerTakesAllCaller{contract: contract}, WinnerTakesAllTransactor: WinnerTakesAllTransactor{contract: contract}, WinnerTakesAllFilterer: WinnerTakesAllFilterer{contract: contract}}, nil
}

// NewWinnerTakesAllCaller creates a new read-only instance of WinnerTakesAll, bound to a specific deployed contract.
func NewWinnerTakesAllCaller(address common.Address, caller bind.ContractCaller) (*WinnerTakesAllCaller, error) {
	contract, err := bindWinnerTakesAll(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WinnerTakesAllCaller{contract: contract}, nil
}

// NewWinnerTakesAllTransactor creates a new write-only instance of WinnerTakesAll, bound to a specific deployed contract.
func NewWinnerTakesAllTransactor(address common.Address, transactor bind.ContractTransactor) (*WinnerTakesAllTransactor, error) {
	contract, err := bindWinnerTakesAll(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WinnerTakesAllTransactor{contract: contract}, nil
}

// NewWinnerTakesAllFilterer creates a new log filterer instance of WinnerTakesAll, bound to a specific deployed contract.
func NewWinnerTakesAllFilterer(address common.Address, filterer bind.ContractFilterer) (*WinnerTakesAllFilterer, error) {
	contract, err := bindWinnerTakesAll(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WinnerTakesAllFilterer{contract: contract}, nil
}

// bindWinnerTakesAll binds a generic wrapper to an already deployed contract.
func bindWinnerTakesAll(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WinnerTakesAllABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WinnerTakesAll *WinnerTakesAllRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WinnerTakesAll.Contract.WinnerTakesAllCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WinnerTakesAll *WinnerTakesAllRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.WinnerTakesAllTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WinnerTakesAll *WinnerTakesAllRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.WinnerTakesAllTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WinnerTakesAll *WinnerTakesAllCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WinnerTakesAll.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WinnerTakesAll *WinnerTakesAllTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WinnerTakesAll *WinnerTakesAllTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.contract.Transact(opts, method, params...)
}

// DeadlineCampaign is a free data retrieval call binding the contract method 0x3356331d.
//
// Solidity: function deadlineCampaign() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCaller) DeadlineCampaign(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WinnerTakesAll.contract.Call(opts, out, "deadlineCampaign")
	return *ret0, err
}

// DeadlineCampaign is a free data retrieval call binding the contract method 0x3356331d.
//
// Solidity: function deadlineCampaign() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllSession) DeadlineCampaign() (*big.Int, error) {
	return _WinnerTakesAll.Contract.DeadlineCampaign(&_WinnerTakesAll.CallOpts)
}

// DeadlineCampaign is a free data retrieval call binding the contract method 0x3356331d.
//
// Solidity: function deadlineCampaign() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCallerSession) DeadlineCampaign() (*big.Int, error) {
	return _WinnerTakesAll.Contract.DeadlineCampaign(&_WinnerTakesAll.CallOpts)
}

// DeadlineProjects is a free data retrieval call binding the contract method 0xeb1f9d39.
//
// Solidity: function deadlineProjects() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCaller) DeadlineProjects(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WinnerTakesAll.contract.Call(opts, out, "deadlineProjects")
	return *ret0, err
}

// DeadlineProjects is a free data retrieval call binding the contract method 0xeb1f9d39.
//
// Solidity: function deadlineProjects() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllSession) DeadlineProjects() (*big.Int, error) {
	return _WinnerTakesAll.Contract.DeadlineProjects(&_WinnerTakesAll.CallOpts)
}

// DeadlineProjects is a free data retrieval call binding the contract method 0xeb1f9d39.
//
// Solidity: function deadlineProjects() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCallerSession) DeadlineProjects() (*big.Int, error) {
	return _WinnerTakesAll.Contract.DeadlineProjects(&_WinnerTakesAll.CallOpts)
}

// GetProjectInfo is a free data retrieval call binding the contract method 0x47ee45b9.
//
// Solidity: function getProjectInfo(addr address) constant returns(name string, url string, funds uint256)
func (_WinnerTakesAll *WinnerTakesAllCaller) GetProjectInfo(opts *bind.CallOpts, addr common.Address) (struct {
	Name  string
	Url   string
	Funds *big.Int
}, error) {
	ret := new(struct {
		Name  string
		Url   string
		Funds *big.Int
	})
	out := ret
	err := _WinnerTakesAll.contract.Call(opts, out, "getProjectInfo", addr)
	return *ret, err
}

// GetProjectInfo is a free data retrieval call binding the contract method 0x47ee45b9.
//
// Solidity: function getProjectInfo(addr address) constant returns(name string, url string, funds uint256)
func (_WinnerTakesAll *WinnerTakesAllSession) GetProjectInfo(addr common.Address) (struct {
	Name  string
	Url   string
	Funds *big.Int
}, error) {
	return _WinnerTakesAll.Contract.GetProjectInfo(&_WinnerTakesAll.CallOpts, addr)
}

// GetProjectInfo is a free data retrieval call binding the contract method 0x47ee45b9.
//
// Solidity: function getProjectInfo(addr address) constant returns(name string, url string, funds uint256)
func (_WinnerTakesAll *WinnerTakesAllCallerSession) GetProjectInfo(addr common.Address) (struct {
	Name  string
	Url   string
	Funds *big.Int
}, error) {
	return _WinnerTakesAll.Contract.GetProjectInfo(&_WinnerTakesAll.CallOpts, addr)
}

// NumberOfProjects is a free data retrieval call binding the contract method 0xdb5e4c7f.
//
// Solidity: function numberOfProjects() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCaller) NumberOfProjects(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WinnerTakesAll.contract.Call(opts, out, "numberOfProjects")
	return *ret0, err
}

// NumberOfProjects is a free data retrieval call binding the contract method 0xdb5e4c7f.
//
// Solidity: function numberOfProjects() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllSession) NumberOfProjects() (*big.Int, error) {
	return _WinnerTakesAll.Contract.NumberOfProjects(&_WinnerTakesAll.CallOpts)
}

// NumberOfProjects is a free data retrieval call binding the contract method 0xdb5e4c7f.
//
// Solidity: function numberOfProjects() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCallerSession) NumberOfProjects() (*big.Int, error) {
	return _WinnerTakesAll.Contract.NumberOfProjects(&_WinnerTakesAll.CallOpts)
}

// ProjectAddresses is a free data retrieval call binding the contract method 0xcc13fd3b.
//
// Solidity: function projectAddresses( uint256) constant returns(address)
func (_WinnerTakesAll *WinnerTakesAllCaller) ProjectAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WinnerTakesAll.contract.Call(opts, out, "projectAddresses", arg0)
	return *ret0, err
}

// ProjectAddresses is a free data retrieval call binding the contract method 0xcc13fd3b.
//
// Solidity: function projectAddresses( uint256) constant returns(address)
func (_WinnerTakesAll *WinnerTakesAllSession) ProjectAddresses(arg0 *big.Int) (common.Address, error) {
	return _WinnerTakesAll.Contract.ProjectAddresses(&_WinnerTakesAll.CallOpts, arg0)
}

// ProjectAddresses is a free data retrieval call binding the contract method 0xcc13fd3b.
//
// Solidity: function projectAddresses( uint256) constant returns(address)
func (_WinnerTakesAll *WinnerTakesAllCallerSession) ProjectAddresses(arg0 *big.Int) (common.Address, error) {
	return _WinnerTakesAll.Contract.ProjectAddresses(&_WinnerTakesAll.CallOpts, arg0)
}

// WinningAddress is a free data retrieval call binding the contract method 0x3ea8b3bc.
//
// Solidity: function winningAddress() constant returns(address)
func (_WinnerTakesAll *WinnerTakesAllCaller) WinningAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WinnerTakesAll.contract.Call(opts, out, "winningAddress")
	return *ret0, err
}

// WinningAddress is a free data retrieval call binding the contract method 0x3ea8b3bc.
//
// Solidity: function winningAddress() constant returns(address)
func (_WinnerTakesAll *WinnerTakesAllSession) WinningAddress() (common.Address, error) {
	return _WinnerTakesAll.Contract.WinningAddress(&_WinnerTakesAll.CallOpts)
}

// WinningAddress is a free data retrieval call binding the contract method 0x3ea8b3bc.
//
// Solidity: function winningAddress() constant returns(address)
func (_WinnerTakesAll *WinnerTakesAllCallerSession) WinningAddress() (common.Address, error) {
	return _WinnerTakesAll.Contract.WinningAddress(&_WinnerTakesAll.CallOpts)
}

// WinningFunds is a free data retrieval call binding the contract method 0x7ecdf746.
//
// Solidity: function winningFunds() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCaller) WinningFunds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WinnerTakesAll.contract.Call(opts, out, "winningFunds")
	return *ret0, err
}

// WinningFunds is a free data retrieval call binding the contract method 0x7ecdf746.
//
// Solidity: function winningFunds() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllSession) WinningFunds() (*big.Int, error) {
	return _WinnerTakesAll.Contract.WinningFunds(&_WinnerTakesAll.CallOpts)
}

// WinningFunds is a free data retrieval call binding the contract method 0x7ecdf746.
//
// Solidity: function winningFunds() constant returns(uint256)
func (_WinnerTakesAll *WinnerTakesAllCallerSession) WinningFunds() (*big.Int, error) {
	return _WinnerTakesAll.Contract.WinningFunds(&_WinnerTakesAll.CallOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_WinnerTakesAll *WinnerTakesAllTransactor) Finish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WinnerTakesAll.contract.Transact(opts, "finish")
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_WinnerTakesAll *WinnerTakesAllSession) Finish() (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.Finish(&_WinnerTakesAll.TransactOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_WinnerTakesAll *WinnerTakesAllTransactorSession) Finish() (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.Finish(&_WinnerTakesAll.TransactOpts)
}

// SubmitProject is a paid mutator transaction binding the contract method 0x7ee3e74f.
//
// Solidity: function submitProject(name string, url string) returns(success bool)
func (_WinnerTakesAll *WinnerTakesAllTransactor) SubmitProject(opts *bind.TransactOpts, name string, url string) (*types.Transaction, error) {
	return _WinnerTakesAll.contract.Transact(opts, "submitProject", name, url)
}

// SubmitProject is a paid mutator transaction binding the contract method 0x7ee3e74f.
//
// Solidity: function submitProject(name string, url string) returns(success bool)
func (_WinnerTakesAll *WinnerTakesAllSession) SubmitProject(name string, url string) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.SubmitProject(&_WinnerTakesAll.TransactOpts, name, url)
}

// SubmitProject is a paid mutator transaction binding the contract method 0x7ee3e74f.
//
// Solidity: function submitProject(name string, url string) returns(success bool)
func (_WinnerTakesAll *WinnerTakesAllTransactorSession) SubmitProject(name string, url string) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.SubmitProject(&_WinnerTakesAll.TransactOpts, name, url)
}

// SupportProject is a paid mutator transaction binding the contract method 0xcf5e9747.
//
// Solidity: function supportProject(addr address) returns(success bool)
func (_WinnerTakesAll *WinnerTakesAllTransactor) SupportProject(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _WinnerTakesAll.contract.Transact(opts, "supportProject", addr)
}

// SupportProject is a paid mutator transaction binding the contract method 0xcf5e9747.
//
// Solidity: function supportProject(addr address) returns(success bool)
func (_WinnerTakesAll *WinnerTakesAllSession) SupportProject(addr common.Address) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.SupportProject(&_WinnerTakesAll.TransactOpts, addr)
}

// SupportProject is a paid mutator transaction binding the contract method 0xcf5e9747.
//
// Solidity: function supportProject(addr address) returns(success bool)
func (_WinnerTakesAll *WinnerTakesAllTransactorSession) SupportProject(addr common.Address) (*types.Transaction, error) {
	return _WinnerTakesAll.Contract.SupportProject(&_WinnerTakesAll.TransactOpts, addr)
}

// WinnerTakesAllPayedOutToIterator is returned from FilterPayedOutTo and is used to iterate over the raw logs and unpacked data for PayedOutTo events raised by the WinnerTakesAll contract.
type WinnerTakesAllPayedOutToIterator struct {
	Event *WinnerTakesAllPayedOutTo // Event containing the contract specifics and raw log

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
func (it *WinnerTakesAllPayedOutToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WinnerTakesAllPayedOutTo)
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
		it.Event = new(WinnerTakesAllPayedOutTo)
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
func (it *WinnerTakesAllPayedOutToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WinnerTakesAllPayedOutToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WinnerTakesAllPayedOutTo represents a PayedOutTo event raised by the WinnerTakesAll contract.
type WinnerTakesAllPayedOutTo struct {
	Addr         common.Address
	WinningFunds *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPayedOutTo is a free log retrieval operation binding the contract event 0x8fba92fdc99409a91fdfb6b6b1da7ddf633a9c3cc3fc11f237866e54321d39a9.
//
// Solidity: event PayedOutTo(addr address, winningFunds uint256)
func (_WinnerTakesAll *WinnerTakesAllFilterer) FilterPayedOutTo(opts *bind.FilterOpts) (*WinnerTakesAllPayedOutToIterator, error) {

	logs, sub, err := _WinnerTakesAll.contract.FilterLogs(opts, "PayedOutTo")
	if err != nil {
		return nil, err
	}
	return &WinnerTakesAllPayedOutToIterator{contract: _WinnerTakesAll.contract, event: "PayedOutTo", logs: logs, sub: sub}, nil
}

// WatchPayedOutTo is a free log subscription operation binding the contract event 0x8fba92fdc99409a91fdfb6b6b1da7ddf633a9c3cc3fc11f237866e54321d39a9.
//
// Solidity: event PayedOutTo(addr address, winningFunds uint256)
func (_WinnerTakesAll *WinnerTakesAllFilterer) WatchPayedOutTo(opts *bind.WatchOpts, sink chan<- *WinnerTakesAllPayedOutTo) (event.Subscription, error) {

	logs, sub, err := _WinnerTakesAll.contract.WatchLogs(opts, "PayedOutTo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WinnerTakesAllPayedOutTo)
				if err := _WinnerTakesAll.contract.UnpackLog(event, "PayedOutTo", log); err != nil {
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

// WinnerTakesAllProjectSubmittedIterator is returned from FilterProjectSubmitted and is used to iterate over the raw logs and unpacked data for ProjectSubmitted events raised by the WinnerTakesAll contract.
type WinnerTakesAllProjectSubmittedIterator struct {
	Event *WinnerTakesAllProjectSubmitted // Event containing the contract specifics and raw log

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
func (it *WinnerTakesAllProjectSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WinnerTakesAllProjectSubmitted)
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
		it.Event = new(WinnerTakesAllProjectSubmitted)
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
func (it *WinnerTakesAllProjectSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WinnerTakesAllProjectSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WinnerTakesAllProjectSubmitted represents a ProjectSubmitted event raised by the WinnerTakesAll contract.
type WinnerTakesAllProjectSubmitted struct {
	Addr        common.Address
	Name        string
	Url         string
	Initialized bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProjectSubmitted is a free log retrieval operation binding the contract event 0xa31614ce68f0443c0017f17939642084995e6a8b91736cac18520b829b157c33.
//
// Solidity: event ProjectSubmitted(addr address, name string, url string, initialized bool)
func (_WinnerTakesAll *WinnerTakesAllFilterer) FilterProjectSubmitted(opts *bind.FilterOpts) (*WinnerTakesAllProjectSubmittedIterator, error) {

	logs, sub, err := _WinnerTakesAll.contract.FilterLogs(opts, "ProjectSubmitted")
	if err != nil {
		return nil, err
	}
	return &WinnerTakesAllProjectSubmittedIterator{contract: _WinnerTakesAll.contract, event: "ProjectSubmitted", logs: logs, sub: sub}, nil
}

// WatchProjectSubmitted is a free log subscription operation binding the contract event 0xa31614ce68f0443c0017f17939642084995e6a8b91736cac18520b829b157c33.
//
// Solidity: event ProjectSubmitted(addr address, name string, url string, initialized bool)
func (_WinnerTakesAll *WinnerTakesAllFilterer) WatchProjectSubmitted(opts *bind.WatchOpts, sink chan<- *WinnerTakesAllProjectSubmitted) (event.Subscription, error) {

	logs, sub, err := _WinnerTakesAll.contract.WatchLogs(opts, "ProjectSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WinnerTakesAllProjectSubmitted)
				if err := _WinnerTakesAll.contract.UnpackLog(event, "ProjectSubmitted", log); err != nil {
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

// WinnerTakesAllProjectSupportedIterator is returned from FilterProjectSupported and is used to iterate over the raw logs and unpacked data for ProjectSupported events raised by the WinnerTakesAll contract.
type WinnerTakesAllProjectSupportedIterator struct {
	Event *WinnerTakesAllProjectSupported // Event containing the contract specifics and raw log

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
func (it *WinnerTakesAllProjectSupportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WinnerTakesAllProjectSupported)
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
		it.Event = new(WinnerTakesAllProjectSupported)
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
func (it *WinnerTakesAllProjectSupportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WinnerTakesAllProjectSupportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WinnerTakesAllProjectSupported represents a ProjectSupported event raised by the WinnerTakesAll contract.
type WinnerTakesAllProjectSupported struct {
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProjectSupported is a free log retrieval operation binding the contract event 0xdaf9d988cff8579fb76d980affd31da41317fddf02e5134b4c00220b2cb2f374.
//
// Solidity: event ProjectSupported(addr address, amount uint256)
func (_WinnerTakesAll *WinnerTakesAllFilterer) FilterProjectSupported(opts *bind.FilterOpts) (*WinnerTakesAllProjectSupportedIterator, error) {

	logs, sub, err := _WinnerTakesAll.contract.FilterLogs(opts, "ProjectSupported")
	if err != nil {
		return nil, err
	}
	return &WinnerTakesAllProjectSupportedIterator{contract: _WinnerTakesAll.contract, event: "ProjectSupported", logs: logs, sub: sub}, nil
}

// WatchProjectSupported is a free log subscription operation binding the contract event 0xdaf9d988cff8579fb76d980affd31da41317fddf02e5134b4c00220b2cb2f374.
//
// Solidity: event ProjectSupported(addr address, amount uint256)
func (_WinnerTakesAll *WinnerTakesAllFilterer) WatchProjectSupported(opts *bind.WatchOpts, sink chan<- *WinnerTakesAllProjectSupported) (event.Subscription, error) {

	logs, sub, err := _WinnerTakesAll.contract.WatchLogs(opts, "ProjectSupported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WinnerTakesAllProjectSupported)
				if err := _WinnerTakesAll.contract.UnpackLog(event, "ProjectSupported", log); err != nil {
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

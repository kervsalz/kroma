// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ProposerRewardVaultMetaData contains all meta data concerning the ProposerRewardVault contract.
var ProposerRewardVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_WITHDRAWAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECIPIENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalProcessed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x61012060405234801561001157600080fd5b5060405161086c38038061086c8339810160408190526100309161005d565b678ac7230489e800006080526001600160a01b031660a052600160c052600060e08190526101005261008d565b60006020828403121561006f57600080fd5b81516001600160a01b038116811461008657600080fd5b9392505050565b60805160a05160c05160e051610100516107876100e560003960006103d3015260006103aa01526000610381015260008181607c015281816102570152610319015260008181610137015261015b01526107876000f3fe60806040526004361061005e5760003560e01c806354fd4d501161004357806354fd4d50146100df57806384411d6514610101578063d3e5792b1461012557600080fd5b80630d9019e11461006a5780633ccfd60b146100c857600080fd5b3661006557005b600080fd5b34801561007657600080fd5b5061009e7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156100d457600080fd5b506100dd610159565b005b3480156100eb57600080fd5b506100f461037a565b6040516100bf9190610638565b34801561010d57600080fd5b5061011760005481565b6040519081526020016100bf565b34801561013157600080fd5b506101177f000000000000000000000000000000000000000000000000000000000000000081565b7f0000000000000000000000000000000000000000000000000000000000000000471015610233576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f4665655661756c743a207769746864726177616c20616d6f756e74206d75737460448201527f2062652067726561746572207468616e206d696e696d756d207769746864726160648201527f77616c20616d6f756e7400000000000000000000000000000000000000000000608482015260a40160405180910390fd5b6000479050806000808282546102499190610652565b9091555050604080518281527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166020820152338183015290517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a1604080516020810182526000815290517fe11013dd0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000099163e11013dd918491610345917f0000000000000000000000000000000000000000000000000000000000000000916188b891600401610691565b6000604051808303818588803b15801561035e57600080fd5b505af1158015610372573d6000803e3d6000fd5b505050505050565b60606103a57f000000000000000000000000000000000000000000000000000000000000000061041d565b6103ce7f000000000000000000000000000000000000000000000000000000000000000061041d565b6103f77f000000000000000000000000000000000000000000000000000000000000000061041d565b604051602001610409939291906106d5565b604051602081830303815290604052905090565b6060600061042a836104db565b600101905060008167ffffffffffffffff81111561044a5761044a61074b565b6040519080825280601f01601f191660200182016040528015610474576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461047e57509392505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310610524577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310610550576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061056e57662386f26fc10000830492506010015b6305f5e1008310610586576305f5e100830492506008015b612710831061059a57612710830492506004015b606483106105ac576064830492506002015b600a83106105b8576001015b92915050565b60005b838110156105d95781810151838201526020016105c1565b838111156105e8576000848401525b50505050565b600081518084526106068160208601602086016105be565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061064b60208301846105ee565b9392505050565b6000821982111561068c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500190565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff831660208201526060604082015260006106cc60608301846105ee565b95945050505050565b600084516106e78184602089016105be565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610723816001850160208a016105be565b6001920191820152835161073e8160028401602088016105be565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c634300080f000a",
}

// ProposerRewardVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ProposerRewardVaultMetaData.ABI instead.
var ProposerRewardVaultABI = ProposerRewardVaultMetaData.ABI

// ProposerRewardVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProposerRewardVaultMetaData.Bin instead.
var ProposerRewardVaultBin = ProposerRewardVaultMetaData.Bin

// DeployProposerRewardVault deploys a new Ethereum contract, binding an instance of ProposerRewardVault to it.
func DeployProposerRewardVault(auth *bind.TransactOpts, backend bind.ContractBackend, _recipient common.Address) (common.Address, *types.Transaction, *ProposerRewardVault, error) {
	parsed, err := ProposerRewardVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProposerRewardVaultBin), backend, _recipient)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProposerRewardVault{ProposerRewardVaultCaller: ProposerRewardVaultCaller{contract: contract}, ProposerRewardVaultTransactor: ProposerRewardVaultTransactor{contract: contract}, ProposerRewardVaultFilterer: ProposerRewardVaultFilterer{contract: contract}}, nil
}

// ProposerRewardVault is an auto generated Go binding around an Ethereum contract.
type ProposerRewardVault struct {
	ProposerRewardVaultCaller     // Read-only binding to the contract
	ProposerRewardVaultTransactor // Write-only binding to the contract
	ProposerRewardVaultFilterer   // Log filterer for contract events
}

// ProposerRewardVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProposerRewardVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposerRewardVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProposerRewardVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposerRewardVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProposerRewardVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposerRewardVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProposerRewardVaultSession struct {
	Contract     *ProposerRewardVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ProposerRewardVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProposerRewardVaultCallerSession struct {
	Contract *ProposerRewardVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ProposerRewardVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProposerRewardVaultTransactorSession struct {
	Contract     *ProposerRewardVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ProposerRewardVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProposerRewardVaultRaw struct {
	Contract *ProposerRewardVault // Generic contract binding to access the raw methods on
}

// ProposerRewardVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProposerRewardVaultCallerRaw struct {
	Contract *ProposerRewardVaultCaller // Generic read-only contract binding to access the raw methods on
}

// ProposerRewardVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProposerRewardVaultTransactorRaw struct {
	Contract *ProposerRewardVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProposerRewardVault creates a new instance of ProposerRewardVault, bound to a specific deployed contract.
func NewProposerRewardVault(address common.Address, backend bind.ContractBackend) (*ProposerRewardVault, error) {
	contract, err := bindProposerRewardVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProposerRewardVault{ProposerRewardVaultCaller: ProposerRewardVaultCaller{contract: contract}, ProposerRewardVaultTransactor: ProposerRewardVaultTransactor{contract: contract}, ProposerRewardVaultFilterer: ProposerRewardVaultFilterer{contract: contract}}, nil
}

// NewProposerRewardVaultCaller creates a new read-only instance of ProposerRewardVault, bound to a specific deployed contract.
func NewProposerRewardVaultCaller(address common.Address, caller bind.ContractCaller) (*ProposerRewardVaultCaller, error) {
	contract, err := bindProposerRewardVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProposerRewardVaultCaller{contract: contract}, nil
}

// NewProposerRewardVaultTransactor creates a new write-only instance of ProposerRewardVault, bound to a specific deployed contract.
func NewProposerRewardVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ProposerRewardVaultTransactor, error) {
	contract, err := bindProposerRewardVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProposerRewardVaultTransactor{contract: contract}, nil
}

// NewProposerRewardVaultFilterer creates a new log filterer instance of ProposerRewardVault, bound to a specific deployed contract.
func NewProposerRewardVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ProposerRewardVaultFilterer, error) {
	contract, err := bindProposerRewardVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProposerRewardVaultFilterer{contract: contract}, nil
}

// bindProposerRewardVault binds a generic wrapper to an already deployed contract.
func bindProposerRewardVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProposerRewardVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposerRewardVault *ProposerRewardVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposerRewardVault.Contract.ProposerRewardVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposerRewardVault *ProposerRewardVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposerRewardVault.Contract.ProposerRewardVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposerRewardVault *ProposerRewardVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposerRewardVault.Contract.ProposerRewardVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposerRewardVault *ProposerRewardVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposerRewardVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposerRewardVault *ProposerRewardVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposerRewardVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposerRewardVault *ProposerRewardVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposerRewardVault.Contract.contract.Transact(opts, method, params...)
}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ProposerRewardVault *ProposerRewardVaultCaller) MINWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProposerRewardVault.contract.Call(opts, &out, "MIN_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ProposerRewardVault *ProposerRewardVaultSession) MINWITHDRAWALAMOUNT() (*big.Int, error) {
	return _ProposerRewardVault.Contract.MINWITHDRAWALAMOUNT(&_ProposerRewardVault.CallOpts)
}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ProposerRewardVault *ProposerRewardVaultCallerSession) MINWITHDRAWALAMOUNT() (*big.Int, error) {
	return _ProposerRewardVault.Contract.MINWITHDRAWALAMOUNT(&_ProposerRewardVault.CallOpts)
}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ProposerRewardVault *ProposerRewardVaultCaller) RECIPIENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProposerRewardVault.contract.Call(opts, &out, "RECIPIENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ProposerRewardVault *ProposerRewardVaultSession) RECIPIENT() (common.Address, error) {
	return _ProposerRewardVault.Contract.RECIPIENT(&_ProposerRewardVault.CallOpts)
}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ProposerRewardVault *ProposerRewardVaultCallerSession) RECIPIENT() (common.Address, error) {
	return _ProposerRewardVault.Contract.RECIPIENT(&_ProposerRewardVault.CallOpts)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ProposerRewardVault *ProposerRewardVaultCaller) TotalProcessed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProposerRewardVault.contract.Call(opts, &out, "totalProcessed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ProposerRewardVault *ProposerRewardVaultSession) TotalProcessed() (*big.Int, error) {
	return _ProposerRewardVault.Contract.TotalProcessed(&_ProposerRewardVault.CallOpts)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ProposerRewardVault *ProposerRewardVaultCallerSession) TotalProcessed() (*big.Int, error) {
	return _ProposerRewardVault.Contract.TotalProcessed(&_ProposerRewardVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProposerRewardVault *ProposerRewardVaultCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProposerRewardVault.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProposerRewardVault *ProposerRewardVaultSession) Version() (string, error) {
	return _ProposerRewardVault.Contract.Version(&_ProposerRewardVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProposerRewardVault *ProposerRewardVaultCallerSession) Version() (string, error) {
	return _ProposerRewardVault.Contract.Version(&_ProposerRewardVault.CallOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProposerRewardVault *ProposerRewardVaultTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposerRewardVault.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProposerRewardVault *ProposerRewardVaultSession) Withdraw() (*types.Transaction, error) {
	return _ProposerRewardVault.Contract.Withdraw(&_ProposerRewardVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProposerRewardVault *ProposerRewardVaultTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ProposerRewardVault.Contract.Withdraw(&_ProposerRewardVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProposerRewardVault *ProposerRewardVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposerRewardVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProposerRewardVault *ProposerRewardVaultSession) Receive() (*types.Transaction, error) {
	return _ProposerRewardVault.Contract.Receive(&_ProposerRewardVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProposerRewardVault *ProposerRewardVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _ProposerRewardVault.Contract.Receive(&_ProposerRewardVault.TransactOpts)
}

// ProposerRewardVaultWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ProposerRewardVault contract.
type ProposerRewardVaultWithdrawalIterator struct {
	Event *ProposerRewardVaultWithdrawal // Event containing the contract specifics and raw log

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
func (it *ProposerRewardVaultWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposerRewardVaultWithdrawal)
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
		it.Event = new(ProposerRewardVaultWithdrawal)
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
func (it *ProposerRewardVaultWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposerRewardVaultWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposerRewardVaultWithdrawal represents a Withdrawal event raised by the ProposerRewardVault contract.
type ProposerRewardVaultWithdrawal struct {
	Value *big.Int
	To    common.Address
	From  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_ProposerRewardVault *ProposerRewardVaultFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*ProposerRewardVaultWithdrawalIterator, error) {

	logs, sub, err := _ProposerRewardVault.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &ProposerRewardVaultWithdrawalIterator{contract: _ProposerRewardVault.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_ProposerRewardVault *ProposerRewardVaultFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ProposerRewardVaultWithdrawal) (event.Subscription, error) {

	logs, sub, err := _ProposerRewardVault.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposerRewardVaultWithdrawal)
				if err := _ProposerRewardVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_ProposerRewardVault *ProposerRewardVaultFilterer) ParseWithdrawal(log types.Log) (*ProposerRewardVaultWithdrawal, error) {
	event := new(ProposerRewardVaultWithdrawal)
	if err := _ProposerRewardVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

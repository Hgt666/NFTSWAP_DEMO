// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// NFTTradeMarketBidOrder is an auto generated low-level Go binding around an user-defined struct.
type NFTTradeMarketBidOrder struct {
	Bidder      common.Address
	NftContract common.Address
	TokenId     *big.Int
	Price       *big.Int
	Active      bool
}

// NFTTradeMarketListOrder is an auto generated low-level Go binding around an user-defined struct.
type NFTTradeMarketListOrder struct {
	Seller      common.Address
	NftContract common.Address
	TokenId     *big.Int
	Price       *big.Int
	Active      bool
}

// NftMarketMetaData contains all meta data concerning the NftMarket contract.
var NftMarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptBid\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"buyMarket\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"cancelBid\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelListOrder\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"editListOrder\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getHighestBid\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structNFTTradeMarket.BidOrder\",\"components\":[{\"name\":\"bidder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getListOrder\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structNFTTradeMarket.ListOrder\",\"components\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"highestBid\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"bidder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"listNFT\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"listOrders\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"matchLimitOrder\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"placeBid\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"BidCancelled\",\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidPlaced\",\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MatchSuccess\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"tradeType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderCancelled\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderEdited\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderListed\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// NftMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use NftMarketMetaData.ABI instead.
var NftMarketABI = NftMarketMetaData.ABI

// NftMarket is an auto generated Go binding around an Ethereum contract.
type NftMarket struct {
	NftMarketCaller     // Read-only binding to the contract
	NftMarketTransactor // Write-only binding to the contract
	NftMarketFilterer   // Log filterer for contract events
}

// NftMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftMarketSession struct {
	Contract     *NftMarket        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftMarketCallerSession struct {
	Contract *NftMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NftMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftMarketTransactorSession struct {
	Contract     *NftMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NftMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftMarketRaw struct {
	Contract *NftMarket // Generic contract binding to access the raw methods on
}

// NftMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftMarketCallerRaw struct {
	Contract *NftMarketCaller // Generic read-only contract binding to access the raw methods on
}

// NftMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftMarketTransactorRaw struct {
	Contract *NftMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNftMarket creates a new instance of NftMarket, bound to a specific deployed contract.
func NewNftMarket(address common.Address, backend bind.ContractBackend) (*NftMarket, error) {
	contract, err := bindNftMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NftMarket{NftMarketCaller: NftMarketCaller{contract: contract}, NftMarketTransactor: NftMarketTransactor{contract: contract}, NftMarketFilterer: NftMarketFilterer{contract: contract}}, nil
}

// NewNftMarketCaller creates a new read-only instance of NftMarket, bound to a specific deployed contract.
func NewNftMarketCaller(address common.Address, caller bind.ContractCaller) (*NftMarketCaller, error) {
	contract, err := bindNftMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftMarketCaller{contract: contract}, nil
}

// NewNftMarketTransactor creates a new write-only instance of NftMarket, bound to a specific deployed contract.
func NewNftMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*NftMarketTransactor, error) {
	contract, err := bindNftMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftMarketTransactor{contract: contract}, nil
}

// NewNftMarketFilterer creates a new log filterer instance of NftMarket, bound to a specific deployed contract.
func NewNftMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*NftMarketFilterer, error) {
	contract, err := bindNftMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftMarketFilterer{contract: contract}, nil
}

// bindNftMarket binds a generic wrapper to an already deployed contract.
func bindNftMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NftMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftMarket *NftMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftMarket.Contract.NftMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftMarket *NftMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftMarket.Contract.NftMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftMarket *NftMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftMarket.Contract.NftMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftMarket *NftMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftMarket *NftMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftMarket *NftMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftMarket.Contract.contract.Transact(opts, method, params...)
}

// GetHighestBid is a free data retrieval call binding the contract method 0x8226508d.
//
// Solidity: function getHighestBid(address nftContract, uint256 tokenId) view returns((address,address,uint256,uint256,bool))
func (_NftMarket *NftMarketCaller) GetHighestBid(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (NFTTradeMarketBidOrder, error) {
	var out []interface{}
	err := _NftMarket.contract.Call(opts, &out, "getHighestBid", nftContract, tokenId)

	if err != nil {
		return *new(NFTTradeMarketBidOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(NFTTradeMarketBidOrder)).(*NFTTradeMarketBidOrder)

	return out0, err

}

// GetHighestBid is a free data retrieval call binding the contract method 0x8226508d.
//
// Solidity: function getHighestBid(address nftContract, uint256 tokenId) view returns((address,address,uint256,uint256,bool))
func (_NftMarket *NftMarketSession) GetHighestBid(nftContract common.Address, tokenId *big.Int) (NFTTradeMarketBidOrder, error) {
	return _NftMarket.Contract.GetHighestBid(&_NftMarket.CallOpts, nftContract, tokenId)
}

// GetHighestBid is a free data retrieval call binding the contract method 0x8226508d.
//
// Solidity: function getHighestBid(address nftContract, uint256 tokenId) view returns((address,address,uint256,uint256,bool))
func (_NftMarket *NftMarketCallerSession) GetHighestBid(nftContract common.Address, tokenId *big.Int) (NFTTradeMarketBidOrder, error) {
	return _NftMarket.Contract.GetHighestBid(&_NftMarket.CallOpts, nftContract, tokenId)
}

// GetListOrder is a free data retrieval call binding the contract method 0xa00a74c0.
//
// Solidity: function getListOrder(address nftContract, uint256 tokenId) view returns((address,address,uint256,uint256,bool))
func (_NftMarket *NftMarketCaller) GetListOrder(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (NFTTradeMarketListOrder, error) {
	var out []interface{}
	err := _NftMarket.contract.Call(opts, &out, "getListOrder", nftContract, tokenId)

	if err != nil {
		return *new(NFTTradeMarketListOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(NFTTradeMarketListOrder)).(*NFTTradeMarketListOrder)

	return out0, err

}

// GetListOrder is a free data retrieval call binding the contract method 0xa00a74c0.
//
// Solidity: function getListOrder(address nftContract, uint256 tokenId) view returns((address,address,uint256,uint256,bool))
func (_NftMarket *NftMarketSession) GetListOrder(nftContract common.Address, tokenId *big.Int) (NFTTradeMarketListOrder, error) {
	return _NftMarket.Contract.GetListOrder(&_NftMarket.CallOpts, nftContract, tokenId)
}

// GetListOrder is a free data retrieval call binding the contract method 0xa00a74c0.
//
// Solidity: function getListOrder(address nftContract, uint256 tokenId) view returns((address,address,uint256,uint256,bool))
func (_NftMarket *NftMarketCallerSession) GetListOrder(nftContract common.Address, tokenId *big.Int) (NFTTradeMarketListOrder, error) {
	return _NftMarket.Contract.GetListOrder(&_NftMarket.CallOpts, nftContract, tokenId)
}

// HighestBid is a free data retrieval call binding the contract method 0x72aa0240.
//
// Solidity: function highestBid(address , uint256 ) view returns(address bidder, address nftContract, uint256 tokenId, uint256 price, bool active)
func (_NftMarket *NftMarketCaller) HighestBid(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	NftContract common.Address
	TokenId     *big.Int
	Price       *big.Int
	Active      bool
}, error) {
	var out []interface{}
	err := _NftMarket.contract.Call(opts, &out, "highestBid", arg0, arg1)

	outstruct := new(struct {
		Bidder      common.Address
		NftContract common.Address
		TokenId     *big.Int
		Price       *big.Int
		Active      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bidder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// HighestBid is a free data retrieval call binding the contract method 0x72aa0240.
//
// Solidity: function highestBid(address , uint256 ) view returns(address bidder, address nftContract, uint256 tokenId, uint256 price, bool active)
func (_NftMarket *NftMarketSession) HighestBid(arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	NftContract common.Address
	TokenId     *big.Int
	Price       *big.Int
	Active      bool
}, error) {
	return _NftMarket.Contract.HighestBid(&_NftMarket.CallOpts, arg0, arg1)
}

// HighestBid is a free data retrieval call binding the contract method 0x72aa0240.
//
// Solidity: function highestBid(address , uint256 ) view returns(address bidder, address nftContract, uint256 tokenId, uint256 price, bool active)
func (_NftMarket *NftMarketCallerSession) HighestBid(arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	NftContract common.Address
	TokenId     *big.Int
	Price       *big.Int
	Active      bool
}, error) {
	return _NftMarket.Contract.HighestBid(&_NftMarket.CallOpts, arg0, arg1)
}

// ListOrders is a free data retrieval call binding the contract method 0x24e297cc.
//
// Solidity: function listOrders(address , uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 price, bool active)
func (_NftMarket *NftMarketCaller) ListOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Seller      common.Address
	NftContract common.Address
	TokenId     *big.Int
	Price       *big.Int
	Active      bool
}, error) {
	var out []interface{}
	err := _NftMarket.contract.Call(opts, &out, "listOrders", arg0, arg1)

	outstruct := new(struct {
		Seller      common.Address
		NftContract common.Address
		TokenId     *big.Int
		Price       *big.Int
		Active      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// ListOrders is a free data retrieval call binding the contract method 0x24e297cc.
//
// Solidity: function listOrders(address , uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 price, bool active)
func (_NftMarket *NftMarketSession) ListOrders(arg0 common.Address, arg1 *big.Int) (struct {
	Seller      common.Address
	NftContract common.Address
	TokenId     *big.Int
	Price       *big.Int
	Active      bool
}, error) {
	return _NftMarket.Contract.ListOrders(&_NftMarket.CallOpts, arg0, arg1)
}

// ListOrders is a free data retrieval call binding the contract method 0x24e297cc.
//
// Solidity: function listOrders(address , uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 price, bool active)
func (_NftMarket *NftMarketCallerSession) ListOrders(arg0 common.Address, arg1 *big.Int) (struct {
	Seller      common.Address
	NftContract common.Address
	TokenId     *big.Int
	Price       *big.Int
	Active      bool
}, error) {
	return _NftMarket.Contract.ListOrders(&_NftMarket.CallOpts, arg0, arg1)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_NftMarket *NftMarketCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _NftMarket.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_NftMarket *NftMarketSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _NftMarket.Contract.OnERC721Received(&_NftMarket.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_NftMarket *NftMarketCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _NftMarket.Contract.OnERC721Received(&_NftMarket.CallOpts, arg0, arg1, arg2, arg3)
}

// AcceptBid is a paid mutator transaction binding the contract method 0x955a5a76.
//
// Solidity: function acceptBid(address nftContract, uint256 tokenId) returns()
func (_NftMarket *NftMarketTransactor) AcceptBid(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.contract.Transact(opts, "acceptBid", nftContract, tokenId)
}

// AcceptBid is a paid mutator transaction binding the contract method 0x955a5a76.
//
// Solidity: function acceptBid(address nftContract, uint256 tokenId) returns()
func (_NftMarket *NftMarketSession) AcceptBid(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.AcceptBid(&_NftMarket.TransactOpts, nftContract, tokenId)
}

// AcceptBid is a paid mutator transaction binding the contract method 0x955a5a76.
//
// Solidity: function acceptBid(address nftContract, uint256 tokenId) returns()
func (_NftMarket *NftMarketTransactorSession) AcceptBid(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.AcceptBid(&_NftMarket.TransactOpts, nftContract, tokenId)
}

// BuyMarket is a paid mutator transaction binding the contract method 0xa63d7859.
//
// Solidity: function buyMarket(address nftContract, uint256 tokenId) payable returns()
func (_NftMarket *NftMarketTransactor) BuyMarket(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.contract.Transact(opts, "buyMarket", nftContract, tokenId)
}

// BuyMarket is a paid mutator transaction binding the contract method 0xa63d7859.
//
// Solidity: function buyMarket(address nftContract, uint256 tokenId) payable returns()
func (_NftMarket *NftMarketSession) BuyMarket(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.BuyMarket(&_NftMarket.TransactOpts, nftContract, tokenId)
}

// BuyMarket is a paid mutator transaction binding the contract method 0xa63d7859.
//
// Solidity: function buyMarket(address nftContract, uint256 tokenId) payable returns()
func (_NftMarket *NftMarketTransactorSession) BuyMarket(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.BuyMarket(&_NftMarket.TransactOpts, nftContract, tokenId)
}

// CancelBid is a paid mutator transaction binding the contract method 0x39b6b1e5.
//
// Solidity: function cancelBid(address nftContract, uint256 tokenId) returns()
func (_NftMarket *NftMarketTransactor) CancelBid(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.contract.Transact(opts, "cancelBid", nftContract, tokenId)
}

// CancelBid is a paid mutator transaction binding the contract method 0x39b6b1e5.
//
// Solidity: function cancelBid(address nftContract, uint256 tokenId) returns()
func (_NftMarket *NftMarketSession) CancelBid(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.CancelBid(&_NftMarket.TransactOpts, nftContract, tokenId)
}

// CancelBid is a paid mutator transaction binding the contract method 0x39b6b1e5.
//
// Solidity: function cancelBid(address nftContract, uint256 tokenId) returns()
func (_NftMarket *NftMarketTransactorSession) CancelBid(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.CancelBid(&_NftMarket.TransactOpts, nftContract, tokenId)
}

// CancelListOrder is a paid mutator transaction binding the contract method 0x664fd551.
//
// Solidity: function cancelListOrder(address nftContract, uint256 tokenId) returns()
func (_NftMarket *NftMarketTransactor) CancelListOrder(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.contract.Transact(opts, "cancelListOrder", nftContract, tokenId)
}

// CancelListOrder is a paid mutator transaction binding the contract method 0x664fd551.
//
// Solidity: function cancelListOrder(address nftContract, uint256 tokenId) returns()
func (_NftMarket *NftMarketSession) CancelListOrder(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.CancelListOrder(&_NftMarket.TransactOpts, nftContract, tokenId)
}

// CancelListOrder is a paid mutator transaction binding the contract method 0x664fd551.
//
// Solidity: function cancelListOrder(address nftContract, uint256 tokenId) returns()
func (_NftMarket *NftMarketTransactorSession) CancelListOrder(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.CancelListOrder(&_NftMarket.TransactOpts, nftContract, tokenId)
}

// EditListOrder is a paid mutator transaction binding the contract method 0x323ef340.
//
// Solidity: function editListOrder(address nftContract, uint256 tokenId, uint256 newPrice) returns()
func (_NftMarket *NftMarketTransactor) EditListOrder(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _NftMarket.contract.Transact(opts, "editListOrder", nftContract, tokenId, newPrice)
}

// EditListOrder is a paid mutator transaction binding the contract method 0x323ef340.
//
// Solidity: function editListOrder(address nftContract, uint256 tokenId, uint256 newPrice) returns()
func (_NftMarket *NftMarketSession) EditListOrder(nftContract common.Address, tokenId *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.EditListOrder(&_NftMarket.TransactOpts, nftContract, tokenId, newPrice)
}

// EditListOrder is a paid mutator transaction binding the contract method 0x323ef340.
//
// Solidity: function editListOrder(address nftContract, uint256 tokenId, uint256 newPrice) returns()
func (_NftMarket *NftMarketTransactorSession) EditListOrder(nftContract common.Address, tokenId *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.EditListOrder(&_NftMarket.TransactOpts, nftContract, tokenId, newPrice)
}

// ListNFT is a paid mutator transaction binding the contract method 0xad05f1b4.
//
// Solidity: function listNFT(address nftContract, uint256 tokenId, uint256 price) returns()
func (_NftMarket *NftMarketTransactor) ListNFT(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _NftMarket.contract.Transact(opts, "listNFT", nftContract, tokenId, price)
}

// ListNFT is a paid mutator transaction binding the contract method 0xad05f1b4.
//
// Solidity: function listNFT(address nftContract, uint256 tokenId, uint256 price) returns()
func (_NftMarket *NftMarketSession) ListNFT(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.ListNFT(&_NftMarket.TransactOpts, nftContract, tokenId, price)
}

// ListNFT is a paid mutator transaction binding the contract method 0xad05f1b4.
//
// Solidity: function listNFT(address nftContract, uint256 tokenId, uint256 price) returns()
func (_NftMarket *NftMarketTransactorSession) ListNFT(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.ListNFT(&_NftMarket.TransactOpts, nftContract, tokenId, price)
}

// MatchLimitOrder is a paid mutator transaction binding the contract method 0xda506609.
//
// Solidity: function matchLimitOrder(address nftContract, uint256 tokenId) returns()
func (_NftMarket *NftMarketTransactor) MatchLimitOrder(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.contract.Transact(opts, "matchLimitOrder", nftContract, tokenId)
}

// MatchLimitOrder is a paid mutator transaction binding the contract method 0xda506609.
//
// Solidity: function matchLimitOrder(address nftContract, uint256 tokenId) returns()
func (_NftMarket *NftMarketSession) MatchLimitOrder(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.MatchLimitOrder(&_NftMarket.TransactOpts, nftContract, tokenId)
}

// MatchLimitOrder is a paid mutator transaction binding the contract method 0xda506609.
//
// Solidity: function matchLimitOrder(address nftContract, uint256 tokenId) returns()
func (_NftMarket *NftMarketTransactorSession) MatchLimitOrder(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.MatchLimitOrder(&_NftMarket.TransactOpts, nftContract, tokenId)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xd98b9bb5.
//
// Solidity: function placeBid(address nftContract, uint256 tokenId) payable returns()
func (_NftMarket *NftMarketTransactor) PlaceBid(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.contract.Transact(opts, "placeBid", nftContract, tokenId)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xd98b9bb5.
//
// Solidity: function placeBid(address nftContract, uint256 tokenId) payable returns()
func (_NftMarket *NftMarketSession) PlaceBid(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.PlaceBid(&_NftMarket.TransactOpts, nftContract, tokenId)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xd98b9bb5.
//
// Solidity: function placeBid(address nftContract, uint256 tokenId) payable returns()
func (_NftMarket *NftMarketTransactorSession) PlaceBid(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftMarket.Contract.PlaceBid(&_NftMarket.TransactOpts, nftContract, tokenId)
}

// NftMarketBidCancelledIterator is returned from FilterBidCancelled and is used to iterate over the raw logs and unpacked data for BidCancelled events raised by the NftMarket contract.
type NftMarketBidCancelledIterator struct {
	Event *NftMarketBidCancelled // Event containing the contract specifics and raw log

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
func (it *NftMarketBidCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftMarketBidCancelled)
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
		it.Event = new(NftMarketBidCancelled)
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
func (it *NftMarketBidCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftMarketBidCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftMarketBidCancelled represents a BidCancelled event raised by the NftMarket contract.
type NftMarketBidCancelled struct {
	Bidder      common.Address
	NftContract common.Address
	TokenId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBidCancelled is a free log retrieval operation binding the contract event 0x9db7b8c3b2c40d21fc215cbb9be237a63b2fdf404ef91029c3955951a968e5b7.
//
// Solidity: event BidCancelled(address indexed bidder, address indexed nftContract, uint256 indexed tokenId)
func (_NftMarket *NftMarketFilterer) FilterBidCancelled(opts *bind.FilterOpts, bidder []common.Address, nftContract []common.Address, tokenId []*big.Int) (*NftMarketBidCancelledIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftMarket.contract.FilterLogs(opts, "BidCancelled", bidderRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftMarketBidCancelledIterator{contract: _NftMarket.contract, event: "BidCancelled", logs: logs, sub: sub}, nil
}

// WatchBidCancelled is a free log subscription operation binding the contract event 0x9db7b8c3b2c40d21fc215cbb9be237a63b2fdf404ef91029c3955951a968e5b7.
//
// Solidity: event BidCancelled(address indexed bidder, address indexed nftContract, uint256 indexed tokenId)
func (_NftMarket *NftMarketFilterer) WatchBidCancelled(opts *bind.WatchOpts, sink chan<- *NftMarketBidCancelled, bidder []common.Address, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftMarket.contract.WatchLogs(opts, "BidCancelled", bidderRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftMarketBidCancelled)
				if err := _NftMarket.contract.UnpackLog(event, "BidCancelled", log); err != nil {
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

// ParseBidCancelled is a log parse operation binding the contract event 0x9db7b8c3b2c40d21fc215cbb9be237a63b2fdf404ef91029c3955951a968e5b7.
//
// Solidity: event BidCancelled(address indexed bidder, address indexed nftContract, uint256 indexed tokenId)
func (_NftMarket *NftMarketFilterer) ParseBidCancelled(log types.Log) (*NftMarketBidCancelled, error) {
	event := new(NftMarketBidCancelled)
	if err := _NftMarket.contract.UnpackLog(event, "BidCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftMarketBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the NftMarket contract.
type NftMarketBidPlacedIterator struct {
	Event *NftMarketBidPlaced // Event containing the contract specifics and raw log

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
func (it *NftMarketBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftMarketBidPlaced)
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
		it.Event = new(NftMarketBidPlaced)
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
func (it *NftMarketBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftMarketBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftMarketBidPlaced represents a BidPlaced event raised by the NftMarket contract.
type NftMarketBidPlaced struct {
	Bidder      common.Address
	NftContract common.Address
	TokenId     *big.Int
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0xdd49bbb40d47a514dddcd458e9718364143bc24a0cca58439ee6f4f45e4ce10d.
//
// Solidity: event BidPlaced(address indexed bidder, address indexed nftContract, uint256 indexed tokenId, uint256 price)
func (_NftMarket *NftMarketFilterer) FilterBidPlaced(opts *bind.FilterOpts, bidder []common.Address, nftContract []common.Address, tokenId []*big.Int) (*NftMarketBidPlacedIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftMarket.contract.FilterLogs(opts, "BidPlaced", bidderRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftMarketBidPlacedIterator{contract: _NftMarket.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0xdd49bbb40d47a514dddcd458e9718364143bc24a0cca58439ee6f4f45e4ce10d.
//
// Solidity: event BidPlaced(address indexed bidder, address indexed nftContract, uint256 indexed tokenId, uint256 price)
func (_NftMarket *NftMarketFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *NftMarketBidPlaced, bidder []common.Address, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftMarket.contract.WatchLogs(opts, "BidPlaced", bidderRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftMarketBidPlaced)
				if err := _NftMarket.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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

// ParseBidPlaced is a log parse operation binding the contract event 0xdd49bbb40d47a514dddcd458e9718364143bc24a0cca58439ee6f4f45e4ce10d.
//
// Solidity: event BidPlaced(address indexed bidder, address indexed nftContract, uint256 indexed tokenId, uint256 price)
func (_NftMarket *NftMarketFilterer) ParseBidPlaced(log types.Log) (*NftMarketBidPlaced, error) {
	event := new(NftMarketBidPlaced)
	if err := _NftMarket.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftMarketMatchSuccessIterator is returned from FilterMatchSuccess and is used to iterate over the raw logs and unpacked data for MatchSuccess events raised by the NftMarket contract.
type NftMarketMatchSuccessIterator struct {
	Event *NftMarketMatchSuccess // Event containing the contract specifics and raw log

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
func (it *NftMarketMatchSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftMarketMatchSuccess)
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
		it.Event = new(NftMarketMatchSuccess)
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
func (it *NftMarketMatchSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftMarketMatchSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftMarketMatchSuccess represents a MatchSuccess event raised by the NftMarket contract.
type NftMarketMatchSuccess struct {
	Seller      common.Address
	Buyer       common.Address
	NftContract common.Address
	TokenId     *big.Int
	Price       *big.Int
	TradeType   uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMatchSuccess is a free log retrieval operation binding the contract event 0x7e769bd821f1c5a635e0bd4c3654e8630e529600e771ec4b9018b0c0bc5a928d.
//
// Solidity: event MatchSuccess(address indexed seller, address indexed buyer, address indexed nftContract, uint256 tokenId, uint256 price, uint8 tradeType)
func (_NftMarket *NftMarketFilterer) FilterMatchSuccess(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, nftContract []common.Address) (*NftMarketMatchSuccessIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _NftMarket.contract.FilterLogs(opts, "MatchSuccess", sellerRule, buyerRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return &NftMarketMatchSuccessIterator{contract: _NftMarket.contract, event: "MatchSuccess", logs: logs, sub: sub}, nil
}

// WatchMatchSuccess is a free log subscription operation binding the contract event 0x7e769bd821f1c5a635e0bd4c3654e8630e529600e771ec4b9018b0c0bc5a928d.
//
// Solidity: event MatchSuccess(address indexed seller, address indexed buyer, address indexed nftContract, uint256 tokenId, uint256 price, uint8 tradeType)
func (_NftMarket *NftMarketFilterer) WatchMatchSuccess(opts *bind.WatchOpts, sink chan<- *NftMarketMatchSuccess, seller []common.Address, buyer []common.Address, nftContract []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _NftMarket.contract.WatchLogs(opts, "MatchSuccess", sellerRule, buyerRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftMarketMatchSuccess)
				if err := _NftMarket.contract.UnpackLog(event, "MatchSuccess", log); err != nil {
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

// ParseMatchSuccess is a log parse operation binding the contract event 0x7e769bd821f1c5a635e0bd4c3654e8630e529600e771ec4b9018b0c0bc5a928d.
//
// Solidity: event MatchSuccess(address indexed seller, address indexed buyer, address indexed nftContract, uint256 tokenId, uint256 price, uint8 tradeType)
func (_NftMarket *NftMarketFilterer) ParseMatchSuccess(log types.Log) (*NftMarketMatchSuccess, error) {
	event := new(NftMarketMatchSuccess)
	if err := _NftMarket.contract.UnpackLog(event, "MatchSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftMarketOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the NftMarket contract.
type NftMarketOrderCancelledIterator struct {
	Event *NftMarketOrderCancelled // Event containing the contract specifics and raw log

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
func (it *NftMarketOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftMarketOrderCancelled)
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
		it.Event = new(NftMarketOrderCancelled)
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
func (it *NftMarketOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftMarketOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftMarketOrderCancelled represents a OrderCancelled event raised by the NftMarket contract.
type NftMarketOrderCancelled struct {
	Seller      common.Address
	NftContract common.Address
	TokenId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0xd23ac4476ccf3eeb78c76d4b7bb113597bcd16c15d47656627ea6221d685ad73.
//
// Solidity: event OrderCancelled(address indexed seller, address indexed nftContract, uint256 indexed tokenId)
func (_NftMarket *NftMarketFilterer) FilterOrderCancelled(opts *bind.FilterOpts, seller []common.Address, nftContract []common.Address, tokenId []*big.Int) (*NftMarketOrderCancelledIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftMarket.contract.FilterLogs(opts, "OrderCancelled", sellerRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftMarketOrderCancelledIterator{contract: _NftMarket.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0xd23ac4476ccf3eeb78c76d4b7bb113597bcd16c15d47656627ea6221d685ad73.
//
// Solidity: event OrderCancelled(address indexed seller, address indexed nftContract, uint256 indexed tokenId)
func (_NftMarket *NftMarketFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *NftMarketOrderCancelled, seller []common.Address, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftMarket.contract.WatchLogs(opts, "OrderCancelled", sellerRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftMarketOrderCancelled)
				if err := _NftMarket.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0xd23ac4476ccf3eeb78c76d4b7bb113597bcd16c15d47656627ea6221d685ad73.
//
// Solidity: event OrderCancelled(address indexed seller, address indexed nftContract, uint256 indexed tokenId)
func (_NftMarket *NftMarketFilterer) ParseOrderCancelled(log types.Log) (*NftMarketOrderCancelled, error) {
	event := new(NftMarketOrderCancelled)
	if err := _NftMarket.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftMarketOrderEditedIterator is returned from FilterOrderEdited and is used to iterate over the raw logs and unpacked data for OrderEdited events raised by the NftMarket contract.
type NftMarketOrderEditedIterator struct {
	Event *NftMarketOrderEdited // Event containing the contract specifics and raw log

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
func (it *NftMarketOrderEditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftMarketOrderEdited)
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
		it.Event = new(NftMarketOrderEdited)
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
func (it *NftMarketOrderEditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftMarketOrderEditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftMarketOrderEdited represents a OrderEdited event raised by the NftMarket contract.
type NftMarketOrderEdited struct {
	Seller      common.Address
	NftContract common.Address
	TokenId     *big.Int
	NewPrice    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrderEdited is a free log retrieval operation binding the contract event 0x22c66d322025d97a762010bee890db9b4c7aa47fac636a660ffb905bf9ff6927.
//
// Solidity: event OrderEdited(address indexed seller, address indexed nftContract, uint256 indexed tokenId, uint256 newPrice)
func (_NftMarket *NftMarketFilterer) FilterOrderEdited(opts *bind.FilterOpts, seller []common.Address, nftContract []common.Address, tokenId []*big.Int) (*NftMarketOrderEditedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftMarket.contract.FilterLogs(opts, "OrderEdited", sellerRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftMarketOrderEditedIterator{contract: _NftMarket.contract, event: "OrderEdited", logs: logs, sub: sub}, nil
}

// WatchOrderEdited is a free log subscription operation binding the contract event 0x22c66d322025d97a762010bee890db9b4c7aa47fac636a660ffb905bf9ff6927.
//
// Solidity: event OrderEdited(address indexed seller, address indexed nftContract, uint256 indexed tokenId, uint256 newPrice)
func (_NftMarket *NftMarketFilterer) WatchOrderEdited(opts *bind.WatchOpts, sink chan<- *NftMarketOrderEdited, seller []common.Address, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftMarket.contract.WatchLogs(opts, "OrderEdited", sellerRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftMarketOrderEdited)
				if err := _NftMarket.contract.UnpackLog(event, "OrderEdited", log); err != nil {
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

// ParseOrderEdited is a log parse operation binding the contract event 0x22c66d322025d97a762010bee890db9b4c7aa47fac636a660ffb905bf9ff6927.
//
// Solidity: event OrderEdited(address indexed seller, address indexed nftContract, uint256 indexed tokenId, uint256 newPrice)
func (_NftMarket *NftMarketFilterer) ParseOrderEdited(log types.Log) (*NftMarketOrderEdited, error) {
	event := new(NftMarketOrderEdited)
	if err := _NftMarket.contract.UnpackLog(event, "OrderEdited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftMarketOrderListedIterator is returned from FilterOrderListed and is used to iterate over the raw logs and unpacked data for OrderListed events raised by the NftMarket contract.
type NftMarketOrderListedIterator struct {
	Event *NftMarketOrderListed // Event containing the contract specifics and raw log

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
func (it *NftMarketOrderListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftMarketOrderListed)
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
		it.Event = new(NftMarketOrderListed)
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
func (it *NftMarketOrderListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftMarketOrderListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftMarketOrderListed represents a OrderListed event raised by the NftMarket contract.
type NftMarketOrderListed struct {
	Seller      common.Address
	NftContract common.Address
	TokenId     *big.Int
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrderListed is a free log retrieval operation binding the contract event 0xb4aa437e32593e335dd4eb6069bc5e3225445ac88d9e797c5b9b99c404d9dcdd.
//
// Solidity: event OrderListed(address indexed seller, address indexed nftContract, uint256 indexed tokenId, uint256 price)
func (_NftMarket *NftMarketFilterer) FilterOrderListed(opts *bind.FilterOpts, seller []common.Address, nftContract []common.Address, tokenId []*big.Int) (*NftMarketOrderListedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftMarket.contract.FilterLogs(opts, "OrderListed", sellerRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftMarketOrderListedIterator{contract: _NftMarket.contract, event: "OrderListed", logs: logs, sub: sub}, nil
}

// WatchOrderListed is a free log subscription operation binding the contract event 0xb4aa437e32593e335dd4eb6069bc5e3225445ac88d9e797c5b9b99c404d9dcdd.
//
// Solidity: event OrderListed(address indexed seller, address indexed nftContract, uint256 indexed tokenId, uint256 price)
func (_NftMarket *NftMarketFilterer) WatchOrderListed(opts *bind.WatchOpts, sink chan<- *NftMarketOrderListed, seller []common.Address, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftMarket.contract.WatchLogs(opts, "OrderListed", sellerRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftMarketOrderListed)
				if err := _NftMarket.contract.UnpackLog(event, "OrderListed", log); err != nil {
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

// ParseOrderListed is a log parse operation binding the contract event 0xb4aa437e32593e335dd4eb6069bc5e3225445ac88d9e797c5b9b99c404d9dcdd.
//
// Solidity: event OrderListed(address indexed seller, address indexed nftContract, uint256 indexed tokenId, uint256 price)
func (_NftMarket *NftMarketFilterer) ParseOrderListed(log types.Log) (*NftMarketOrderListed, error) {
	event := new(NftMarketOrderListed)
	if err := _NftMarket.contract.UnpackLog(event, "OrderListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

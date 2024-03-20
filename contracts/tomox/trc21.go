package titanx

import (
	"github.com/titanchain/titanchain/accounts/abi/bind"
	"github.com/titanchain/titanchain/common"
	"github.com/titanchain/titanchain/contracts/titanx/contract"
	"math/big"
)

type MyTRC21 struct {
	*contract.MyTRC21Session
	contractBackend bind.ContractBackend
}

func NewTRC21(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*MyTRC21, error) {
	smartContract, err := contract.NewMyTRC21(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &MyTRC21{
		&contract.MyTRC21Session{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployTRC21(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, owners []common.Address, required *big.Int, name string, symbol string, decimals uint8, cap, fee, depositFee, withdrawFee *big.Int) (common.Address, *MyTRC21, error) {
	contractAddr, _, _, err := contract.DeployMyTRC21(transactOpts, contractBackend, owners, required, name, symbol, decimals, cap, fee, depositFee, withdrawFee)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewTRC21(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}

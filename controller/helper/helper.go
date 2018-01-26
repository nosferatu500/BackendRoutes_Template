package helper

import (
	"BackendRoutes_Template/gocontracts"
	"BackendRoutes_Template/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateAuth() (*bind.TransactOpts, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactor(key), err
}

func CreateSession(auth *bind.TransactOpts, connection bind.ContractBackend) (*token.PowerTokenSession, error) {
	contract, err := token.NewPowerToken(common.HexToAddress(model.ContractAddress), connection)
	if err != nil {
		return nil, err
	}

	session := token.PowerTokenSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    auth.From,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: model.Gaslimit,
		},
	}
	return &session, err
}

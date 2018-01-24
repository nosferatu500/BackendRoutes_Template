package helper

import (
	"BackendRoutes_Template/gocontracts"
	"BackendRoutes_Template/jwt"
	"BackendRoutes_Template/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"strings"
)

func GenerateAuth() (*bind.TransactOpts, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactor(key), err
}

func CreateAuth(token string) (*bind.TransactOpts, error) {
	privateKey := strings.TrimPrefix(token, "0x")

	key, err := crypto.HexToECDSA(privateKey)
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

func ParseContext(ctx context.Context) (*model.ApiCustomClaims, error) {
	tok := ctx.Value("access_token")
	if tok == nil {
		return nil, errors.New("access_token is null")
	}
	claims, err, valid := jwt.ParseToken(tok.(string))
	if err != nil || !valid {
		return nil, errors.New("access_token not valid")
	}
	return claims, nil
}

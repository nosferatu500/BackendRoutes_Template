package action

import (
	auth_gw "BackendRoutes_Template/api/auth"
	token_gw "BackendRoutes_Template/api/token"
	"BackendRoutes_Template/controller/helper"
	jwt_manager "BackendRoutes_Template/jwt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"log"
	"strings"
)

type Rpc struct {
	Connect bind.ContractBackend
}

func (s *Rpc) LoginKey(ctx context.Context, in *auth_gw.LoginKeyReq) (*auth_gw.LoginResp, error) {
	if in.Key == "" {
		log.Printf("Error LoginKey: %s", errors.Errorf("Key is null").Error())
		return nil, errors.Errorf("Error LoginKey Key is null")
	}
	privateKey := strings.TrimPrefix(in.Key, "0x")

	_, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Printf("Error LoginKey HexToECDSA: %s", err.Error())
		return nil, err
	}

	token, err := jwt_manager.CreateToken(privateKey)
	if err != nil {
		log.Printf("Error LoginKey jwt_manager.CreateToken: %s", err.Error())
		return nil, err
	}

	return &auth_gw.LoginResp{Token: token}, nil
}

func (s *Rpc) GetName(ctx context.Context, in *token_gw.DummyReq) (*token_gw.GetNameResp, error) {
	auth, err := helper.GenerateAuth()
	if err != nil {
		log.Printf("Error GetName GenerateAuth: %s\n", err.Error())
		return nil, err
	}

	session, err := helper.CreateSession(auth, s.Connect)
	if err != nil {
		log.Printf("Error GetName CreateSession: %s\n", err.Error())
		return nil, err
	}

	result, err := session.Name()
	if err != nil {
		log.Printf("Error GetName session.Name: %s\n", err.Error())
		return nil, err
	}

	return &token_gw.GetNameResp{Result: result}, nil
}

func (s *Rpc) GetSymbol(ctx context.Context, in *token_gw.DummyReq) (*token_gw.GetSymbolResp, error) {
	claims, err := helper.ParseContext(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := helper.CreateAuth(claims.AccessToken)
	if err != nil {
		log.Printf("Error GetSymbol CreateAuth: %s\n", err.Error())
		return nil, err
	}

	session, err := helper.CreateSession(auth, s.Connect)
	if err != nil {
		log.Printf("Error GetSymbol CreateSession: %s\n", err.Error())
		return nil, err
	}

	result, err := session.Symbol()
	if err != nil {
		log.Printf("Error GetSymbol session.Name: %s\n", err.Error())
		return nil, err
	}

	return &token_gw.GetSymbolResp{Result: result}, nil
}

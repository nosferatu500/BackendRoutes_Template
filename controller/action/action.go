package action

import (
	token_gw "BackendRoutes_Template/api/token"
	"BackendRoutes_Template/controller/helper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"golang.org/x/net/context"
	"log"
)

type Rpc struct {
	Connect bind.ContractBackend
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
	auth, err := helper.GenerateAuth()
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

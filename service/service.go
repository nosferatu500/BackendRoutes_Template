package service

import (
	auth_gw "BackendRoutes_Template/api/auth"
	token_gw "BackendRoutes_Template/api/token"
	"BackendRoutes_Template/controller/action"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	natssrv "github.com/nats-io/gnatsd/server"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	Connect bind.ContractBackend
}

func authfunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	newCtx := context.WithValue(ctx, "access_token", token)
	return newCtx, nil
}

func BrokerSrv() *natssrv.Server {
	opts := natssrv.Options{
		HTTPPort: 8222,
	}
	return natssrv.New(&opts)
}

func BrokerSrvStart(broker *natssrv.Server) {
	broker.Start()
}

func (srv *Server) StartGRPCServer(grpc_port int) {

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(grpc_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_auth.StreamServerInterceptor(
				authfunc,
			),
		),
		grpc.UnaryInterceptor(
			grpc_auth.UnaryServerInterceptor(
				authfunc,
			),
		),
	)

	auth_gw.RegisterAuthServer(
		s,
		&action.Rpc{Connect: srv.Connect},
	)

	token_gw.RegisterTokenServer(
		s,
		&action.Rpc{Connect: srv.Connect},
	)

	bro := BrokerSrv()
	go BrokerSrvStart(bro)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (srv *Server) StartRESTGatewey(rest_port, grpc_port int) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := token_gw.RegisterTokenHandlerFromEndpoint(ctx, mux, ":"+strconv.Itoa(grpc_port), opts)
	if err != nil {
		log.Fatalf("Fatal RegisterTokenHandlerFromEndpoint: %s", err.Error())
	}

	err = auth_gw.RegisterAuthHandlerFromEndpoint(ctx, mux, ":"+strconv.Itoa(grpc_port), opts)
	if err != nil {
		log.Fatalf("Fatal RegisterAuthHandlerFromEndpoint: %s", err.Error())
	}

	if err := http.ListenAndServe(":"+strconv.Itoa(rest_port), handlers.LoggingHandler(os.Stdout, mux)); err != nil {
		log.Fatalf("Fatal ListenAndServe: %s", err.Error())
	}

}

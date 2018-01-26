package service

import (
	token_gw "BackendRoutes_Template/api/token"
	"BackendRoutes_Template/controller/action"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gorilla/handlers"
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

	s := grpc.NewServer()

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

	if err := http.ListenAndServe(":"+strconv.Itoa(rest_port), handlers.LoggingHandler(os.Stdout, mux)); err != nil {
		log.Fatalf("Fatal ListenAndServe: %s", err.Error())
	}

}

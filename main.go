//go:generate abigen --sol ./contracts/TestToken.sol --pkg token  --out ./gocontracts/TestToken.go

//go:generate protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --go_out=plugins=grpc:.                 api/auth/auth.proto
//go:generate protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --grpc-gateway_out=logtostderr=true:.   api/auth/auth.proto
//go:generate protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --swagger_out=logtostderr=true:.        api/auth/auth.proto

//go:generate protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --go_out=plugins=grpc:.                 api/token/token.proto
//go:generate protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --grpc-gateway_out=logtostderr=true:.   api/token/token.proto
//go:generate protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --swagger_out=logtostderr=true:.        api/token/token.proto

package main

import (
	"BackendRoutes_Template/controller/helper"
	"BackendRoutes_Template/gocontracts"
	"BackendRoutes_Template/model"
	"BackendRoutes_Template/service"
	"flag"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/golang/glog"
	"os"
	"fmt"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	s := &service.Server{}

	for _, arg := range os.Args {
		if arg == "simulator" {
			s.StartSimulatedConnection()
			auth, _ := helper.CreateAuth(s.Auth)
			address, _, _, _ := token.DeployPowerToken(auth, s.Connect)
			model.ContractAddress = address.String()
			fmt.Println("Simulator mode: on")
			break
		} else {
			fmt.Println("Simulator mode: off")
			s.Connect, _ = ethclient.Dial("https://rinkeby.infura.io/dUQe3kE7aGgdnkBmEAny")
			model.ContractAddress = "0xef40f21f9c84fd5c83070875a76a2b86c35ae7ea"
		}
	}

	go s.StartGRPCServer(9092)

	s.StartRESTGatewey(9091, 9092)
}

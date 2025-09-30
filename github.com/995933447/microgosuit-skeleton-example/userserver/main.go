package main

import (
	"context"
	"fmt"
	"log"

	"github.com/995933447/microgosuit"
	"github.com/995933447/microgosuit-skeleton-example/user"
	"github.com/995933447/microgosuit-skeleton-example/userserver/boot"
	"github.com/995933447/microgosuit-skeleton-example/userserver/handler"
	"github.com/995933447/microgosuit/discovery"

	"google.golang.org/grpc"
)

func main() {
	err := microgosuit.InitSuitWithGrpc(context.TODO(), "", "microgosuit", "microgosuit/")
	if err != nil {
		log.Fatal(err)
	}

	port, err := GetListenRpcPort()
	if err != nil {
		log.Fatal(err)
	}

	if err = boot.OnBoot(); err != nil {
		log.Fatal(err)
	}

	err = microgosuit.ServeGrpc(context.TODO(), &microgosuit.ServeGrpcReq{
		RegDiscoverKeyPrefix: "microgosuit/",
		SrvNames:             ServiceNames,
		IpVar:                "$inner_ip",
		Port:                 port,
		EnabledHealth:        true,
		RegisterCustomServiceServerFunc: func(server *grpc.Server) error {
			user.RegisterUserServer(server, &handler.User{})
			user.RegisterUserGatewayServer(server, &handler.UserGateway{})
			return nil
		},
		OnReady: func(server *grpc.Server, node *discovery.Node) {
			fmt.Printf("up node %s:%d!\n", node.Host, node.Port)
			fmt.Printf(">>>>>>>>>>>>>>> run %+v at 2025-09-29 10:16:12 success <<<<<<<<<<<<<<<", ServiceNames)
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

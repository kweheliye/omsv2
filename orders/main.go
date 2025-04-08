package main

import (
	"context"
	"github.com/kweheliye/omsv2/common"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)
	NewGrpcHandler(grpcServer)

	svc.CreateOrder(context.Background())

	log.Println("starting grpc server", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}

}

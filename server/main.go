package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hello-grpc/common"
	"hello-grpc/protos"
	"log"
	"net"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(context context.Context, in *protos.HelloRequest) (*protos.HelloResponse, error) {
	log.Printf("Message received is: %v", in.Name)
	return &protos.HelloResponse{Message: "Hello " + in.Name}, nil
}
func main() {
	listen, err := net.Listen("tcp", common.GRPCPORT)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()

	//here we need to register a service
	protos.RegisterHelloWorldServer(grpcServer, &server{})
	//reflection.Register(grpcServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

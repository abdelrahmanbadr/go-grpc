package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hello-grpc/common"
	pb "hello-grpc/protos"
	"log"
	"net"
)

// server is used to implement helloworld.GreeterServer.
type Server struct{}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(context context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Message received is: %v", in.Name)
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}
func main() {
	listen, err := net.Listen("tcp", common.GRPCPORT)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHelloWorldServer(grpcServer, &Server{})
	//reflection.Register(grpcServer)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

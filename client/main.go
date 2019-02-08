package main

import (
	"context"
	"hello-grpc/common"
	"log"
	"time"

	"google.golang.org/grpc"
	"hello-grpc/protos"
)

func main() {
	// Set up a connection to the server.
	//by default grpc has ssl
	//but now we don't have ssl certificate, so we will use insecure
	conn, err := grpc.Dial(common.GRPCPORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//service client
	c := protos.NewHelloWorldClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &protos.HelloRequest{Name: "Test"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Message sent from client is : %s", r.Message)
}

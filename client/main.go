package main

import (
	"context"
	"hello-grpc/common"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "hello-grpc/protos"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(common.GRPCPORT, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := pb.NewHelloWorldClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Test"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Message sent from client is : %s", r.Message)
}

package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"

	pb "github.com/uma-co82/go-web-standard"
)

// gRPC通信チェック用

func main() {
	addr := "localhost:50051"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := os.Args[1]

	md := metadata.Pairs("timestamp", time.Now().Format(time.Stamp))

	ctx := context.Background()

	ctx = metadata.NewOutgoingContext(ctx, md)

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name}, grpc.Trailer(&md))

	if err != nil {
		log.Fatalf("could not greeter: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

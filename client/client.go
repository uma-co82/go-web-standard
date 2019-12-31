package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

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
		// FromErrorを呼び出す事でgRPCのStatusに変換
		s, ok := status.FromError(err)
		if ok {
			log.Printf("gRPC Error (message: %s)", s.Message())
			// エラーの詳細情報のスライス
			for _, d := range s.Details() {
				// エラーの型にキャストする事でcaseで個別にハンドリングできる
				switch info := d.(type) {
				case *errdetails.RetryInfo:
					log.Printf("RetryInfo: %v", info)
				}
			}
			os.Exit(1)
		} else {
			log.Fatalf("could not greeter: %v", err)
		}
	}
	log.Printf("Greeting: %s", r.Message)
}

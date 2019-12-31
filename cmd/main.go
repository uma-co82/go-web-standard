package main

// 起動、終了処理

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/uma-co82/go-web-standard"
)

type server struct{}

// 正常系
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	time.Sleep(3 * time.Second)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

// errorパターン
// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	log.Printf("Received: %v", in.Name)
// 	st, _ := status.New(codes.Aborted, "aborted").WithDetails(&errdetails.RetryInfo{
// 		RetryDelay: &duration.Duration{
// 			Seconds: 3,
// 			Nanos:   0,
// 		},
// 	})
// 	return nil, st.Err()
// }

func main() {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("gRPC server listening on " + addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

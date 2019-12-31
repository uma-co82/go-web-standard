package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"

	pb "github.com/uma-co82/go-web-standard"
)

// gRPC通信チェック用

func main() {
	resolver.Register(&exampleResolverBuilder{})

	// loadbarancing
	addr := "testScheme:///example"
	// addr := "localhost:50051"

	// TLS
	// creds, err := credentials.NewClientTLSFromFile("server.crt", "")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// loadbarancing
	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(),
		grpc.WithBalancerName("round_robin"),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds), grpc.WithUnaryInterceptor(unaryInterceptor))

	// grpc.WithInsecure()は認証しない時
	// conn, err := grpc.Dial(addr, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := os.Args[1]

	// metadata
	// md := metadata.Pairs("timestamp", time.Now().Format(time.Stamp))

	ctx := context.Background()
	// キャンセルやタイムアウトを考慮せず、ずっと待ち続ける
	// ctx := context.Background()
	// キャンセル考慮
	// ctx, cancel := context.WithCancel(context.Background())
	// キャンセル実行
	// defer cancel()

	// 1秒後にキャンセル実行
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cancel()
	// }()

	// metadataをheaderにつめる
	// ctx = metadata.NewOutgoingContext(ctx, md)

	// r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name}, grpc.Trailer(&md))
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})

	// errorhandling
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

func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("before call: %s, request: %+v", method, req)
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("after call: %s, response: %+v", method, reply)
	return err
}

type exampleResolverBuilder struct{}
type exampleResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *exampleResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

// interface実装
func (*exampleResolver) ResolveNow(o resolver.ResolveNowOption) {}
func (*exampleResolver) Close()                                 {}

func (*exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	r := &exampleResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			"example": {"localhost:50051", "localhost:50052"},
		},
	}
	r.start()
	return r, nil
}

func (*exampleResolverBuilder) Scheme() string { return "testScheme" }

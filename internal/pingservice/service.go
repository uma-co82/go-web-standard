package service

import (
	"context"
	"fmt"

	"github.com/uma-co82/go-web-standard/proto/ping"
)

type PingService struct {
}

func (s *PingService) Hello(ctx context.Context, req *ping.HelloRequest) (*ping.HelloResponse, error) {
	toMessage := req.GetToMessage()
	fmt.Println(toMessage)
	response := ping.HelloResponse{
		ResMessage: "I hear " + toMessage,
	}
	return &response, nil
}

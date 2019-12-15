package controller

import (
	service "github.com/uma-co82/go-web-standard/internal/pingservice"
	"github.com/uma-co82/go-web-standard/proto/ping"
	"google.golang.org/grpc"
)

type Controller struct {
	service service.PingService
}

func (ctrl *Controller) GetMessage(server *grpc.Server) {
	pingService := ctrl.service
	ping.RegisterPingServer(server, &pingService)
}

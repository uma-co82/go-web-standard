package main

import (
	"fmt"
	"log"
	"net"

	"github.com/uma-co82/go-web-standard/cmd/http/controller"

	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	ctrl := controller.Controller{}
	ctrl.GetMessage(server)

	fmt.Printf("[server started] localhost%s", ":19003")
	err = server.Serve(listenPort)
	if err != nil {
		log.Fatalln(err)
	}
}

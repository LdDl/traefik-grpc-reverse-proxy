package main

import (
	"fmt"
	"net"
	service "service_two/rpc"
	"service_two/rpc/proto"

	"google.golang.org/grpc"
)

var (
	host = "localhost"
	port = 35002
)

func main() {
	// New gRPC server instance
	fullServer := grpc.NewServer()
	// Construct custom server with proper implementation
	customServer := service.ServiceTwo{}
	// Register custom server
	proto.RegisterServiceServer(fullServer, &customServer)

	// Init STD TCP-listener
	url := fmt.Sprintf("%s:%d", host, port)
	stdListener, err := net.Listen("tcp", url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Starting gRPC service #2...", url)
	if err := fullServer.Serve(stdListener); err != nil {
		fmt.Println(err)
		return
	}
}

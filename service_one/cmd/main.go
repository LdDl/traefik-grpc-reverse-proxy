package main

import (
	"fmt"
	"net"
	service "service_one/rpc"
	"service_one/rpc/proto"

	"google.golang.org/grpc"
)

var (
	host = "localhost"
	port = 35001
)

func main() {
	// New gRPC server instance
	fullServer := grpc.NewServer()
	// Construct custom server with proper implementation
	customServer := service.ServiceOne{}
	// Register custom server
	proto.RegisterServiceServer(fullServer, &customServer)

	// Init STD TCP-listener
	stdListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Starting gRPC service #1...", fmt.Sprintf("%s:%d", host, port))
	if err := fullServer.Serve(stdListener); err != nil {
		fmt.Println(err)
		return
	}
}

package service

import (
	"context"
	"log"
	"service_two/rpc/proto"
)

// ServiceTwo Wrap Service
type ServiceTwo struct {
	proto.ServiceServer
}

// GetSomeResponse Implement GetSomeResponse() to match interface of service.pb.go
func (server *ServiceTwo) GetSomeResponse(ctx context.Context, in *proto.NoArguments) (*proto.Response, error) {
	log.Println("Service #2 has been called")
	return &proto.Response{
		Code: 200,
		Text: "Message from service #2",
	}, nil
}

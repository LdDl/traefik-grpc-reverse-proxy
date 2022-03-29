package service

import (
	"context"
	"log"
	"service_one/rpc/proto"
)

// ServiceOne Wrap Service
type ServiceOne struct {
	proto.ServiceServer
}

// GetHiddenData Implement GetHiddenData() to match interface of service.pb.go
func (server *ServiceOne) GetSomeResponse(ctx context.Context, in *proto.NoArguments) (*proto.Response, error) {
	log.Println("Service #1 has been called")
	return &proto.Response{
		Code: 200,
		Text: "Message from service #1",
	}, nil
}

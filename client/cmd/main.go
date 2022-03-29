package main

import (
	"client/proto_one"
	"client/proto_two"
	"context"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func main() {
	serviceDirectOne := "localhost:35001"
	serviceDirectTwo := "localhost:35002"
	serviceProxy := "localhost:37000"

	// Call first service by direct port
	err := callServiceOneDirect(serviceDirectOne)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Call second service by direct port
	err = callServiceTwoDirect(serviceDirectTwo)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Call first service by reproxy port of Traefik
	err = callServiceOneProxy(serviceProxy)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Call second service by reproxy port of Traefik
	err = callServiceTwoProxy(serviceProxy)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func callServiceOneDirect(url string) error {
	// Connect to server
	grpcConn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "Can't connect to service #1 by direct port")
	}
	defer grpcConn.Close()

	// Create client
	grpcClient := proto_one.NewServiceClient(grpcConn)

	// Do request for data
	resp, err := grpcClient.GetSomeResponse(context.Background(), &proto_one.NoArguments{})
	if err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "Can't execute procedure on service #1 by direct port")
	}
	if resp == nil {
		return fmt.Errorf("Response is nil from service #1")
	}
	fmt.Println("Response from service #1 is:", resp)
	return nil
}

func callServiceTwoDirect(url string) error {
	// Connect to server
	grpcConn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "Can't connect to service #2 by direct port")
	}
	defer grpcConn.Close()

	// Create client
	grpcClient := proto_two.NewServiceClient(grpcConn)

	// Do request for data
	resp, err := grpcClient.GetSomeResponse(context.Background(), &proto_two.NoArguments{})
	if err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "Can't execute procedure on service #2 by direct port")
	}
	if resp == nil {
		return fmt.Errorf("Response is nil from service #2")
	}
	fmt.Println("Response from service #2 is:", resp)
	return nil
}

func callServiceOneProxy(url string) error {
	// Connect to server
	grpcConn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "Can't connect to service #1 by reverse proxy port")
	}
	defer grpcConn.Close()

	// Create client
	grpcClient := proto_one.NewServiceClient(grpcConn)

	// Do request for data
	resp, err := grpcClient.GetSomeResponse(context.Background(), &proto_one.NoArguments{})
	if err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "Can't execute procedure on service #1 by reverse proxy port")
	}
	if resp == nil {
		return fmt.Errorf("Response is nil from service #1 (reverse proxy)")
	}
	fmt.Println("Response from service #1 (reverse proxy) is:", resp)
	return nil
}
func callServiceTwoProxy(url string) error {
	// Connect to server
	grpcConn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "Can't connect to service #2 by reverse proxy port")
	}
	defer grpcConn.Close()

	// Create client
	grpcClient := proto_two.NewServiceClient(grpcConn)

	// Do request for data
	resp, err := grpcClient.GetSomeResponse(context.Background(), &proto_two.NoArguments{})
	if err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "Can't execute procedure on service #2 by reverse proxy port")
	}
	if resp == nil {
		return fmt.Errorf("Response is nil from service #2 (reverse proxy)")
	}
	fmt.Println("Response from service #2 (reverse proxy) is:", resp)
	return nil
}

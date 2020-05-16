package service

import (
	"fmt"

	"golang.org/x/net/context"

	pb "github.com/zoido/twirp-grpc/api/v1"
)

// NewGreeterService creates new instance of the corresponding GRPC request handler.
func NewGreeterService() pb.GreeterServiceServer {
	return &greeterService{}
}

type greeterService struct{}

func (*greeterService) GetGreeting(_ context.Context, req *pb.GetGreetingRequest) (*pb.GetGreetingResponse, error) {
	return &pb.GetGreetingResponse{
		Greeting: &pb.Greeting{
			Message: fmt.Sprintf("Hello %s", req.Name),
		},
	}, nil
}

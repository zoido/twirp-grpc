package service

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/twitchtv/twirp"
	pb "github.com/zoido/twirp-grpc/api/v1"
)

// NewGreeterService creates new instance of the corresponding Greeter request handler.
func NewGreeterService() pb.GreeterServiceServer {
	return &greeterService{}
}

type greeterService struct{}

func (*greeterService) GetGreeting(_ context.Context, req *pb.GetGreetingRequest) (*pb.GetGreetingResponse, error) {
	if req.Name == "Dr.Evil" {
		return nil, twirp.NewError(twirp.PermissionDenied, "Dr.Evil's not welcome here.")
	}
	return &pb.GetGreetingResponse{
		Greeting: &pb.Greeting{
			Message: fmt.Sprintf("Hello %s!", req.Name),
		},
	}, nil
}

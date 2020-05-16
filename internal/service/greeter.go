package service

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/rs/zerolog/log"
	"github.com/twitchtv/twirp"
	pb "github.com/zoido/twirp-grpc/api/v1"
)

// NewGreeterService creates new instance of the corresponding GRPC request handler.
func NewGreeterService() pb.GreeterServiceServer {
	return &greeterService{}
}

type greeterService struct{}

func (*greeterService) GetGreeting(_ context.Context, req *pb.GetGreetingRequest) (*pb.GetGreetingResponse, error) {
	if req.Name == "Dr.Evil" {
		err := twirp.NewError(twirp.PermissionDenied, "Dr.Evil's not welcome here.")
		log.Error().
			Err(err).
			Send()
		return nil, err
	}
	return &pb.GetGreetingResponse{
		Greeting: &pb.Greeting{
			Message: fmt.Sprintf("Hello %s!", req.Name),
		},
	}, nil
}

type err struct {
}

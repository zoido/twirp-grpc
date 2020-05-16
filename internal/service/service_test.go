package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/twitchtv/twirp"

	pb "github.com/zoido/twirp-grpc/api/v1"
	greeter "github.com/zoido/twirp-grpc/internal/service"
)

type GreeterTestSuite struct {
	suite.Suite

	service pb.GreeterServiceServer
}

func TestGreeterTestSuite(t *testing.T) {
	suite.Run(t, new(GreeterTestSuite))
}

func (s *GreeterTestSuite) SetupTest() {
	s.service = greeter.NewGreeterService()
}

func (s *GreeterTestSuite) TestGreet_Ok() {
	// When
	resp, err := s.service.GetGreeting(context.TODO(),
		&pb.GetGreetingRequest{
			Name: "Test Name",
		},
	)

	// Then
	s.Require().NoError(err)
	s.Require().Equal(resp.Greeting.Message, "Hello Test Name!")
}

func (s *GreeterTestSuite) TestGreet_DrEvilNotAllowed() {
	// When
	_, err := s.service.GetGreeting(context.TODO(),
		&pb.GetGreetingRequest{
			Name: "Dr.Evil",
		},
	)

	// Then
	s.Require().Error(err)
	s.Require().Equal(err.(twirp.Error).Code(), twirp.PermissionDenied)
}

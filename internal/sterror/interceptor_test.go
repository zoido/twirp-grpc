package sterror_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/twitchtv/twirp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zoido/twirp-grpc/internal/sterror"
)

type SterrorTestSuite struct {
	suite.Suite
}

func TestSterrorTestSuite(t *testing.T) {
	suite.Run(t, new(SterrorTestSuite))
}

func (s *SterrorTestSuite) TestStream_Conversion() {
	// Given
	interceptor := sterror.StreamInterceptor()

	twErr := twirp.NewError(twirp.NotFound, "NOT FOUND TEST")

	// When
	err := interceptor(context.TODO(), nil, &grpc.StreamServerInfo{}, errStreamHandler(twErr))

	// Then
	s.Require().Error(err)
	s.Require().Equal(codes.NotFound, status.Code(err))
	s.Require().Contains(err.Error(), "NOT FOUND TEST")
}

func (s *SterrorTestSuite) TestStream_PassStatus() {
	// Given
	interceptor := sterror.StreamInterceptor()
	sErr := status.Error(codes.NotFound, "TEST NOT FOUND")

	// When
	err := interceptor(context.TODO(), nil, &grpc.StreamServerInfo{}, errStreamHandler(sErr))

	// Then
	s.Require().Error(err)
	s.Require().Equal(codes.NotFound, status.Code(err))
	s.Require().Contains(err.Error(), "TEST NOT FOUND")
}

func (s *SterrorTestSuite) TestStream_PassNative() {
	// Given
	interceptor := sterror.StreamInterceptor()
	nErr := errors.New("TEST NATIVE ERROR")

	// When
	err := interceptor(context.TODO(), nil, &grpc.StreamServerInfo{}, errStreamHandler(nErr))

	// Then
	s.Require().Error(err)
	s.Require().Equal(codes.Unknown, status.Code(err))
	s.Require().Contains(err.Error(), "TEST NATIVE ERROR")

}

func (s *SterrorTestSuite) TestUnary_Conversion() {
	// Given
	interceptor := sterror.UnaryInterceptor()

	twErr := twirp.NewError(twirp.NotFound, "NOT FOUND TEST")

	// When
	_, err := interceptor(context.TODO(), nil, &grpc.UnaryServerInfo{}, errUnaryHandler(twErr))

	// Then
	s.Require().Error(err)
	s.Require().Equal(codes.NotFound, status.Code(err))
	s.Require().Contains(err.Error(), "NOT FOUND TEST")
}

func (s *SterrorTestSuite) TestUnary_PassStatus() {
	// Given
	interceptor := sterror.UnaryInterceptor()
	sErr := status.Error(codes.NotFound, "TEST NOT FOUND")

	// When
	_, err := interceptor(context.TODO(), nil, &grpc.UnaryServerInfo{}, errUnaryHandler(sErr))

	// Then
	s.Require().Error(err)
	s.Require().Equal(codes.NotFound, status.Code(err))
	s.Require().Contains(err.Error(), "TEST NOT FOUND")
}

func (s *SterrorTestSuite) TestUnary_PassResponse() {
	// Given
	interceptor := sterror.UnaryInterceptor()
	response := "TEST RESPONSE STRING"

	// When
	resp, err := interceptor(context.TODO(), nil, &grpc.UnaryServerInfo{}, responseUnaryHandler(response))

	// Then
	s.Require().NoError(err)
	s.Require().Equal(response, resp)

}

func errUnaryHandler(err error) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, err
	}
}

func errStreamHandler(err error) grpc.StreamHandler {
	return func(srv interface{}, stream grpc.ServerStream) error {
		return err
	}
}

func responseUnaryHandler(resp interface{}) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return resp, nil
	}
}

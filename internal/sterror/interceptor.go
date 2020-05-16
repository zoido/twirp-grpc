package sterror

import (
	"context"

	"github.com/twitchtv/twirp"
	"google.golang.org/grpc"
)

func StreamInterceptor() grpc.StreamServerInterceptor {
	return convertStreamError
}
func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return convertUnaryError
}

func convertStreamError(
	srv interface{},
	ss grpc.ServerStream,
	_ *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	err := handler(srv, ss)
	return toStatus(err)
}

func convertUnaryError(
	ctx context.Context,
	req interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	return resp, toStatus(err)
}

func toStatus(err error) error {
	twe, ok := err.(twirp.Error)
	if ok {
		return FromTwirpError(twe)
	}
	return err
}

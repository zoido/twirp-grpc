package sterror

import (
	"github.com/twitchtv/twirp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errCodeMap = map[twirp.ErrorCode]codes.Code{
	twirp.Canceled:           codes.Canceled,
	twirp.Unknown:            codes.Unknown,
	twirp.InvalidArgument:    codes.InvalidArgument,
	twirp.DeadlineExceeded:   codes.DeadlineExceeded,
	twirp.NotFound:           codes.NotFound,
	twirp.AlreadyExists:      codes.AlreadyExists,
	twirp.PermissionDenied:   codes.PermissionDenied,
	twirp.ResourceExhausted:  codes.ResourceExhausted,
	twirp.FailedPrecondition: codes.FailedPrecondition,
	twirp.Aborted:            codes.Aborted,
	twirp.OutOfRange:         codes.OutOfRange,
	twirp.Unimplemented:      codes.Unimplemented,
	twirp.Internal:           codes.Internal,
	twirp.Unavailable:        codes.Unavailable,
	twirp.DataLoss:           codes.DataLoss,
	twirp.Unauthenticated:    codes.Unauthenticated,
}

func FromTwirpError(err twirp.Error) error {
	var c codes.Code
	c, ok := errCodeMap[err.Code()]
	if !ok {
		c = codes.Unknown
	}
	return status.Error(c, err.Msg())
}

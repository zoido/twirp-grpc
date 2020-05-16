# twirp-grpc

Proof of concept  serving Twirp and gRPC from single implementation.

- built by Bazel
- gRPC interceptors converting `twirp.Error` to `status.Status`

package main

import (
	"net"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zoido/yag-config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/zoido/twirp-grpc/api/v1"
	"github.com/zoido/twirp-grpc/internal/service"
)

func main() {
	bind := "localhost:9911"
	y := yag.New()
	y.String(&bind, "bind", "address to listen on")

	if err := y.Parse(os.Args[1:]); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed parsing configuration")
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		},
	)

	listen, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("bind", bind).
			Msg("failed to listen")
	}

	address := listen.Addr().(*net.TCPAddr)
	log.Info().
		Int("port", address.Port).
		IPAddr("ip_address", address.IP).
		Msg("Listening")

	srv := grpc.NewServer()
	pb.RegisterGreeterServiceServer(srv, service.NewGreeterService())
	reflection.Register(srv)

	if err := srv.Serve(listen); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start the gRPC server")
	}
}

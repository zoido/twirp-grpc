package main

import (
	"net"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zoido/yag-config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/zoido/twirp-grpc/api/v1"
	"github.com/zoido/twirp-grpc/internal/service"
	"github.com/zoido/twirp-grpc/internal/sterror"
)

func main() {
	bindGrpc := "localhost:9911"
	bindTwirp := "localhost:8011"
	y := yag.New()
	y.String(&bindGrpc, "bind_grpc", "address to listen on")
	y.String(&bindTwirp, "bind_twirp", "address to listen on")

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

	svc := service.NewGreeterService()

	srv := grpc.NewServer(
		grpc.StreamInterceptor(sterror.StreamInterceptor()),
		grpc.UnaryInterceptor(sterror.UnaryInterceptor()),
	)

	pb.RegisterGreeterServiceServer(srv, svc)
	reflection.Register(srv)

	twirpHandler := pb.NewGreeterServiceServer(svc, nil)

	shutdown := make(chan bool)

	go func() {
		log.Info().
			Str("bind", bindGrpc).
			Msg("Listening")
		listen, err := net.Listen("tcp", bindGrpc)
		if err != nil {
			log.Fatal().
				Err(err).
				Str("bind", bindGrpc).
				Msg("failed to listen")
		}
		log.Info().
			Str("bind", bindGrpc).
			Msg("Starting gRPC server")
		if err := srv.Serve(listen); err != nil {
			log.Error().
				Err(err).
				Msg("Failed to start the gRPC server")
		}
		shutdown <- true
	}()

	go func() {
		log.Info().
			Str("bind", bindTwirp).
			Msg("Starting Twirp HTTP server")

		if err := http.ListenAndServe(bindTwirp, twirpHandler); err != nil {
			log.Error().
				Err(err).
				Send()
		}

		shutdown <- true
	}()

	<-shutdown
}

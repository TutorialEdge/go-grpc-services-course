package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

// RocketService - defines the rocket service interface
// which any service passed in to our Handler will need to
// conform to
type RocketService interface{}

// Handler -
type Handler struct {
	RocketService RocketService
}

// New - returns a new gRPC handler
func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

// Serve - starts out gRPC listeners
func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return err
	}

	return nil
}

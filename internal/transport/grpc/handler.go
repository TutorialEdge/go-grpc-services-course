package grpc

import (
	"context"
	"log"
	"net"

	rkt "github.com/TutorialEdge/tutorial-protos/rocket/v1"
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
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return err
	}

	return nil
}

func (h Handler) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	log.Print("Get Rocket gRPC Endpoint Hit")
	return &rkt.GetRocketResponse{}, nil
}

func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	log.Print("Get Rocket gRPC Endpoint Hit")
	return &rkt.AddRocketResponse{}, nil
}

func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	log.Print("Get Rocket gRPC Endpoint Hit")
	return &rkt.DeleteRocketResponse{}, nil
}

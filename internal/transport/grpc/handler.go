package grpc

import (
	"context"
	"log"
	"net"

	"github.com/TutorialEdge/go-grpc-services-course/internal/rocket"
	rkt "github.com/TutorialEdge/tutorial-protos/rocket/v1"
	"google.golang.org/grpc"
)

// RocketService - defines the rocket service interface
// which any service passed in to our Handler will need to
// conform to
type RocketService interface {
	GetRocketByID(id string) (rocket.Rocket, error)
	AddRocket(rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(id string) error
}

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
	rocket, err := h.RocketService.GetRocketByID(req.Id)
	if err != nil {
		return &rkt.GetRocketResponse{}, err
	}
	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   rocket.ID,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	log.Print("Add Rocket gRPC Endpoint Hit")
	newRkt, err := h.RocketService.AddRocket(rocket.Rocket{
		ID:   req.Rocket.Id,
		Name: req.Rocket.Name,
		Type: req.Rocket.Type,
	})
	if err != nil {
		return &rkt.AddRocketResponse{}, err
	}
	return &rkt.AddRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   newRkt.ID,
			Type: newRkt.Type,
			Name: newRkt.Name,
		},
	}, nil
}

func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	log.Print("Delete Rocket gRPC Endpoint Hit")
	err := h.RocketService.DeleteRocket(req.Id)
	if err != nil {
		return &rkt.DeleteRocketResponse{
			Status: err.Error(),
		}, err
	}

	return &rkt.DeleteRocketResponse{
		Status: "successfully deleted",
	}, nil

}

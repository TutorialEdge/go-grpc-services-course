package main

import (
	"log"

	"github.com/TutorialEdge/go-grpc-services-course/internal/db"
	"github.com/TutorialEdge/go-grpc-services-course/internal/rocket"
	"github.com/TutorialEdge/go-grpc-services-course/internal/transport/grpc"
)

// Run - handles the setup and starting of our application
// using this approach makes testing easier and we can more
// gracefully handle errors
func Run() error {
	log.Println("Starting up Rocket gRPC Service")

	// rktStore - the store responsible for holding
	// our rocket inventory
	rktStore := db.New()
	// if err := rktStore.Migrate(); err != nil {
	// 	return err
	// }

	// rktService the service responsible for updating our
	// rocket inventory
	rktService := rocket.New(rktStore)
	// rktHandler instantiates a new gRPC handler
	// which we pass our rktService into
	rktHandler := grpc.New(rktService)

	// Start our gRPC listener, this is a blocking
	// function call so it should be the last thing
	// we run in this function
	if err := rktHandler.Serve(); err != nil {
		return err
	}
	return nil
}

func main() {
	// our main function is super small, only responsible
	// for calling Run and then handling the error
	if err := Run(); err != nil {
		log.Fatal(err.Error())
	}
}

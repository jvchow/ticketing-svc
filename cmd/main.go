package main

import (
	"log"
	"net"
	"ticketing-svc/config"
	train "ticketing-svc/proto"
	"ticketing-svc/service"

	"google.golang.org/grpc"
)

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the train service to the server
	train.RegisterTicketServiceServer(s, service.NewServer())

	log.Printf("server listening at %v", lis.Addr())

	// Start serving requests via the listener
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

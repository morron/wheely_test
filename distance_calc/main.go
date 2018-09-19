package main

import (
	"log"
	"net"

	"distance_calc/calc"
	p "distance_calc/proto"
	"distance_calc/services"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Calculate(ctx context.Context, in *p.Ride) (*p.DirectionResult, error) {
	service := new(services.Directions)
	response, _ := calc.Calc(service, in.Origin, in.Destination)
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	p.RegisterDistanceCalcultaionServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

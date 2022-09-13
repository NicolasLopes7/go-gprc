package main

import (
	"log"
	"net"

	"github.com/NicolasLopes7/gprc-go/pb"
	"github.com/NicolasLopes7/gprc-go/services"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}

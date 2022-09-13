package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/NicolasLopes7/gprc-go/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	AddUserVerbose(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "bati a cabeça no teclado e saiu: afdlsjkasdfljk",
		Name:  "Nicolau",
		Email: "n@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	log.Printf("Response from server: %v", res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "bati a cabeça no teclado e saiu: afdlsjkasdfljk",
		Name:  "Nicolau",
		Email: "n@gmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}

		fmt.Println("Status: ", stream.Status)
	}
}

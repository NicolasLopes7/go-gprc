package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	AddUsers(client)
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

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "nicolino",
			Email: "nicolino@gmail.com",
		},
		{
			Id:    "2",
			Name:  "nicolau",
			Email: "nicolau@gmail.com",
		},
		{
			Id:    "3",
			Name:  "nicolove",
			Email: "nicolove@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for i := range reqs {
		fmt.Println("Sending user: ", reqs[i].GetName())
		stream.Send(reqs[i])
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println("Server Response: ", res)
}

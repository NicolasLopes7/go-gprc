package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/NicolasLopes7/gprc-go/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	// Insert - Database
	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

/*
That method emulates a long user register process.
Let's imagine that on your client registration you need to check some informations on external services: KYC, Document OCR...
We want to be able to see the progress in real time of that registration. Luckily, gRPC has a solution for that. Streams.
*/
func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	var user = pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &user,
	})

	time.Sleep(time.Second * 2)

	stream.Send(&pb.UserResultStream{
		Status: "User has been saved on our DB",
		User:   &user,
	})

	time.Sleep(time.Second * 2)

	stream.Send(&pb.UserResultStream{
		Status: "Document OCR approved",
		User:   &user,
	})

	time.Sleep(time.Second * 1)

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User:   &user,
	})

	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	users := []*pb.User{}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}

		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})

		fmt.Println("Adding: ", req.GetName())
	}
}

func (*UserService) AddUsersStreamBoth(stream pb.UserService_AddUserStreamBothServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving stream from the client: %v", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User:   req,
		})
		if err != nil {
			log.Fatalf("Error sending stream to the client: %v", err)
		}
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/datpham912000/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

// Implement function GreetServiceServer interface (greet_grpc.pb.go)
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet function was invoked with ", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	result := "Hello " + firstName + " " + lastName

	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello from server")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v ", err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/datpham912000/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Println("Sum function was invoked with ", req)
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()

	result := firstNumber + secondNumber

	res := &calculatorpb.SumResponse{
		SumResult: result,
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
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v ", err)
	}
}

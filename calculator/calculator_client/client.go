package main

import (
	"context"
	"fmt"
	"log"

	"github.com/datpham912000/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client")
	//create connection to server
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	//Close connection when program exit
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	//fmt.Println("Created client: ", c)

	doUnary(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Sum RPC...")

	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling Sum RPC", err)
	}
	log.Println("Response from Sum: ", res.SumResult)
}

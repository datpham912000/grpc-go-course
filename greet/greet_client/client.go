package main

import (
	"context"
	"fmt"
	"log"

	"github.com/datpham912000/grpc-go-course/greet/greetpb"
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

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Println("Created client: ", c)

	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Pham",
			LastName:  "Dat",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling Greet RPC", err)
	}
	log.Println("Response from Greet: ", res.Result)
}

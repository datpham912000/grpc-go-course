package main

import (
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
	fmt.Println("Created client: ", c)

}

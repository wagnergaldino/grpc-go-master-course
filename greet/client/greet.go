package main

import (
	"context"
	"log"

	pb "github.com/wagnergaldino/grpc-go-master-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Wagner",
	})

	if err != nil {
		log.Fatalf("Failed to Greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}

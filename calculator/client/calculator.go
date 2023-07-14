package main

import (
	"context"
	"log"

	pb "github.com/wagnergaldino/grpc-go-master-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.CalculatorRequest{
		FirstNumber:  2,
		SecondNumber: 5,
	})

	if err != nil {
		log.Fatalf("Failed to Sum: %v\n", err)
	}

	log.Printf("Sum = %d\n", res.Result)
}

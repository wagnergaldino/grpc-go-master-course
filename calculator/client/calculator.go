package main

import (
	"context"
	"io"
	"log"

	pb "github.com/wagnergaldino/grpc-go-master-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  2,
		SecondNumber: 5,
	})

	if err != nil {
		log.Fatalf("Failed to Sum: %v\n", err)
	}

	log.Printf("Sum = %d\n", res.Result)
}

func doPrime(c pb.CalculatorServiceClient) {
	log.Println("doPrime was invoked")
	req := &pb.PrimeRequest{
		Number: 120,
	}

	stream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to Prime: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Prime: %d\n", msg.Result)
	}

}

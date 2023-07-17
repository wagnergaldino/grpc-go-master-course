package main

import (
	"context"
	"io"
	"log"
	"time"

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

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Failed to Avg: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to get response from Avg: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	waitch := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			msg, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while reading the stream: %v\n", err)
				break
			}

			log.Printf("Max: %d\n", msg.Result)
		}
		close(waitch)
	}()
	<-waitch
}

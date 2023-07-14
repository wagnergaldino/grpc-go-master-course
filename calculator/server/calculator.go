package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/wagnergaldino/grpc-go-master-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v + %v\n", in.FirstNumber, in.SecondNumber)
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Printf("Prime function was invoked with %v\n", in)

	nr := in.Number
	div := int32(2)

	for nr > 1 {
		if nr%div == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: div,
			})
			nr = nr / div
		} else {
			div++
			fmt.Printf("Divisor has increased to %d\n", div)
		}
	}

	return nil
}

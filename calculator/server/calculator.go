package main

import (
	"context"
	"log"

	pb "github.com/wagnergaldino/grpc-go-master-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Sum function was invoked with %v + %v\n", in.FirstNumber, in.SecondNumber)
	return &pb.CalculatorResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}

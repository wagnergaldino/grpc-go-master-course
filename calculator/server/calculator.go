package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"

	pb "github.com/wagnergaldino/grpc-go-master-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Printf("Avg function was invoked")

	var cont, soma int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(soma) / float64(cont),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving request %v\n", req)

		cont++
		soma += int32(req.Number)

		fmt.Printf("Avg = %f\n", float64(soma)/float64(cont))
	}
}

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max function was invoked")

	var max int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving request %v\n", req)
		if req.Number > max {
			max = req.Number
			err = stream.Send(&pb.MaxResponse{
				Result: max,
			})
			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt function was invoked with %v\n", in)

	number := in.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d\n", number),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}

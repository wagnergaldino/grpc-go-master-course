package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/wagnergaldino/grpc-go-master-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	log.Printf("Connected on: %s\n", addr)

	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
	doGreetManyTimes(c)
	doLongGreet(c)
	doGreetEveryone(c)
}

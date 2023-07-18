package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/wagnergaldino/grpc-go-master-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true // change to false if needed
	opts := []grpc.DialOption{}

	if tls {
		certfile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certfile, "")
		if err != nil {
			log.Fatalf("Error while loading CA trust cvertificate: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	// if tls == false use this one below
	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	log.Printf("Connected on: %s\n", addr)

	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	// doGreetWithDeadline(c, 5*time.Second)
	// doGreetWithDeadline(c, 2*time.Second)
}

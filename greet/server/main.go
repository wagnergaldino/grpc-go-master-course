package main

import (
	"log"
	"net"

	pb "github.com/wagnergaldino/grpc-go-master-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listeninig on: %s\n", addr)

	opts := []grpc.ServerOption{}
	tls := true // change that to false if needed

	if tls {
		certfile := "ssl/server.crt"
		keyfile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certfile, keyfile)
		if err != nil {
			log.Fatalf("Failed to load certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v\n", err)
	}

}

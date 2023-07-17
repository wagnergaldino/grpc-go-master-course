package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/wagnergaldino/grpc-go-master-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "WaGal1"},
		{FirstName: "WaGal2"},
		{FirstName: "WaGal3"},
		{FirstName: "WaGal4"},
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

			log.Printf("GreetEveryone: %s\n", msg.Result)
		}
		close(waitch)
	}()
	<-waitch
}

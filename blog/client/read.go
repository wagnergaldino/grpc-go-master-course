package main

import (
	"context"
	"log"

	pb "github.com/wagnergaldino/grpc-go-master-course/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Printf("readBlog function was invoked with id %s\n", id)

	req := &pb.BlogId{
		Id: id,
	}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error while reading data: %v\n", err)
	}

	log.Printf("Blog data: %v\n", res)
	return res
}

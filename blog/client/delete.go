package main

import (
	"context"
	"log"

	pb "github.com/wagnergaldino/grpc-go-master-course/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Printf("deleteBlog function was invoked with id %s\n", id)

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error while deleting data: %v\n", err)
	}

	log.Println("Blog successfully deleted")
}

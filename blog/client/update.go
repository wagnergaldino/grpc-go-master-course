package main

import (
	"context"
	"log"

	pb "github.com/wagnergaldino/grpc-go-master-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog function was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not WaGal",
		Title:    "New Title",
		Content:  "New Content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}

package main

import (
	"context"
	"log"

	pb "github.com/wagnergaldino/grpc-go-master-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog function was invoked")

	blog := pb.Blog{
		AuthorId: "WaGal",
		Title:    "1st blog",
		Content:  "1st blog content",
	}

	res, err := c.CreateBlog(context.Background(), &blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog successfully created with id %s\n", res.Id)
	return res.Id
}

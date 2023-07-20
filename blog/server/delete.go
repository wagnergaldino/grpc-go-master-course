package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/wagnergaldino/grpc-go-master-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*empty.Empty, error) {
	log.Printf("DeleteBlog function was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		log.Fatalf("Error while converting to ObjectID: %v\n", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID: %v\n", err),
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		log.Fatalf("Error while deleting blog: %v\n", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete blog: %v\n", err),
		)
	}

	if res.DeletedCount == 0 {
		log.Fatalf("Error while deleting blog: %v\n", err)
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Blog not found: %v\n", err),
		)
	}

	return &emptypb.Empty{}, nil
}

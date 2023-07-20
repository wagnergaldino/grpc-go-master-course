package main

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/wagnergaldino/grpc-go-master-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*empty.Empty, error) {
	log.Printf("UpdateBlog function was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot parse id",
		)
	}

	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)

	if err != nil {
		log.Println("Error while updating data")
		return nil, status.Errorf(
			codes.Internal,
			"Cannot update",
		)
	}

	if res.MatchedCount == 0 {
		log.Println("Error while updating data")
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog with id",
		)
	}

	return &emptypb.Empty{}, nil
}

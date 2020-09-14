package main

import (
	"context"
	"log"
	"time"

	pb "github.com/takuya911/gRPC_sample/user/proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetUser(ctx, &pb.GetUserRequest{Id: "1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf(r.GetUser())
}

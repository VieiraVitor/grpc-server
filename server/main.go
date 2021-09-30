package main

import (
	"context"
	"log"
	"net"

	pb "github.com/vieiravitor/grpc-server/protobuf/pb"
	"google.golang.org/grpc"
)

type TransactionService struct {
	pb.UnimplementedTransactionServiceServer
}

func (t *TransactionService) SaveTransaction(ctx context.Context, in *pb.Payment) (*pb.Empty, error) {

	return &pb.Empty{}, nil
}

const port = ":50051"

func main() {
	log.Printf("Listening")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	log.Printf("Listening")
	grpcServer := grpc.NewServer()

	pb.RegisterTransactionServiceServer(grpcServer, &TransactionService{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to start server!")
	}
}

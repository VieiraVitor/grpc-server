package main

import (
	"context"
	"log"
	"testing"

	pb "github.com/vieiravitor/grpc-server/protobuf/pb"
	"google.golang.org/grpc"
)

func BenchmarkGRPC(b *testing.B) {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	transactionServiceClient := pb.NewTransactionServiceClient(conn)
	for i := 0; i < b.N; i++ {
		paymentRequest := &pb.PaymentRequest{
			Id:          int32(i),
			Amount:      1000,
			Description: "Novo produto",
			CreditCard: &pb.CreditCard{
				Number: "1234",
				Name:   "Vitor",
				Cvv:    123,
			},
		}
		SaveTransaction(transactionServiceClient, paymentRequest, b)
	}
}

func SaveTransaction(transactionServiceClient pb.TransactionServiceClient, paymentRequest *pb.PaymentRequest, b *testing.B) {
	_, err := transactionServiceClient.SaveTransaction(context.Background(), paymentRequest)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// fmt.Printf("%s - ID: %d\n", paymentResponse.Response, paymentResponse.Id)
}

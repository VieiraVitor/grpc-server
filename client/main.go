package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vieiravitor/grpc-server/protobuf/pb"
	"google.golang.org/grpc"
)

type CreditCard struct {
	Number string
	Name   string
	CVV    int32
}

type Payment struct {
	ID          string
	Amount      float64
	Description string
	CreditCard  CreditCard
}

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	TransactionServiceClient := pb.NewTransactionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Payment.ID = 1
	// Payment.CreditCard.Number = "1234"
	// Payment.CreditCard.Name = "Vitor"
	// Payment.CreditCard.CVV = 123
	// Payment.Amount = 1000
	// Payment.Description = "Novo produto"
	_, err = TransactionServiceClient.SaveTransaction(ctx,
		&pb.Payment{
			Id:          1,
			Amount:      1000,
			Description: "Novo produto",
			CreditCard: &pb.CreditCard{
				Number: "1234",
				Name:   "Vitor",
				Cvv:    123,
			},
		})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Sucess: !")
}

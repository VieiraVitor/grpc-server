package main

import (
	"context"
	"fmt"
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

	transactionServiceClient := pb.NewTransactionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var i int32
	var amount float64
	for i = 0; i < 100; i++ {
		amount++
		PaymentRequest := &pb.PaymentRequest{
			Id:          i,
			Amount:      amount,
			Description: "Novo produto",
			CreditCard: &pb.CreditCard{
				Number: "1234",
				Name:   "Vitor",
				Cvv:    123,
			},
		}
		paymentResponse, err := transactionServiceClient.SaveTransaction(ctx, PaymentRequest)
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Printf("%s - ID: %d\n", paymentResponse.Response, paymentResponse.Id)
	}

	log.Printf("Sucess: !")
}

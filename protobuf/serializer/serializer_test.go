package serializer_test

import (
	"log"
	"testing"

	pb "github.com/vieiravitor/grpc-server/protobuf/pb"
	"github.com/vieiravitor/grpc-server/protobuf/serializer"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "./binaries/payment.bin"

	payment1 := &pb.PaymentRequest{
		Id:          1,
		Amount:      1000,
		Description: "Novo produto",
		CreditCard: &pb.CreditCard{
			Number: "1234",
			Name:   "Vitor",
			Cvv:    123,
		},
	}
	err := serializer.WriteProtobufToBinaryFile(payment1, binaryFile)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	payment2 := &pb.PaymentRequest{
		Id:          1,
		Amount:      1000,
		Description: "Novo produto",
		CreditCard: &pb.CreditCard{
			Number: "1234",
			Name:   "Vitor",
			Cvv:    123,
		},
	}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, payment2)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}

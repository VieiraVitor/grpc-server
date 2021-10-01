Protoc usage:
protoc --proto_path=protobuf/protofile protobuf/protofile/*.proto --go_out=protobuf --go-grpc_out=protobuf

ghz usage:
./ghz --insecure --proto ./protobuf/protofile/transaction.proto --call transaction.TransactionService.SaveTransaction -d '{"id": 1, "amount": 1000, "description": "novo produto", "creditCard": {"number": "1234", "name": "Vitor", "cvv": 123}}' 0.0.0.0:50051

To get a html output:
./ghz --insecure --proto ./protobuf/protofile/transaction.proto --call transaction.TransactionService.SaveTransaction -d '{"id": 1, "amount": 1000, "description": "novo produto", "creditCard": {"number": "1234", "name": "Vitor", "cvv": 123}}' 0.0.0.0:50051 -O html -o html
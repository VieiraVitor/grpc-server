Protoc usage:
protoc --proto_path=protobuf/protofile protobuf/protofile/*.proto --go_out=protobuf --go-grpc_out=protobuf

ghz usage:
./ghz --insecure --proto ./protobuf/protofile/transaction.proto --call transaction.TransactionService.SaveTransaction -d '{"id": 1, "amount": 1000, "description": "novo produto", "creditCard": {"number": "1234", "name": "Vitor", "cvv": 123}}' 0.0.0.0:50051

To get a html output:
./ghz --insecure --proto ./protobuf/protofile/transaction.proto --call transaction.TransactionService.SaveTransaction -d '{"id": 1, "amount": 1000, "description": "novo produto", "creditCard": {"number": "1234", "name": "Vitor", "cvv": 123}}' 0.0.0.0:50051 -O html -o html

exemplo:
./ghz --insecure -n 100 -c 10 --proto ./protobuf/protofile/transaction.proto --call transaction.TransactionService.SaveTransaction -d '{"id": 1, "amount": 1000, "description": "novo produto", "creditCard": {"number": "1234", "name": "Vitor", "cvv": 123}}' localhost:50051

teste1:
./ghz --insecure -n 1000 -c 100 --proto ./protobuf/protofile/transaction.proto --call transaction.TransactionService.SaveTransaction -d '{"id": 1, "amount": 1000, "description": "novo produto", "creditCard": {"number": "1234", "name": "Vitor", "cvv": 123}}' localhost:50051

teste2:
./ghz --insecure -n 10000 -c 1000 --proto ./protobuf/protofile/transaction.proto --call transaction.TransactionService.SaveTransaction -d '{"id": 1, "amount": 1000, "description": "novo produto", "creditCard": {"number": "1234", "name": "Vitor", "cvv": 123}}' localhost:50051

teste3:
./ghz --insecure -n 100000 -c 10000 --proto ./protobuf/protofile/transaction.proto --call transaction.TransactionService.SaveTransaction -d '{"id": 1, "amount": 1000, "description": "novo produto", "creditCard": {"number": "1234", "name": "Vitor", "cvv": 123}}' localhost:50051

teste4;
./ghz --insecure -z 10s -c 1000 --proto ./protobuf/protofile/transaction.proto --call transaction.TransactionService.SaveTransaction -d '{"id": 1, "amount": 1000, "description": "novo produto", "creditCard": {"number": "1234", "name": "Vitor", "cvv": 123}}' localhost:50051
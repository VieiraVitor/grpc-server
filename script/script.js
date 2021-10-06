import grpc from 'k6/net/grpc';
import { sleep } from 'k6';
export let options = {
	vus: 5000,
	duration: '10s',
}

const client = new grpc.Client();
client.load(['../protobuf/protofile'], 'transaction.proto');

export default () => {
	client.connect('localhost:50051', {
		plaintext: true
	});

	const data = {
		id: 1,
		amount: 1000,
		description: "Novo produto",
		creditCard: {
			number: "1234",
			name: "Vitor",
			cvv: 123,
		}
	};
	client.invoke('transaction.TransactionService/SaveTransaction', data);

	client.close();
	sleep(1);
};

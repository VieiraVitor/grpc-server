import grpc from 'k6/net/grpc';
import { sleep } from 'k6';
import { check } from 'k6';
export let options = {
	vus: 10000,
	iterations: 100000,
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
	const res = client.invoke('transaction.TransactionService/SaveTransaction', data);
	check(res, {
		'is status 200': (r) => r && r.status === grpc.StatusOK,
	});

	client.close();
	sleep(1);
};

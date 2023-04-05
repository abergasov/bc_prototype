import CircleContract = models.CircleContract;
import DeployedCircleContract = models.DeployedCircleContract;

const mockedCircle = "372e3c13-a552-4ebe-96fb-7f1bf45cf68c"

export async function getCircleContract(): Promise<DeployedCircleContract> {
	const res = await fetch(`/api/circle/${mockedCircle}/contract`, {
		method: "GET",
		headers: {
			"Content-Type": "application/json",
		},
	})
	return res.json()
}

export async function deployCircleContract(): Promise<CircleContract> {
	const res = await fetch(`/api/circle/${mockedCircle}/contract`, {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
	})
	return res.json()
}

export async function saveContract(contractAddress: string, transactionHash: string): Promise<CircleContract> {
	const res = await fetch(`/api/circle/save_contract`, {
		method: "POST",
		body: JSON.stringify({
			circle_id: mockedCircle,
			transaction_hash: transactionHash,
			contract_address: contractAddress,
		}),
		headers: {
			"Content-Type": "application/json",
		},
	})
	return res.json()
}
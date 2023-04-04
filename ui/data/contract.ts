const mockedCircle = "372e3c13-a552-4ebe-96fb-7f1bf45cf68c"

export async function getCircleContract(): Promise<string> {
	const res = await fetch(`/api/circle/${mockedCircle}/contract`, {
		method: "GET",
		headers: {
			"Content-Type": "application/json",
		},
	})
	return res.json()
}

export async function deployCircleContract(): Promise<string> {
	const res = await fetch(`/api/circle/${mockedCircle}/contract`, {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
	})
	return res.json()
}

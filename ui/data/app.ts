export async function resetAppData(): Promise<string> {
	const res = await fetch(`/api/reset`, {
		method: "GET",
		headers: {
			"Content-Type": "application/json",
		},
	})
	return res.json()
}

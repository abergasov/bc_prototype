export async function resetAppData(): Promise<string> {
	const res = await fetch(`/api/erase_all`, {
		method: "GET",
		headers: {
			"Content-Type": "application/json",
		},
	})
	return res.json()
}

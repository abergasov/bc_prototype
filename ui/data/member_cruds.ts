export async function setMemberAddress(memberID: string, address: string): Promise<number> {
	const res = await fetch(`/api/members`, {
		method: "POST",
		body: JSON.stringify({
			member: memberID,
			address: address,
		}),
		headers: {
			"Content-Type": "application/json",
		},
	})
	return res.status
}

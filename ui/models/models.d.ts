declare module models {
	export interface Member {
		circle_id: string
		member_id: string
		address: string
		role: number
	}
	export interface CircleContract {
		contract_abi: string
		contract_code: string
		params: string[]
	}
	export interface DeployedCircleContract {
		contract_address: string
		transaction_hash: string
		date: string
	}
}

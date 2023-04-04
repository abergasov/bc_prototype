import { Heading, Button, Skeleton, Stack, VStack } from "@chakra-ui/react"
import { useRouter } from "next/router"
import React, { useMemo } from "react"
import { useMoralis, useMoralisWeb3Api } from "react-moralis"
import { useInfiniteQuery } from "react-query"
import Member = models.Member
import Layout from "../components/Layout"
import SingleMember from "../components/circle/SingleMember"
import { setMemberAddress } from "../data/member_cruds"
import {deployCircleContract, deployContract, getCircleContract} from "../data/contract"
import CircleContract from "../components/circle/Contract"

const Home = () => {
	const { chainId: chainIdHex, account: account, isWeb3Enabled, web3: provider } = useMoralis()
	const chainId: string = parseInt(chainIdHex!).toString()

	// if (isWeb3Enabled) {
	// 	console.log("web3 enabled", account)
	// 	let balacne = ""
	// 	provider?.getBalance(account).then((balance) => {
	// 		balacne = balance
	// 		console.log("balance", balance.toString())
	// 	})
	// 	console.log("balance", balacne)
	// }
	//if (!isWeb3Enabled) return <Button onClick={() => provider.enable()}>Connect</Button>
	//const balance = provider.account.getNativeBalance()
	//const address = "0x1..."
	//const { data: nativeBalance } = useEvmNativeBalance({ address })
	const { push } = useRouter()
	const { data: membersList, refetch: refetchMembers } = useInfiniteQuery<Member[]>(
		"feeds",
		async () => {
			const data: Member[] = await fetch(`/api/members`).then((res) => res.json())
			return data
		},
		{}
	)

	const { data: contract, refetch: refetchContract } = useInfiniteQuery<string>(
		"contract",
		async () => getCircleContract(),
		{}
	)

	const deployContract = async () => {
		console.log("deploy_")
		await deployCircleContract()
		await refetchContract()
	}

	const onConnect = async (id: string) => {
		if (!isWeb3Enabled) {
			return
		}
		if (!account) {
			return
		}
		await setMemberAddress(id, account)
		await refetchMembers()
	}

	const allData = useMemo(() => {
		return membersList?.pages.flatMap((page) => page)
	}, [membersList])

	const circleContract = useMemo(() => {
		console.log("contract", contract)
		return contract?.address?.toString() ?? ""
	}, [contract])

	return (
		<Layout title="Home" backRoute="/">
			<VStack spacing={6} width="100%">
				<CircleContract contractAddress={circleContract} deploy={deployContract} />
				<VStack boxShadow="0px 2px 8px #ccc" p={4} borderRadius={6} width="100%" align="flex-start">
					{!membersList && (
						<Stack width="100%">
							<Skeleton height="20px" />
						</Stack>
					)}
					{allData && allData.length == 0 && (
						<Stack width="100%">
							<h1>No members found</h1>
						</Stack>
					)}
					{allData?.map((member) => (
						<SingleMember
							key={member.member_id}
							loggedUser={account ?? ""}
							onConnect={onConnect}
							{...member}
						/>
					))}
				</VStack>
			</VStack>
		</Layout>
	)
}

export default Home

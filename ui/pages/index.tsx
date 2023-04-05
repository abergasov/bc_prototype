import { Heading, Button, Skeleton, Stack, VStack } from "@chakra-ui/react"
import { useRouter } from "next/router"
import React, { useMemo } from "react"
import { useMoralis, useMoralisWeb3Api } from "react-moralis"
import { useInfiniteQuery } from "react-query"
import Member = models.Member
import Layout from "../components/Layout"
import SingleMember from "../components/circle/SingleMember"
import { setMemberAddress } from "../data/member_cruds"
import {deployCircleContract, getCircleContract, saveContract} from "../data/contract"
import CircleContract from "../components/circle/Contract"
import { resetAppData } from "../data/app"
import Web3 from "web3";
import DeployedCircleContract = models.DeployedCircleContract;

const Home = () => {
	const { chainId: chainIdHex, account: account, isWeb3Enabled, web3: provider, Moralis: moralis } = useMoralis()
	const { push } = useRouter()
	const { data: membersList, refetch: refetchMembers } = useInfiniteQuery<Member[]>(
		"feeds",
		async () => {
			const data: Member[] = await fetch(`/api/members`).then((res) => res.json())
			return data
		},
		{}
	)

	const { data: contract, refetch: refetchContract } = useInfiniteQuery<DeployedCircleContract>(
		"contract",
		async () => {
			const data: DeployedCircleContract = await getCircleContract()
			return data
		},
		{}
	)

	const deployContract = async () => {
		console.log("deploy_")
		if (!provider) {
			return
		}
		if (!account) {
			return
		}
		let web3 = new Web3(provider.provider);
		const deployParams = await deployCircleContract()


		let contract = new web3.eth.Contract(JSON.parse(deployParams.contract_abi));
		const params = web3.utils.fromAscii("123")

		contract.options.data = deployParams.contract_code;
		const deployTx = await contract.deploy({
			data: deployParams.contract_code,
			arguments: [params],
		})
		const gas = await deployTx.estimateGas()
		console.log("gas", gas)
		let txHash = ""
		const deployedContract = await deployTx
			.send({
				from: account,
				gas: gas,
			})
			.once("transactionHash", (txhash) => {
				console.log(`Mining deployment transaction ...`);
				console.log(`https://goerli.etherscan.io/tx/${txhash}`);
				txHash = txhash
			});
		console.log(deployedContract.options.address)
		await saveContract(deployedContract.options.address, txHash)
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

	const resetApp = async () => {
		await resetAppData()
		await refetchMembers()
		await refetchContract()
	}

	const allData = useMemo(() => {
		return membersList?.pages.flatMap((page) => page)
	}, [membersList])

	const deployedContract = useMemo(() => {
		if (!contract) {
			return {}
		}
		return contract?.pages?.length > 0 ? contract?.pages[0] : null
	}, [contract])

	return (
		<Layout title="Home" backRoute="/">
			<VStack spacing={6} width="100%">
				<CircleContract {...deployedContract} deploy={deployContract} reset={resetApp} />
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

import {Alert, AlertIcon, Flex, Button, HStack, Text, Heading, Link} from "@chakra-ui/react"
import React from "react"
import DeployedCircleContract = models.DeployedCircleContract;

interface IProps extends DeployedCircleContract {
	deploy: () => void
	reset: () => void
}

const CircleContract = (params: IProps) => {
	const linkContract = `https://goerli.etherscan.io/address/${params.contract_address}`
	return (
		<Flex direction="row" align="center" justify="space-between" py={2} width="100%">
			<Button colorScheme="pink" onClick={() => params.reset()}>
				Reset app
			</Button>
			<Heading>circle id : 372e3c13-a552-4ebe-96fb-7f1bf45cf68c</Heading>
			<HStack spacing={8} align="center">
				{params.contract_address && (
					<Alert status="warning">
						<AlertIcon />
						Contract 
						<Link target={"_blank"} color='teal.500' href={linkContract}>
							{params.contract_address}
						</Link>
					</Alert>
				)}
				{!params.contract_address && <Button onClick={() => params.deploy()}>deploy contract</Button>}
			</HStack>
		</Flex>
	)
}

export default CircleContract

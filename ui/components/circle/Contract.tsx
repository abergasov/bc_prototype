import { Alert, AlertIcon, Flex, Button, HStack, Text, Heading } from "@chakra-ui/react"
import React from "react"

interface IProps {
	deploy: () => void
	reset: () => void
	contractAddress: string
}

const CircleContract = (params: IProps) => {
	return (
		<Flex direction="row" align="center" justify="space-between" py={2} width="100%">
			<Button colorScheme="pink" onClick={() => params.reset()}>
				Reset app
			</Button>
			<Heading>circle id : 372e3c13-a552-4ebe-96fb-7f1bf45cf68c</Heading>
			<HStack spacing={8} align="center">
				{params.contractAddress && (
					<Alert status="warning">
						<AlertIcon />
						Contract deployed at {params.contractAddress}
					</Alert>
				)}
				{!params.contractAddress && <Button onClick={() => params.deploy()}>deploy contract</Button>}
			</HStack>
		</Flex>
	)
}

export default CircleContract

import { ViewIcon } from "@chakra-ui/icons"
import { Alert, AlertIcon, ButtonGroup, Flex, HStack, IconButton, Text } from "@chakra-ui/react"
import React from "react"
import Member = models.Member

interface IProps extends Member {
	onConnect: (id: string) => void
	loggedUser: string
}

const SingleMember = (member: IProps) => {
	const memberType = member.role == 3 ? "DIRECTOR" : "Member"
	const wallet = member.address ? `(${member.address})` : ""
	const isUserLogged = member.address !== "" && member.loggedUser !== "" && member.address === member.loggedUser
	return (
		<Flex direction="row" align="center" justify="space-between" py={2} width="100%">
			<HStack spacing={8} align="center">
				{isUserLogged && (
					<Alert status="success">
						<AlertIcon />
						Logged
					</Alert>
				)}
				<Text>
					<b>{memberType}</b>: {member.member_id} {wallet}
				</Text>
			</HStack>
			<ButtonGroup spacing={2}>
				{!member.address && (
					<IconButton
						aria-label="View"
						icon={<ViewIcon color="white" />}
						colorScheme="green"
						size="sm"
						onClick={() => member.onConnect(member.member_id)}
					/>
				)}
			</ButtonGroup>
		</Flex>
	)
}

export default SingleMember

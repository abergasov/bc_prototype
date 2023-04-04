import { Box, Center, ChakraProvider, Container } from "@chakra-ui/react"
import React from "react"
import Header from "./Header"

type IProps = React.PropsWithChildren<{
	maxWidth?: string
	title: string
	description?: string
	backRoute?: string
}>

const Layout: React.FC<IProps> = ({ title, backRoute, children, maxWidth }) => {
	return (
		<Container maxW="container.xl" w="100%">
			{/*<Box height="100vh">*/}
			{/*	/!*<Header title={title} backRoute={backRoute} />*!/*/}
			{/*	/!*<Container maxW='container.xl' w='100%'  >*!/*/}
			{/*	/!*	<Center>{children}</Center>*!/*/}
			{/*	/!*</Container>*!/*/}
			{/*</Box>*/}
			<Header title={title} backRoute={backRoute} />
			<Container maxW="container.xl" w="100%">
				<Center>{children}</Center>
			</Container>
		</Container>
	)
}

export default Layout

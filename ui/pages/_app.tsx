import React from "react"
import type { AppProps } from "next/app"
import { ChakraProvider } from "@chakra-ui/react"
import { QueryClient, QueryClientProvider } from "react-query"
import { MoralisProvider } from "react-moralis"
import { NotificationProvider } from "web3uikit"

const queryClient = new QueryClient()

const App = ({ Component, pageProps }: AppProps) => {
	return (
		<MoralisProvider initializeOnMount={false}>
			<NotificationProvider>
				<QueryClientProvider client={queryClient}>
					<ChakraProvider resetCSS>
						<Component {...pageProps} />
					</ChakraProvider>
				</QueryClientProvider>
			</NotificationProvider>
		</MoralisProvider>
	)
}

export default App

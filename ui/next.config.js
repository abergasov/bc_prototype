const path = require("path")
/** @type {import('next').NextConfig} */
const nextConfig = {
	reactStrictMode: true,
	sassOptions: {
		includePaths: [path.join(__dirname, "styles")],
	},
	rewrites: async () => {
		return [
			{
				source: "/:path*",
				destination: "http://127.0.0.1:8000/:path*",
			},
		]
	},
}

module.exports = nextConfig

/** @type {import('next').NextConfig} */
const API_URL = "http://localhost:8080/api"

const nextConfig = {
	async rewrites() {
		return [
			{
				source: '/api/:path*',
				destination: `${API_URL}/:path*`,
			},
		]
	},
}


export default nextConfig;

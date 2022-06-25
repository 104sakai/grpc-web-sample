/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: false,
  env: {
    GO_HOST: process.env.GO_HOST,
  }
}

module.exports = nextConfig

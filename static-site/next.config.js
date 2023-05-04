/** @type {import('next').NextConfig} */
const nextConfig = {
  output: "export",
  distDir: "dist",
  trailingSlash: true,
  experimental: {
    appDir: true,
  },
};

module.exports = nextConfig;

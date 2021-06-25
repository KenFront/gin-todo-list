module.exports = {
  reactStrictMode: true,
  rewrites: () => [
    { source: '/api/:path*', destination: 'http://localhost/:path*' }
  ]
}

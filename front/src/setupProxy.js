const { createProxyMiddleware } = require('http-proxy-middleware')

module.exports = function (app) {
  app.use(
    '/api',
    createProxyMiddleware({
      target: 'http://127.0.0.1:8088',
      changeOrigin: true
    })
  )
  app.use(
    '/static/upload',
    createProxyMiddleware({
      target: 'http://127.0.0.1:8088',
      changeOrigin: true
    })
  )
}

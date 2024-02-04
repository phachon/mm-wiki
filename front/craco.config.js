const path = require('path')

module.exports = {
  webpack: {
    // 别名
    alias: {
      '@': path.join(__dirname, 'src')
    },
    optimization: {
      splitChunks: {
        chunks: 'async',
        minSize: 40000,
        maxSize: 1000000,
        name: true,
        maxAsyncRequests: 5, // 最大异步请求数
        maxInitialRequests: 4, // 页面初始化最大异步请求数
        automaticNameDelimiter: '~', // 解决命名冲突
        cacheGroups: {
          vender: {
            name: 'vendor',
            test: /[\\/]node_modules[\\/]/,
            chunks: 'async',
            priority: 10,
            enforce: true
          },
          antd: {
            name: 'antd',
            test: (module) => {
              return /ant/.test(module.context)
            },
            chunks: 'async',
            priority: 11,
            enforce: true
          }
        }
      }
    }
  }
}

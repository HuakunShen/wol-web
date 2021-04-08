module.exports = {
  devServer: {
    proxy: {
      '^/api': {
        target:
          process.env.NODE_ENV === 'production'
            ? 'http://backend:9090'
            : 'http://localhost:9090',
        changeOrigin: true,
      },
    },
  },
};

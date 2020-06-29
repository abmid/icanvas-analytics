// vue.config.js
module.exports = {
    configureWebpack: {
        module: {
            rules: [
              {
                test: require.resolve('jquery'),
                loader: 'expose-loader',
                options: {
                  exposes: ['$', 'jQuery'],
                },
              },
            ],
        },    
    }
  }
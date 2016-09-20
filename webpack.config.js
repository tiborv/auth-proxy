const webpack = require('webpack')
const ExtractTextPlugin = require('extract-text-webpack-plugin')

module.exports = {
  module: {
    loaders: [{
      test: /\.css$/,
      loader: ExtractTextPlugin.extract('style', 'css?-minimize!')
    }, {
      test: /\.png|\.jpe?g|\.gif|\.svg|\.woff|\.woff2|\.ttf|\.eot|\.ico|\.svg$/,
      loader: 'file?name=fonts/[name].[hash].[ext]?'
    }, {
      test: /\.js$/,
      exclude: /node_modules/,
      loader: 'babel'
    }]
  },
  extensions: ['', '.js', '.css'],
  entry: {
    app: './client/app.js',
  },
  devtool: 'eval-source-map',
  output: {
    path: './static/assets/',
    filename: '[name].bundle.js',
    publicPath: '/static/assets/',
  },
  devServer: {
    contentBase: './static',
    historyApiFallback: true,
    hot: true,
    proxy: {
      '/api': {
        target: 'http://localhost:3000/',
        secure: false
      }
    }
  },
  plugins: [
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: JSON.stringify('development')
      }
    }),
    new ExtractTextPlugin('style.min.css')
  ],
}

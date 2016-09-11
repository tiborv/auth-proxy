const ExtractTextPlugin = require('extract-text-webpack-plugin')

module.exports = {
  module: {
    loaders: [{
      test: /\.css$/,
      loader: ExtractTextPlugin.extract('style', 'css?minimize')
    }, {
      test: /\.js$/,
      exclude: /node_modules/,
      loaders: ['babel']
    }]
  },
  extensions: ['', '.js', '.css', '.json'],
  entry: {
    app: './client/index.js',
  },
  devtool: 'source-map',
  output: {
    path: './static/js/',
    filename: '[name].bundle.js',
    publicPath: '/js/',
  },
  plugins: [],
}

const webpack = require('webpack')
const LessPluginCleanCSS = require('less-plugin-clean-css')

module.exports = {
  module: {
    loaders: [{
      test: /\.css$/,
      loader: 'style!css'
    }, {
      test: /\.less?$/,
      loader: 'style!css!less'
    }, {
      test: /\.(eot|woff|woff2|ttf|svg|png|jpe?g|gif)(\?\S*)?$/,
      loader: 'url?limit=100000&name=[name].[ext]'
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
  lessLoader: {
    lessPlugins: [
      new LessPluginCleanCSS({ advanced: true })
    ]
  },
  plugins: [
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: JSON.stringify('development')
      }
    })
  ],
}

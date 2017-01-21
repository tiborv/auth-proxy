const config = require('./webpack.config')
const webpack = require('webpack')
const OptimizeCssAssetsPlugin = require('optimize-css-assets-webpack-plugin')

config.plugins[0] = new webpack.DefinePlugin({
  'process.env': {
    NODE_ENV: JSON.stringify('production')
  }
})
config.plugins.push(new webpack.optimize.DedupePlugin())
config.plugins.push(new webpack.optimize.UglifyJsPlugin({
  minimize: true,
  compressor: { warnings: false },
  output: { comments: false }
}))
config.plugins.push(new webpack.optimize.OccurrenceOrderPlugin())
config.plugins.push(new webpack.DefinePlugin({ __DEV__: false }))
config.plugins.push(new OptimizeCssAssetsPlugin({
  cssProcessorOptions: { discardComments: { removeAll: true } },
  canPrint: true
}))

config.devtool = ''
module.exports = config

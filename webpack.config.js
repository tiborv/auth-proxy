
module.exports = {
  module: {
    loaders: [{
      test: /\.css$/,
      loaders: ['style', 'css']
    },
    {
      test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/,
      loader: 'url-loader?limit=10000&mimetype=application/font-woff' },
    {
      test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
      loader: 'file-loader'
    },
    {
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

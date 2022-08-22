const path = require('path')
const HtmlWebpackPlugin = require('html-webpack-plugin')

module.exports = {
  mode: 'development',
  resolve: {
    extensions: ['.ts', '.tsx', '.js']
  },
  entry: path.resolve(__dirname, 'web/src/index.tsx'),
  devtool: 'source-map',
  module: {
    rules: [
      {
        test: /\.(tsx?)/i,
        exclude: /node_modules/,
        use: 'babel-loader'
      }, {
        test: /\.(ttf|eot|woff2?)$/,
        type: 'asset/resource',
      }, {
        test: /\.svg$/,
        type: 'asset/source'
      }, {
        test: /\.(jpg|png|ico)$/,
        type: 'asset/resource'
      }
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({
      inject: 'body',
      template: path.resolve(__dirname, 'web/index.html')
    })
  ]
}
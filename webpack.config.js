const path = require('path')
const HtmlWebpackPlugin = require('html-webpack-plugin')

const isDev = process.env.NODE_ENV === 'development'

module.exports = {
  mode: isDev ? 'development' : 'production',
  resolve: {
    extensions: ['.ts', '.tsx', '.js']
  },
  entry: path.resolve(__dirname, 'web/src/index.tsx'),
  devtool: isDev ? 'source-map' : undefined,
  output: {
    publicPath: '/ui',
    filename: '[name].bundle.js'
  },
  module: {
    rules: [
      {
        test: /\.(tsx?)/i,
        exclude: /node_modules/,
        use: 'babel-loader'
      }, {
        test: /.css$/i,
        exclude: /node_modules/,
        use: ['style-loader', 'css-loader', 'postcss-loader']
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
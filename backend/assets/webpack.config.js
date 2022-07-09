const path = require('path');
const glob = require('glob');

const scss = {
  test: /\.s[ac]ss$/i,
  use: [
    { loader: 'style-loader' },
    { loader: 'css-loader' },
    { loader: 'resolve-url-loader' },
    { loader: 'sass-loader', options: { sourceMap: true } },
  ]
}

const js = {
  test: /\.js$/,
  exclude: /node_modules/,
  loader: 'swc-loader',
}

module.exports = {
  entry: {
    './js/app.js': glob.sync('./vendor/**/*.js').concat(['./js/app.js'])
  },
  output: {
    filename: 'app.js',
    path: path.resolve(__dirname, '../priv/static/assets')
  },
  module: {
    rules: [
      scss,
      js
    ]
  }
}
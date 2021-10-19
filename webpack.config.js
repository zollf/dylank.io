/* eslint-disable @typescript-eslint/no-var-requires */
const path = require('path');
const glob = require('glob');

const files = {};
glob.sync('./src/*.tsx').forEach((s) => {
  files[s.split('/').slice(-1)[0].replace('.tsx', '')] = s;
});

const ts = {
  test: /\.tsx?$/,
  loader: 'babel-loader',
};

const svg = {
  test: /\.svg$/,
  use: [
    {
      loader: 'babel-loader',
    },
    {
      loader: 'react-svg-loader',
      options: {
        jsx: true,
      },
    },
  ],
};

const scss = {
  test: /\.s[ac]ss$/i,
  use: [
    'style-loader',
    {
      loader: 'css-loader',
      options: {
        import: false,
        modules: {
          localIdentName: '[local]_[hash:base64:5]',
        },
      },
    },
    {
      loader: 'sass-loader',
      options: {
        additionalData: '@import "~/src/styles/globals.scss";',
      },
    },
  ],
  include: /\.module\.s[ac]ss$/,
};

module.exports = {
  mode: 'development',
  entry: {
    ...files,
  },
  output: {
    path: path.resolve(__dirname, 'public/static'),
    publicPath: '/static/',
    filename: '[name].bundle.js',
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js'],
    alias: {
      '@': __dirname,
    },
  },
  module: {
    rules: [ts, scss, svg],
  },
  optimization: {
    minimize: true,
  },
  devServer: {
    static: {
      directory: path.join(__dirname, 'public'),
    },
    compress: true,
    port: 3000,
    allowedHosts: 'all',
  },
};

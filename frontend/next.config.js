/* eslint-disable @typescript-eslint/no-var-requires */
const path = require('path');

module.exports = {
  env: {
    NGINX_HOST: process.env.NGINX_HOST,
  },
  sassOptions: {
    prependData: '@import "globals.scss";',
    includePaths: [path.join(__dirname, 'src/styles')],
  },
  webpack(config, options) {
    config.module.rules.push({
      test: /\.(graphql|gql)$/,
      include: [options.dir],
      use: [{ loader: 'graphql-tag/loader' }],
    });

    config.module.rules.push({
      test: /\.svg$/,
      issuer: /\.tsx?$/,
      include: [options.dir],
      use: [
        'next-swc-loader',
        {
          loader: '@svgr/webpack',
          options: { babel: false },
        },
      ],
    });

    return config;
  },
};

/* eslint-disable @typescript-eslint/no-var-requires */
const path = require('path');

module.exports = {
  presets: [
    [
      '@babel/preset-env',
      {
        targets: {
          esmodules: true,
        },
      },
    ],
    '@babel/preset-react',
    '@babel/preset-typescript',
  ],
  plugins: [
    [
      '@oscarbarrett/babel-plugin-inline-react-svg',
      {
        root: path.resolve(__dirname, 'src'),
        alias: {
          '@': [path.resolve(__dirname, 'src')],
        },
        svgo: {
          plugins: [{ cleanupIDs: false }],
        },
      },
    ],
  ],
};

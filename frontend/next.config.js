/* eslint-disable @typescript-eslint/no-var-requires */
const path = require('path');

module.exports = {
  sassOptions: {
    prependData: '@import "globals.scss";',
    includePaths: [path.join(__dirname, 'src/styles')],
  },
};

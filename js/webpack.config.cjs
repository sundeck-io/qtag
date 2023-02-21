const path = require('path');

module.exports = {
  entry: './src/qtag/entry.js',
  output: {
    filename: 'bundle.cjs',
    path: path.resolve(__dirname, 'dist'),
    library: {
      name: 'qtag',
      type: 'var',
      export: 'default',
    },
  },
  mode: 'production',
  optimization: {
    usedExports: true,
  },
  module: {
    rules: [{
      test: /\.yaml$/,
      use: 'js-yaml-loader',
    }]
  }
};

const path = require('path');
let minify = true;
if (process.env.NO_MINIFY) {
    minify = false;
}

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
    minimize: minify,
  },
  module: {
    rules: [{
      test: /\.yaml$/,
      use: 'js-yaml-loader',
    }]
  }
};

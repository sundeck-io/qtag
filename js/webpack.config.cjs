const path = require('path');

module.exports = {
  entry: './src/ql/extract.js',
  output: {
    filename: 'bundle.cjs',
    path: path.resolve(__dirname, 'dist'),
    library: {
      name: 'CommentQL',
      type: 'var',
      export: 'default',
    },
  },
  mode: 'production',
  optimization: {
    usedExports: true,
  }
};

export default {
  verbose: true,
  testMatch: ['<rootDir>/**/*.test.js'],
  transform: {
    '\\.yaml$': 'yaml-jest-transform',
  },

}

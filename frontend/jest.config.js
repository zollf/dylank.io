/* eslint-disable @typescript-eslint/no-var-requires */
const nextJest = require('next/jest');

const createJestConfig = nextJest({ dir: '.' });

const customJestConfig = {
  cacheDirectory: 'node_modules/.cache/jest',
  collectCoverage: true,
  collectCoverageFrom: ['src/**/*.tsx', '!src/index.tsx'],
  setupFiles: ['<rootDir>/tests/config.ts'],
  setupFilesAfterEnv: ['<rootDir>/tests/setupTests.ts'],
  testEnvironment: 'jest-environment-jsdom',
  testPathIgnorePatterns: ['node_modules'],
  moduleDirectories: ['node_modules'],
  moduleNameMapper: {
    '^@/(.*)$': '<rootDir>/src/$1',
    '\\.(scss)': 'identity-obj-proxy',
  },
  transform: {
    '\\.(gql|graphql)$': 'jest-transform-graphql',
  },
};

module.exports = createJestConfig(customJestConfig);

const VUE_JEST_TRANSFORMER = "@vue/vue2-jest";

module.exports = {
  testMatch: ["<rootDir>/__tests__/**/*_spec.js"],
  moduleFileExtensions: ["js", "json", "vue"],
  cacheDirectory: "<rootDir>/tmp/cache/jest",
  restoreMocks: true,
  transform: {
    "^.+\\.js$": "babel-jest",
    "^.+\\.vue$": VUE_JEST_TRANSFORMER,
    "^.+\\.svg$": VUE_JEST_TRANSFORMER,
  },
  transformIgnorePatterns: [
    "node_modules/(?!(@gitlab/(ui|svgs)|bootstrap-vue)/)",
  ],
};

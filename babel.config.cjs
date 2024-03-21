module.exports = {
  env: {
    development: {
      presets: [["@babel/preset-env"]],
    },
    test: {
      presets: [
        [
          "@babel/preset-env",
          {
            targets: {
              node: "current",
            },
          },
        ],
      ],
    },
  },
};

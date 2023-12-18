const THEME_PATH = "themes/gitlab-docs";

module.exports = {
  publicDir: false,
  build: {
    outDir: `${THEME_PATH}/static/js`,
    rollupOptions: {
      input: {
        main: `${THEME_PATH}/src/main.js`,
      },
      output: {
        entryFileNames: "[name].js",
      },
    },
  },
};

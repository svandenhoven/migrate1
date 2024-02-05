import { defineConfig } from "vite";
import copy from "rollup-plugin-copy";

const THEME_PATH = "themes/gitlab-docs";
const THEME_VENDOR_DIR = `${THEME_PATH}/static/gitlab_ui`;

const GITLAB_THEME_ASSETS = [
  "@gitlab/ui/dist/index.css*",
  "@gitlab/fonts/gitlab-sans/GitLabSans*.woff2",
  "@gitlab/fonts/gitlab-mono/GitLabMono*.woff2",
  "@gitlab/svgs/dist/sprite_icons/*",
];

export default defineConfig({
  publicDir: false,
  build: {
    outDir: `${THEME_PATH}/static/vite`,
    rollupOptions: {
      input: {
        main: `${THEME_PATH}/src/main.js`,
      },
      output: {
        entryFileNames: "[name].js",
        assetFileNames: `[name].[ext]`,
      },
      plugins: [
        copy({
          targets: GITLAB_THEME_ASSETS.map((file) => {
            const { 1: packageName } = file.split("/");
            return {
              src: `./node_modules/${file}`,
              dest: `${THEME_VENDOR_DIR}/${packageName}`,
            };
          }),
        }),
      ],
    },
  },
});

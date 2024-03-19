import { defineConfig } from "vite";
import { globSync } from "glob";
import copy from "rollup-plugin-copy";
import vue2 from "@vitejs/plugin-vue2";
import resolve from "@rollup/plugin-node-resolve";
import fs from "fs";
import path from "path";

const THEME_PATH = "themes/gitlab-docs";
const THEME_VENDOR_DIR = `${THEME_PATH}/static/gitlab_ui`;

const GITLAB_THEME_ASSETS = [
  "@gitlab/ui/dist/index.css*",
  "@gitlab/fonts/gitlab-sans/GitLabSans*.woff2",
  "@gitlab/fonts/gitlab-mono/GitLabMono*.woff2",
  "@gitlab/svgs/dist/sprite_icons/*",
];

// Process all JS files in the theme src directory.
// See https://rollupjs.org/configuration-options/#input
const inputs = globSync(`${THEME_PATH}/src/*.js`).reduce((entries, file) => {
  const name = path.basename(file);
  return { ...entries, [name]: file };
}, {});

export default defineConfig({
  publicDir: false,
  build: {
    outDir: `${THEME_PATH}/static/vite`,
    rollupOptions: {
      input: inputs,
      output: {
        format: "es",
        inlineDynamicImports: false,
        entryFileNames: "[name]",
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
        vue2(),
        resolve(),
      ],
      onwarn(warning, warn) {
        // Ignore known issues from third-party code.
        const ignore = {
          INVALID_ANNOTATION: ["bootstrap-vue"],
        };
        const ignoredKeywords = ignore[warning.code] || [];
        if (
          !ignoredKeywords.some((keyword) => warning.message.includes(keyword))
        ) {
          warn(warning);
        }
      },
    },
  },
  plugins: [
    {
      // Move icons.svg from GitLab UI out of the Vite directory.
      // Otherwise it will 404 when viewing components with icons.
      name: "move-gitlab-ui-icons-svg",
      closeBundle: () => {
        const sourcePath = path.join(THEME_PATH, "static/vite/icons.svg");
        const destinationPath = path.join(THEME_PATH, "static/icons.svg");

        if (fs.existsSync(sourcePath)) {
          fs.rename(sourcePath, destinationPath, (err) => {
            if (err) {
              console.error("Error moving icons.svg:", err);
            }
          });
        }
      },
    },
  ],
});

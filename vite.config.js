import { defineConfig } from "vite";
import { globSync } from "glob";
import copy from "rollup-plugin-copy";
import vue2 from "@vitejs/plugin-vue2";
import resolve from "@rollup/plugin-node-resolve";
import fs from "fs";
import path from "path";

const THEME_PATH = "themes/gitlab-docs";
const THEME_GITLAB_UI_DIR = `${THEME_PATH}/static/gitlab_ui`;

const GITLAB_THEME_ASSETS = [
  "@gitlab/ui/dist/index.css*",
  "@gitlab/fonts/gitlab-sans/GitLabSans*.woff2",
  "@gitlab/fonts/gitlab-mono/GitLabMono*.woff2",
  "@gitlab/svgs/dist/sprite_icons/*",
  "@gitlab/svgs/dist/icons.json",
];

const isWatchMode = process.argv.includes("--watch");

// Process all JS files in the theme src directory.
// See https://rollupjs.org/configuration-options/#input
const inputs = globSync(`${THEME_PATH}/src/*.js`).reduce((entries, file) => {
  const name = path.basename(file);
  return { ...entries, [name]: file };
}, {});

// Define static assets we reference from CSS as external.
// These are copied into the project outside of Vite's CSS compilation step.
const replaceBasePath = (file) => file.replace(`${THEME_PATH}/static`, "..");
const externalAssets = globSync([
  `${THEME_PATH}/static/fa-icons/*.svg`,
  `${THEME_GITLAB_UI_DIR}/fonts/*.woff2`,
  `${THEME_GITLAB_UI_DIR}/svgs/*.svg`,
]).map(replaceBasePath);

// Vite configuration
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
      external: externalAssets,
      plugins: [
        copy({
          targets: [
            ...GITLAB_THEME_ASSETS.map((file) => {
              const { 1: packageName } = file.split("/");
              return {
                src: `./node_modules/${file}`,
                dest: `${THEME_GITLAB_UI_DIR}/${packageName}`,
              };
            }),
            {
              src: "./node_modules/mermaid/dist/mermaid.min.js",
              dest: `${THEME_PATH}/static/vite`,
            },
          ],
          hook: "writeBundle",
          copyOnce: true,
        }),
        vue2(),
        resolve(),
      ],
      // Ignore known issues from third-party code.
      onwarn(warning, warn) {
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
      name: "move-gitlab-ui-icons-svg",
      closeBundle: () => {
        const copy = [
          // Copy icons.svg into the static directory.
          // GitLab UI components look for it at /, not /vite/.
          {
            src: path.join(THEME_PATH, "static/vite/icons.svg"),
            dest: path.join(THEME_PATH, "static/icons.svg"),
          },
          // Copy icons.json to the data directory. We need this
          // for validating icons in shortcodes/icon.html.
          {
            src: path.join(THEME_PATH, "static/gitlab_ui/svgs/icons.json"),
            dest: "./data/icons.json",
          },
        ];

        for (const { src, dest } of copy) {
          // Skip copy if we're in watch mode
          if (isWatchMode && fs.existsSync(dest)) {
            continue;
          }

          if (fs.existsSync(src)) {
            fs.copyFile(src, dest, (err) => {
              if (err) {
                console.error("Error moving icons.svg:", err);
              }
            });
          }
        }
      },
    },
  ],
});

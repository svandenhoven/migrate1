import { defineConfig } from "vite";
import { globSync } from "glob";
import copy from "rollup-plugin-copy";
import yaml from "@rollup/plugin-yaml";
import vue2 from "@vitejs/plugin-vue2";
import resolve from "@rollup/plugin-node-resolve";
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

// Process all JS files in the theme src directory.
// See https://rollupjs.org/configuration-options/#input
const inputs = globSync(`${THEME_PATH}/src/*.js`).reduce((entries, file) => {
  const name = path.basename(file);
  return { ...entries, [name]: file };
}, {});

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
      // Define static assets we reference from CSS as external.
      // Otherwise Vite flags these as unable to resolve
      external: [
        // Theme icons
        "../fa-icons/x-twitter.svg",
        "../fa-icons/facebook.svg",
        "../fa-icons/linkedin.svg",
        "../fa-icons/youtube.svg",
        // Fonts
        "../gitlab_ui/fonts/GitLabSans.woff2",
        "../gitlab_ui/fonts/GitLabMono.woff2",
        "../gitlab_ui/fonts/GitLabSans-Italic.woff2",
        "../gitlab_ui/fonts/GitLabMono-Italic.woff2",
        // GitLab SVG icons
        "../gitlab_ui/svgs/external-link.svg",
        "../gitlab_ui/svgs/link.svg",
        "../gitlab_ui/svgs/information-o.svg",
        "../gitlab_ui/svgs/warning.svg",
        "../gitlab_ui/svgs/flag.svg",
        "../gitlab_ui/svgs/review-warning.svg",
        "../gitlab_ui/svgs/ellipsis_h.svg",
        "../gitlab_ui/svgs/error.svg",
      ],
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
            {
              src: "./node_modules/@gitlab/svgs/dist/icons.svg",
              dest: `${THEME_PATH}/static`,
            },
            // Copy icons.json to the data directory so we can
            // use it to validate the icons shortcode.
            {
              src: "./node_modules/@gitlab/svgs/dist/icons.json",
              dest: "./data",
            },
          ],
          copyOnce: true,
        }),
        yaml(),
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
});

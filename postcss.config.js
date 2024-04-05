import postcssImport from "postcss-import";

export default {
  plugins: {
    "postcss-import": postcssImport({
      path: ["node_modules"],
    }),
    "tailwindcss/nesting": {},
    tailwindcss: {},
    autoprefixer: {},
  },
};

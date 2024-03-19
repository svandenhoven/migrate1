/* eslint-disable import/no-default-export */
import defaultPreset from "@gitlab/ui/tailwind.defaults";

/** @type {import('tailwindcss').Config} */
export default {
  presets: [defaultPreset],
  content: ["themes/**/*.html", "themes/**/*.vue"],
};

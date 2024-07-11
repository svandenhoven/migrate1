#!/usr/bin/env node

/**
 * Detects legacy CSS utility classes
 *
 * These classes should be updated to use
 * their Tailwind equivalents.
 *
 * We can drop this once we've copied all
 * templates (HTML, Vue) over from the Nanoc
 * site.
 *
 * Usage:
 * make lint-frontend
 *
 * To fix any found classes automatically, run the script
 * with the --fix flag:
 *
 * ./scripts/frontend/find_deprecated_classes.js --fix
 */

import fs from "fs";
import path from "path";
import https from "https";

const jsonUrl =
  "https://gitlab.com/gitlab-org/gitlab/-/raw/master/scripts/frontend/tailwind_equivalents.json?ref_type=heads";
const searchDir = "themes/gitlab-docs";

const shouldFix = process.argv.includes("--fix");

// No need to flag third-party or compiled code
const ignorePaths = [
  "themes/gitlab-docs/static/gitlab_ui",
  "themes/gitlab-docs/static/vite",
  "themes/gitlab-docs/assets/css",
];

function fetchJson(url) {
  return new Promise((resolve, reject) => {
    https
      .get(url, (res) => {
        let data = "";
        res.on("data", (chunk) => {
          data += chunk;
        });
        res.on("end", () => {
          try {
            resolve(JSON.parse(data));
          } catch (e) {
            reject(e);
          }
        });
      })
      .on("error", reject);
  });
}

function shouldIgnore(filePath) {
  return ignorePaths.some(
    (ignorePath) =>
      filePath === ignorePath || filePath.startsWith(ignorePath + path.sep),
  );
}

function processFiles(dir, equivalents) {
  let foundDeprecated = false;

  function process(currentDir) {
    fs.readdirSync(currentDir, { withFileTypes: true }).forEach((dirent) => {
      const filePath = path.join(currentDir, dirent.name);

      if (shouldIgnore(filePath)) {
        return;
      }

      if (dirent.isDirectory()) {
        process(filePath);
      } else {
        let content = fs.readFileSync(filePath, "utf8");
        let fileChanged = false;

        Object.entries(equivalents).forEach(([oldClass, newClass]) => {
          if (content.includes(oldClass)) {
            foundDeprecated = true;
            if (shouldFix && newClass !== null) {
              content = content.split(oldClass).join(newClass);
              fileChanged = true;
              console.info(
                `Fixed in ${filePath}: Replaced ${oldClass} with ${newClass}`,
              );
            } else {
              console.info(
                `Deprecated utility class found in ${filePath}. Replace ${oldClass} with ${newClass || "null"}`,
              );
            }
          }
        });

        if (fileChanged) {
          fs.writeFileSync(filePath, content, "utf8");
        }
      }
    });
  }

  process(dir);
  return foundDeprecated;
}

async function main() {
  try {
    const equivalents = await fetchJson(jsonUrl);
    const foundDeprecated = processFiles(searchDir, equivalents);

    if (foundDeprecated) {
      if (shouldFix) {
        console.info(
          "Deprecated utility classes were found and fixed where possible.",
        );
        process.exit(0);
      } else {
        console.error(
          "Deprecated utility classes were found. Run with --fix to automatically update them.",
        );
        process.exit(1);
      }
    } else {
      console.info("No deprecated utility classes were found.");
      process.exit(0);
    }
  } catch (error) {
    console.error("An error occurred:", error);
    process.exit(1);
  }
}

main();

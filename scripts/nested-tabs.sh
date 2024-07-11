#!/bin/bash

# Report on instances of tabs nested 2+ levels deep in lists
# These work on the Nanoc site, but not on the Hugo site.
# Until we resolve https://gitlab.com/gitlab-org/technical-writing-group/gitlab-docs-hugo/-/issues/87,
# we should flag these and plan to refactor them prior to launching
# the Hugo site.

DOCS_DIRS=("../gitlab/doc" "../gitlab-operator/doc" "../gitlab-runner/docs" "../charts-gitlab/doc" "../omnibus-gitlab/doc")

search_directory() {
  local dir="$1"
  echo "Searching in directory: $dir"
  find "$dir" -type f -name "*.md" | while read -r file; do
    # Find tabs with 6+ spaces before them.
    # This indicates they're in an element that's nested 2+ levels deep.
    sed -n '/^      {{< tabs >}}/=' "$file" | while read -r line_number; do
      echo "File: $file"
      echo "Line $line_number: $(sed -n "${line_number}p" "$file")"
      echo
    done
  done
}

# Loop through each directory and search
for dir in "${DOCS_DIRS[@]}"; do
  if [ -d "$dir" ]; then
    search_directory "$dir"
  else
    echo "Directory not found: $dir"
  fi
done

#!/bin/bash

# Verify the asdf installation, then
# install and update asdf plugins.

if ! command -v asdf 2> /dev/null; then
  printf "ERROR: asdf not found. See https://asdf-vm.com/guide/getting-started.html for installation directions.\n"
  exit 1
fi

if [ ! -f .tool-versions ]; then
  printf "ERROR: .tool-versions not found.\n"
  exit 1
fi

while IFS= read -r LINE; do
  # Skip code comments
  if [[ $LINE == \#* ]]; then
    continue
  fi

  DEPENDENCY=$(echo "$LINE" | awk '{print $1}')
  asdf plugin add "$DEPENDENCY" || true
  asdf plugin update "$DEPENDENCY"
done < .tool-versions

asdf install

# Use corepack to add Yarn
# - https://yarnpkg.com/getting-started/install
# - https://github.com/asdf-vm/asdf-nodejs#corepack
corepack enable && asdf reshim nodejs

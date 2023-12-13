#!/bin/bash

source scripts/helpers/colors.sh

# Verify the asdf installation, then
# install and update asdf plugins.

if ! command -v asdf 2> /dev/null; then
  printf "${ERROR}ERROR: asdf not found. See https://asdf-vm.com/guide/getting-started.html for installation directions.${END}\n"
  exit 1
fi

if [ ! -f .tool-versions ]; then
  printf "${ERROR}ERROR: .tool-versions not found.${END}\n"
  exit 1
fi

while IFS= read -r line; do
  dependency=$(echo "$line" | awk '{print $1}')
  asdf plugin add "$dependency" || true
  asdf plugin update "$dependency"
done < .tool-versions

asdf install

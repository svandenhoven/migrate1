#!/usr/bin/env bash

COLOR_RED="\e[31m"
COLOR_YELLOW="\e[33m"
COLOR_GREEN="\e[32m"
COLOR_RESET="\e[39m"

if [ ! -f .tool-versions ]; then
  printf "${COLOR_RED}ERROR: .tool-versions file not found!${COLOR_RESET}\n"
  exit 1
elif [ ! -f Brewfile ]; then
  printf "${COLOR_RED}ERROR: Brewfile file not found!${COLOR_RESET}\n"
  exit 1
fi

if command -v asdf &> /dev/null; then
  printf "${COLOR_GREEN}INFO: asdf found! Using asdf to install dependencies...${COLOR_RESET}\n"

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
  corepack enable && asdf reshim nodejs

elif command -v mise &> /dev/null; then
    printf "${COLOR_GREEN}INFO: mise found! Using mise to install dependencies...${COLOR_RESET}\n"
    mise install
    corepack enable
else
  printf "${COLOR_YELLOW}WARNING: asdf and mise not found! For more information, see doc/setup.md.${COLOR_RESET}\n"
fi

if command -v brew &> /dev/null; then
  printf "${COLOR_GREEN}INFO: Homebrew found! Using Homebrew to install dependencies...${COLOR_RESET}\n"
  brew bundle
else
  printf "${COLOR_YELLOW}WARNING: Homebrew not found! For more information, see doc/setup.md.${COLOR_RESET}\n"
fi

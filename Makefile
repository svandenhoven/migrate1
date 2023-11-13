.PHONY: all setup

INFO = \033[32m
END = \033[0m

setup: install-asdf-dependencies

install-asdf-dependencies:
	@printf "\n$(INFO)INFO: Installing and updating asdf dependencies...$(END)\n"
	@sh ./scripts/install-asdf-dependencies.sh

shellcheck-tests:
	@printf "\n$(INFO)INFO: Running shellcheck tests...$(END)\n"
	@shellcheck scripts/*.sh

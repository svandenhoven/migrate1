.PHONY: all setup go-tests test clone-docs-projects

INFO = \033[32m
END = \033[0m

setup: install-asdf-dependencies

install-asdf-dependencies:
	@printf "\n$(INFO)INFO: Installing and updating asdf dependencies...$(END)\n"
	@sh ./scripts/install-asdf-dependencies.sh

shellcheck-tests:
	@printf "\n$(INFO)INFO: Running shellcheck tests...$(END)\n"
	@shellcheck scripts/*.sh

go-tests:
	@printf "\n$(INFO)INFO: Running Go tests...$(END)\n"
	@go test ./scripts/...

test: shellcheck-tests go-tests

clone-docs-projects:
	@printf "\n$(INFO)INFO: Fetching docs content sources...$(END)\n"
	@go run ./scripts/clone_projects.go

.PHONY: all setup go-tests test clone-docs-projects view

INFO = \033[32m
END = \033[0m

setup: install-asdf-dependencies

install-asdf-dependencies:
	@printf "\n$(INFO)INFO: Installing and updating asdf dependencies...$(END)\n"
	@scripts/install-asdf-dependencies.sh

test: shellcheck-tests go-tests

shellcheck-tests:
	@printf "\n$(INFO)INFO: Running shellcheck tests...$(END)\n"
	@shellcheck scripts/*.sh

go-tests:
	@printf "\n$(INFO)INFO: Running Go tests...$(END)\n"
	@go test ./scripts/...

clone-docs-projects:
	@printf "\n$(INFO)INFO: Fetching docs content sources...$(END)\n"
	@go run scripts/clone_projects.go
	@printf "\n$(INFO)INFO: Running content update scripts...$(END)\n"
	@scripts/content-post-process.sh

view:
	@hugo serve

.PHONY: all clean test setup

INFO = \033[32m
END = \033[0m

setup: install-asdf-dependencies install-nodejs-dependencies

install-asdf-dependencies:
	@printf "\n$(INFO)INFO: Installing and updating asdf dependencies...$(END)\n"
	@scripts/install-asdf-dependencies.sh

install-nodejs-dependencies:
	@yarn install

test: shellcheck-tests lint-go go-tests jest-tests lint-frontend yaml-tests

shellcheck-tests:
	@printf "\n$(INFO)INFO: Running shellcheck tests...$(END)\n"
	@shellcheck scripts/*.sh

lint-go:
	@printf "\n$(INFO)INFO: Running Go linting...$(END)\n"
	@golangci-lint run

go-tests:
	@printf "\n$(INFO)INFO: Running Go tests...$(END)\n"
	@go test ./scripts/...

jest-tests:
	@printf "\n$(INFO)INFO: Running JavaScript tests...$(END)\n"
	@yarn test --passWithNoTests

lint-frontend:
	@printf "\n$(INFO)INFO: Running frontend linters...$(END)\n"
	@yarn eslint
	@yarn prettier
	@yarn stylelint

yaml-tests:
	@printf "\n$(INFO)INFO: Running YAML tests...$(END)\n"
	@yamllint .

clone-docs-projects:
	@printf "\n$(INFO)INFO: Fetching docs content sources...$(END)\n"
	@go run scripts/clone_projects.go
	@printf "\n$(INFO)INFO: Running content update scripts...$(END)\n"
	@scripts/content-post-process.sh

view:
	@yarn build
	@hugo serve

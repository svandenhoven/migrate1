INFO = \033[32m
END = \033[0m

.PHONY: setup
setup: install-asdf-dependencies install-nodejs-dependencies

.PHONY: install-asdf-dependencies
install-asdf-dependencies:
	@printf "\n$(INFO)INFO: Installing and updating asdf dependencies...$(END)\n"
	@scripts/install-asdf-dependencies.sh

.PHONY: install-nodejs-dependencies
install-nodejs-dependencies:
	@printf "\n$(INFO)INFO: Installing Node.js packages...$(END)\n"
	@yarn install

.PHONY: test
test: lint-shell-scripts lint-go go-tests jest-tests lint-frontend lint-yaml

.PHONY: lint-shell-scripts
lint-shell-scripts:
	@printf "\n$(INFO)INFO: Running shell script linting...$(END)\n"
	@shellcheck scripts/*.sh

.PHONY: lint-go
lint-go:
	@printf "\n$(INFO)INFO: Running Go linting...$(END)\n"
	@golangci-lint run

.PHONY: go-tests
go-tests:
	@printf "\n$(INFO)INFO: Running Go tests...$(END)\n"
	@go test ./scripts/...

.PHONY: jest-tests
jest-tests: install-nodejs-dependencies
	@printf "\n$(INFO)INFO: Running JavaScript tests...$(END)\n"
	@yarn test --passWithNoTests

.PHONY: lint-frontend
lint-frontend: install-nodejs-dependencies
	@printf "\n$(INFO)INFO: Running frontend linters...$(END)\n"
	@yarn eslint
	@yarn prettier
	@yarn stylelint

.PHONY: lint-yaml
lint-yaml:
	@printf "\n$(INFO)INFO: Running YAML tests...$(END)\n"
	@yamllint .

.PHONY: clone-docs-projects
clone-docs-projects:
	@printf "\n$(INFO)INFO: Fetching docs content sources...$(END)\n"
	@go run scripts/clone_projects.go
	@printf "\n$(INFO)INFO: Running content update scripts...$(END)\n"
	@scripts/content-post-process.sh

.PHONY: view
view:
	@yarn build
	@hugo serve

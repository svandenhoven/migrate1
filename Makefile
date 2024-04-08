INFO = \033[32m
END = \033[0m

include makefiles/*.mk

.PHONY: install-dependencies
install-dependencies:
	@printf "\n$(INFO)INFO: Installing and updating dependencies...$(END)\n"
	@scripts/install-dependencies.sh

.PHONY: install-nodejs-dependencies
install-nodejs-dependencies:
	@printf "\n$(INFO)INFO: Installing Node.js packages...$(END)\n"
	@yarn install

.PHONY: clone-docs-projects
clone-docs-projects:
	@printf "\n$(INFO)INFO: Fetching docs content sources...$(END)\n"
	@go run cmd/gldocs/main.go clone
	@printf "\n$(INFO)INFO: Running content update scripts...$(END)\n"
	@scripts/content-post-process.sh

.PHONY: all
all: clean setup test

.PHONY: clean
clean:
	@printf "\n$(INFO)INFO: Removing ephemeral directories...$(END)\n"
	@rm -rfv public resources node_modules

.PHONY: setup
setup: install-dependencies install-nodejs-dependencies

.PHONY: test
test: lint-markdown lint-shell-scripts lint-go lint-frontend lint-yaml go-tests jest-tests markdown-link-tests

.PHONY: view
view:
	@yarn build
	@hugo serve

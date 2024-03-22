.PHONY: lint-markdown
lint-markdown:
	@printf "\n$(INFO)INFO: Running Markdown linting...$(END)\n"
	@yarn markdownlint

.PHONY: lint-shell-scripts
lint-shell-scripts:
	@printf "\n$(INFO)INFO: Running shell script linting...$(END)\n"
	@shellcheck scripts/*.sh

.PHONY: lint-go
lint-go:
	@printf "\n$(INFO)INFO: Running Go linting...$(END)\n"
	@golangci-lint run

.PHONY: lint-frontend
lint-frontend: install-nodejs-dependencies
	@printf "\n$(INFO)INFO: Running frontend linters...$(END)\n"
	@yarn eslint
	@yarn prettier
	@yarn stylelint

.PHONY: lint-yaml
lint-yaml:
	@printf "\n$(INFO)INFO: Running YAML tests...$(END)\n"
	@yarn yamllint "**/*.(yaml|yml)" --ignore=public

.PHONY: all clean test

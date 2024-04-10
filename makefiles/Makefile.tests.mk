.PHONY: go-tests
go-tests:
	@printf "\n$(INFO)INFO: Running Go tests...$(END)\n"
	@go test ./cmd/gldocs/tasks/...

.PHONY: jest-tests
jest-tests: install-nodejs-dependencies
	@printf "\n$(INFO)INFO: Running JavaScript tests...$(END)\n"
	@yarn test

# See also test:markdown-links job in ../.gitlab-ci.yml
.PHONY: markdown-link-tests
markdown-link-tests:
	@printf "\n$(INFO)INFO: Running Markdown link tests...$(END)\n"
	@lychee --offline --include-fragments README.md **/*.md

.PHONY: kramdown-audit
kramdown-audit: clone-docs-projects
	@printf "${INFO}Auditing for Kramdown markup...${END}\n"
	@scripts/kramdown-audit.sh

.PHONY: all clean test

.PHONY: go-tests
go-tests:
	@printf "\n$(INFO)INFO: Running Go tests...$(END)\n"
	@go test ./scripts/...

.PHONY: jest-tests
jest-tests: install-nodejs-dependencies
	@printf "\n$(INFO)INFO: Running JavaScript tests...$(END)\n"
	@yarn test --passWithNoTests

.PHONY: all clean test

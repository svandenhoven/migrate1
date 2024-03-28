.PHONY: go-tests
go-tests:
	@printf "\n$(INFO)INFO: Running Go tests...$(END)\n"
	@go test ./cmd/gldocs/tasks/...

.PHONY: jest-tests
jest-tests: install-nodejs-dependencies
	@printf "\n$(INFO)INFO: Running JavaScript tests...$(END)\n"
	@yarn test

.PHONY: all clean test

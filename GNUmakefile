TEST?=$$(go list ./... | grep -v 'vendor')
PKG_NAME=azurermfutures


fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -w -s ./$(PKG_NAME)

fmtcheck:
	echo $(TEST)
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

testacc: fmtcheck
	@echo "Running acceptance tests"
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 10s

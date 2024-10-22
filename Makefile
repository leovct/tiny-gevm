##@ Help

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


##@ Lint

# Add missing and remove unused modules.
.PHONY: tidy
tidy:
	go mod tidy

# Run `go fmt` against code.
.PHONY: fmt
fmt:
	go fmt ./...

# Run `go vet` and `shadow` (which reports shadowed variables) against code.
# https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/shadow
# `go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest`
.PHONY: vet
vet:
	go vet ./...
	shadow ./...

# Run `golangci-lint` against code.
# `golangci-lint` runs `gofmt`, `govet`, `staticcheck` and other linters.
# https://golangci-lint.run/usage/install/#local-installation
.PHONY: golangci-lint
golangci-lint:
	golangci-lint run --fix --timeout 5m

.PHONY: lint
lint: tidy vet golangci-lint ## Run linters.

##@ Test

.PHONY: test
test: ## Run tests.
	go test -race -v -coverprofile=coverage.out go-evm/evm
	go tool cover -func coverage.out

.PHONY: coverage
coverage: test ## Open HTML coverage report.
	go tool cover -html=coverage.out

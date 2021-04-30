

PKG_VERSION := 0.1.0
GIT_COMMIT  ?= $(shell git rev-parse --short HEAD 2> /dev/null || true)
BUILD_DATE  := $(shell date -u +%Y-%m-%dT%T 2> /dev/null)


.PHONY: clean
clean:  ## Remove temporary files and build artifacts
	go clean -v ./...
	rm -rf bin
	rm -f coverage.out

.PHONY: cover
cover: test-unit  ## Run unit tests and open the coverage report
	go tool cover -html=coverage.out

.PHONY: fmt
fmt:  ## Run gofmt on all files
	gofmt -s -w .

.PHONY: lint
lint:  ## Lint project source files
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.39.0 golangci-lint run

.PHONY: test-unit
test-unit:  ## Run unit tests
	go test -race -cover -run Unit -coverprofile=coverage.out -covermode=atomic ./...

.PHONY: version
version: ## Print the version
	@echo "${PKG_VERSION}"

.PHONY: help
help:  ## Print usage information
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help

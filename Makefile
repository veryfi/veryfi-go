PKG_VERSION := v0.1.12
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

.PHONY: github-tag
github-tag:  ## Create and push a tag with the current client version
	git tag -a ${PKG_VERSION} -m "Veryfi Go Client ${PKG_VERSION}"
	git push -u origin ${PKG_VERSION}

.PHONY: lint
lint:  ## Lint project source files
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.39.0 golangci-lint run

.PHONY: test-unit
test-unit:  ## Run unit tests
	go test -race -cover -run Unit -coverprofile=coverage.out -covermode=atomic ./...

.PHONY: test-integration
test-integration:  ## Run integration tests
	# CLIENT_ID=FIXME USERNAME=FIXME API_KEY=FIXME go test -race -cover -run Integration -coverprofile=coverage.out -covermode=atomic ./...
	CLIENT_ID=vrf1pF9CSF19wGpasWeTsfVueKUPesnMJTMVp58 CLIENT_SECRET=JgG6jB2oJMLbHwBN5jEd6h3f3Lj1o9gK5kcz2xAf8Kyi2X1PNaJ6F6B0Sv4YcOsSllGkMofgUiZV5ZTNsjkk6bXyctGGAdfcratu4Dl2CA2XtU6En5icHxjVRUNoSFGP USERNAME=hoanhan.api API_KEY=2e53375fb12d450e60f486ec7a61b3w20 go test -race -cover -run Integration -coverprofile=coverage.out -covermode=atomic ./...

.PHONY: version
version: ## Print the version
	@echo "${PKG_VERSION}"

.PHONY: help
help:  ## Print usage information
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help

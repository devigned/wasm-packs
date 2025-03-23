SHELL:=/usr/bin/env bash
ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
.DEFAULT_GOAL:=help

## --------------------------------------
## Build
## --------------------------------------

##@ Build:

.PHONY: build
build: build-base-images build-builder ## Build all of the things \o/

.PHONY: build-base-images
build-base-images: ## Build the base images.
	./base-images/build.sh wasm

.PHONY: build-builder
build-builder: ## Build the builder image.
	pack builder create wasm/demo-builder:wasm --config ./builders/wasm/builder.toml


## --------------------------------------
## Test
## --------------------------------------

##@ Test:

.PHONY: test
test: test-js test-compose ## Run all tests.

.PHONY: test-compose
test-compose: ## Test the compose project.
	pack build test-wasm-compose --builder wasm/demo-builder:wasm --path apps/compose

.PHONY: test-js
test-js: ## Test the JS app.
	pack build test-wasm-js --builder wasm/demo-builder:wasm --path apps/js

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

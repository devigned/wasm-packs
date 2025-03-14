SHELL:=/usr/bin/env bash
ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
SERVER_COMPONENT_NAME := server.component.wasm
GO_COMPONENT_NAME := add.wasm
SERVER_COMPONENT_PATH := $(ROOT_DIR)/app-js/$(SERVER_COMPONENT_NAME)
GO_COMPONENT_PATH := $(ROOT_DIR)/service-go/$(GO_COMPONENT_NAME)
COMPOSED_COMPONENT_PATH := $(ROOT_DIR)/composed.wasm
.DEFAULT_GOAL:=help

## --------------------------------------
## Build
## --------------------------------------

##@ Build:

.PHONY: build
build: build-js build-go build-compose ## Run all build steps.

.PHONY: build-js
build-js: ## Build the JavaScript app.
	$(info "Building JavaScript Component...")
	@cd "$(ROOT_DIR)/app-js" && npm install && npm run build

.PHONY: build-go
build-go: ## Build the Go app.
	$(info "Building Golang Component...")
	@cd $(ROOT_DIR)/service-go && go generate && tinygo build -target=wasip2 -o add.wasm --wit-package ./wit --wit-world adder service.go

.PHONY: build-compose
build-compose: ## Build the Docker Compose app.
	$(info "Fusing components together with wac...")
	@wac plug --plug $(GO_COMPONENT_PATH) -o composed.wasm $(SERVER_COMPONENT_PATH)

.PHONY: run
run: build ## Run the JavaScript app.
	$(info "Running Component Application Using Wasmtime...")
	wasmtime serve -S cli $(COMPOSED_COMPONENT_PATH)

.PHONY: build-and-run
build-and-run: build ## Build and run the JavaScript app.
	$(info "Running Component Application Using Wasmtime...")
	wasmtime serve -S cli $(COMPOSED_COMPONENT_PATH)

##@ Cleanup:

.PHONY: clean
clean: ## Clean up the build artifacts.
	@echo "Cleaning up build artifacts..."
	@rm -rf $(SERVER_COMPONENT_PATH) $(GO_COMPONENT_PATH) $(COMPOSED_COMPONENT_PATH)

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

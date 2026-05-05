# GLOBAL BUILD VARS
BIN_NAME_PREFIX?=spinner
GOROOT=$(shell go env GOROOT)

# API BUILD VARS
API_BIN_OUT=./bin/$(BIN_NAME_PREFIX)-api

# SPINNER BUILD VARS
APP_WASM_OUT=./web/static/main.wasm
APP_WASM_EXEC=./web/static/wasm_exec.js
APP_LINUX_BIN_OUT = ./bin/$(BIN_NAME_PREFIX)-app
# APP_API_* variables are injected into the WASM spinner binary to specify where the API is
APP_API_HOST?=https://badmovie2.api.acid1.xyz
APP_API_PORT?=443
APP_WASM_BUILD_FLAGS=-tags="js wasm" \
					 -ldflags="-X 'main.APIHost=$(APP_API_HOST)' -X 'main.APIPort=$(APP_API_PORT)'"
APP_LINUX_BUILD_FLAGS="-tags=native"

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all
all: build-wasm

.PHONY: build-api
build-api:  ## Build the spinner's backend API
	go build -o $(API_BIN_OUT) ./cmd/api

.PHONY: build-linux
build-linux:  ## Build the spinner as a Linux binary
	go build $(APP_LINUX_BUILD_FLAGS) -o $(APP_LINUX_BIN_OUT) ./cmd/spinner

.PHONY: copy-wasm-exec
copy-wasm-exec:  ## Copy the wasm_exec.js dependency into the html app build directory
	cp $(GOROOT)/lib/wasm/wasm_exec.js $(APP_WASM_EXEC)

.PHONY: build-wasm
build-wasm: copy-wasm-exec  ## Build the spinner WASM binary
	GOOS=js GOARCH=wasm go build $(APP_WASM_BUILD_FLAGS) -o $(APP_WASM_OUT) ./cmd/spinner/main.go

.PHONY: debug
debug: ## Run the spinner with live reload as a linux binary through delve (see `.air-spinner.toml` for debugger connection details)
	bash scripts/run_dlv.sh spinner

.PHONY: run
run:  ## Run the spinner linux binary with live reload but no debugger
	go tool air -c .air-spinner.toml -build.entrypoint bin/$(BIN_NAME)-spinner

.PHONY: run-api
run-api:  ## Run the spinner API with live reload with delve debugger
	go tool air -c .air-api.toml -build.entrypoint bin/$(BIN_NAME)-api

.PHONY: gen-docs
gen-docs:  ## Generate API documentation
	go tool swag init -g cmd/api/main.go

.PHONY: format
format:  ## Format codebase
	go fmt ./...
	go tool swag fmt

.PHONY: clean
clean:  ## Clean up builds and reset to a clean state
	rm -f $(APP_WASM_OUT) $(APP_WASM_EXEC)

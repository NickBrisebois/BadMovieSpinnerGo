# OUTPUT BUILD ROOTS
GOROOT=$(shell go env GOROOT)
BUILD_DIR=./bin
WASM_BUILD_DIR=./web/static/

# CONFIGURATION
BIN_NAME_PREFIX?=spinner
APP_API_HOST?=https://badmovie2.api.acid1.xyz
APP_API_PORT?=443
APP_WASM_BUILD_FLAGS=-tags="js wasm" \
					 -ldflags="-X 'main.APIHost=$(APP_API_HOST)' -X 'main.APIPort=$(APP_API_PORT)'"
APP_LINUX_BUILD_FLAGS="-tags=native"

# BUILD OUTPUT BINARY FILENAMES
API_BIN=$(BIN_NAME_PREFIX)-api
APP_LINUX_BIN=$(BIN_NAME_PREFIX)-app
APP_WASM_BIN=$(BIN_NAME_PREFIX)-wasm

# BUILD OUTPUT FULL PATHS
APP_WASM_EXEC_OUT=$(WASM_BUILD_DIR)/wasm_exec.js
APP_WASM_OUT = $(WASM_BUILD_DIR)/$(APP_WASM_BIN)
APP_LINUX_BIN_OUT = $(BUILD_DIR)/$(BIN_NAME_PREFIX)-app
API_BIN_OUT = $(BUILD_DIR)/$(API_BIN)

# INSTALL PATHS
API_INSTALL_DIR?=/opt/spinner/api/
WASM_INSTALL_DIR?=/opt/spinner/web/static
LINUX_INSTALL_DIR?=/opt/spinner/linux


.PHONY: help
help:  ## Print this help
	@echo "Usage: make [target]"
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all
all: build-wasm build-api

.PHONY: build-api
build-api:  ## Build the spinner's backend API
	go build -o $(API_BIN) ./cmd/api

.PHONY: install-api
install-api:  ## Copy API binary to install target directory
	mkdir -p $(LINUX_INSTALL_DIR)
	cp $(API_BIN) $(LINUX_INSTALL_DIR)/$(API_BIN)

.PHONY: build-linux
build-linux:  ## Build the spinner as a Linux binary
	go build $(APP_LINUX_BUILD_FLAGS) -o $(APP_LINUX_BIN_OUT) ./cmd/spinner

.PHONY: install-linux
install-linux:  ## Copy Linux binary to install target directory
	mkdir -p $(LINUX_INSTALL_DIR)
	cp $(APP_LINUX_BIN_OUT) $(LINUX_INSTALL_DIR)/$(APP_LINUX_BIN)

.PHONY: copy-wasm-exec
copy-wasm-exec:  ## Copy the wasm_exec.js dependency into the html app build directory
	cp $(GOROOT)/lib/wasm/wasm_exec.js $(APP_WASM_EXEC_OUT)

.PHONY: build-wasm
build-wasm: copy-wasm-exec  ## Build the spinner WASM binary
	GOOS=js GOARCH=wasm go build $(APP_WASM_BUILD_FLAGS) -o $(APP_WASM_OUT) ./cmd/spinner/main.go

.PHONY: install-wasm
install-wasm:  ## Copy WASM binary to install target directory
	mkdir -p $(WASM_INSTALL_DIR)
	cp $(APP_WASM_BIN) $(WASM_INSTALL_DIR)/$(APP_WASM_BIN)

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

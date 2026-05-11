# OUTPUT BUILD ROOTS
GOROOT=$(shell go env GOROOT)
BUILD_DIR=./bin
WASM_BUILD_DIR=./internal/web/html/static

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
APP_WASM_BIN=$(BIN_NAME_PREFIX)
APP_WEB_BIN=$(BIN_NAME_PREFIX)-web

# BUILD OUTPUT FULL PATHS
APP_WASM_EXEC_OUT=$(WASM_BUILD_DIR)/wasm_exec.js
APP_WASM_OUT = $(WASM_BUILD_DIR)/$(APP_WASM_BIN).wasm
APP_LINUX_BIN_OUT = $(BUILD_DIR)/$(BIN_NAME_PREFIX)
APP_WEB_BIN_OUT = $(BUILD_DIR)/$(APP_WEB_BIN)
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

# ===============
# ---- BUILD ----
# ===============

.PHONY: clean
clean:  ## Clean up builds and reset to a clean state
	rm -f $(APP_WASM_OUT) $(APP_WASM_EXEC_OUT)
	rm -f $(BUILD_DIR)/*

###### API
.PHONY: build-api
build-api:  ## Build the spinner's backend API
	go build -o $(API_BIN_OUT) ./cmd/api

###### SPINNER
.PHONY: build-linux
build-linux:  ## Build the spinner as a Linux binary
	go build $(APP_LINUX_BUILD_FLAGS) -o $(APP_LINUX_BIN_OUT) ./cmd/spinner

.PHONY: build-wasm
build-wasm: ## Build the spinner WASM binary
	cp $(GOROOT)/lib/wasm/wasm_exec.js $(APP_WASM_EXEC_OUT)
	GOOS=js GOARCH=wasm go build $(APP_WASM_BUILD_FLAGS) -o $(APP_WASM_OUT) ./cmd/spinner/main.go

.PHONY: build-web
build-web: build-wasm  ## Build fully self-contained front-end web server binary with WASM spinner
	go build -o $(APP_WEB_BIN_OUT) ./cmd/web


# =================
# ---- INSTALL ----
# =================
.PHONY: install-linux
install-linux:  ## Copy Linux binary to install target directory
	mkdir -p $(LINUX_INSTALL_DIR)
	cp $(APP_LINUX_BIN_OUT) $(LINUX_INSTALL_DIR)/$(APP_LINUX_BIN)

.PHONY: install-api
install-api:  ## Copy API binary to install target directory
		mkdir -p $(API_INSTALL_DIR)
		cp $(API_BIN) $(API_INSTALL_DIR)/$(API_BIN)

# ===============
# ---- DEBUG ----
# ===============
.PHONY: debug
debug: ## Run the spinner with live reload as a linux binary through delve (see `.air-spinner.toml` for debugger connection details)
	bash scripts/run_dlv.sh spinner

.PHONY: run
run:  ## Run the spinner linux binary with live reload but no debugger
	go tool air -c .air-spinner.toml -build.entrypoint $(APP_LINUX_BIN_OUT)

.PHONY: run-api
run-api:  ## Run the spinner API with live reload with delve debugger
	go tool air -c .air-api.toml -build.entrypoint $(API_BIN_OUT)

.PHONY: run-web
run-web:  ## Run the full web-based frontend
	go tool air -c .air-web.toml -build.entrypoint $(APP_WEB_BIN_OUT)

# ======================
# ---- CODE QUALITY ----
# ======================

.PHONY: gen-docs
gen-docs:  ## Generate API documentation
	go tool swag init -g cmd/api/main.go

.PHONY: format
format:  ## Format codebase
	go fmt ./...
	go tool swag fmt

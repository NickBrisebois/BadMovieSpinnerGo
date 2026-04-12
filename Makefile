SERVER_BIN=./bin/server
WASM_OUT=./web/static/main.wasm
WASM_EXEC=./web/static/wasm_exec.js

LINUX_BIN = ./bin/badmoviespinner

GOROOT=$(shell go env GOROOT)

.PHONY: all build-wasm build-server build-linux copy-wasm-exec

all: build-wasm

build-server:
	go build -o $(SERVER_BIN) ./cmd/server/main.go

copy-wasm-exec:
		cp $(GOROOT)/misc/wasm/wasm_exec.js $(WASM_EXEC)

build-wasm: copy-wasm-exec
	GOOS=js GOARCH=wasm go build -o $(WASM_OUT) ./cmd/spinner/main.go

build-linux:
	CGO_ENABLED=1 go build -o $(LINUX_BIN) ./cmd/spinner/main.go

debug:
	go tool air -c .air-spinner.toml

run:
	go tool air -c .air-spinner.toml -build.entrypoint bin/badmoviespinner

clean:
	rm -f $(WASM_OUT) $(WASM_EXEC)

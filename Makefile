BIN_NAME := BadMovieSpinnerGo
WASM_DIR := web
WASM_FILE := ${WASM_DIR}/app.wasm

.PHONY: build run wasm server dev clean

wasm:
	GOARCH=wasm GOOS=js go build -o ${WASM_FILE}

server:
	go build -o ${BIN_NAME}

build: wasm server

run: build
	./${BIN_NAME}

dev:
	go tool air

clean:
	rm -f ${WASM_FILE} ${BIN_NAME}
	rm -rf tmp/

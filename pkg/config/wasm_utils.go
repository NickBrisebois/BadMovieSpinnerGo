//go:build js && wasm

// WASM-specific implementation of retrieving variable values from environment
// This requires values be set to the JS global `appConfig`
// wasm_utils.go is only built for WASM targets (-tags=wasm)

package config

import "syscall/js"

func getEnvValue(keyName string) (string, bool) {
	return js.Global().Get("appConfig").Get(keyName).String(), true
}

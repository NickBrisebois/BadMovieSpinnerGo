//go:build !wasm

// Native desktop implementation for retrieving variable values from environment
// `wasm_utils.go` is only built for native targets (-tags=native)

package config

import "os"

func getEnvValue(keyName string) (string, bool) {
	return os.LookupEnv(keyName)
}

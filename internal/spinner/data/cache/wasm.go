//go:build wasm

package cache

type wasmCache struct {
}

func NewCache() *wasmCache {
	return &wasmCache{}
}

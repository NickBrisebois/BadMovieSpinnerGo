//go:build !wasm

package cache

type nativeCache struct {
}

func NewCache() *nativeCache {
	return &nativeCache{}
}

//go:build js && wasm

// WASM-specific implementation of image cache
// `wasm.go` is only built for WASM targets (-tags=wasm)

package cache

import (
	"log/slog"

	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
)

type wasmCache struct {
	logger *slog.Logger
}

func NewCache(logger *slog.Logger) (Cache, error) {
	return &wasmCache{
		logger: logger,
	}, nil
}

func (c *wasmCache) GetMoviePoster(key int) ([]byte, bool) {
	return nil, false
}

func (c *wasmCache) PutMoviePoster(key int, data []byte) error {
	return nil
}

func (c *wasmCache) GetMovieList() ([]models.MovieMeta, error) {
	return nil, nil
}

func (c *wasmCache) PutMovieList(movies []models.MovieMeta) error {
	return nil
}

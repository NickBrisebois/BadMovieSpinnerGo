//go:build js && wasm

// WASM-specific implementation of image cache
// Stores cached images and movie list in browser's localStorage
// `wasm.go` is only built for WASM targets (-tags=wasm)

package cache

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"syscall/js"

	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
)

const (
	movieListLocalStoreKey = "spinner:movie_list"
	posterKeyFmt           = "spinner:poster:%d"
)

type wasmCache struct {
	logger     *slog.Logger
	localStore js.Value
}

func NewCache(logger *slog.Logger) (Cache, error) {
	localStore := js.Global().Get("localStorage")
	if localStore.IsUndefined() || localStore.IsNull() {
		return nil, fmt.Errorf("localStorage could not be accessed for caching")
	}
	logger.Debug("wasm cache initialised in browser localStorage")
	return &wasmCache{
		logger:     logger,
		localStore: localStore,
	}, nil
}

func (c *wasmCache) getCacheItem(key string) (string, bool) {
	cacheVal := c.localStore.Call("getItem", key)
	if cacheVal.IsUndefined() || cacheVal.IsNull() {
		return "", false
	}
	return cacheVal.String(), true
}

func (c *wasmCache) setCacheItem(key, value string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if jsError, ok := r.(js.Error); ok {
				err = fmt.Errorf("failed to set cache item: %w", jsError.Error())
			} else {
				err = fmt.Errorf("unknown panic while writing to localStorage: %v", r)
			}
		}
	}()
	c.localStore.Call("setItem", key, value)
	return nil
}

func (c *wasmCache) GetMoviePoster(key int) ([]byte, bool) {
	storeKey := fmt.Sprintf(posterKeyFmt, key)
	encodedVal, ok := c.getCacheItem(storeKey)
	if !ok {
		c.logger.Debug("poster not found in cache", "key", key)
		return nil, false
	}
	data, err := base64.StdEncoding.DecodeString(encodedVal)
	if err != nil {
		c.logger.Error("failed to decode cached poster data from base64", "key", key, "error", err)
		return nil, false
	}

	c.logger.Debug("poster found in cache", "key", key)
	return data, true
}

func (c *wasmCache) PutMoviePoster(key int, data []byte) error {
	storeKey := fmt.Sprintf(posterKeyFmt, key)
	encodedVal := base64.StdEncoding.EncodeToString(data)
	if err := c.setCacheItem(storeKey, encodedVal); err != nil {
		c.logger.Error("failed to set poster in cache", "key", key, "error", err)
		return err
	}
	c.logger.Debug("poster set in cache", "key", key)
	return nil
}

func (c *wasmCache) GetMovieList() ([]models.MovieMeta, error) {
	rawMovies, ok := c.getCacheItem(movieListLocalStoreKey)
	if !ok {
		return nil, fmt.Errorf("movie list not found in cache")
	}
	var payload movieListCachePayload
	if err := json.Unmarshal([]byte(rawMovies), &payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal movie list from cache: %w", err)
	}
	return payload.Movies, nil
}

func (c *wasmCache) PutMovieList(movies []models.MovieMeta) error {
	payload := movieListCachePayload{Movies: movies}
	encoded, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal movie list for cache: %w", err)
	}
	if err := c.setCacheItem(movieListLocalStoreKey, string(encoded)); err != nil {
		return fmt.Errorf("failed to set movie list in cache: %w", err)
	}
	return nil
}

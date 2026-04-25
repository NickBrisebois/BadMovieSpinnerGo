//go:build !wasm

// Native desktop implementation of image cache
// `native.go` is only built for native targets (-tags=native)

package cache

import (
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
)

type fileCache struct {
	cacheDir string
	logger   *slog.Logger
}

func NewCache(logger *slog.Logger) (Cache, error) {
	rootCacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}
	imagesCacheDir := filepath.Join(rootCacheDir, "moviespinner", "images")
	logger.Info("initialising cache directory", "path", imagesCacheDir)
	if err := os.MkdirAll(imagesCacheDir, 0755); err != nil {
		return nil, err
	}
	return &fileCache{
		cacheDir: rootCacheDir,
		logger:   logger,
	}, nil
}

func (c *fileCache) getFilePath(fileName string, fileType string, ext string) (*string, error) {
	cacheDir := filepath.Join(c.cacheDir, "moviespinner")
	if err := os.MkdirAll(filepath.Join(cacheDir, fileType), 0755); err != nil {
		c.logger.Error("failed to create cache directory", "path", filepath.Join(cacheDir, fileType), "error", err)
		return nil, err
	}
	path := filepath.Join(cacheDir, fileType, fmt.Sprintf("%s.%s", fileName, ext))
	return &path, nil
}

func (c *fileCache) atomicWriteFile(filePath string, fileData []byte) error {
	// write cache to a tmp file before replacing any existing caches to prevent corruption
	tmpFile := fmt.Sprintf("%s.tmp", filePath)
	if err := os.WriteFile(tmpFile, fileData, 0644); err != nil {
		return err
	}
	if err := os.Rename(tmpFile, filePath); err != nil {
		return err
	}
	return nil
}

func (c *fileCache) GetMoviePoster(key int) ([]byte, bool) {
	filePath, err := c.getFilePath(strconv.Itoa(key), "image", "jpg")
	if err != nil {
		return nil, false
	}
	data, err := os.ReadFile(*filePath)
	if err != nil {
		return nil, false
	}
	c.logger.Info("cache item found", "key", key)
	return data, true
}

func (c *fileCache) PutMoviePoster(key int, data []byte) error {
	filePath, err := c.getFilePath(strconv.Itoa(key), "image", "jpg")
	if err != nil {
		return err
	}
	if err := c.atomicWriteFile(*filePath, data); err != nil {
		return err
	}

	c.logger.Info("cache item stored", "key", key)
	return nil
}

func (c *fileCache) GetMovieList() ([]models.MovieMeta, error) {
	filePath, err := c.getFilePath("movies", "json", "json")
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(*filePath)
	if err != nil {
		c.logger.Error("cache item not found", "key", "movies")
		return nil, err
	}
	var movieListCache movieListCachePayload
	if err := json.Unmarshal(data, &movieListCache); err != nil {
		return nil, err
	}
	return movieListCache.Movies, nil
}

func (c *fileCache) PutMovieList(movies []models.MovieMeta) error {
	filePath, err := c.getFilePath("movies", "json", "json")
	if err != nil {
		return err
	}
	movieListCache := movieListCachePayload{
		Movies: movies,
	}
	rawMovies, err := json.Marshal(movieListCache)
	if err != nil {
		return err
	}

	return c.atomicWriteFile(*filePath, rawMovies)
}

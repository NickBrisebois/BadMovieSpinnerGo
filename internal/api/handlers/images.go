package handlers

import (
	"log/slog"
	"time"
)

type memCachedImage struct {
	data     []byte
	cachedOn time.Time
}

type ImageHandler struct {
	// TODO: Cache images to disk
	diskImageCacheDir string
	imageMemCache     map[int][]byte
	logger            *slog.Logger
}

func NewImageHandler(diskCacheDir string, logger *slog.Logger) *ImageHandler {
	return &ImageHandler{
		diskImageCacheDir: diskCacheDir,
		imageMemCache:     make(map[int][]byte),
		logger:            logger,
	}
}

func (h *ImageHandler) GetImage(imgKey int) ([]byte, bool) {
	img, found := h.imageMemCache[imgKey]
	return img, found
}

func (h *ImageHandler) CacheImage(image []byte, imgKey int) {
	h.imageMemCache[imgKey] = image
	h.logger.Info("inserted new image into cache", "key", imgKey)
}

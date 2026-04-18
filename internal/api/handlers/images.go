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
	imageMemCache     map[string][]byte
	logger            *slog.Logger
}

func NewImageHandler(diskCacheDir string, logger *slog.Logger) *ImageHandler {
	return &ImageHandler{
		diskImageCacheDir: diskCacheDir,
		imageMemCache:     make(map[string][]byte),
		logger:            logger,
	}
}

func (h *ImageHandler) GetImage(imgKey string) ([]byte, bool) {
	img, found := h.imageMemCache[imgKey]
	return img, found
}

func (h *ImageHandler) CacheImage(image []byte, imgKey string) {
	h.imageMemCache[imgKey] = image
	h.logger.Info("cached image", slog.String("key", imgKey))
}

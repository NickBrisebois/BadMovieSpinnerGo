package render

import "image"

type ImageCache struct {
	images map[string]image.Image
}

func NewImageCache() *ImageCache {
	return &ImageCache{
		images: make(map[string]image.Image),
	}
}

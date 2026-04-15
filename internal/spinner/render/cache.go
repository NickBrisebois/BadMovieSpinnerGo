package render

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type ImageCache struct {
	images map[string]*ebiten.Image
}

func NewImageCache() *ImageCache {
	cache := &ImageCache{
		images: make(map[string]*ebiten.Image),
	}

	cache.images["test.png"] = ebiten.NewImageFromImage(image.NewRGBA(image.Rect(0, 0, 1, 1)))

	return cache
}

func (c *ImageCache) CacheImage(rawImage *image.Image) {

}

func (c *ImageCache) GetImage() *ebiten.Image {
	return c.images["test.png"]
}

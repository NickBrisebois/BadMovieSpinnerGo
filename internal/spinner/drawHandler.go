package spinner

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type DrawHandler struct {
	NumParts int
}

var sliceColours = [...]color.RGBA{
	color.RGBA{242, 145, 153, 255},
	color.RGBA{107, 167, 191, 255},
	color.RGBA{242, 242, 56, 255},
	color.RGBA{242, 90, 56, 255},
	color.RGBA{64, 1, 1, 255},

	// put more unique colours down here
	color.RGBA{242, 145, 153, 255},
	color.RGBA{107, 167, 191, 255},
	color.RGBA{242, 242, 56, 255},
	color.RGBA{242, 90, 56, 255},
	color.RGBA{64, 1, 1, 255},
}

func NewDrawHandler(numParts int) *DrawHandler {
	return &DrawHandler{NumParts: numParts}
}

func (s *DrawHandler) Draw(screen *ebiten.Image, x, y, sizeX, sizeY int) {
	lastEnd := float32(0)

	// strokeOptions := vector.StrokeOptions{Width: 5}
	drawPathOptions := vector.DrawPathOptions{AntiAlias: true}
	fillOptions := vector.FillOptions{}

	for i := 0; i < s.NumParts; i++ {
		path := vector.Path{}
		path.MoveTo(float32(sizeX), float32(sizeY))
		path.Arc(
			float32(sizeX),
			float32(sizeY),
			float32(sizeX),
			lastEnd,
			2*math.Pi*float32(i)/float32(s.NumParts),
			vector.Clockwise,
		)
		path.LineTo(float32(sizeX), float32(sizeY))
		lastEnd += 2 * math.Pi / float32(s.NumParts/2)

		// drawPathOptions.ColorScale.ScaleWithColor(sliceColours[i])
		// vector.StrokePath(screen, &path, &strokeOptions, &drawPathOptions)

		drawPathOptions.ColorScale.ScaleWithColor(sliceColours[i])
		drawPathOptions.Blend = ebiten.BlendClear
		vector.FillPath(screen, &path, &fillOptions, &drawPathOptions)
		drawPathOptions.ColorScale.Reset()
	}

}

package render

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data"
)

type DrawHandler struct {
	Slices     *[]data.Slice
	SliceAngle float32
	CentreX    float32
	CentreY    float32
	RadiusX    float32
	RadiusY    float32

	maskRenderTarget  *ebiten.Image
	sliceRenderTarget *ebiten.Image

	imageCache *ImageCache
}

// outerSliceArcSegments defines how many segments to draw for the outer arc connecting the two corners
// higher values == smoother arc but worsening performance
const outerSliceArcSegments = 20

func NewDrawHandler(
	slices *[]data.Slice,
	sliceAngle, centreX, centreY, radiusX, radiusY float32,
) *DrawHandler {
	return &DrawHandler{
		Slices:            slices,
		SliceAngle:        sliceAngle,
		CentreX:           centreX,
		CentreY:           centreY,
		RadiusX:           radiusX,
		RadiusY:           radiusY,
		maskRenderTarget:  nil,
		sliceRenderTarget: nil,
		imageCache:        NewImageCache(),
	}
}

// addOuterArc draws an arc connecting the two outer corners in an arc shape
func addOuterArc(path *vector.Path, centreX, centreY, radiusX, radiusY, start, end float32) {
	span := end - start
	if span < 0 {
		span = -span
	}

	segments := float32(math.Max(math.Ceil(float64(span/(2*math.Pi/outerSliceArcSegments))), 1))

	for j := 1; j <= int(segments); j++ {
		t := float32(j) / segments
		angle := start + (end-start)*t

		pointX, pointY := GetEllipsePoint(centreX, centreY, radiusX, radiusY, angle)
		path.LineTo(pointX, pointY)
	}

}

func (s *DrawHandler) getSlicePath(
	slice *data.Slice,
) *vector.Path {
	path := vector.Path{}

	// move to the centre of the spinner wheel
	path.MoveTo(s.CentreX, s.CentreY)

	// given the start angle and the circle properties, get the outer corner coords
	ePointX, ePointY := GetEllipsePoint(s.CentreX, s.CentreY, s.RadiusX, s.RadiusY, slice.DrawProperties.StartAngle)
	path.LineTo(ePointX, ePointY)

	// draw a fancy arc to the slice's other outer corner
	addOuterArc(&path, s.CentreX, s.CentreY, s.RadiusX, s.RadiusY, slice.DrawProperties.StartAngle, slice.DrawProperties.EndAngle)

	// back to the middle again
	path.LineTo(s.CentreX, s.CentreY)
	path.Close()
	return &path
}

func (s *DrawHandler) drawMaskToSlice(
	screen *ebiten.Image,
	slice *data.Slice,
) {

}

func (s *DrawHandler) drawSlice(screen *ebiten.Image, slice *data.Slice) {
	s.maskRenderTarget.Clear()

	// generate the vector slice and then create a mask to fill in with a terrible movie's poster
	slicePath := s.getSlicePath(slice)

	// apply the the slice's vector shape to the masking layer
	maskFillOptions := &vector.DrawPathOptions{AntiAlias: true}
	maskFillOptions.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})
	vector.FillPath(s.maskRenderTarget, slicePath, nil, maskFillOptions)

	// Draw the movie poster image to render target to be masked over mask target
	s.sliceRenderTarget.Clear()
	drawImageOptions := &ebiten.DrawImageOptions{}
	s.sliceRenderTarget.DrawImage(s.imageCache.GetImage(), drawImageOptions)

	// Draw poster RT to mask RT
	s.maskRenderTarget.Clear()
	maskCopyOptions := &ebiten.DrawImageOptions{}
	s.maskRenderTarget.DrawImage(s.sliceRenderTarget, maskCopyOptions)

	// Apply mask blend
	maskDrawOptions := &ebiten.DrawImageOptions{
		Blend: ebiten.BlendDestinationIn,
	}
	screen.DrawImage(s.maskRenderTarget, maskDrawOptions)

	// Draw prepared slice to screen
	screenDrawImageOptions := &ebiten.DrawImageOptions{}
	screen.DrawImage(s.maskRenderTarget, screenDrawImageOptions)
}

func (s *DrawHandler) EnsureRenderTargets(screenW, screenH int) {
	if s.maskRenderTarget == nil || s.maskRenderTarget.Bounds().Dx() != screenW || s.maskRenderTarget.Bounds().Dy() != screenH {
		s.maskRenderTarget = ebiten.NewImage(screenW, screenH)
		s.sliceRenderTarget = ebiten.NewImage(screenW, screenH)
	}
}

func (s *DrawHandler) Draw(screen *ebiten.Image) {

	for i := range *s.Slices {
		slice := &(*s.Slices)[i]
		if slice.DrawProperties == nil {
			return
		}

		s.drawSlice(screen, slice)
	}
}

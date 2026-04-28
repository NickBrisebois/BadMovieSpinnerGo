package render

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data"
)

type DrawHandler struct {
	slices     *[]data.Slice
	sliceAngle float32

	maskRenderTarget  *ebiten.Image
	sliceRenderTarget *ebiten.Image
}

// outerSliceArcSegments defines how many segments to draw for the outer arc connecting the two corners
// higher values == smoother arc but worsening performance
const outerSliceArcSegments = 20

func NewDrawHandler(
	slices *[]data.Slice,
	sliceAngle float32,
) *DrawHandler {
	return &DrawHandler{
		slices:            slices,
		sliceAngle:        sliceAngle,
		maskRenderTarget:  nil,
		sliceRenderTarget: nil,
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
	centreX, centreY, radius float32,
) *vector.Path {
	path := vector.Path{}

	// move to the centre of the spinner wheel
	path.MoveTo(centreX, centreY)

	// given the start angle and the circle properties, get the outer corner coords
	ePointX, ePointY := GetEllipsePoint(centreX, centreY, radius, radius, slice.DrawProperties.StartAngle)
	path.LineTo(ePointX, ePointY)

	// draw a fancy arc to the slice's other outer corner
	addOuterArc(&path, centreX, centreY, radius, radius, slice.DrawProperties.StartAngle, slice.DrawProperties.EndAngle)

	// back to the middle again
	path.LineTo(centreX, centreY)
	path.Close()
	return &path
}

// getPosterScale returns the scale factor to apply to the poster image to fit it within the spinner wheel's slice bounds
func (s *DrawHandler) getPosterScale(radius float32, poster *ebiten.Image) float64 {
	initialWidth := float64(poster.Bounds().Dx())
	initialHeight := float64(poster.Bounds().Dy())

	targetWidth := float64(radius * 2)
	targetHeight := float64(radius * 2)

	return math.Max(targetWidth/initialWidth, targetHeight/initialHeight)
}

func (s *DrawHandler) getPosterTranslate(posterScale float64, poster *ebiten.Image) (float64, float64) {
	initialWidth := float64(poster.Bounds().Dx())
	initialHeight := float64(poster.Bounds().Dy())
	targetWidth := float64(posterScale * initialWidth)
	targetHeight := float64(posterScale * initialHeight)
	return (targetWidth / 2), (targetHeight / 2)
}

func (s *DrawHandler) drawSlice(screen *ebiten.Image, spinnerRect image.Rectangle, slice *data.Slice, centreX, centreY, radius float32) {

	// generate the vector slice and then create a mask to fill in with a terrible movie's poster
	s.maskRenderTarget.Clear()
	slicePath := s.getSlicePath(slice, centreX, centreY, radius)

	// apply the the slice's vector shape to the masking layer
	maskFillOptions := &vector.DrawPathOptions{AntiAlias: true}
	maskFillOptions.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})
	vector.FillPath(s.maskRenderTarget, slicePath, nil, maskFillOptions)

	s.sliceRenderTarget.Clear()

	// Draw the movie poster image to render target to be masked over mask target
	posterScale := s.getPosterScale(radius, slice.DrawProperties.SliceImage)
	posterXOffset, posterYOffset := s.getPosterTranslate(posterScale, slice.DrawProperties.SliceImage)
	drawImageOptions := &ebiten.DrawImageOptions{}
	drawImageOptions.GeoM.Scale(posterScale, posterScale)
	drawImageOptions.GeoM.Translate(
		float64(centreX)-posterXOffset,
		float64(centreY)-posterYOffset,
	)
	s.sliceRenderTarget.DrawImage(slice.DrawProperties.SliceImage, drawImageOptions)

	// Draw prepared mask to the render target
	maskCopyOptions := &ebiten.DrawImageOptions{Blend: ebiten.BlendDestinationIn}
	s.sliceRenderTarget.DrawImage(s.maskRenderTarget, maskCopyOptions)

	// Draw rendering to the screen
	screenDrawOptions := &ebiten.DrawImageOptions{}
	screenDrawOptions.GeoM.Translate(
		float64(spinnerRect.Min.X),
		float64(spinnerRect.Min.Y),
	)
	screen.DrawImage(s.sliceRenderTarget, screenDrawOptions)
}

func (s *DrawHandler) EnsureRenderTargets(screenW, screenH int) {
	if s.maskRenderTarget == nil || s.maskRenderTarget.Bounds().Dx() != screenW || s.maskRenderTarget.Bounds().Dy() != screenH {
		s.maskRenderTarget = ebiten.NewImage(screenW, screenH)
		s.sliceRenderTarget = ebiten.NewImage(screenW, screenH)
	}
}

func (s *DrawHandler) Draw(screen *ebiten.Image, spinnerRect image.Rectangle) {
	width := spinnerRect.Dx()
	height := spinnerRect.Dy()
	radius := float32(min(float64(width), float64(height))) / 2
	centreX := float32(width / 2)
	centreY := float32(height / 2)
	s.EnsureRenderTargets(width, height)

	for i := range *s.slices {
		slice := &(*s.slices)[i]
		if slice.DrawProperties == nil {
			return
		}

		s.drawSlice(screen, spinnerRect, slice, centreX, centreY, radius)
	}
}

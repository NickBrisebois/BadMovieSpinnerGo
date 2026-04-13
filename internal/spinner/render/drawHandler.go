package render

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data"
)

type DrawHandler struct {
	Screen     *ebiten.Image
	Slices     *[]data.Slice
	SliceAngle float32
	CentreX    float32
	CentreY    float32
	RadiusX    float32
	RadiusY    float32

	// render targets
	maskRenderTarget  *ebiten.Image
	sliceRenderTarget *ebiten.Image
}

// outerSliceArcSegments defines how many segments to draw for the outer arc connecting the two corners
// higher values == smoother arc but worsening performance
const outerSliceArcSegments = 20

func NewDrawHandler(
	screen *ebiten.Image,
	slices *[]data.Slice,
	sliceAngle, centreX, centreY, radiusX, radiusY float32,
) *DrawHandler {
	return &DrawHandler{screen, slices, sliceAngle, centreX, centreY, radiusX, radiusY, nil, nil}
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

func (s *DrawHandler) drawSlice(slice *data.Slice) {
	slicePath := s.getSlicePath(slice)

	drawPathOptions := &vector.DrawPathOptions{AntiAlias: true}
	drawPathOptions.ColorScale.ScaleWithColor(slice.FillColour)

	vector.FillPath(s.maskRenderTarget, slicePath, nil, drawPathOptions)
}

func (s *DrawHandler) ResetScreen(screen *ebiten.Image) {
	s.Screen = screen
}

func (s *DrawHandler) Draw() {
	// split the spinner wheel into equal parts
	for i := 0; i < len(*s.Slices); i++ {
		slice := (*s.Slices)[i]
		if slice.DrawProperties == nil {
		}

		s.drawSlice(&slice)
	}

}

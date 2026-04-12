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
}

// outerSliceArcSegments defines how many segments to draw for the outer arc connecting the two corners
// higher values == smoother arc but worsening performance
const outerSliceArcSegments = 20

func NewDrawHandler(
	slices *[]data.Slice,
	sliceAngle, centreX, centreY, radiusX, radiusY float32) *DrawHandler {
	return &DrawHandler{slices, sliceAngle, centreX, centreY, radiusX, radiusY}
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

func (s *DrawHandler) drawSlice(
	screen *ebiten.Image,
	centreX, centreY, radiusX, radiusY float32,
	start, end float32,
	fillColour color.RGBA,
	outlineColour color.RGBA,
) {
	path := vector.Path{}

	// move to the centre of the spinner wheel
	path.MoveTo(centreX, centreY)

	// line out to the outer corner of the slice
	ePointX, ePointY := GetEllipsePoint(centreX, centreY, radiusX, radiusY, start)
	path.LineTo(ePointX, ePointY)

	// fancy math to draw a line connecting the two corners in an arc shape
	addOuterArc(&path, centreX, centreY, radiusX, radiusY, start, end)

	// back to the middle again
	path.LineTo(centreX, centreY)

	drawPathOptions := &vector.DrawPathOptions{AntiAlias: true}
	drawPathOptions.ColorScale.ScaleWithColor(fillColour)

	vector.FillPath(screen, &path, nil, drawPathOptions)
}

func (s *DrawHandler) Draw(screen *ebiten.Image, x, y, radiusX, radiusY float32) {
	// split the spinner wheel into equal parts
	for i := 0; i < len(*s.Slices); i++ {
		slice := (*s.Slices)[i]
		if slice.DrawProperties == nil {
		}

		f32i := float32(i)
		start := f32i * s.SliceAngle
		end := (f32i + 1) * s.SliceAngle

		s.drawSlice(
			screen,
			x,
			y,
			radiusX,
			radiusY,
			start,
			end,
			slice.FillColour,
			slice.OutlineColour,
		)
	}

}

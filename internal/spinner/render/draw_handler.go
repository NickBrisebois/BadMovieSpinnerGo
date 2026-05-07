package render

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data"
)

type DrawHandler struct {
	slices     *[]data.Slice
	sliceAngle float32
}

// outerSliceArcSegments defines how many segments to draw for the outer arc connecting the two corners
// higher values == smoother arc but worsening performance
const outerSliceArcSegments = 20

func NewDrawHandler(
	slices *[]data.Slice,
	sliceAngle float32,
) *DrawHandler {
	return &DrawHandler{
		slices:     slices,
		sliceAngle: sliceAngle,
	}
}

// getEbitenRGBAFromColor converts a color.Color to an ebiten compatible RGBA tuple (useful for ebiten vertices)
func getEbitenRGBAFromColor(toConvert color.Color) (cr, cg, cb, ca float32) {
	r, g, b, a := toConvert.RGBA()
	return float32(r) / 0xffff, float32(g) / 0xffff, float32(b) / 0xffff, float32(a) / 0xffff
}

func getVertex(x, y float32, vertexColour color.Color) ebiten.Vertex {
	r, g, b, a := getEbitenRGBAFromColor(vertexColour)
	return ebiten.Vertex{
		DstX:   x,
		DstY:   y,
		ColorR: r,
		ColorG: g,
		ColorB: b,
		ColorA: a,
	}
}

func addOuterArc(vertices *[]ebiten.Vertex, slice *data.Slice, centreX, centreY, radius float32, numSegments int) {
	startAngle := slice.DrawProperties.StartAngle
	endAngle := slice.DrawProperties.EndAngle
	for j := 0; j <= numSegments; j++ {
		t := float32(j) / float32(numSegments)
		angle := startAngle + (endAngle-startAngle)*t
		x, y := GetEllipsePoint(centreX, centreY, radius, radius, angle)
		newVertex := getVertex(x, y, color.White)
		*vertices = append(*vertices, newVertex)
	}
}

func (s *DrawHandler) addVertexUV(vertex *ebiten.Vertex, slice *data.Slice, centreX, centreY, radius float32) {
	drawOpts := slice.DrawProperties

	// get offset from the centre of the spinner
	offsetX := vertex.DstX - centreX
	offsetY := vertex.DstY - centreY

	// determine the rotated offset based on the mid angle
	midAngle := (drawOpts.StartAngle + drawOpts.EndAngle) / 2
	cos := float32(math.Cos(float64(-midAngle)))
	sin := float32(math.Sin(float64(-midAngle)))
	rotationX := offsetX*cos - offsetY*sin
	rotationY := offsetX*sin + offsetY*cos

	// calculate poster coordinates
	posterScale := getPosterScale(radius, drawOpts.SliceImage)
	posterX := rotationX / float32(posterScale)
	posterY := rotationY / float32(posterScale)

	vertex.SrcX = posterX + float32(drawOpts.SliceImage.Bounds().Dx())/2
	vertex.SrcY = posterY + float32(drawOpts.SliceImage.Bounds().Dy())/2
}

func (s *DrawHandler) getSliceVertices(slice *data.Slice, centreX, centreY, radius float32) []ebiten.Vertex {
	vertices := make([]ebiten.Vertex, 0)

	// centre vertex for slice
	centreVertex := getVertex(centreX, centreY, color.White)
	vertices = append(vertices, centreVertex)

	// jump out to outer side and begin arc
	addOuterArc(&vertices, slice, centreX, centreY, radius, outerSliceArcSegments)

	for i := range vertices {
		s.addVertexUV(&vertices[i], slice, centreX, centreY, radius)
	}
	return vertices
}

func (s *DrawHandler) drawSlice(screen *ebiten.Image, slice *data.Slice, centreX, centreY, radius float32) {
	vertices := s.getSliceVertices(slice, centreX, centreY, radius)

	// indices represent the three corner points of the slice triangle
	indices := make([]uint16, outerSliceArcSegments*3)
	for i := 0; i < outerSliceArcSegments; i++ {
		indices[i*3] = 0
		indices[i*3+1] = uint16(i + 1)
		indices[i*3+2] = uint16(i + 2)
	}

	options := &ebiten.DrawTrianglesOptions{Filter: ebiten.FilterLinear}
	screen.DrawTriangles(vertices, indices, slice.DrawProperties.SliceImage, options)
}

func (s *DrawHandler) Draw(screen *ebiten.Image, spinnerRect image.Rectangle) {
	width := spinnerRect.Dx()
	height := spinnerRect.Dy()
	radius := float32(min(float64(width), float64(height))) / 2
	centreX := float32(spinnerRect.Min.X) + float32(width/2)
	centreY := float32(spinnerRect.Min.Y) + float32(height/2)

	for i := range *s.slices {
		slice := &(*s.slices)[i]
		if slice.DrawProperties == nil {
			return
		}

		s.drawSlice(screen, slice, centreX, centreY, radius)
	}
}

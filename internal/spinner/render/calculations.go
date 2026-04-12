package render

import (
	"math"
)

func GetEllipsePoint(centreX, centreY, radiusX, radiusY float32, angle float32) (float32, float32) {
	// return the point on an ellipse perimeter at `angle` (radians)
	x := centreX + radiusX*float32(math.Cos(float64(angle)))
	y := centreY + radiusY*float32(math.Sin(float64(angle)))
	return x, y
}

func GetAnglePerSlice(numSlices int) float32 {
	return float32(2 * math.Pi / float64(numSlices))
}

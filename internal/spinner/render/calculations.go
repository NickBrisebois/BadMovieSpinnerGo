package render

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetDeltaTime() float32 {
	return 1.0 / float32(ebiten.TPS())
}

func GetEllipsePoint(centreX, centreY, radiusX, radiusY float32, angle float32) (float32, float32) {
	// return the point on an ellipse perimeter at `angle` (radians)
	x := centreX + radiusX*float32(math.Cos(float64(angle)))
	y := centreY + radiusY*float32(math.Sin(float64(angle)))
	return x, y
}

func GetSliceAngle(numSlices int) float32 {
	return float32(2 * math.Pi / float64(numSlices))
}

func GetSliceStartAngle(step int, sliceAngle float32) float32 {
	// get the *start* angle (which is where the draw starts for this slice)
	// based on the general sliceAngle
	return float32(step) * sliceAngle
}

func GetSliceEndAngle(step int, sliceAngle float32) float32 {
	// get the *end* angle (which is where the draw ends for this slice)
	// based on the general sliceAngle
	return GetSliceStartAngle(step+1, sliceAngle)
}

func UpdateAngularVelocityFromAcceleration(angularVelocity *float32, angularAcceleration float32, maxVelocity float32) {
	if *angularVelocity >= maxVelocity {
		*angularVelocity = maxVelocity
		return
	}
	*angularVelocity += angularAcceleration * GetDeltaTime()
}

func UpdateRotationFromAngularVelocity(rotation *float32, angularVelocity float32) {
	*rotation += angularVelocity * GetDeltaTime()
}

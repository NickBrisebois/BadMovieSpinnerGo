package drawprop

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type WheelDrawProperties struct {
	SliceAngle float32 // radians
	Rotation   float32 // radians
}

type SliceDrawProperties struct {
	Step       int
	StartAngle float32
	EndAngle   float32
	SliceImage *ebiten.Image
}

func NewSliceDrawProperties(step int, startAngle float32, endAngle float32, sliceImage *ebiten.Image) *SliceDrawProperties {
	return &SliceDrawProperties{
		Step:       step,
		StartAngle: startAngle,
		EndAngle:   endAngle,
		SliceImage: sliceImage,
	}
}

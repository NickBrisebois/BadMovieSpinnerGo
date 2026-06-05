package state

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/render/drawprop"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"

	"github.com/hajimehoshi/ebiten/v2"
)

type Wheel struct {
	IsSpinning       bool
	DrawProperties   *drawprop.WheelDrawProperties
	MotionProperties *WheelMotionProperties
	Slices           *[]Slice
}

type WheelMotionProperties struct {
	AngularVelocity     float32 // radians/sec
	AngularAcceleration float32 // radians/sec^2
	MaxVelocity         float32 // radians/sec
}

type Slice struct {
	ID             int
	Movie          models.MovieMeta
	Label          string
	DrawProperties *drawprop.SliceDrawProperties
}

func NewSlice(id int, step int, startAngle float32, endAngle float32, movie models.MovieMeta, sliceImage *ebiten.Image) *Slice {
	return &Slice{
		ID:             id,
		Movie:          movie,
		Label:          movie.Title,
		DrawProperties: drawprop.NewSliceDrawProperties(step, startAngle, endAngle, sliceImage),
	}
}

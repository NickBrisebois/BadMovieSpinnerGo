package data

import (
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"

	"github.com/hajimehoshi/ebiten/v2"
)

type Wheel struct {
	IsSpinning     bool
	DrawProperties *WheelDrawProperties
	Slices         *[]Slice
}

type WheelDrawProperties struct {
	SliceAngle          float32 // radians
	Rotation            float32 // radians
	AngularVelocity     float32 // radians/sec
	AngularAcceleration float32 // radians/sec^2
	MaxVelocity         float32 // radians/sec
}

type Slice struct {
	ID             int
	Movie          models.MovieMeta
	Label          string
	DrawProperties *SliceDrawProperties
}

type SliceDrawProperties struct {
	Step       int
	StartAngle float32
	EndAngle   float32
	SliceImage *ebiten.Image
}

func NewSlice(id int, step int, startAngle float32, endAngle float32, movie models.MovieMeta, sliceImage *ebiten.Image) *Slice {
	return &Slice{
		ID:             id,
		Movie:          movie,
		Label:          movie.Title,
		DrawProperties: NewSliceDrawProperties(step, startAngle, endAngle, sliceImage),
	}
}

func NewSliceDrawProperties(step int, startAngle float32, endAngle float32, sliceImage *ebiten.Image) *SliceDrawProperties {
	return &SliceDrawProperties{
		Step:       step,
		StartAngle: startAngle,
		EndAngle:   endAngle,
		SliceImage: sliceImage,
	}
}

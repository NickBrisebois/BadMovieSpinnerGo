package data

import (
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"image/color"
)

type Wheel struct {
	IsSpinning     bool
	DrawProperties *WheelDrawProperties
	Slices         *[]Slice
}

type WheelDrawProperties struct {
	Rotation            float32 // radians
	AngularVelocity     float32 // radians/sec
	AngularAcceleration float32 // radians/sec^2
}

type Slice struct {
	ID             int
	Movie          models.MovieMeta
	Label          string
	FillColour     color.RGBA
	DrawProperties *SliceDrawProperties
}

type SliceDrawProperties struct {
	Step       int
	StartAngle float32
	EndAngle   float32
}

func NewSlice(id int, step int, sliceAngle float32, movie models.MovieMeta, fillColour color.RGBA) *Slice {
	return &Slice{
		ID:             id,
		Movie:          movie,
		Label:          movie.Title,
		FillColour:     fillColour,
		DrawProperties: NewSliceDrawProperties(step, sliceAngle),
	}
}

func NewSliceDrawProperties(step int, sliceAngle float32) *SliceDrawProperties {
	return &SliceDrawProperties{
		Step:       step,
		StartAngle: float32(step) * sliceAngle,
		EndAngle:   float32(step+1) * sliceAngle,
	}
}

func GetNextStateSliceDrawProperties(currentDrawProperties *SliceDrawProperties) *SliceDrawProperties {
	return nil
}

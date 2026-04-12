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
	SliceAngle          float32 // radians
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

func NewSlice(id int, step int, startAngle float32, endAngle float32, movie models.MovieMeta, fillColour color.RGBA) *Slice {
	return &Slice{
		ID:             id,
		Movie:          movie,
		Label:          movie.Title,
		FillColour:     fillColour,
		DrawProperties: NewSliceDrawProperties(step, startAngle, endAngle),
	}
}

func NewSliceDrawProperties(step int, startAngle float32, endAngle float32) *SliceDrawProperties {
	return &SliceDrawProperties{
		Step:       step,
		StartAngle: startAngle,
		EndAngle:   endAngle,
	}
}

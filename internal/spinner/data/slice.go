package data

import (
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"image/color"
)

type Slice struct {
	ID             int
	Movie          models.MovieMeta
	Label          string
	FillColour     color.RGBA
	OutlineColour  color.RGBA
	DrawProperties *DrawProperties
}

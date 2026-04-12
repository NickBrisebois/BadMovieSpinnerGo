package spinner

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/render"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"image/color"
)

var (
	outlineColour = color.RGBA{0, 0, 0, 255}
	sliceColours  = [...]color.RGBA{
		{242, 145, 153, 255},
		{107, 167, 191, 255},
		{242, 242, 56, 255},
		{242, 90, 56, 255},
		{64, 1, 1, 255},
		{42, 45, 53, 255},
		{7, 67, 91, 255},
		{42, 242, 56, 255},
		{42, 90, 56, 255},
		{4, 1, 1, 255},
	}
)

type Spinner struct {
	DrawHandler *render.DrawHandler
	Wheel       *data.Wheel
}

func (s *Spinner) Init(centreX, centreY, radiusX, radiusY float32) {
	movies := data.PullMovieList()
	sliceAngle := render.GetAnglePerSlice(len(movies))

	s.Wheel = &data.Wheel{
		IsSpinning: false,
		DrawProperties: &data.WheelDrawProperties{
			Rotation:            0,
			AngularVelocity:     0,
			AngularAcceleration: 0,
		},
		Slices: s.genSlices(sliceAngle, movies),
	}
	s.DrawHandler = render.NewDrawHandler(s.Wheel.Slices, sliceAngle, centreX, centreY, radiusX, radiusY)
}

func (s *Spinner) genSlices(sliceAngle float32, movies []models.MovieMeta) *[]data.Slice {
	slices := make([]data.Slice, 0, len(movies))

	for i, movie := range movies {
		slices = append(slices, *data.NewSlice(
			i,
			i,
			sliceAngle,
			movie,
			sliceColours[i%len(sliceColours)],
		))
	}

	return &slices
}

func (s *Spinner) Update() {
	if s.Wheel.Slices == nil {
		return
	}

	for i := range *s.Wheel.Slices {
		drawProperties := (*s.Wheel.Slices)[i].DrawProperties
		if drawProperties == nil {
			(*s.Wheel.Slices)[i].DrawProperties = data.GetNextStateSliceDrawProperties(drawProperties)
		}
	}
}

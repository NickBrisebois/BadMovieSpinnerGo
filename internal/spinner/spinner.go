package spinner

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/render"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
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
	sliceAngle := render.GetSliceAngle(len(movies))

	// Initialise the wheel with 0'd out animation properties
	wheelDrawProperties := &data.WheelDrawProperties{
		Rotation:            0,
		AngularVelocity:     0.01,
		AngularAcceleration: 0.00001,
		SliceAngle:          sliceAngle,
	}

	s.Wheel = &data.Wheel{
		IsSpinning:     false,
		DrawProperties: wheelDrawProperties,
		Slices:         s.genSlices(sliceAngle, movies),
	}
	s.DrawHandler = render.NewDrawHandler(s.Wheel.Slices, sliceAngle, centreX, centreY, radiusX, radiusY)
}

func (s *Spinner) genSlices(sliceAngle float32, movies []models.MovieMeta) *[]data.Slice {
	slices := make([]data.Slice, 0, len(movies))

	for i, movie := range movies {
		slices = append(slices, *data.NewSlice(
			i,
			i,
			render.GetSliceStartAngle(i, sliceAngle),
			render.GetSliceEndAngle(i, sliceAngle),
			movie,
			sliceColours[i%len(sliceColours)],
		))
	}

	return &slices
}

func (s *Spinner) updateWheelState() {

	render.UpdateAngularVelocityFromAcceleration(&s.Wheel.DrawProperties.AngularVelocity, s.Wheel.DrawProperties.AngularAcceleration)
	render.UpdateRotationFromAngularVelocity(&s.Wheel.DrawProperties.Rotation, s.Wheel.DrawProperties.AngularVelocity)
	rotation := s.Wheel.DrawProperties.Rotation

	// Update the wheel's rotation based on the updated slice angles
	for i := range *s.Wheel.Slices {
		sliceDrawProperties := (*s.Wheel.Slices)[i].DrawProperties
		if sliceDrawProperties != nil {
			sliceDrawProperties.StartAngle += rotation
			sliceDrawProperties.EndAngle += rotation
		}
	}
}

func (s *Spinner) Update() {
	if s.Wheel.Slices == nil || !s.Wheel.IsSpinning {
		return
	}

	s.updateWheelState()
}

func (s *Spinner) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	s.DrawHandler.EnsureRenderTargets(screen.Bounds().Dx(), screen.Bounds().Dy())
	s.DrawHandler.Draw(screen)
}

package spinner

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/render"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"image/color"
	"log/slog"

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

type SpinnerHandler struct {
	drawHandler *render.DrawHandler
	movieData   *data.MovieDataHandler
	wheel       *data.Wheel
	logger      *slog.Logger
}

func NewSpinner(movieDataHandler *data.MovieDataHandler, screenWidth, screenHeight int, logger *slog.Logger) *SpinnerHandler {
	spinner := &SpinnerHandler{
		movieData: movieDataHandler,
		logger:    logger,
	}

	spinner.Init(
		float32(screenWidth)/2,
		float32(screenHeight)/2,
		float32(screenWidth)/2,
		float32(screenHeight)/2,
	)

	return spinner
}

func (s *SpinnerHandler) Init(centreX, centreY, radiusX, radiusY float32) {
	movies := s.movieData.GetMovieList()
	movies = movies[:7]
	sliceAngle := render.GetSliceAngle(len(movies))

	// Initialise the wheel with 0'd out animation properties
	wheelDrawProperties := &data.WheelDrawProperties{
		SliceAngle:          sliceAngle,
		Rotation:            0,
		AngularVelocity:     0.01,
		AngularAcceleration: 0.01,
		MaxVelocity:         0.01,
	}

	s.wheel = &data.Wheel{
		IsSpinning:     true,
		DrawProperties: wheelDrawProperties,
		Slices:         s.genSlices(sliceAngle, movies),
	}
	s.drawHandler = render.NewDrawHandler(s.wheel.Slices, sliceAngle, centreX, centreY, radiusX, radiusY)
}

func (s *SpinnerHandler) genSlices(sliceAngle float32, movies []models.MovieMeta) *[]data.Slice {
	slices := make([]data.Slice, 0, len(movies))

	for i, movie := range movies {
		moviePoster := s.movieData.GetMoviePoster(movie.TMDBId)
		if moviePoster == nil {
			s.logger.Error("failed to get movie poster, not adding to spinner", "TMDBId", movie.TMDBId)
			continue
		}
		slices = append(slices, *data.NewSlice(
			i,
			i,
			render.GetSliceStartAngle(i, sliceAngle),
			render.GetSliceEndAngle(i, sliceAngle),
			movie,
			sliceColours[i%len(sliceColours)],
			moviePoster,
		))
	}

	return &slices
}

func (s *SpinnerHandler) updateWheelState() {
	render.UpdateAngularVelocityFromAcceleration(
		&s.wheel.DrawProperties.AngularVelocity,
		s.wheel.DrawProperties.AngularAcceleration,
		s.wheel.DrawProperties.MaxVelocity,
	)
	render.UpdateRotationFromAngularVelocity(&s.wheel.DrawProperties.Rotation, s.wheel.DrawProperties.AngularVelocity)
	rotation := s.wheel.DrawProperties.Rotation

	// Update the wheel's rotation based on the updated slice angles
	for i := range *s.wheel.Slices {
		sliceDrawProperties := (*s.wheel.Slices)[i].DrawProperties
		if sliceDrawProperties != nil {
			sliceDrawProperties.StartAngle += rotation
			sliceDrawProperties.EndAngle += rotation
		}
	}
}

func (s *SpinnerHandler) Update() {
	if s.wheel.Slices == nil || !s.wheel.IsSpinning {
		return
	}

	s.updateWheelState()
}

func (s *SpinnerHandler) Draw(screen *ebiten.Image) {

	s.drawHandler.Draw(screen)
}

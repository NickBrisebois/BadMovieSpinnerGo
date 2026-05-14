package spinner

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data/external"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data/filters"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/render"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpinnerHandler struct {
	initialised  bool
	config       *SpinnerConfig
	screenWidth  int
	screenHeight int
	drawHandler  *render.DrawHandler
	uiHandler    *ui.UIHandler
	movieData    *data.MovieDataHandler
	wheel        *data.Wheel
	logger       *slog.Logger
}

func NewSpinner(
	config *SpinnerConfig,
	screenWidth, screenHeight int,
	logger *slog.Logger,
) (*SpinnerHandler, error) {
	// APIs!
	apiBaseURL, err := config.ServerURL()
	if err != nil {
		logger.Error("failed to parse API server URL", "error", err)
		return nil, err
	}
	spinnerAPI := external.NewSpinnerAPI(apiBaseURL, logger)

	// Data!
	moviesDataHandler := data.NewMovieDataHandler(spinnerAPI, logger)

	// Logic!
	return &SpinnerHandler{
		uiHandler:    ui.NewUIHandler(screenWidth, screenHeight, logger),
		config:       config,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		movieData:    moviesDataHandler,
		drawHandler:  nil, // dependent on UI so initialised during first draw
		logger:       logger,
	}, nil
}

func (s *SpinnerHandler) initDrawHandler() {
	movies := s.movieData.GetMovieList(
		&data.GetMovieListOptions{
			Filters: &filters.MovieFilters{
				Watched: filters.WatchedStatusUnwatched,
			},
		},
	)
	sliceAngle := render.GetSliceAngle(len(movies))

	// Initialise the wheel with 0'd out animation properties
	wheelDrawProperties := &data.WheelDrawProperties{
		SliceAngle:          sliceAngle,
		Rotation:            0,
		AngularVelocity:     0.05,
		AngularAcceleration: 0.002,
		MaxVelocity:         0.1,
	}

	s.wheel = &data.Wheel{
		IsSpinning:     false,
		DrawProperties: wheelDrawProperties,
		Slices:         s.genSlices(sliceAngle, movies),
	}
	s.drawHandler = render.NewDrawHandler(s.wheel.Slices, sliceAngle)
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

func (s *SpinnerHandler) HandleScreenResize() {

}

func (s *SpinnerHandler) Update() error {
	s.uiHandler.Update()

	if !s.initialised || s.wheel.Slices == nil || !s.wheel.IsSpinning {
		return nil
	}
	s.updateWheelState()
	return nil
}

func (s *SpinnerHandler) Draw(screen *ebiten.Image) {
	spinnerRect := s.uiHandler.GetSpinnerBoxRect()
	if !s.initialised && !spinnerRect.Empty() {
		// The spinner has to be initialised after the first UI draw since
		// the spinner box widget's dimensions are only calculated during that
		s.initDrawHandler()
		s.initialised = true
	}

	s.uiHandler.DrawUI(screen)

	if s.initialised {
		// The actual "game" or spinner is drawn to a subimage that's sized to the container provided by the UI handler
		spinnerScreen := screen.SubImage(spinnerRect).(*ebiten.Image)
		s.drawHandler.Draw(spinnerScreen, spinnerRect)
	}
}

func (s *SpinnerHandler) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.screenWidth, s.screenHeight
}

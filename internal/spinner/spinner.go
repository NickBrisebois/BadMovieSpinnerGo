package spinner

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data/external"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data/processing"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/render"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpinnerHandler struct {
	initialised bool
	config      *SpinnerConfig

	logicalScreenWidth  int
	logicalScreenHeight int

	drawHandler *render.DrawHandler
	uiHandler   *ui.UIHandler
	movieData   *data.MovieDataHandler
	wheel       *data.Wheel
	logger      *slog.Logger
}

func NewSpinner(
	config *SpinnerConfig,
	logicalScreenWidth, logicalScreenHeight int,
	logger *slog.Logger,
) (*SpinnerHandler, error) {
	spinnerHandler := &SpinnerHandler{
		uiHandler:           nil,
		config:              config,
		logicalScreenWidth:  logicalScreenWidth,
		logicalScreenHeight: logicalScreenHeight,
		movieData:           nil,
		drawHandler:         nil, // dependent on UI so initialised during first draw
		logger:              logger,
	}
	apiBaseURL, err := config.ServerURL()
	if err != nil {
		logger.Error("failed to parse API server URL", "error", err)
		return nil, err
	}
	spinnerAPI := external.NewSpinnerAPI(apiBaseURL, logger)
	moviesDataHandler := data.NewMovieDataHandler(spinnerAPI, logger)
	spinnerHandler.movieData = moviesDataHandler
	spinnerHandler.uiHandler = ui.NewUIHandler(logicalScreenWidth, logicalScreenHeight, logger, spinnerHandler.uiEventCallback)

	moviesBySuggestedBy := moviesDataHandler.GetMoviesBySuggestedBy(nil)
	spinnerHandler.uiHandler.SetMovies(&moviesBySuggestedBy)

	return spinnerHandler, nil
}

func (s *SpinnerHandler) uiEventCallback(data *ui.UIEventCallbackData) {
	switch data.EventType {
	case ui.UIIEventTypeSuggestedByChanged:
		s.logger.Info("suggested by changed, reinventing the wheel", "suggestedBy", data.SuggestedUsers)
		s.rebuildWheel(data.SuggestedUsers)
	default:
	}
}

func (s *SpinnerHandler) rebuildWheel(suggestedBy *[]models.PersonName) {
	movies := s.movieData.GetMovieList(
		&data.GetMovieListOptions{
			Filters: &processing.MovieFilters{
				Watched:     processing.WatchedStatusUnwatched,
				SuggestedBy: suggestedBy,
			},
		},
	)
	if len(movies) == 0 {
		s.logger.Warn("no movies match filters and can't build a wheel with no slices")
		return
	}

	sliceAngle := render.GetSliceAngle(len(movies))
	s.wheel = &data.Wheel{
		IsSpinning: false,
		DrawProperties: &data.WheelDrawProperties{
			SliceAngle:          sliceAngle,
			Rotation:            0,
			AngularVelocity:     0.05,
			AngularAcceleration: 0.002,
			MaxVelocity:         0.1,
		},
		Slices: s.genSlices(sliceAngle, movies),
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
		s.rebuildWheel(nil)
		s.initialised = true
		return
	}

	if s.initialised {
		// The actual "game" or spinner is drawn to a subimage that's sized to the container provided by the UI handler
		spinnerScreen := screen.SubImage(spinnerRect).(*ebiten.Image)
		s.drawHandler.Draw(spinnerScreen)
	}

	s.uiHandler.DrawUI(screen)
}

func (s *SpinnerHandler) Layout(outsideWidth, outsideHeight int) (int, int) {
	scale := ebiten.Monitor().DeviceScaleFactor()
	scaledW := int(float64(outsideWidth) * scale)
	scaledH := int(float64(outsideHeight) * scale)

	if outsideWidth != s.logicalScreenWidth || outsideHeight != s.logicalScreenHeight {
		s.logger.Info("Screen resize detected", "width", outsideWidth, "height", outsideHeight)
		s.uiHandler.SetDimensions(scaledW, scaledH)
	}

	s.logicalScreenWidth = outsideWidth
	s.logicalScreenHeight = outsideHeight
	return scaledW, scaledH
}

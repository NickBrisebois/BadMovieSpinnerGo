package main

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data/external"
	"NickBrisebois/BadMovieSpinnerGo/pkg/config"
	"log"
	"log/slog"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 600
	screenHeight = 600

	winTitle = "Bad Movie Spinner"
)

// SpinnerGame is the main ebitengine "game" object
// It implements the functions from the ebitengine.Game interface
type SpinnerGame struct {
	spinner *spinner.SpinnerHandler
}

func NewSpinnerGame(spinnerConfig *spinner.SpinnerConfig, logger *slog.Logger) (*SpinnerGame, error) {
	// APIs!
	apiBaseURL, err := spinnerConfig.ServerURL()
	if err != nil {
		logger.Error("failed to parse API server URL", "error", err)
		return nil, err
	}
	spinnerAPI := external.NewSpinnerAPI(apiBaseURL, logger)

	// Data!
	moviesDataHandler := data.NewMovieDataHandler(spinnerAPI, logger)

	// Logic!
	spinner := spinner.NewSpinner(moviesDataHandler, screenWidth, screenHeight, logger)

	return &SpinnerGame{spinner: spinner}, nil
}

func (g *SpinnerGame) Update() error {
	g.spinner.Update()
	return nil
}

func (g *SpinnerGame) Draw(screen *ebiten.Image) {
	g.spinner.Draw(screen)
}

func (g *SpinnerGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	spinnerConfig := &spinner.SpinnerConfig{}
	err := config.LoadConfig(spinnerConfig)
	if err != nil {
		logger.Error("failed to load spinner config", "error", err)
		return
	}

	game, err := NewSpinnerGame(spinnerConfig, logger)
	if err != nil {
		logger.Error("failed to initialise spinner game instance", "error", err)
		return
	}

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(winTitle)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

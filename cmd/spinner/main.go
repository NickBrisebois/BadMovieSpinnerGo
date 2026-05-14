package main

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner"
	"NickBrisebois/BadMovieSpinnerGo/pkg/config"
	"log"
	"log/slog"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	initScreenWidth  = 1200
	initScreenHeight = 600

	winTitle = "Bad Movie Spinner"
	binName  = "moviespinner"
)

// Config values injected during build via `-ldflags` and then mapped
// to allow for string-based lookup to match native-build ENV var lookups
// Note: if these aren't passed during compilation, they're grabbed from environment variables
var (
	APIHost           string
	APIPort           string
	LookupOverrideMap = map[string]*string{
		"API_HOST": &APIHost,
		"API_PORT": &APIPort,
	}
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	logger.Debug("starting spinner", "Backend Host", APIHost, "Backend Port", APIPort)

	spinnerConfig := &spinner.SpinnerConfig{}
	err := config.LoadConfig(spinnerConfig, &config.ConfigOptions{
		EnvOverrideMap: &LookupOverrideMap,
		Logger:         logger,
	})
	if err != nil {
		logger.Error("failed to load spinner config", "error", err)
		return
	}

	game, err := spinner.NewSpinner(
		spinnerConfig,
		initScreenWidth,
		initScreenHeight,
		logger,
	)
	if err != nil {
		logger.Error("failed to initialise spinner game instance", "error", err)
		return
	}

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(initScreenWidth, initScreenHeight)
	ebiten.SetWindowTitle(winTitle)
	ebiten.SetScreenClearedEveryFrame(true)
	ebiten.SetVsyncEnabled(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

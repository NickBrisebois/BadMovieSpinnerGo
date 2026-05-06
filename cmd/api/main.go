package main

import (
	"log/slog"
	"os"

	"NickBrisebois/BadMovieSpinnerGo/internal/api"
	"NickBrisebois/BadMovieSpinnerGo/pkg/config"

	_ "NickBrisebois/BadMovieSpinnerGo/docs"
)

// MovieSpinner API godoc
// note on docs: https://github.com/swaggo/swag/issues/2045#issuecomment-3892744315
//
//	@title			Bad Movie Spinner API
//	@version		1.0
//	@description	Backend API for communication with Google Sheets list of movies and TMDB
//	@host			localhost:8080
//	@BasePath		/
func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	apiConfig := &api.Config{}
	err := config.LoadConfig(apiConfig, &config.ConfigOptions{
		Logger:         logger,
		ParsingOptions: &config.ParsingOptions{SliceDelimiter: ","},
	})
	if err != nil {
		logger.Error("failed to load config", "err", err)
		os.Exit(1)
	}

	movie_api, err := api.NewAPIServer(apiConfig, logger)
	if err != nil {
		logger.Error("failed to create API server", "err", err)
		os.Exit(1)
	}

	logger.Info("starting movie spinner API", "addr", movie_api.Addr)

	if err := movie_api.ListenAndServe(); err != nil {
		logger.Error("server is dead", "err", err)
		os.Exit(1)
	}
}

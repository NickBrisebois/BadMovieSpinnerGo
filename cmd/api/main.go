package main

import (
	"log/slog"
	"os"

	"NickBrisebois/BadMovieSpinnerGo/internal/api"

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
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	config := &api.Config{}
	err := api.LoadConfig(config)
	if err != nil {
		logger.Error("Failed to load config", "err", err)
		os.Exit(1)
	}

	movie_api := api.NewAPIServer(config, logger)

	logger.Info("Starting movie spinner API", "addr", movie_api.Addr)

	if err := movie_api.ListenAndServe(); err != nil {
		logger.Error("Server is dead", "err", err)
		os.Exit(1)
	}
}

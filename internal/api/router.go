package api

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/api/views"
	"log/slog"
	"net/http"
)

func NewRouter(config *Config, logger *slog.Logger) (http.Handler, error) {
	mux := http.NewServeMux()

	gsMoviesView, err := views.NewGSheetsView(
		config.Auth.GCPServiceAccountKeyPath,
		config.GSheets.SheetID,
		config.Auth.TMDBReadAccessToken,
		config.Cache.ImageCacheDir,
		logger,
	)
	if err != nil {
		logger.Error("failed to create gsheets handler", slog.Any("error", err))
		return nil, err
	}

	mux.HandleFunc("/healthz", views.GetHealthz)
	mux.HandleFunc("/sheets/movies", gsMoviesView.GetMovieList)
	mux.HandleFunc("/sheets/movies/{tmdbID}/poster", gsMoviesView.GetMoviePoster)

	return mux, nil
}

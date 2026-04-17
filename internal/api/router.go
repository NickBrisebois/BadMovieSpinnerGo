package api

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/api/handlers"
	"NickBrisebois/BadMovieSpinnerGo/internal/api/views"
	"log/slog"
	"net/http"
)

func NewRouter(config *Config, logger *slog.Logger) (http.Handler, error) {
	mux := http.NewServeMux()

	gSheetsHandler, err := views.NewGSheetsView(
		config.Auth.GCPServiceAccountKeyPath,
		config.GSheets.SheetID,
		config.Auth.TMDBReadAccessToken,
		logger,
	)
	if err != nil {
		logger.Error("failed to create gsheets handler", slog.Any("error", err))
		return nil, err
	}

	mux.HandleFunc("/healthz", handlers.GetHealthz)
	mux.HandleFunc("/sheets/movies", gSheetsHandler.GetMovieList)

	return mux, nil
}

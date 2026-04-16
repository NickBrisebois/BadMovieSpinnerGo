package api

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/api/handlers"
	"log/slog"
	"net/http"
)

func NewRouter(config *Config, logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()

	gSheetsHandler, _ := handlers.NewGSheetsReqHandler(
		config.Auth.GCPServiceAccountKeyPath,
		config.Auth.GCPServiceAccountKeyPath,
		logger,
	)

	mux.HandleFunc("/healthz", handlers.GetHealthz)
	mux.HandleFunc("/sheets/movies", gSheetsHandler.GetMovieList)

	return mux
}
